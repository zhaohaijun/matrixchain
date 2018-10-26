package common

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zhaohaijun/matrixchain/common/log"
	"github.com/zhaohaijun/matrixchain/core/payload"
	"github.com/zhaohaijun/matrixchain/core/types"
)

var (
	txn *types.Transaction
)

func init() {
	log.Init(log.PATH, log.Stdout)

	mutable := &types.MutableTransaction{
		TxType:  types.Invoke,
		Nonce:   uint32(time.Now().Unix()),
		Payload: &payload.InvokeCode{Code: []byte{}},
	}

	txn, _ = mutable.IntoImmutable()
}

func TestTxPool(t *testing.T) {
	txPool := &TXPool{}
	txPool.Init()

	txEntry := &TXEntry{
		Tx:    txn,
		Attrs: []*TXAttr{},
	}

	ret := txPool.AddTxList(txEntry)
	if ret == false {
		t.Error("Failed to add tx to the pool")
		return
	}

	ret = txPool.AddTxList(txEntry)
	if ret == true {
		t.Error("Failed to add tx to the pool")
		return
	}

	txList, oldTxList := txPool.GetTxPool(true, 0)
	for _, v := range txList {
		assert.NotNil(t, v)
	}

	for _, v := range oldTxList {
		assert.NotNil(t, v)
	}

	entry := txPool.GetTransaction(txn.Hash())
	if entry == nil {
		t.Error("Failed to get the transaction")
		return
	}

	assert.Equal(t, txn.Hash(), entry.Hash())

	status := txPool.GetTxStatus(txn.Hash())
	if status == nil {
		t.Error("failed to get the status")
		return
	}

	assert.Equal(t, txn.Hash(), status.Hash)

	count := txPool.GetTransactionCount()
	assert.Equal(t, count, 1)

	err := txPool.CleanTransactionList([]*types.Transaction{txn})
	if err != nil {
		t.Error("Failed to clean transaction list")
		return
	}
}
