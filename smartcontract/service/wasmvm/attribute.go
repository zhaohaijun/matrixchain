package wasmvm

import (
	"bytes"

	"github.com/zhaohaijun/matrixchain/core/types"
	"github.com/zhaohaijun/matrixchain/errors"
	"github.com/zhaohaijun/matrixchain/vm/wasmvm/exec"
)

func (this *WasmVmService) attributeGetUsage(engine *exec.ExecutionEngine) (bool, error) {
	vm := engine.GetVM()
	envCall := vm.GetEnvCall()
	params := envCall.GetParams()
	if len(params) != 1 {
		return false, errors.NewErr("[transactionGetHash] parameter count error")
	}

	attributebytes, err := vm.GetPointerMemory(params[0])
	if err != nil {
		return false, nil
	}

	attr := types.TxAttribute{}
	err = attr.Deserialize(bytes.NewBuffer(attributebytes))
	if err != nil {
		return false, nil
	}
	vm.RestoreCtx()
	vm.PushResult(uint64(attr.Usage))
	return true, nil
}
func (this *WasmVmService) attributeGetData(engine *exec.ExecutionEngine) (bool, error) {
	vm := engine.GetVM()
	envCall := vm.GetEnvCall()
	params := envCall.GetParams()
	if len(params) != 1 {
		return false, errors.NewErr("[transactionGetHash] parameter count error")
	}

	attributebytes, err := vm.GetPointerMemory(params[0])
	if err != nil {
		return false, nil
	}

	attr := types.TxAttribute{}
	err = attr.Deserialize(bytes.NewBuffer(attributebytes))
	if err != nil {
		return false, nil
	}

	idx, err := vm.SetPointerMemory(attr.Data)
	if err != nil {
		return false, nil
	}

	vm.RestoreCtx()
	vm.PushResult(uint64(idx))
	return true, nil
}
