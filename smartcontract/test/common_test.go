package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhaohaijun/matrixchain/common"
	"github.com/zhaohaijun/matrixchain/smartcontract"
)

func TestConvertNeoVmTypeHexString(t *testing.T) {
	code := `00c57676c8681553797374656d2e52756e74696d652e4e6f74696679`

	hex, err := common.HexToBytes(code)

	if err != nil {
		t.Fatal("hex to byte error:", err)
	}

	config := &smartcontract.Config{
		Time:   10,
		Height: 10,
		Tx:     nil,
	}
	//cache := storage.NewCloneCache(testBatch)
	sc := smartcontract.SmartContract{
		Config: config,
		Gas:    100000,
	}
	engine, err := sc.NewExecuteEngine(hex)

	_, err = engine.Invoke()

	assert.Error(t, err, "over max parameters convert length")
}
