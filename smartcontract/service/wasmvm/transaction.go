package wasmvm

import (
	"github.com/zhaohaijun/matrixchain/core/types"
	"github.com/zhaohaijun/matrixchain/errors"
	"github.com/zhaohaijun/matrixchain/vm/wasmvm/exec"
)

func (this *WasmVmService) transactionGetHash(engine *exec.ExecutionEngine) (bool, error) {
	vm := engine.GetVM()
	envCall := vm.GetEnvCall()
	params := envCall.GetParams()
	if len(params) != 1 {
		return false, errors.NewErr("[transactionGetHash] parameter count error")
	}

	transbytes, err := vm.GetPointerMemory(params[0])
	if err != nil {
		return false, err
	}

	trans, err := types.TransactionFromRawBytes(transbytes)
	if err != nil {
		return false, err
	}
	hash := trans.Hash()
	idx, err := vm.SetPointerMemory(hash.ToArray())
	if err != nil {
		return false, err
	}
	vm.RestoreCtx()
	vm.PushResult(uint64(idx))
	return true, nil
}
func (this *WasmVmService) transactionGetType(engine *exec.ExecutionEngine) (bool, error) {
	vm := engine.GetVM()
	envCall := vm.GetEnvCall()
	params := envCall.GetParams()
	if len(params) != 1 {
		return false, errors.NewErr("[transactionGetType] parameter count error")
	}

	transbytes, err := vm.GetPointerMemory(params[0])
	if err != nil {
		return false, err
	}

	trans, err := types.TransactionFromRawBytes(transbytes)
	if err != nil {
		return false, err
	}
	txtype := trans.TxType

	vm.RestoreCtx()
	vm.PushResult(uint64(txtype))
	return true, nil
}
func (this *WasmVmService) transactionGetAttributes(engine *exec.ExecutionEngine) (bool, error) {
	vm := engine.GetVM()
	envCall := vm.GetEnvCall()
	params := envCall.GetParams()
	if len(params) != 1 {
		return false, errors.NewErr("[transactionGetAttributes] parameter count error")
	}

	transbytes, err := vm.GetPointerMemory(params[0])
	if err != nil {
		return false, err
	}

	_, err = types.TransactionFromRawBytes(transbytes)
	if err != nil {
		return false, err
	}
	attributes := make([][]byte, 0)

	idx, err := vm.SetPointerMemory(attributes)
	if err != nil {
		return false, err
	}
	vm.RestoreCtx()
	vm.PushResult(uint64(idx))
	return true, nil
}
