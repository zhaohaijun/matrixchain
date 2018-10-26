package storage

import (
	"github.com/zhaohaijun/matrixchain/core/states"
	"github.com/zhaohaijun/matrixchain/core/store/common"
)

// StateItem describe smart contract cache item element
type StateItem struct {
	Prefix common.DataEntryPrefix
	Key    string
	Value  states.StateValue
	State  common.ItemState
}

type Memory map[string]*StateItem

// CacheDB is smart contract execute cache, it contain transaction cache and block cache
// When smart contract execute finish, need to commit transaction cache to block cache
type CloneCache struct {
	Memory Memory
	Store  common.StateStore
}

// NewCloneCache return a new contract cache
func NewCloneCache(store common.StateStore) *CloneCache {
	return &CloneCache{
		Memory: make(Memory),
		Store:  store,
	}
}

// Commit current transaction cache to block cache
func (this *CloneCache) Commit() {
	for _, v := range this.Memory {
		vk := []byte(v.Key)
		if v.State == common.Deleted {
			this.Store.TryDelete(v.Prefix, vk)
		} else if v.State == common.Changed {
			this.Store.TryAdd(v.Prefix, vk, v.Value)
		}
	}
}

// Add item to cache
func (this *CloneCache) Add(prefix common.DataEntryPrefix, key []byte, value states.StateValue) {
	pk := string(append([]byte{byte(prefix)}, key...))
	this.Memory[pk] = &StateItem{
		Prefix: prefix,
		Key:    string(key),
		Value:  value,
		State:  common.Changed,
	}
}

// GetOrAdd item
// If item has existed, return it
// Else add it to cache
func (this *CloneCache) GetOrAdd(prefix common.DataEntryPrefix, key []byte, value states.StateValue) (states.StateValue, error) {
	pk := string(append([]byte{byte(prefix)}, key...))
	if v, ok := this.Memory[pk]; ok {
		if v.State == common.Deleted {
			this.Memory[pk] = &StateItem{Prefix: prefix, Key: string(key), Value: value, State: common.Changed}
			return value, nil
		}
		return v.Value, nil
	}
	item, err := this.Store.TryGet(prefix, key)
	if err != nil {
		return nil, err
	}
	if item != nil && item.State != common.Deleted {
		return item.Value, nil
	}
	this.Memory[pk] = &StateItem{Prefix: prefix, Key: string(key), Value: value, State: common.Changed}
	return value, nil
}

// Get item by key
func (this *CloneCache) Get(prefix common.DataEntryPrefix, key []byte) (states.StateValue, error) {
	pk := string(append([]byte{byte(prefix)}, key...))
	if v, ok := this.Memory[pk]; ok {
		if v.State == common.Deleted {
			return nil, nil
		}
		return v.Value, nil
	}
	item, err := this.Store.TryGet(prefix, key)
	if err != nil {
		return nil, err
	}
	if item == nil || item.State == common.Deleted {
		return nil, nil
	}
	return item.Value, nil
}

// Delete item from cache
func (this *CloneCache) Delete(prefix common.DataEntryPrefix, key []byte) {
	pk := string(append([]byte{byte(prefix)}, key...))
	if v, ok := this.Memory[pk]; ok {
		v.State = common.Deleted
	} else {
		this.Memory[pk] = &StateItem{
			Prefix: prefix,
			Key:    string(key),
			State:  common.Deleted,
		}
	}
}
