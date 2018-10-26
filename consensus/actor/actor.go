package actor

import (
	"errors"
	"time"

	"github.com/zhaohaijun/go-async-queue/actor"
	"github.com/zhaohaijun/matrixchain/core/types"
	ontErrors "github.com/zhaohaijun/matrixchain/errors"
	netActor "github.com/zhaohaijun/matrixchain/p2pserver/actor/server"
	ptypes "github.com/zhaohaijun/matrixchain/p2pserver/message/types"
	txpool "github.com/zhaohaijun/matrixchain/txnpool/common"
)

type TxPoolActor struct {
	Pool *actor.PID
}

func (self *TxPoolActor) GetTxnPool(byCount bool, height uint32) []*txpool.TXEntry {
	poolmsg := &txpool.GetTxnPoolReq{ByCount: byCount, Height: height}
	future := self.Pool.RequestFuture(poolmsg, time.Second*10)
	entry, err := future.Result()
	if err != nil {
		return nil
	}

	txs := entry.(*txpool.GetTxnPoolRsp).TxnPool
	return txs
}

func (self *TxPoolActor) VerifyBlock(txs []*types.Transaction, height uint32) error {
	poolmsg := &txpool.VerifyBlockReq{Txs: txs, Height: height}
	future := self.Pool.RequestFuture(poolmsg, time.Second*10)
	entry, err := future.Result()
	if err != nil {
		return err
	}

	txentry := entry.(*txpool.VerifyBlockRsp).TxnPool
	for _, entry := range txentry {
		if entry.ErrCode != ontErrors.ErrNoError {
			return errors.New(entry.ErrCode.Error())
		}
	}

	return nil
}

type P2PActor struct {
	P2P *actor.PID
}

func (self *P2PActor) Broadcast(msg interface{}) {
	self.P2P.Tell(msg)
}

func (self *P2PActor) Transmit(target uint64, msg ptypes.Message) {
	self.P2P.Tell(&netActor.TransmitConsensusMsgReq{
		Target: target,
		Msg:    msg,
	})
}

type LedgerActor struct {
	Ledger *actor.PID
}
