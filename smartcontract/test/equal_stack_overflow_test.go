package test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhaohaijun/matrixchain/common/log"
	"github.com/zhaohaijun/matrixchain/core/types"
	. "github.com/zhaohaijun/matrixchain/smartcontract"
	"github.com/zhaohaijun/matrixchain/vm/neovm"
)

func TestEqualStackOverflow(t *testing.T) {
	log.InitLog(4)
	defer func() {
		os.RemoveAll("./Log")
	}()

	code := []byte{
		byte(neovm.PUSH1),    // {1}
		byte(neovm.NEWARRAY), // {[]}
		byte(neovm.DUP),      // {[],[]}
		byte(neovm.DUP),      // {[],[],[]}
		byte(neovm.PUSH0),    // {[],[],[],0}
		byte(neovm.ROT),      // {[],[],0,[]}
		byte(neovm.SETITEM),  // {[[]]}
		byte(neovm.DUP),      // {[[]],[[]]}
		byte(neovm.EQUAL),
	}

	config := &Config{
		Time:   10,
		Height: 10,
		Tx:     &types.Transaction{},
	}
	sc := SmartContract{
		Config:  config,
		Gas:     10000,
		CacheDB: nil,
	}
	engine, _ := sc.NewExecuteEngine(code)
	_, err := engine.Invoke()

	assert.Nil(t, err)
}
