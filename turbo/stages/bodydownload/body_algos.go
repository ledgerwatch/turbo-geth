package bodydownload

import (
	"context"
	"fmt"
	"math/big"

	"github.com/holiman/uint256"
	"github.com/ledgerwatch/turbo-geth/common"
	"github.com/ledgerwatch/turbo-geth/common/debug"
	"github.com/ledgerwatch/turbo-geth/consensus"
	"github.com/ledgerwatch/turbo-geth/core"
	"github.com/ledgerwatch/turbo-geth/core/rawdb"
	"github.com/ledgerwatch/turbo-geth/core/types"
	"github.com/ledgerwatch/turbo-geth/eth/stagedsync/stages"
	"github.com/ledgerwatch/turbo-geth/ethdb"
	"github.com/ledgerwatch/turbo-geth/log"
	"github.com/ledgerwatch/turbo-geth/turbo/adapter"
	"github.com/ledgerwatch/turbo-geth/turbo/stages/headerdownload"
)

const BlockBufferSize = 128

// UpdateFromDb reads the state of the database and refreshes the state of the body download
func (bd *BodyDownload) UpdateFromDb(db ethdb.RwTx) (headHeight uint64, headHash common.Hash, headTd256 *uint256.Int, err error) {
	var headerProgress, bodyProgress uint64
	headerProgress, err = stages.GetStageProgress(db, stages.Headers)
	if err != nil {
		return 0, common.Hash{}, nil, err
	}
	bodyProgress, err = stages.GetStageProgress(db, stages.Bodies)
	if err != nil {
		return 0, common.Hash{}, nil, err
	}
	bd.lock.Lock()
	defer bd.lock.Unlock()
	bd.maxProgress = headerProgress + 1
	// Resetting for requesting a new range of blocks
	bd.requestedLow = bodyProgress + 1
	bd.lowWaitUntil = 0
	bd.requestHigh = bd.requestedLow + (bd.outstandingLimit / 2)
	bd.requestedMap = make(map[DoubleHash]uint64)
	bd.delivered.Clear()
	bd.deliveredCount = 0
	bd.wastedCount = 0
	for i := 0; i < len(bd.deliveries); i++ {
		bd.deliveries[i] = nil
		bd.requests[i] = nil
	}
	bd.peerMap = make(map[string]int)
	headHeight = bodyProgress
	headHash, err = rawdb.ReadCanonicalHash(db, headHeight)
	if err != nil {
		return 0, common.Hash{}, nil, err
	}
	var headTd *big.Int
	headTd, err = rawdb.ReadTd(db, headHash, headHeight)
	if err != nil {
		return 0, common.Hash{}, nil, fmt.Errorf("reading total difficulty for head height %d and hash %x: %d, %w", headHeight, headHash, headTd, err)
	}
	if headTd == nil {
		headTd = new(big.Int)
	}
	headTd256 = new(uint256.Int)
	headTd256.SetFromBig(headTd)
	return headHeight, headHash, headTd256, nil
}

