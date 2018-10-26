package wasmvm

import (
	"bytes"

	"github.com/zhaohaijun/matrixchain/core/types"
	"github.com/zhaohaijun/matrixchain/errors"
	"github.com/zhaohaijun/matrixchain/vm/wasmvm/exec"
)

func (this *WasmVmService) headerGetHash(engine *exec.ExecutionEngine) (bool, error) {
	vm := engine.GetVM()
	envCall := vm.GetEnvCall()
	params := envCall.GetParams()
	if len(params) != 1 {
		return false, errors.NewErr("[transactionGetHash] parameter count error")
	}

	headerbytes, err := vm.GetPointerMemory(params[0])
	if err != nil {
		return false, err
	}
	header := &types.Header{}
	err = header.Deserialize(bytes.NewBuffer(headerbytes))
	if err != nil {
		return false, err
	}

	hash := header.Hash()
	idx, err := vm.SetPointerMemory(hash.ToArray())
	if err != nil {
		return false, err
	}
	vm.RestoreCtx()
	vm.PushResult(uint64(idx))
	return true, nil
}

func (this *WasmVmService) headerGetVersion(engine *exec.ExecutionEngine) (bool, error) {
	vm := engine.GetVM()
	envCall := vm.GetEnvCall()
	params := envCall.GetParams()
	if len(params) != 1 {
		return false, errors.NewErr("[transactionGetHash] parameter count error")
	}

	headerbytes, err := vm.GetPointerMemory(params[0])
	if err != nil {
		return false, err
	}
	header := &types.Header{}
	err = header.Deserialize(bytes.NewBuffer(headerbytes))
	if err != nil {
		return false, err
	}

	version := header.Version

	vm.RestoreCtx()
	vm.PushResult(uint64(version))
	return true, nil
}

func (this *WasmVmService) headerGetPrevHash(engine *exec.ExecutionEngine) (bool, error) {
	vm := engine.GetVM()
	envCall := vm.GetEnvCall()
	params := envCall.GetParams()
	if len(params) != 1 {
		return false, errors.NewErr("[transactionGetHash] parameter count error")
	}

	headerbytes, err := vm.GetPointerMemory(params[0])
	if err != nil {
		return false, err
	}
	header := &types.Header{}
	err = header.Deserialize(bytes.NewBuffer(headerbytes))
	if err != nil {
		return false, err
	}

	hash := header.PrevBlockHash.ToArray()
	idx, err := vm.SetPointerMemory(hash)
	if err != nil {
		return false, err
	}
	vm.RestoreCtx()
	vm.PushResult(uint64(idx))
	return true, nil
}

func (this *WasmVmService) headerGetMerkleRoot(engine *exec.ExecutionEngine) (bool, error) {
	vm := engine.GetVM()
	envCall := vm.GetEnvCall()
	params := envCall.GetParams()
	if len(params) != 1 {
		return false, errors.NewErr("[transactionGetHash] parameter count error")
	}

	headerbytes, err := vm.GetPointerMemory(params[0])
	if err != nil {
		return false, err
	}
	header := &types.Header{}
	err = header.Deserialize(bytes.NewBuffer(headerbytes))
	if err != nil {
		return false, err
	}

	hash := header.TransactionsRoot.ToArray()
	idx, err := vm.SetPointerMemory(hash)
	if err != nil {
		return false, err
	}
	vm.RestoreCtx()
	vm.PushResult(uint64(idx))
	return true, nil
}

func (this *WasmVmService) headerGetIndex(engine *exec.ExecutionEngine) (bool, error) {
	vm := engine.GetVM()
	envCall := vm.GetEnvCall()
	params := envCall.GetParams()
	if len(params) != 1 {
		return false, errors.NewErr("[transactionGetHash] parameter count error")
	}

	headerbytes, err := vm.GetPointerMemory(params[0])
	if err != nil {
		return false, err
	}
	header := &types.Header{}
	err = header.Deserialize(bytes.NewBuffer(headerbytes))
	if err != nil {
		return false, err
	}

	height := header.Height

	vm.RestoreCtx()
	vm.PushResult(uint64(height))
	return true, nil
}

func (this *WasmVmService) headerGetTimestamp(engine *exec.ExecutionEngine) (bool, error) {
	vm := engine.GetVM()
	envCall := vm.GetEnvCall()
	params := envCall.GetParams()
	if len(params) != 1 {
		return false, errors.NewErr("[transactionGetHash] parameter count error")
	}

	headerbytes, err := vm.GetPointerMemory(params[0])
	if err != nil {
		return false, err
	}
	header := &types.Header{}
	err = header.Deserialize(bytes.NewBuffer(headerbytes))
	if err != nil {
		return false, err
	}

	tm := header.Timestamp

	vm.RestoreCtx()
	vm.PushResult(uint64(tm))
	return true, nil
}

func (this *WasmVmService) headerGetConsensusData(engine *exec.ExecutionEngine) (bool, error) {
	vm := engine.GetVM()
	envCall := vm.GetEnvCall()
	params := envCall.GetParams()
	if len(params) != 1 {
		return false, errors.NewErr("[transactionGetHash] parameter count error")
	}

	headerbytes, err := vm.GetPointerMemory(params[0])
	if err != nil {
		return false, err
	}
	header := &types.Header{}
	err = header.Deserialize(bytes.NewBuffer(headerbytes))
	if err != nil {
		return false, err
	}

	cd := header.ConsensusData

	vm.RestoreCtx()
	vm.PushResult(cd)
	return true, nil
}

func (this *WasmVmService) headerGetNextConsensus(engine *exec.ExecutionEngine) (bool, error) {
	vm := engine.GetVM()
	envCall := vm.GetEnvCall()
	params := envCall.GetParams()
	if len(params) != 1 {
		return false, errors.NewErr("[transactionGetHash] parameter count error")
	}

	headerbytes, err := vm.GetPointerMemory(params[0])
	if err != nil {
		return false, err
	}
	header := &types.Header{}
	err = header.Deserialize(bytes.NewBuffer(headerbytes))
	if err != nil {
		return false, err
	}

	cd := header.NextBookkeeper[:]
	idx, err := vm.SetPointerMemory(cd)
	if err != nil {
		return false, err
	}
	vm.RestoreCtx()
	vm.PushResult(uint64(idx))
	return true, nil
}
