package storage

import (
	"github.com/syndtr/goleveldb/leveldb/util"
	comm "github.com/zhaohaijun/matrixchain/common"
	"github.com/zhaohaijun/matrixchain/core/payload"
	"github.com/zhaohaijun/matrixchain/core/store/common"
	"github.com/zhaohaijun/matrixchain/core/store/overlaydb"
)

// CacheDB is smart contract execute cache, it contain transaction cache and block cache
// When smart contract execute finish, need to commit transaction cache to block cache
type CacheDB struct {
	memdb      *overlaydb.MemDB
	backend    *overlaydb.OverlayDB
	keyScratch []byte
}

// NewCacheDB return a new contract cache
func NewCacheDB(store *overlaydb.OverlayDB) *CacheDB {
	return &CacheDB{
		backend: store,
		memdb:   overlaydb.NewMemDB(0),
	}
}

func ensureBuffer(b []byte, n int) []byte {
	if cap(b) < n {
		return make([]byte, n)
	}
	return b[:n]
}

func makePrefixedKey(dst []byte, prefix byte, key []byte) []byte {
	dst = ensureBuffer(dst, len(key)+1)
	dst[0] = prefix
	copy(dst[1:], key)
	return dst
}

// Commit current transaction cache to block cache
func (self *CacheDB) Commit() {
	self.memdb.ForEach(func(key, val []byte) {
		if len(val) == 0 {
			self.backend.Delete(key)
		} else {
			self.backend.Put(key, val)
		}
	})
}

func (self *CacheDB) Put(key []byte, value []byte) {
	self.put(common.ST_STORAGE, key, value)
}

func (self *CacheDB) put(prefix common.DataEntryPrefix, key []byte, value []byte) {
	self.keyScratch = makePrefixedKey(self.keyScratch, byte(prefix), key)
	self.memdb.Put(self.keyScratch, value)
}

func (self *CacheDB) GetContract(addr comm.Address) (*payload.DeployCode, error) {
	value, err := self.get(common.ST_CONTRACT, addr[:])
	if err != nil {
		return nil, err
	}

	if len(value) == 0 {
		return nil, nil
	}

	contract := new(payload.DeployCode)
	if err := contract.Deserialization(comm.NewZeroCopySource(value)); err != nil {
		return nil, err
	}
	return contract, nil
}

func (self *CacheDB) PutContract(contract *payload.DeployCode) error {
	address := contract.Address()

	sink := comm.NewZeroCopySink(nil)
	err := contract.Serialization(sink)
	if err != nil {
		return err
	}

	value := sink.Bytes()
	self.put(common.ST_CONTRACT, address[:], value)
	return nil
}

func (self *CacheDB) DeleteContract(address comm.Address) {
	self.delete(common.ST_CONTRACT, address[:])
}

func (self *CacheDB) Get(key []byte) ([]byte, error) {
	return self.get(common.ST_STORAGE, key)
}

func (self *CacheDB) get(prefix common.DataEntryPrefix, key []byte) ([]byte, error) {
	self.keyScratch = makePrefixedKey(self.keyScratch, byte(prefix), key)
	value, unknown := self.memdb.Get(self.keyScratch)
	if unknown {
		v, err := self.backend.Get(self.keyScratch)
		if err != nil {
			return nil, err
		}
		value = v
	}

	return value, nil
}

func (self *CacheDB) Delete(key []byte) {
	self.delete(common.ST_STORAGE, key)
}

// Delete item from cache
func (self *CacheDB) delete(prefix common.DataEntryPrefix, key []byte) {
	self.keyScratch = makePrefixedKey(self.keyScratch, byte(prefix), key)
	self.memdb.Delete(self.keyScratch)
}

func (self *CacheDB) NewIterator(key []byte) common.StoreIterator {
	pkey := make([]byte, 1+len(key))
	pkey[0] = byte(common.ST_STORAGE)
	copy(pkey[1:], key)
	prefixRange := util.BytesPrefix(pkey)
	backIter := self.backend.NewIterator(pkey)
	memIter := self.memdb.NewIterator(prefixRange)

	return overlaydb.NewJoinIter(memIter, backIter)
}

type Iter struct {
	*overlaydb.JoinIter
}

func (self *Iter) Key() []byte {
	key := self.JoinIter.Key()
	if len(key) != 0 {
		key = key[1:] // remove the first prefix
	}
	return key
}