// RequestMoreBodies - returns nil if nothing to request
func (bd *BodyDownload) RequestMoreBodies(db ethdb.Tx, blockNum uint64, currentTime uint64, blockPropagator adapter.BlockPropagator) (*BodyRequest, uint64, error) {
	bd.lock.Lock()
	defer bd.lock.Unlock()
	if blockNum < bd.requestedLow {
		blockNum = bd.requestedLow
	}
	var bodyReq *BodyRequest
	blockNums := make([]uint64, 0, BlockBufferSize)
	hashes := make([]common.Hash, 0, BlockBufferSize)
	for ; len(blockNums) < BlockBufferSize && bd.requestedLow <= bd.maxProgress; blockNum++ {
		// Check if we reached highest allowed request block number, and turn back
		if blockNum >= bd.requestedLow+bd.outstandingLimit || blockNum >= bd.maxProgress {
			blockNum = 0
			break // Avoid tight loop
		}
		if bd.delivered.Contains(blockNum) {
			// Already delivered, no need to request
			continue
		}
		req := bd.requests[blockNum-bd.requestedLow]
		if req != nil {
			if currentTime < req.waitUntil {
				continue
			}
			bd.peerMap[string(req.peerID)]++
			bd.requests[blockNum-bd.requestedLow] = nil
		}
		var hash common.Hash
		var header *types.Header
		var err error
		request := true
		if bd.deliveries[blockNum-bd.requestedLow] != nil {
			// If this block was requested before, we don't need to fetch the headers from the database the second time
			header = bd.deliveries[blockNum-bd.requestedLow].Header()
			if header == nil {
				return nil, 0, fmt.Errorf("header not found: %w, blockNum=%d, trace=%s", err, blockNum, debug.Callers(7))
			}
			hash = header.Hash()
		} else {
			hash, err = rawdb.ReadCanonicalHash(db, blockNum)
			if err != nil {
				return nil, 0, fmt.Errorf("could not find canonical header: %w, blockNum=%d, trace=%s", err, blockNum, debug.Callers(7))
			}
			header = rawdb.ReadHeader(db, hash, blockNum)
			if header == nil {
				return nil, 0, fmt.Errorf("header not found: %w, blockNum=%d, trace=%s", err, blockNum, debug.Callers(7))
			}

			if block := bd.prefetchedBlocks.Pop(hash); block != nil {
				// Block is prefetched, no need to request
				bd.deliveries[blockNum-bd.requestedLow] = block

				// Calculate the TD of the block (it's not imported yet, so block.Td is not valid)
				var td *big.Int
				if parent, err := rawdb.ReadTd(db, block.ParentHash(), block.NumberU64()-1); err != nil {
					log.Error("Failed to ReadTd", "err", err, "number", block.NumberU64()-1, "hash", block.ParentHash())
				} else if parent != nil {
					td = new(big.Int).Add(block.Difficulty(), parent)
					go blockPropagator.BroadcastNewBlock(context.Background(), block, td)
				} else {
					log.Error("Propagating dangling block", "number", block.Number(), "hash", hash)
				}
				request = false
			} else {
				bd.deliveries[blockNum-bd.requestedLow] = types.NewBlockWithHeader(header) // Block without uncles and transactions
				if header.UncleHash != types.EmptyUncleHash || header.TxHash != types.EmptyRootHash {
					var doubleHash DoubleHash
					copy(doubleHash[:], header.UncleHash.Bytes())
					copy(doubleHash[common.HashLength:], header.TxHash.Bytes())
					bd.requestedMap[doubleHash] = blockNum
				} else {
					request = false
				}
			}
		}
		if request {
			blockNums = append(blockNums, blockNum)
			hashes = append(hashes, hash)
		} else {
			// Both uncleHash and txHash are empty (or block is prefetched), no need to request
			bd.delivered.Add(blockNum)
		}
	}
	if len(blockNums) > 0 {
		bodyReq = &BodyRequest{BlockNums: blockNums, Hashes: hashes}
		for _, blockNum := range blockNums {
			bd.requests[blockNum-bd.requestedLow] = bodyReq
		}
	}
	return bodyReq, blockNum, nil
}

func (bd *BodyDownload) RequestSent(bodyReq *BodyRequest, timeWithTimeout uint64, peer []byte) {
	bd.lock.Lock()
	defer bd.lock.Unlock()
	for _, blockNum := range bodyReq.BlockNums {
		if blockNum < bd.requestedLow {
			continue
		}
		req := bd.requests[blockNum-bd.requestedLow]
		if req != nil {
			bd.requests[blockNum-bd.requestedLow].waitUntil = timeWithTimeout
			bd.requests[blockNum-bd.requestedLow].peerID = peer
		}
	}
}

// DeliverBodies takes the block body received from a peer and adds it to the various data structures
func (bd *BodyDownload) DeliverBodies(txs [][]types.Transaction, uncles [][]*types.Header) (int, int) {
	bd.lock.Lock()
	defer bd.lock.Unlock()
	reqMap := make(map[uint64]*BodyRequest)
	var delivered, undelivered int
	for i := range txs {
		uncleHash := types.CalcUncleHash(uncles[i])
		txHash := types.DeriveSha(types.Transactions(txs[i]))
		var doubleHash DoubleHash
		copy(doubleHash[:], uncleHash.Bytes())
		copy(doubleHash[common.HashLength:], txHash.Bytes())
		// Block numbers are added to the bd.delivered bitmap here, only for blocks for which the body has been received, and their double hashes are present in the bd.requesredMap
		// Also, block numbers can be added to bd.delivered for empty blocks, above
		if blockNum, ok := bd.requestedMap[doubleHash]; ok {
			bd.delivered.Add(blockNum)
			block := bd.deliveries[blockNum-bd.requestedLow].WithBody(txs[i], uncles[i])
			bd.deliveries[blockNum-bd.requestedLow] = block
			req := bd.requests[blockNum-bd.requestedLow]
			if req != nil {
				if _, ok := reqMap[req.BlockNums[0]]; !ok {
					reqMap[req.BlockNums[0]] = req
				}
			}
			delete(bd.requestedMap, doubleHash) // Delivered, cleaning up
			delivered++
		} else {
			undelivered++
		}
	}
	// Clean up the requests
	for _, req := range reqMap {
		for _, blockNum := range req.BlockNums {
			bd.requests[blockNum-bd.requestedLow] = nil
		}
	}
	select {
	case bd.DeliveryNotify <- struct{}{}:
	default:
	}
	return delivered, undelivered
}

