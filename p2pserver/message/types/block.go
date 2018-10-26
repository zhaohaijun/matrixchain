package types

import (
	"fmt"

	"github.com/zhaohaijun/matrixchain/common"
	ct "github.com/zhaohaijun/matrixchain/core/types"
	"github.com/zhaohaijun/matrixchain/errors"
	comm "github.com/zhaohaijun/matrixchain/p2pserver/common"
)

type Block struct {
	Blk *ct.Block
}

//Serialize message payload
func (this *Block) Serialization(sink *common.ZeroCopySink) error {
	err := this.Blk.Serialization(sink)
	if err != nil {
		return errors.NewDetailErr(err, errors.ErrNetPackFail, fmt.Sprintf("serialize error. err:%v", err))
	}

	return nil
}

func (this *Block) CmdType() string {
	return comm.BLOCK_TYPE
}

//Deserialize message payload
func (this *Block) Deserialization(source *common.ZeroCopySource) error {
	this.Blk = new(ct.Block)
	err := this.Blk.Deserialization(source)
	if err != nil {
		return errors.NewDetailErr(err, errors.ErrNetUnPackFail, fmt.Sprintf("read Blk error. err:%v", err))
	}

	return nil
}
