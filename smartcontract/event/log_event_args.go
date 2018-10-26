package event

import (
	"github.com/zhaohaijun/matrixchain/common"
)

// LogEventArgs describe smart contract event log struct
type LogEventArgs struct {
	TxHash          common.Uint256
	ContractAddress common.Address
	Message         string
}
