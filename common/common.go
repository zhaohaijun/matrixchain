package common

import (
	"encoding/hex"
	"math/rand"
	"os"
)

func GetNonce() uint64 {
	nonce := uint64(rand.Uint32())<<32 + uint64(rand.Uint32())
	return nonce
}
func ToHexString(data []byte) string {
	return hex.EncodeToString(data)
}
func HexToBytes(value string) ([]byte, error) {
	return hex.DecodeString(value)
}
func ToArrayReverse(arr []byte) []byte {
	l := len(arr)
	x := make([]byte, 0)
	for i := l - 1; i >= 0; i-- {
		x = append(x, arr[i])
	}
	return x
}
func FileExisted(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
