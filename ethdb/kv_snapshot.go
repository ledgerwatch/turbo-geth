package ethdb

import (
	"context"
	"errors"
	"fmt"
	"github.com/ledgerwatch/turbo-geth/common"
	"github.com/ledgerwatch/turbo-geth/common/dbutils"
	"github.com/ledgerwatch/turbo-geth/log"
	"sync"
)

var (
	_ KV             = &SnapshotKV{}
	_ BucketMigrator = &snapshotTX{}
	_ Tx             = &snapshotTX{}
	_ Tx             = &lazyTx{}
	_ Cursor         = &snapshotCursor{}
)

func (s *snapshotTX) CursorDupSort(bucket string) CursorDupSort {
	panic("implement me")
}

func (s *snapshotTX) CursorDupFixed(bucket string) CursorDupFixed {
	panic("implement me")
}

func (s *snapshotTX) Comparator(bucket string) dbutils.CmpFunc {
	panic("implement me")
}

func (s *snapshotTX) Cmp(bucket string, a, b []byte) int {
	panic("implement me")
}

func (s *snapshotTX) DCmp(bucket string, a, b []byte) int {
	panic("implement me")
}

func (v *lazyTx) CursorDupSort(bucket string) CursorDupSort {
	panic("implement me")
}

func (v *lazyTx) CursorDupFixed(bucket string) CursorDupFixed {
	panic("implement me")
}

func (v *lazyTx) Comparator(bucket string) dbutils.CmpFunc {
	panic("implement me")
}

func (v *lazyTx) Cmp(bucket string, a, b []byte) int {
	panic("implement me")
}

func (v *lazyTx) DCmp(bucket string, a, b []byte) int {
	panic("implement me")
}

func (s *snapshotCursor) Prev() ([]byte, []byte, error) {
	panic("implement me")
}

func (s *snapshotCursor) Current() ([]byte, []byte, error) {
	panic("implement me")
}

func (s *snapshotCursor) DeleteCurrent() error {
	panic("implement me")
}

func (s *snapshotCursor) Count() (uint64, error) {
	panic("implement me")
}

func (s *SnapshotKV) AllBuckets() dbutils.BucketsCfg {
	return s.db.AllBuckets()
}

func (s *snapshotTX) DropBucket(bucket string) error {
	db, ok := s.dbTX.(BucketMigrator)
	if !ok {
		return fmt.Errorf("%T doesn't implement ethdb.TxMigrator interface", s.dbTX)
	}
	return db.DropBucket(bucket)
}

func (s *snapshotTX) CreateBucket(bucket string) error {
	db, ok := s.dbTX.(BucketMigrator)
	if !ok {
		return fmt.Errorf("%T doesn't implement ethdb.TxMigrator interface", s.dbTX)
	}
	return db.CreateBucket(bucket)
}

func (s *snapshotTX) ExistsBucket(bucket string) bool {
	db, ok := s.dbTX.(BucketMigrator)
	if !ok {
		return false
	}
	return db.ExistsBucket(bucket)
}

func (s *snapshotTX) ClearBucket(bucket string) error {
	db, ok := s.dbTX.(BucketMigrator)
	if !ok {
		return fmt.Errorf("%T doesn't implement ethdb.TxMigrator interface", s.dbTX)
	}
	return db.ClearBucket(bucket)
}

func (s *snapshotTX) ExistingBuckets() ([]string, error) {
	db, ok := s.dbTX.(BucketMigrator)
	if !ok {
		return []string{}, fmt.Errorf("%T doesn't implement ethdb.TxMigrator interface", s.dbTX)
	}
	return db.ExistingBuckets()
}

type SnapshotUsageOpt struct {
	Path       string
	ForBuckets map[string]struct{}
}

func NewSnapshotKV() snapshotOpts {
	return snapshotOpts{
		forBuckets: make(map[string]dbutils.BucketConfigItem),
	}
}

type SnapshotKV struct {
	db           KV
	snapshotDB   KV
	snapshotPath string
	forBuckets   map[string]dbutils.BucketConfigItem
}

type snapshotOpts struct {
	db         KV
	snapshot   KV
	forBuckets map[string]dbutils.BucketConfigItem
}

func (opts snapshotOpts) SnapshotDB(db KV) snapshotOpts {
	opts.snapshot = db
	return opts
}

func (opts snapshotOpts) DB(db KV) snapshotOpts {
	opts.db = db
	return opts
}

func (opts snapshotOpts) For(bucket string, config dbutils.BucketConfigItem) snapshotOpts {
	opts.forBuckets[bucket] = config
	return opts
}

func (opts snapshotOpts) MustOpen() KV {
	return &SnapshotKV{
		snapshotDB: opts.snapshot,
		db:         opts.db,
		forBuckets: opts.forBuckets,
	}
}

