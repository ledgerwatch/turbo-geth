package adapter

import (
	"github.com/ledgerwatch/erigon/common"
	"github.com/ledgerwatch/erigon/core/rawdb"
	"github.com/ledgerwatch/erigon/core/types"
	"github.com/ledgerwatch/erigon/ethdb"
)

func NewBlockGetter(tx ethdb.Tx) *blockGetter {
	return &blockGetter{tx}
}

type blockGetter struct {
	tx ethdb.Tx
}

func (g *blockGetter) GetBlockByHash(hash common.Hash) (*types.Block, error) {
	return rawdb.ReadBlockByHash(g.tx, hash)
}

func (g *blockGetter) GetBlock(hash common.Hash, number uint64) *types.Block {
	return rawdb.ReadBlock(g.tx, hash, number)
}
