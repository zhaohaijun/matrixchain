package states

import (
	"bytes"
	"testing"

	"github.com/zhaohaijun/matrixchain/common"
)

func TestContract_Serialize_Deserialize(t *testing.T) {
	addr := common.AddressFromVmCode([]byte{1})

	c := &ContractInvokeParam{
		Version: 0,
		Address: addr,
		Method:  "init",
		Args:    []byte{2},
	}
	bf := new(bytes.Buffer)
	if err := c.Serialize(bf); err != nil {
		t.Fatalf("ContractInvokeParam serialize error: %v", err)
	}

	v := new(ContractInvokeParam)
	if err := v.Deserialize(bf); err != nil {
		t.Fatalf("ContractInvokeParam deserialize error: %v", err)
	}
}
