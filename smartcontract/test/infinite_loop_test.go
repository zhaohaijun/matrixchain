package test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhaohaijun/matrixchain/core/types"
	. "github.com/zhaohaijun/matrixchain/smartcontract"
)

func TestInfiniteLoopCrash(t *testing.T) {
	evilBytecode := []byte(" e\xff\u007f\xffhm\xb7%\xa7AAAAAAAAAAAAAAAC\xef\xed\x04INVERT\x95ve")
	dbFile := "test"
	defer func() {
		os.RemoveAll(dbFile)
	}()
	//testLevelDB, err := leveldbstore.NewLevelDBStore(dbFile)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//store := statestore.NewMemDatabase()
	//testBatch := statestore.NewStateStoreBatch(store, testLevelDB)
	config := &Config{
		Time:   10,
		Height: 10,
		Tx:     &types.Transaction{},
	}
	//cache := storage.NewCloneCache(testBatch)
	sc := SmartContract{
		Config:  config,
		Gas:     10000,
		CacheDB: nil,
	}
	engine, err := sc.NewExecuteEngine(evilBytecode)
	if err != nil {
		t.Fatal(err)
	}
	_, err = engine.Invoke()
	assert.Equal(t, "[NeoVmService] vm execute error!: the biginteger over max size 32bit", err.Error())
}