func (s *SnapshotKV) View(ctx context.Context, f func(tx Tx) error) error {
	dbTx, err := s.db.Begin(ctx, nil, false)
	if err != nil {
		return err
	}

	t := &snapshotTX{
		dbTX: dbTx,
		snTX: newVirtualTx(func() (Tx, error) {
			return s.snapshotDB.Begin(ctx, nil, false)
		}, s.forBuckets),
		forBuckets: s.forBuckets,
	}
	defer t.Rollback()
	return f(t)
}

func (s *SnapshotKV) Update(ctx context.Context, f func(tx Tx) error) error {
	return s.db.Update(ctx, f)
}

func (s *SnapshotKV) Close() {
	defer s.db.Close()
	if s.snapshotDB != nil {
		defer s.snapshotDB.Close()
	}
}

func (s *SnapshotKV) Begin(ctx context.Context, parentTx Tx, writable bool) (Tx, error) {
	dbTx, err := s.db.Begin(ctx, parentTx, writable)
	if err != nil {
		return nil, err
	}

	t := &snapshotTX{
		dbTX: dbTx,
		snTX: newVirtualTx(func() (Tx, error) {
			return s.snapshotDB.Begin(ctx, parentTx, false)
		}, s.forBuckets),
		forBuckets: s.forBuckets,
	}
	return t, nil
}

func newVirtualTx(construct func() (Tx, error), forBucket map[string]dbutils.BucketConfigItem) *lazyTx {
	return &lazyTx{
		construct:  construct,
		forBuckets: forBucket,
	}
}

var ErrNotSnapshotBucket = errors.New("not snapshot bucket")

//lazyTx is used for lazy transaction creation.
type lazyTx struct {
	construct  func() (Tx, error)
	forBuckets map[string]dbutils.BucketConfigItem
	tx         Tx
	sync.Mutex
}

func (v *lazyTx) getTx() (Tx, error) {
	var err error
	v.Lock()
	defer v.Unlock()
	if v.tx == nil {
		v.tx, err = v.construct()
	}
	return v.tx, err
}

func (v *lazyTx) Cursor(bucket string) Cursor {
	if _, ok := v.forBuckets[bucket]; !ok {
		return nil
	}

	tx, err := v.getTx()
	if err != nil {
		log.Error("Fail to create tx", "err", err)
	}
	if tx != nil {
		return tx.Cursor(bucket)
	}
	return nil
}

func (v *lazyTx) Get(bucket string, key []byte) (val []byte, err error) {
	if _, ok := v.forBuckets[bucket]; !ok {
		return nil, ErrNotSnapshotBucket
	}
	tx, err := v.getTx()
	if err != nil {
		return nil, err
	}
	return tx.Get(bucket, key)
}

func (v *lazyTx) Commit(ctx context.Context) error {
	return nil
}

func (v *lazyTx) Rollback() {
	v.Lock()
	defer v.Unlock()
	if v.tx != nil {
		v.tx.Rollback()
	}
}

func (v *lazyTx) BucketSize(bucket string) (uint64, error) {
	if _, ok := v.forBuckets[bucket]; !ok {
		return 0, nil
	}
	tx, err := v.getTx()
	if err != nil {
		return 0, err
	}
	return tx.BucketSize(bucket)
}

type snapshotTX struct {
	dbTX       Tx
	snTX       Tx
	forBuckets map[string]dbutils.BucketConfigItem
	writable bool
}

func (s *snapshotTX) Commit(ctx context.Context) error {
	defer s.snTX.Rollback()
	return s.dbTX.Commit(ctx)
}

func (s *snapshotTX) Rollback() {
	defer s.snTX.Rollback()
	defer s.dbTX.Rollback()
}

func (s *snapshotTX) Cursor(bucket string) Cursor {
	if _,ok:=s.forBuckets[bucket];!ok {
		return s.dbTX.Cursor(bucket)
	}
	snCursor := s.snTX.Cursor(bucket)
	//check snapshot bucket
	if snCursor == nil {
		return s.dbTX.Cursor(bucket)
	}
	return &snapshotCursor{
		snCursor: snCursor,
		dbCursor: s.dbTX.Cursor(bucket),
	}
}

func (s *snapshotTX) Get(bucket string, key []byte) (val []byte, err error) {
	_, ok := s.forBuckets[bucket]
	if !ok {
		return s.dbTX.Get(bucket, key)
	}
	v, err := s.dbTX.Get(bucket, key)
	switch {
	case err == nil && v != nil:
		return v, nil
	case err != nil && err != ErrKeyNotFound:
		return nil, err
	}
	return s.snTX.Get(bucket, key)
}

func (s *snapshotTX) BucketSize(name string) (uint64, error) {
	dbSize, err := s.snTX.BucketSize(name)
	if err != nil {
		return 0, fmt.Errorf("db err %w", err)
	}
	snSize, err := s.dbTX.BucketSize(name)
	if err != nil {
		return 0, fmt.Errorf("Snapshot db err %w", err)
	}
	return dbSize + snSize, nil
}

type snapshotCursor struct {
	snCursor Cursor
	dbCursor Cursor

	lastDBKey   []byte
	lastSNDBKey []byte
	lastDBVal   []byte
	lastSNDBVal []byte
	keyCmp      int
}

