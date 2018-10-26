package utils

import (
	"bytes"

	"github.com/zhaohaijun/matrixchain/common/serialization"
	cstates "github.com/zhaohaijun/matrixchain/core/states"
	"github.com/zhaohaijun/matrixchain/errors"
	"github.com/zhaohaijun/matrixchain/smartcontract/service/native"
)

func GetStorageItem(native *native.NativeService, key []byte) (*cstates.StorageItem, error) {
	store, err := native.CacheDB.Get(key)
	if err != nil {
		return nil, errors.NewDetailErr(err, errors.ErrNoCode, "[GetStorageItem] storage error!")
	}
	if store == nil {
		return nil, nil
	}
	item := new(cstates.StorageItem)
	err = item.Deserialize(bytes.NewBuffer(store))
	if err != nil {
		return nil, errors.NewDetailErr(err, errors.ErrNoCode, "[GetStorageItem] instance doesn't StorageItem!")
	}
	return item, nil
}

func GetStorageUInt64(native *native.NativeService, key []byte) (uint64, error) {
	item, err := GetStorageItem(native, key)
	if err != nil {
		return 0, err
	}
	if item == nil {
		return 0, nil
	}
	v, err := serialization.ReadUint64(bytes.NewBuffer(item.Value))
	if err != nil {
		return 0, err
	}
	return v, nil
}

func GetStorageUInt32(native *native.NativeService, key []byte) (uint32, error) {
	item, err := GetStorageItem(native, key)
	if err != nil {
		return 0, err
	}
	if item == nil {
		return 0, nil
	}
	v, err := serialization.ReadUint32(bytes.NewBuffer(item.Value))
	if err != nil {
		return 0, err
	}
	return v, nil
}

func GenUInt64StorageItem(value uint64) *cstates.StorageItem {
	bf := new(bytes.Buffer)
	serialization.WriteUint64(bf, value)
	return &cstates.StorageItem{Value: bf.Bytes()}
}

func GenUInt32StorageItem(value uint32) *cstates.StorageItem {
	bf := new(bytes.Buffer)
	serialization.WriteUint32(bf, value)
	return &cstates.StorageItem{Value: bf.Bytes()}
}

func PutBytes(native *native.NativeService, key []byte, value []byte) {
	native.CacheDB.Put(key, cstates.GenRawStorageItem(value))
}
