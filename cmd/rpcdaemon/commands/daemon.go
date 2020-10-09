package commands

import (
	"context"
	"fmt"

	"github.com/ledgerwatch/turbo-geth/cmd/rpcdaemon/cli"
	"github.com/ledgerwatch/turbo-geth/common"
	"github.com/ledgerwatch/turbo-geth/common/hexutil"
	"github.com/ledgerwatch/turbo-geth/core/rawdb"
	"github.com/ledgerwatch/turbo-geth/core/types"
	"github.com/ledgerwatch/turbo-geth/ethdb"
	"github.com/ledgerwatch/turbo-geth/rpc"
	"github.com/ledgerwatch/turbo-geth/turbo/adapter/ethapi"
)

// GetBlockByNumber see https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getblockbynumber
// see internal/ethapi.PublicBlockChainAPI.GetBlockByNumber
func (api *APIImpl) GetBlockByNumber(ctx context.Context, number rpc.BlockNumber, fullTx bool) (map[string]interface{}, error) {
	blockNum, err := getBlockNumber(number, api.dbReader)
	if err != nil {
		return nil, err
	}
	additionalFields := make(map[string]interface{})

	block := rawdb.ReadBlockByNumber(api.dbReader, blockNum)
	if block == nil {
		return nil, fmt.Errorf("block not found: %d", blockNum)
	}

	additionalFields["totalDifficulty"] = (*hexutil.Big)(rawdb.ReadTd(api.dbReader, block.Hash(), blockNum))
	response, err := ethapi.RPCMarshalBlock(block, true, fullTx, additionalFields)

	if err == nil && number == rpc.PendingBlockNumber {
		// Pending blocks need to nil out a few fields
		for _, field := range []string{"hash", "nonce", "miner"} {
			response[field] = nil
		}
	}
	return response, err
}

// GetBlockByHash see https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getblockbyhash
// see internal/ethapi.PublicBlockChainAPI.GetBlockByHash
func (api *APIImpl) GetBlockByHash(ctx context.Context, hash common.Hash, fullTx bool) (map[string]interface{}, error) {
	additionalFields := make(map[string]interface{})

	block := rawdb.ReadBlockByHash(api.dbReader, hash)
	if block == nil {
		return nil, fmt.Errorf("block not found: %x", hash)
	}
	number := block.NumberU64()

	additionalFields["totalDifficulty"] = (*hexutil.Big)(rawdb.ReadTd(api.dbReader, hash, number))
	response, err := ethapi.RPCMarshalBlock(block, true, fullTx, additionalFields)

	if err == nil && int64(number) == rpc.PendingBlockNumber.Int64() {
		// Pending blocks need to nil out a few fields
		for _, field := range []string{"hash", "nonce", "miner"} {
			response[field] = nil
		}
	}
	return response, err
}

// GetHeaderByNumber returns a block's header by number
func (api *APIImpl) GetHeaderByNumber(_ context.Context, number rpc.BlockNumber) (*types.Header, error) {
	header := rawdb.ReadHeaderByNumber(api.dbReader, uint64(number.Int64()))
	if header == nil {
		return nil, fmt.Errorf("block header not found: %d", number.Int64())
	}

	return header, nil
}

// GetHeaderByHash returns a block's header by hash
func (api *APIImpl) GetHeaderByHash(_ context.Context, hash common.Hash) (*types.Header, error) {
	header := rawdb.ReadHeaderByHash(api.dbReader, hash)
	if header == nil {
		return nil, fmt.Errorf("block header not found: %s", hash.String())
	}

	return header, nil
}

// APIList describes the list of available RPC apis
func APIList(db ethdb.KV, eth ethdb.Backend, cfg cli.Flags, customAPIList []rpc.API) []rpc.API {
	var defaultAPIList []rpc.API

	dbReader := ethdb.NewObjectDatabase(db)

	ethImpl := NewEthAPI(db, dbReader, eth, cfg.Gascap)
	netImpl := NewNetAPIImpl(eth)
	debugImpl := NewPrivateDebugAPI(db, dbReader)
	traceImpl := NewTraceAPI(db, dbReader, &cfg)
	web3Impl := NewWeb3APIImpl()
	shhImpl := NewSHHAPIImpl()

	for _, enabledAPI := range cfg.API {
		switch enabledAPI {
		case "eth":
			defaultAPIList = append(defaultAPIList, rpc.API{
				Namespace: "eth",
				Public:    true,
				Service:   EthAPI(ethImpl),
				Version:   "1.0",
			})
		case "debug":
			defaultAPIList = append(defaultAPIList, rpc.API{
				Namespace: "debug",
				Public:    true,
				Service:   PrivateDebugAPI(debugImpl),
				Version:   "1.0",
			})
		case "net":
			defaultAPIList = append(defaultAPIList, rpc.API{
				Namespace: "net",
				Public:    true,
				Service:   NetAPI(netImpl),
				Version:   "1.0",
			})
		case "web3":
			defaultAPIList = append(defaultAPIList, rpc.API{
				Namespace: "web3",
				Public:    true,
				Service:   Web3API(web3Impl),
				Version:   "1.0",
			})
		case "trace":
			defaultAPIList = append(defaultAPIList, rpc.API{
				Namespace: "trace",
				Public:    true,
				Service:   TraceAPI(traceImpl),
				Version:   "1.0",
			})
		case "shh":
			defaultAPIList = append(defaultAPIList, rpc.API{
				Namespace: "shh",
				Public:    true,
				Service:   SHHAPI(shhImpl),
				Version:   "1.0",
			})
		}
	}

	return append(defaultAPIList, customAPIList...)
}
