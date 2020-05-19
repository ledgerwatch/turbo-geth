package ethapi

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"sort"

	"github.com/ledgerwatch/turbo-geth/common"
	"github.com/ledgerwatch/turbo-geth/common/changeset"
	"github.com/ledgerwatch/turbo-geth/common/dbutils"
	"github.com/ledgerwatch/turbo-geth/common/hexutil"
	"github.com/ledgerwatch/turbo-geth/core/rawdb"
	"github.com/ledgerwatch/turbo-geth/core/types/accounts"
	"github.com/ledgerwatch/turbo-geth/rpc"
	"github.com/ledgerwatch/turbo-geth/trie"
)

func (s *PublicBlockChainAPI) GetProof(ctx context.Context, address common.Address, storageKeys []string, blockNr rpc.BlockNumber) (*AccountResult, error) {
	block := uint64(blockNr.Int64())
	db := s.b.ChainDb()
	ts := dbutils.EncodeTimestamp(block)
	accountCs := 0
	accountMap := make(map[string]*accounts.Account)
	if err := db.Walk(dbutils.AccountChangeSetBucket, ts, 0, func(k, v []byte) (bool, error) {
		if changeset.Len(v) > 0 {
			walker := func(kk, vv []byte) error {
				if _, ok := accountMap[string(kk)]; !ok {
					if len(vv) > 0 {
						var a accounts.Account
						if innerErr := a.DecodeForStorage(vv); innerErr != nil {
							return innerErr
						}
						accountMap[string(kk)] = &a
					} else {
						accountMap[string(kk)] = nil
					}
				}
				return nil
			}
			v = common.CopyBytes(v) // Making copy because otherwise it will be invalid after the transaction
			if innerErr := changeset.AccountChangeSetBytes(v).Walk(walker); innerErr != nil {
				return false, innerErr
			}
		}
		accountCs++
		return true, nil
	}); err != nil {
		return nil, err
	}
	storageCs := 0
	storageMap := make(map[string][]byte)
	if err := db.Walk(dbutils.StorageChangeSetBucket, ts, 0, func(k, v []byte) (bool, error) {
		if changeset.Len(v) > 0 {
			walker := func(kk, vv []byte) error {
				if _, ok := storageMap[string(kk)]; !ok {
					storageMap[string(kk)] = vv
				}
				return nil
			}
			v = common.CopyBytes(v) // Making copy because otherwise it will be invalid after the transaction
			if innerErr := changeset.StorageChangeSetBytes(v).Walk(walker); innerErr != nil {
				return false, innerErr
			}
		}
		storageCs++
		return true, nil
	}); err != nil {
		return nil, err
	}
	var unfurlList = make([]string, len(accountMap)+len(storageMap))
	unfurl := trie.NewRetainList(0)
	i := 0
	for ks, acc := range accountMap {
		unfurlList[i] = ks
		i++
		unfurl.AddKey([]byte(ks))
		if acc != nil {
			// Fill the code hashes
			if acc.Incarnation > 0 && acc.IsEmptyCodeHash() {
				if codeHash, err1 := db.Get(dbutils.ContractCodeBucket, dbutils.GenerateStoragePrefix([]byte(ks), acc.Incarnation)); err1 == nil {
					copy(acc.CodeHash[:], codeHash)
				} else {
					return nil, err1
				}
			}
		}
	}
	for ks := range storageMap {
		unfurlList[i] = ks
		i++
		var sk [64]byte
		copy(sk[:], []byte(ks)[:common.HashLength])
		copy(sk[common.HashLength:], []byte(ks)[common.HashLength+common.IncarnationLength:])
		unfurl.AddKey(sk[:])
	}
	rl := trie.NewRetainList(0)
	addrHash, err := common.HashData(address[:])
	if err != nil {
		return nil, err
	}
	rl.AddKey(addrHash[:])
	unfurl.AddKey(addrHash[:])
	for _, key := range storageKeys {
		keyAsHash := common.HexToHash(key)
		if keyHash, err1 := common.HashData(keyAsHash[:]); err1 == nil {
			trieKey := append(addrHash[:], keyHash[:]...)
			rl.AddKey(trieKey)
			unfurl.AddKey(trieKey)
		} else {
			return nil, err1
		}
	}
	sort.Strings(unfurlList)
	fmt.Printf("Account changesets: %d, storage changesets: %d, unfurlList: %d\n", accountCs, storageCs, len(unfurlList))
	loader := trie.NewFlatDbSubTrieLoader()
	if err = loader.Reset(db, unfurl, [][]byte{nil}, []int{0}, false); err != nil {
		return nil, err
	}
	r := &Receiver{defaultReceiver: trie.NewDefaultReceiver(), unfurlList: unfurlList, accountMap: accountMap, storageMap: storageMap}
	r.defaultReceiver.Reset(rl, false)
	loader.SetStreamReceiver(r)
	subTries, err1 := loader.LoadSubTries()
	if err1 != nil {
		return nil, err1
	}
	headHash := rawdb.ReadHeadBlockHash(db)
	headNumber := rawdb.ReadHeaderNumber(db, headHash)
	headHeader := rawdb.ReadHeader(db, headHash, *headNumber)
	fmt.Printf("Head block number: %d, root hash: %x\n", *headNumber, headHeader.Root)
	fmt.Printf("Root hash: %x\n", subTries.Hashes[0])
	hash := rawdb.ReadCanonicalHash(db, block-1)
	header := rawdb.ReadHeader(db, hash, block-1)
	fmt.Printf("Block state root: %x\n", header.Root)
	tr := trie.New(header.Root)
	if err = tr.HookSubTries(subTries, [][]byte{nil}); err != nil {
		return nil, err
	}
	accountProof, err2 := tr.Prove(addrHash[:], 0, false /* storage */)
	if err2 != nil {
		return nil, err2
	}
	storageProof := make([]StorageResult, len(storageKeys))
	for i, key := range storageKeys {
		keyAsHash := common.HexToHash(key)
		if keyHash, err1 := common.HashData(keyAsHash[:]); err1 == nil {
			trieKey := append(addrHash[:], keyHash[:]...)
			if proof, err3 := tr.Prove(trieKey, 64 /* nibbles to get to the storage sub-trie */, true /* storage */); err3 == nil {
				v, _ := tr.Get(trieKey)
				bv := new(big.Int)
				bv.SetBytes(v)
				storageProof[i] = StorageResult{key, (*hexutil.Big)(bv), common.ToHexArray(proof)}
			} else {
				return nil, err3
			}
		} else {
			return nil, err1
		}
	}
	acc, found := tr.GetAccount(addrHash[:])
	if !found {
		return nil, nil
	}
	return &AccountResult{
		Address:      address,
		AccountProof: common.ToHexArray(accountProof),
		Balance:      (*hexutil.Big)(&acc.Balance),
		CodeHash:     acc.CodeHash,
		Nonce:        hexutil.Uint64(acc.Nonce),
		StorageHash:  acc.Root,
		StorageProof: storageProof,
	}, nil
}