func (bd *BodyDownload) DeliverySize(delivered float64, wasted float64) {
	bd.lock.Lock()
	defer bd.lock.Unlock()
	bd.deliveredCount += delivered
	bd.wastedCount += wasted
}

// ValidateBody validates the given block's uncles and verifies the block
// header's transaction and uncle roots. The headers are assumed to be already
// validated at this point.
func (bd *BodyDownload) ValidateBody(block *types.Block, r consensus.ChainReader) (headerdownload.Penalty, error) {
	bd.lock.Lock()
	defer bd.lock.Unlock()

	// Check whether the block's known, and if not, that it's linkable
	foundInDB := r.GetBlock(block.Hash(), block.NumberU64())
	if foundInDB != nil {
		return headerdownload.BadBlockPenalty, core.ErrKnownBlock
	}

	// Header validity is known at this point, check the uncles and transactions
	header := block.Header()
	if err := bd.Engine.VerifyUncles(r, block); err != nil {
		return headerdownload.BadBlockPenalty, err
	}
	if hash := types.CalcUncleHash(block.Uncles()); hash != header.UncleHash {
		return headerdownload.BadBlockPenalty, fmt.Errorf("uncle root hash mismatch: have %x, want %x", hash, header.UncleHash)
	}
	if hash := types.DeriveSha(block.Transactions()); hash != header.TxHash {
		return headerdownload.BadBlockPenalty, fmt.Errorf("transaction root hash mismatch: have %x, want %x", hash, header.TxHash)
	}
	return headerdownload.NoPenalty, nil
}

func (bd *BodyDownload) GetDeliveries(validateBody func(block *types.Block) (headerdownload.Penalty, error)) ([]*types.Block, []headerdownload.PenaltyItem) {
	bd.lock.Lock()
	defer bd.lock.Unlock()
	var i uint64
	for i = 0; !bd.delivered.IsEmpty() && bd.requestedLow+i == bd.delivered.Minimum(); i++ {
		bd.delivered.Remove(bd.requestedLow + i)
	}
	// Move the deliveries back
	// bd.requestedLow can only be moved forward if there are consecutive block numbers present in the bd.delivered map
	var d []*types.Block
	var penalties []headerdownload.PenaltyItem
	if i > 0 {
		d = make([]*types.Block, i)
		copy(d, bd.deliveries[:i])
		copy(bd.deliveries[:], bd.deliveries[i:])
		for j := uint64(0); j < i; j++ {
			penalty, err := validateBody(d[j])
			if penalty == headerdownload.NoPenalty {
				continue
			}
			if err != nil {
				// Log them?
			}
			penalties = append(penalties, headerdownload.PenaltyItem{PeerID: string(bd.requests[j].peerID), Reason: headerdownload.BadBlockPenalty})
		}

		copy(bd.requests[:], bd.requests[i:])
		for j := len(bd.deliveries) - int(i); j < len(bd.deliveries); j++ {
			bd.deliveries[j] = nil
			bd.requests[j] = nil
		}
		bd.requestedLow += i
	}
	return d, penalties
}

func (bd *BodyDownload) DeliveryCounts() (float64, float64) {
	bd.lock.Lock()
	defer bd.lock.Unlock()
	return bd.deliveredCount, bd.wastedCount
}

func (bd *BodyDownload) GetPenaltyPeers() [][]byte {
	bd.lock.Lock()
	defer bd.lock.Unlock()
	peers := make([][]byte, len(bd.peerMap))
	i := 0
	for p := range bd.peerMap {
		peers[i] = []byte(p)
		i++
	}
	return peers
}

func (bd *BodyDownload) PrintPeerMap() {
	bd.lock.Lock()
	defer bd.lock.Unlock()
	fmt.Printf("---------------------------\n")
	for p, n := range bd.peerMap {
		fmt.Printf("%s = %d\n", p, n)
	}
	fmt.Printf("---------------------------\n")
	bd.peerMap = make(map[string]int)
}

func (bd *BodyDownload) AddToPrefetch(block *types.Block) {
	if hash := types.CalcUncleHash(block.Uncles()); hash != block.UncleHash() {
		log.Warn("Propagated block has invalid uncles", "have", hash, "exp", block.UncleHash())
		return
	}
	if hash := types.DeriveSha(block.Transactions()); hash != block.TxHash() {
		log.Warn("Propagated block has invalid body", "have", hash, "exp", block.TxHash())
		return
	}
	bd.prefetchedBlocks.Add(block)
}
