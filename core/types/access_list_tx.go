// Copyright 2020 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package types

import (
	"io"

	"github.com/holiman/uint256"
	"github.com/ledgerwatch/turbo-geth/common"
	"github.com/ledgerwatch/turbo-geth/rlp"
)

// go:generate gencodec -type AccessTuple -out gen_access_tuple.go

// AccessList is an EIP-2930 access list.
type AccessList []AccessTuple

// AccessTuple is the element type of an access list.
type AccessTuple struct {
	Address     common.Address `json:"address"        gencodec:"required"`
	StorageKeys []common.Hash  `json:"storageKeys"    gencodec:"required"`
}

// StorageKeys returns the total number of storage keys in the access list.
func (al AccessList) StorageKeys() int {
	sum := 0
	for _, tuple := range al {
		sum += len(tuple.StorageKeys)
	}
	return sum
}

// AccessListTx is the data of EIP-2930 access list transactions.
type AccessListTx struct {
	LegacyTx
	ChainID    *uint256.Int
	AccessList AccessList // EIP-2930 access list
}

// copy creates a deep copy of the transaction data and initializes all fields.
func (tx *AccessListTx) copy() *AccessListTx {
	cpy := &AccessListTx{
		LegacyTx: LegacyTx{
			CommonTx: CommonTx{
				Nonce: tx.Nonce,
				To:    tx.To, // TODO: copy pointed-to address
				Data:  common.CopyBytes(tx.Data),
				Gas:   tx.Gas,
				// These are copied below.
				Value: new(uint256.Int),
				V:     new(uint256.Int),
				R:     new(uint256.Int),
				S:     new(uint256.Int),
			},
			GasPrice: new(uint256.Int),
		},
		ChainID:    new(uint256.Int),
		AccessList: make(AccessList, len(tx.AccessList)),
	}
	copy(cpy.AccessList, tx.AccessList)
	if tx.Value != nil {
		cpy.Value.Set(tx.Value)
	}
	if tx.ChainID != nil {
		cpy.ChainID.Set(tx.ChainID)
	}
	if tx.GasPrice != nil {
		cpy.GasPrice.Set(tx.GasPrice)
	}
	if tx.V != nil {
		cpy.V.Set(tx.V)
	}
	if tx.R != nil {
		cpy.R.Set(tx.R)
	}
	if tx.S != nil {
		cpy.S.Set(tx.S)
	}
	return cpy
}

func (tx *AccessListTx) EncodeRLP(w io.Writer) error {
	return nil
}

func (tx *AccessListTx) DecodeRLP(s *rlp.Stream) error {
	return nil
}

// AsMessage returns the transaction as a core.Message.
func (tx AccessListTx) AsMessage(s Signer) (Message, error) {
	msg := Message{
		nonce:      tx.Nonce,
		gasLimit:   tx.Gas,
		gasPrice:   *tx.GasPrice,
		tip:        *tx.GasPrice,
		feeCap:     *tx.GasPrice,
		to:         tx.To,
		amount:     *tx.Value,
		data:       tx.Data,
		accessList: tx.AccessList,
		checkNonce: true,
	}

	var err error
	msg.from, err = Sender(s, &tx)
	return msg, err
}

// Hash computes the hash (but not for signatures!)
func (tx AccessListTx) Hash() common.Hash {
	return rlpHash([]interface{}{
		tx.ChainID,
		tx.Nonce,
		tx.GasPrice,
		tx.Gas,
		tx.To,
		tx.Value,
		tx.Data,
		tx.AccessList,
		tx.V, tx.R, tx.S,
	})
}

// accessors for innerTx.

func (tx *AccessListTx) Type() byte             { return AccessListTxType }
func (tx *AccessListTx) chainID() *uint256.Int  { return tx.ChainID }
func (tx *AccessListTx) accessList() AccessList { return tx.AccessList }
func (tx *AccessListTx) data() []byte           { return tx.Data }
func (tx *AccessListTx) gas() uint64            { return tx.Gas }
func (tx *AccessListTx) gasPrice() *uint256.Int { return tx.GasPrice }
func (tx *AccessListTx) value() *uint256.Int    { return tx.Value }
func (tx *AccessListTx) nonce() uint64          { return tx.Nonce }
func (tx *AccessListTx) to() *common.Address    { return tx.To }

func (tx *AccessListTx) rawSignatureValues() (v, r, s *uint256.Int) {
	return tx.V, tx.R, tx.S
}

func (tx *AccessListTx) setSignatureValues(chainID, v, r, s *uint256.Int) {
	tx.ChainID, tx.V, tx.R, tx.S = chainID, v, r, s
}