type Receiver struct {
	defaultReceiver *trie.DefaultReceiver
	accountMap      map[string]*accounts.Account
	storageMap      map[string][]byte
	unfurlList      []string
	currentIdx      int
}

func (r *Receiver) Receive(
	itemType trie.StreamItem,
	accountKey []byte,
	storageKeyPart1 []byte,
	storageKeyPart2 []byte,
	accountValue *accounts.Account,
	storageValue []byte,
	hash []byte,
	cutoff int,
	witnessLen uint64,
) error {
	for r.currentIdx < len(r.unfurlList) {
		ks := r.unfurlList[r.currentIdx]
		k := []byte(ks)
		var c int
		switch itemType {
		case trie.StorageStreamItem, trie.SHashStreamItem:
			if len(k) > common.HashLength {
				c = bytes.Compare(k[:common.HashLength], storageKeyPart1)
				if c == 0 {
					c = bytes.Compare(k[common.HashLength+common.IncarnationLength:], storageKeyPart2)
				}
			} else {
				c = bytes.Compare(k, storageKeyPart1)
			}
		case trie.AccountStreamItem, trie.AHashStreamItem:
			c = bytes.Compare(k, accountKey)
		case trie.CutoffStreamItem:
			c = -1
		}
		if c > 0 {
			return r.defaultReceiver.Receive(itemType, accountKey, storageKeyPart1, storageKeyPart2, accountValue, storageValue, hash, cutoff, witnessLen)
		}
		if len(k) > common.HashLength {
			v := r.storageMap[ks]
			if c <= 0 && len(v) > 0 {
				if err := r.defaultReceiver.Receive(trie.StorageStreamItem, nil, k[:32], k[40:], nil, v, nil, 0, 0); err != nil {
					return err
				}
			}
		} else {
			v := r.accountMap[ks]
			if c <= 0 && v != nil {
				if err := r.defaultReceiver.Receive(trie.AccountStreamItem, k, nil, nil, v, nil, nil, 0, 0); err != nil {
					return err
				}
			}
		}
		r.currentIdx++
		if c == 0 {
			return nil
		}
	}
	// We ran out of modifications, simply pass through
	return r.defaultReceiver.Receive(itemType, accountKey, storageKeyPart1, storageKeyPart2, accountValue, storageValue, hash, cutoff, witnessLen)
}

func (r *Receiver) Result() trie.SubTries {
	return r.defaultReceiver.Result()
}
