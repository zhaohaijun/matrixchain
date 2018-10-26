package types

import (
	"fmt"
	"io"

	"github.com/zhaohaijun/matrixchain/common"
	ct "github.com/zhaohaijun/matrixchain/core/types"
	comm "github.com/zhaohaijun/matrixchain/p2pserver/common"
)

type BlkHeader struct {
	BlkHdr []*ct.Header
}

//Serialize message payload
func (this BlkHeader) Serialization(sink *common.ZeroCopySink) error {
	sink.WriteUint32(uint32(len(this.BlkHdr)))

	for _, header := range this.BlkHdr {
		err := header.Serialization(sink)
		if err != nil {
			return err
		}
	}

	return nil
}

func (this *BlkHeader) CmdType() string {
	return comm.HEADERS_TYPE
}

//Deserialize message payload
func (this *BlkHeader) Deserialization(source *common.ZeroCopySource) error {
	var count uint32
	count, eof := source.NextUint32()
	if eof {
		return io.ErrUnexpectedEOF
	}

	for i := 0; i < int(count); i++ {
		var headers ct.Header
		err := headers.Deserialization(source)
		if err != nil {
			return fmt.Errorf("deserialze BlkHeader error: %v", err)
		}
		this.BlkHdr = append(this.BlkHdr, &headers)
	}
	return nil
}
