package init

import (
	"bytes"
	"math/big"

	"github.com/zhaohaijun/matrixchain/common"
	"github.com/zhaohaijun/matrixchain/smartcontract/service/native/auth"
	params "github.com/zhaohaijun/matrixchain/smartcontract/service/native/global_params"
	"github.com/zhaohaijun/matrixchain/smartcontract/service/native/governance"
	"github.com/zhaohaijun/matrixchain/smartcontract/service/native/ong"
	"github.com/zhaohaijun/matrixchain/smartcontract/service/native/ont"
	"github.com/zhaohaijun/matrixchain/smartcontract/service/native/ontid"
	"github.com/zhaohaijun/matrixchain/smartcontract/service/native/utils"
	"github.com/zhaohaijun/matrixchain/smartcontract/service/neovm"
	vm "github.com/zhaohaijun/matrixchain/vm/neovm"
)

var (
	COMMIT_DPOS_BYTES = InitBytes(utils.GovernanceContractAddress, governance.COMMIT_DPOS)
)

func init() {
	ong.InitOng()
	ont.InitOnt()
	params.InitGlobalParams()
	ontid.Init()
	auth.Init()
	governance.InitGovernance()
}

func InitBytes(addr common.Address, method string) []byte {
	bf := new(bytes.Buffer)
	builder := vm.NewParamsBuilder(bf)
	builder.EmitPushByteArray([]byte{})
	builder.EmitPushByteArray([]byte(method))
	builder.EmitPushByteArray(addr[:])
	builder.EmitPushInteger(big.NewInt(0))
	builder.Emit(vm.SYSCALL)
	builder.EmitPushByteArray([]byte(neovm.NATIVE_INVOKE_NAME))

	return builder.ToArray()
}
