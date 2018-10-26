package states

import (
	"io"

	"github.com/zhaohaijun/matrixchain/common"
	"github.com/zhaohaijun/matrixchain/common/serialization"
	"github.com/zhaohaijun/matrixchain/errors"
)

// Invoke smart contract struct
// Param Version: invoke smart contract version, default 0
// Param Address: invoke on blockchain smart contract by address
// Param Method: invoke smart contract method, default ""
// Param Args: invoke smart contract arguments
type ContractInvokeParam struct {
	Version byte
	Address common.Address
	Method  string
	Args    []byte
}

// Serialize contract
func (this *ContractInvokeParam) Serialize(w io.Writer) error {
	if err := serialization.WriteByte(w, this.Version); err != nil {
		return errors.NewDetailErr(err, errors.ErrNoCode, "[ContractInvokeParam] Version serialize error!")
	}
	if err := this.Address.Serialize(w); err != nil {
		return errors.NewDetailErr(err, errors.ErrNoCode, "[ContractInvokeParam] Address serialize error!")
	}
	if err := serialization.WriteVarBytes(w, []byte(this.Method)); err != nil {
		return errors.NewDetailErr(err, errors.ErrNoCode, "[ContractInvokeParam] Method serialize error!")
	}
	if err := serialization.WriteVarBytes(w, this.Args); err != nil {
		return errors.NewDetailErr(err, errors.ErrNoCode, "[ContractInvokeParam] Args serialize error!")
	}
	return nil
}

func (this *ContractInvokeParam) Serialization(sink *common.ZeroCopySink) {
	sink.WriteByte(this.Version)
	sink.WriteAddress(this.Address)
	sink.WriteVarBytes([]byte(this.Method))
	sink.WriteVarBytes([]byte(this.Args))
}

// Deserialize contract
func (this *ContractInvokeParam) Deserialize(r io.Reader) error {
	var err error
	this.Version, err = serialization.ReadByte(r)
	if err != nil {
		return errors.NewDetailErr(err, errors.ErrNoCode, "[ContractInvokeParam] Version deserialize error!")
	}

	if err := this.Address.Deserialize(r); err != nil {
		return errors.NewDetailErr(err, errors.ErrNoCode, "[ContractInvokeParam] Address deserialize error!")
	}

	method, err := serialization.ReadVarBytes(r)
	if err != nil {
		return errors.NewDetailErr(err, errors.ErrNoCode, "[ContractInvokeParam] Method deserialize error!")
	}
	this.Method = string(method)

	this.Args, err = serialization.ReadVarBytes(r)
	if err != nil {
		return errors.NewDetailErr(err, errors.ErrNoCode, "[ContractInvokeParam] Args deserialize error!")
	}
	return nil
}

// `ContractInvokeParam.Args` has reference of `source`
func (this *ContractInvokeParam) Deserialization(source *common.ZeroCopySource) error {
	var irregular, eof bool
	this.Version, eof = source.NextByte()
	this.Address, eof = source.NextAddress()
	var method []byte
	method, _, irregular, eof = source.NextVarBytes()
	if irregular {
		return common.ErrIrregularData
	}
	this.Method = string(method)

	this.Args, _, irregular, eof = source.NextVarBytes()
	if irregular {
		return common.ErrIrregularData
	}
	if eof {
		return io.ErrUnexpectedEOF
	}
	return nil
}

type PreExecResult struct {
	State  byte
	Gas    uint64
	Result interface{}
}
