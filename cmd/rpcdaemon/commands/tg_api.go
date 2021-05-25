package commands

import (
	"context"

	"github.com/ledgerwatch/erigon/common"
	"github.com/ledgerwatch/erigon/core/types"
	"github.com/ledgerwatch/erigon/ethdb"
	"github.com/ledgerwatch/erigon/rpc"
)

// ErigonAPI Erigon specific routines
type ErigonAPI interface {
	// System related (see ./tg_system.go)
	Forks(ctx context.Context) (Forks, error)

	// Blocks related (see ./tg_blocks.go)
	GetHeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Header, error)
	GetHeaderByHash(_ context.Context, hash common.Hash) (*types.Header, error)

	// Receipt related (see ./tg_receipts.go)
	GetLogsByHash(ctx context.Context, hash common.Hash) ([][]*types.Log, error)
	//GetLogsByNumber(ctx context.Context, number rpc.BlockNumber) ([][]*types.Log, error)

	// Issuance / reward related (see ./erigon_issuance.go)
	// BlockReward(ctx context.Context, blockNr rpc.BlockNumber) (Issuance, error)
	// UncleReward(ctx context.Context, blockNr rpc.BlockNumber) (Issuance, error)
	Issuance(ctx context.Context, blockNr rpc.BlockNumber) (Issuance, error)
}

// ErigonImpl is implementation of the ErigonAPI interface
type ErigonImpl struct {
	*BaseAPI
	db ethdb.RoKV
}

// NewErigonAPI returns ErigonImpl instance
func NewErigonAPI(base *BaseAPI, db ethdb.RoKV) *ErigonImpl {
	return &ErigonImpl{
		BaseAPI: base,
		db:      db,
	}
}
