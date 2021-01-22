package bodydownload

import (
	"sync"

	"github.com/RoaringBitmap/roaring/roaring64"
	"github.com/ledgerwatch/turbo-geth/common"
	"github.com/ledgerwatch/turbo-geth/core/types"
)

// DoubleHash is type to be used for the mapping between TxHash and UncleHash to the block header
type DoubleHash [2 * common.HashLength]byte

const MaxBodiesInRequest = 128

// BodyDownload represents the state of body downloading process
type BodyDownload struct {
	lock             sync.RWMutex
	delivered        *roaring64.Bitmap
	deliveries       []*types.Block
	requestedMap     map[DoubleHash]uint64
	maxProgress      uint64
	requestedLow     uint64 // Lower bound of block number for outstanding requests
	requestHigh      uint64
	outstandingLimit uint64 // Limit of number of outstanding blocks for body requests
	blockChannel     chan *types.Block
}

type RequestQueueItem struct {
	lowestBlockNum uint64
	requested      *roaring64.Bitmap
	waitUntil      uint64
}

type RequestQueue []RequestQueueItem

func (rq RequestQueue) Len() int {
	return len(rq)
}

func (rq RequestQueue) Less(i, j int) bool {
	return rq[i].lowestBlockNum < rq[j].lowestBlockNum
}

func (rq RequestQueue) Swap(i, j int) {
	rq[i], rq[j] = rq[j], rq[i]
}

func (rq *RequestQueue) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*rq = append(*rq, x.(RequestQueueItem))
}

func (rq *RequestQueue) Pop() interface{} {
	old := *rq
	n := len(old)
	x := old[n-1]
	*rq = old[0 : n-1]
	return x
}

// BodyRequest is a sketch of the request for block bodies, meaning that access to the database is required to convert it to the actual BlockBodies request (look up hashes of canonical blocks)
type BodyRequest struct {
	BlockNums []uint64
	Hashes    []common.Hash
	requested *roaring64.Bitmap
}

// NewBodyDownload create a new body download state object
func NewBodyDownload(outstandingLimit int) *BodyDownload {
	bd := &BodyDownload{
		requestedMap:     make(map[DoubleHash]uint64),
		outstandingLimit: uint64(outstandingLimit),
		delivered:        roaring64.New(),
		deliveries:       make([]*types.Block, outstandingLimit+MaxBodiesInRequest),
	}
	return bd
}
