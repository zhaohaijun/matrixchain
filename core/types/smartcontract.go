package types

import "github.com/zhaohaijun/matrixchain/common"

type SmartCodeEvent struct {
	TxHash common.Uint256
	Action string
	Result interface{}
	Error  int64
}
