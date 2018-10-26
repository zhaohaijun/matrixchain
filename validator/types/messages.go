package types

import (
	"github.com/zhaohaijun/go-async-queue/actor"
	"github.com/zhaohaijun/matrixchain/common"
	"github.com/zhaohaijun/matrixchain/core/types"
	"github.com/zhaohaijun/matrixchain/errors"
)

// message
type RegisterValidator struct {
	Sender *actor.PID
	Type   VerifyType
	Id     string
}

type UnRegisterValidator struct {
	Id   string
	Type VerifyType
}

type UnRegisterAck struct {
	Id   string
	Type VerifyType
}

type CheckTx struct {
	WorkerId uint8
	Tx       *types.Transaction
}

type CheckResponse struct {
	WorkerId uint8
	Type     VerifyType
	Hash     common.Uint256
	Height   uint32
	ErrCode  errors.ErrCode
}

// VerifyType of validator
type VerifyType uint8

const (
	Stateless VerifyType = iota
	Stateful  VerifyType = iota
)