func (s *snapshotCursor) Prefix(v []byte) Cursor {
	panic("implement me")
}

func (s *snapshotCursor) MatchBits(u uint) Cursor {
	panic("implement me")
}

func (s *snapshotCursor) Prefetch(v uint) Cursor {
	panic("implement me")
}

func (s *snapshotCursor) First() ([]byte, []byte, error) {
	var err error
	s.lastSNDBKey, s.lastSNDBVal, err = s.snCursor.First()
	if err != nil {
		return nil, nil, err
	}
	s.lastDBKey, s.lastDBVal, err = s.dbCursor.First()
	if err != nil {
		return nil, nil, err
	}
	cmp, br := common.KeyCmp(s.lastDBKey, s.lastSNDBKey)
	if br {
		return nil, nil, nil
	}

	s.keyCmp = cmp
	if cmp <= 0 {
		return s.lastDBKey, s.lastDBVal, nil
	} else {
		return s.lastSNDBKey, s.lastSNDBVal, nil
	}
}

func (s *snapshotCursor) Seek(seek []byte) ([]byte, []byte, error) {
	var err error
	s.lastSNDBKey, s.lastSNDBVal, err = s.snCursor.Seek(seek)
	if err != nil {
		return nil, nil, err
	}
	s.lastDBKey, s.lastDBVal, err = s.dbCursor.Seek(seek)
	if err != nil {
		return nil, nil, err
	}
	cmp, br := common.KeyCmp(s.lastDBKey, s.lastSNDBKey)
	if br {
		return nil, nil, nil
	}

	s.keyCmp = cmp
	if cmp <= 0 {
		return s.lastDBKey, s.lastDBVal, nil
	} else {
		return s.lastSNDBKey, s.lastSNDBVal, nil
	}
}

func (s *snapshotCursor) SeekTo(seek []byte) ([]byte, []byte, error) {
	panic("implement me")
}

func (s *snapshotCursor) Next() ([]byte, []byte, error) {
	var err error
	if s.keyCmp >= 0 {
		s.lastSNDBKey, s.lastSNDBVal, err = s.snCursor.Next()
	}
	if err != nil {
		return nil, nil, err
	}
	if s.keyCmp <= 0 {
		s.lastDBKey, s.lastDBVal, err = s.dbCursor.Next()
	}
	if err != nil {
		return nil, nil, err
	}

	cmp, br := common.KeyCmp(s.lastDBKey, s.lastSNDBKey)
	if br {
		return nil, nil, nil
	}

	s.keyCmp = cmp
	if cmp <= 0 {
		return s.lastDBKey, s.lastDBVal, nil
	} else {
		return s.lastSNDBKey, s.lastSNDBVal, nil
	}
}

func (s *snapshotCursor) Walk(walker func(k []byte, v []byte) (bool, error)) error {
	panic("implement me")
}

func (s *snapshotCursor) Put(key []byte, value []byte) error {
	return s.dbCursor.Put(key, value)
}

func (s *snapshotCursor) Delete(key []byte) error {
	return s.dbCursor.Delete(key)
}

func (s *snapshotCursor) Append(key []byte, value []byte) error {
	return s.dbCursor.Append(key, value)
}

func (s *snapshotCursor) SeekExact(key []byte) ([]byte, error) {
	v,err := s.dbCursor.SeekExact(key)
	if err!=nil {
		return nil, err
	}
	if v==nil {
		return s.snCursor.SeekExact(key)
	}
	return v, err
}

func (s *snapshotCursor) Last() ([]byte, []byte, error) {
	var err error
	s.lastSNDBKey, s.lastSNDBVal, err = s.snCursor.Last()
	if err != nil {
		return nil, nil, err
	}
	s.lastDBKey, s.lastDBVal, err = s.dbCursor.Last()
	if err != nil {
		return nil, nil, err
	}
	cmp, br := common.KeyCmp(s.lastDBKey, s.lastSNDBKey)
	if br {
		return nil, nil, nil
	}

	s.keyCmp = cmp
	if cmp >= 0 {
		return s.lastDBKey, s.lastDBVal, nil
	} else {
		return s.lastSNDBKey, s.lastSNDBVal, nil
	}
}

/*
number=10487424

WriteHeadHeaderHash - db.Put(dbutils.HeadHeaderKey, dbutils.HeadHeaderKey, hash.Bytes())
DeleteCanonicalHash - db.Delete(dbutils.HeaderPrefix, dbutils.HeaderHashKey(number))
WriteCanonicalHash - db.Put(dbutils.HeaderPrefix, dbutils.HeaderHashKey(number), hash.Bytes())
WriteHeader - db.Put(dbutils.HeaderNumberPrefix, hash[:], encoded),  db.Put(dbutils.HeaderPrefix, dbutils.HeaderKey(number, hash), data)
WriteTd - db.Put(dbutils.HeaderPrefix, dbutils.HeaderTDKey(number, hash), data)

ReadHeader
ReadCanonicalHash
ReadHeaderNumber
ReadHeadHeaderHash
ReadTd
*/
