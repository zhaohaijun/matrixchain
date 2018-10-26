package payload

import (
	"fmt"
	"io"

	"github.com/zhaohaijun/matrixchain/common"
	"github.com/zhaohaijun/matrixchain/common/serialization"
)

// InvokeCode is an implementation of transaction payload for invoke smartcontract
type InvokeCode struct {
	Code []byte
}

func (self *InvokeCode) Serialize(w io.Writer) error {
	if err := serialization.WriteVarBytes(w, self.Code); err != nil {
		return fmt.Errorf("InvokeCode Code Serialize failed: %s", err)
	}
	return nil
}

func (self *InvokeCode) Deserialize(r io.Reader) error {
	code, err := serialization.ReadVarBytes(r)
	if err != nil {
		return fmt.Errorf("InvokeCode Code Deserialize failed: %s", err)
	}
	self.Code = code
	return nil
}

//note: InvokeCode.Code has data reference of param source
func (self *InvokeCode) Deserialization(source *common.ZeroCopySource) error {
	code, _, irregular, eof := source.NextVarBytes()
	if eof {
		return io.ErrUnexpectedEOF
	}
	if irregular {
		return common.ErrIrregularData
	}

	self.Code = code
	return nil
}

func (self *InvokeCode) Serialization(sink *common.ZeroCopySink) error {
	sink.WriteVarBytes(self.Code)
	return nil
}
