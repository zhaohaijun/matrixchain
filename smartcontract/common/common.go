package common

import (
	"errors"

	"github.com/zhaohaijun/matrixchain/common"
	"github.com/zhaohaijun/matrixchain/common/log"
	"github.com/zhaohaijun/matrixchain/vm/neovm/types"
)

// ConvertReturnTypes return neovm stack element value
// According item types convert to hex string value
// Now neovm support type contain: ByteArray/Integer/Boolean/Array/Struct/Interop/StackItems
const (
	MAX_COUNT = 1024
)

func ConvertNeoVmTypeHexString(item interface{}) (interface{}, error) {
	var count int
	return convertNeoVmTypeHexString(item, &count)
}

func convertNeoVmTypeHexString(item interface{}, count *int) (interface{}, error) {
	if item == nil {
		return nil, nil
	}
	if *count > MAX_COUNT {
		return nil, errors.New("over max parameters convert length")
	}
	switch v := item.(type) {
	case *types.ByteArray:
		arr, _ := v.GetByteArray()
		return common.ToHexString(arr), nil
	case *types.Integer:
		i, _ := v.GetBigInteger()
		if i.Sign() == 0 {
			return common.ToHexString([]byte{0}), nil
		} else {
			return common.ToHexString(common.BigIntToNeoBytes(i)), nil
		}
	case *types.Boolean:
		b, _ := v.GetBoolean()
		if b {
			return common.ToHexString([]byte{1}), nil
		} else {
			return common.ToHexString([]byte{0}), nil
		}
	case *types.Array:
		var arr []interface{}
		ar, _ := v.GetArray()
		for _, val := range ar {
			*count++
			cv, err := convertNeoVmTypeHexString(val, count)
			if err != nil {
				return nil, err
			}
			arr = append(arr, cv)
		}
		return arr, nil
	case *types.Struct:
		var arr []interface{}
		ar, _ := v.GetStruct()
		for _, val := range ar {
			*count++
			cv, err := convertNeoVmTypeHexString(val, count)
			if err != nil {
				return nil, err
			}
			arr = append(arr, cv)
		}
		return arr, nil
	case *types.Interop:
		it, _ := v.GetInterface()
		return common.ToHexString(it.ToArray()), nil
	default:
		log.Error("[ConvertTypes] Invalid Types!")
		return nil, errors.New("[ConvertTypes] Invalid Types!")
	}
}
