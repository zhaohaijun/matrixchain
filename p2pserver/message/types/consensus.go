package types

import (
	comm "github.com/zhaohaijun/matrixchain/common"
	"github.com/zhaohaijun/matrixchain/p2pserver/common"
)

type Consensus struct {
	Cons ConsensusPayload
}

//Serialize message payload
func (this *Consensus) Serialization(sink *comm.ZeroCopySink) error {
	return this.Cons.Serialization(sink)
}

func (this *Consensus) CmdType() string {
	return common.CONSENSUS_TYPE
}

//Deserialize message payload
func (this *Consensus) Deserialization(source *comm.ZeroCopySource) error {
	return this.Cons.Deserialization(source)
}
