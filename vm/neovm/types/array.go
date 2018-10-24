package types

import (
	"fmt"
	"math/big"

	"github.com/zhaohaijun/matrixchain/vm/neovm/interfaces"
)

type Array struct {
	_array []StackItems
}

func NewArray(value []StackItems) *Array {
	var this Array
	this._array = value
	return &this
}

func (this *Array) Equals(other StackItems) bool {
	return this == other
}

func (this *Array) GetBigInteger() (*big.Int, error) {
	return nil, fmt.Errorf("%s", "Not support array to integer")
}

func (this *Array) GetBoolean() (bool, error) {
	return false, fmt.Errorf("%s", "Not support array to boolean")
}

func (this *Array) GetByteArray() ([]byte, error) {
	return nil, fmt.Errorf("%s", "Not support array to byte array")
}

func (this *Array) GetInterface() (interfaces.Interop, error) {
	return nil, fmt.Errorf("%s", "Not support array to interface")
}

func (this *Array) GetArray() ([]StackItems, error) {
	return this._array, nil
}

func (this *Array) GetStruct() ([]StackItems, error) {
	return this._array, nil
}

func (this *Array) GetMap() (map[StackItems]StackItems, error) {
	return nil, fmt.Errorf("%s", "Not support array to map")
}

func (this *Array) Add(item StackItems) {
	this._array = append(this._array, item)
}

func (this *Array) RemoveAt(index int) {
	this._array = append(this._array[:index-1], this._array[index:]...)
}

func (this *Array) Count() int {
	return len(this._array)
}

func (this *Array) IsMapKey() bool {
	return false
}
