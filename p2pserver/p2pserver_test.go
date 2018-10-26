package p2pserver

import (
	"fmt"
	"testing"

	"github.com/zhaohaijun/matrixchain/common/log"
	"github.com/zhaohaijun/matrixchain/p2pserver/common"
)

func init() {
	log.InitLog(log.InfoLog)
	fmt.Println("Start test the netserver...")

}
func TestNewP2PServer(t *testing.T) {
	log.Init(log.Stdout)
	fmt.Println("Start test new p2pserver...")

	p2p := NewServer()

	if p2p.GetVersion() != common.PROTOCOL_VERSION {
		t.Error("TestNewP2PServer p2p version error", p2p.GetVersion())
	}

	if p2p.GetVersion() != common.PROTOCOL_VERSION {
		t.Error("TestNewP2PServer p2p version error")
	}
	sync, cons := p2p.GetPort()
	if sync != 20338 {
		t.Error("TestNewP2PServer sync port error")
	}

	if cons != 20339 {
		t.Error("TestNewP2PServer consensus port error")
	}
}
