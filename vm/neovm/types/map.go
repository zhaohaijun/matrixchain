package types

import (
	"fmt"
	"math/big"

	"github.com/zhaohaijun/matrixchain/vm/neovm/interfaces"
)

type Map struct {
	_map map[StackItems]StackItems
}

func NewMap() *Map {
	var mp Map
	mp._map = make(map[StackItems]StackItems)
	return &mp
}

func (this *Map) Add(key StackItems, value StackItems) {
	for k := range this._map {
		if k.Equals(key) {
			delete(this._map, k)
			break
		}
	}
	this._map[key] = value
}

func (this *Map) Clear() {
	this._map = make(map[StackItems]StackItems)
}

func (this *Map) Remove(key StackItems) {
	for k := range this._map {
		if k.Equals(key) {
			delete(this._map, k)
			break
		}
	}
}

func (this *Map) Equals(that StackItems) bool {
	return this == that
}

func (this *Map) GetBoolean() (bool, error) {
	return true, nil
}

func (this *Map) GetByteArray() ([]byte, error) {
	return nil, fmt.Errorf("%s", "Not support map to byte array")
}

func (this *Map) GetBigInteger() (*big.Int, error) {
	return nil, fmt.Errorf("%s", "Not support map to integer")
}

func (this *Map) GetInterface() (interfaces.Interop, error) {
	return nil, fmt.Errorf("%s", "Not support map to interface")
}

func (this *Map) GetArray() ([]StackItems, error) {
	return nil, fmt.Errorf("%s", "Not support map to array")
}

func (this *Map) GetStruct() ([]StackItems, error) {
	return nil, fmt.Errorf("%s", "Not support map to struct")
}

func (this *Map) GetMap() (map[StackItems]StackItems, error) {
	return this._map, nil
}

func (this *Map) TryGetValue(key StackItems) StackItems {
	for k, v := range this._map {
		if k.Equals(key) {
			return v
		}
	}
	return nil
}

func (this *Map) IsMapKey() bool {
	return false
}
