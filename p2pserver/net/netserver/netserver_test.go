package netserver

import (
	"fmt"
	"testing"
	"time"

	"github.com/zhaohaijun/matrixchain/common/log"
	"github.com/zhaohaijun/matrixchain/p2pserver/common"
	"github.com/zhaohaijun/matrixchain/p2pserver/peer"
)

func init() {
	log.Init(log.Stdout)
	fmt.Println("Start test the netserver...")
}

func creatPeers(cnt uint16) []*peer.Peer {
	np := []*peer.Peer{}
	var syncport uint16
	var consport uint16
	var id uint64
	var height uint64
	for i := uint16(0); i < cnt; i++ {
		syncport = 20224 + i
		consport = 20335 + i
		id = 0x7533345 + uint64(i)
		height = 434923 + uint64(i)
		p := peer.NewPeer()
		p.UpdateInfo(time.Now(), 2, 3, syncport, consport, id, 0, height)
		p.SetConsState(2)
		p.SetSyncState(4)
		p.SetHttpInfoState(true)
		p.SyncLink.SetAddr("127.0.0.1:10338")
		np = append(np, p)
	}
	return np

}
func TestNewNetServer(t *testing.T) {
	server := NewNetServer()
	server.Start()
	defer server.Halt()

	server.SetHeight(1000)
	if server.GetHeight() != 1000 {
		t.Error("TestNewNetServer set server height error")
	}

	if server.GetRelay() != true {
		t.Error("TestNewNetServer server relay state error", server.GetRelay())
	}
	if server.GetServices() != 1 {
		t.Error("TestNewNetServer server service state error", server.GetServices())
	}
	if server.GetVersion() != common.PROTOCOL_VERSION {
		t.Error("TestNewNetServer server version error", server.GetVersion())
	}
	if server.GetSyncPort() != 20338 {
		t.Error("TestNewNetServer sync port error", server.GetSyncPort())
	}
	if server.GetConsPort() != 20339 {
		t.Error("TestNewNetServer sync port error", server.GetConsPort())
	}

	fmt.Printf("lastest server time is %s\n", time.Unix(server.GetTime()/1e9, 0).String())

}

func TestNetServerNbrPeer(t *testing.T) {
	log.Init(log.Stdout)
	server := NewNetServer()
	server.Start()
	defer server.Halt()

	nm := &peer.NbrPeers{}
	nm.Init()
	np := creatPeers(5)
	for _, v := range np {
		server.AddNbrNode(v)
	}
	if server.GetConnectionCnt() != 5 {
		t.Error("TestNetServerNbrPeer GetConnectionCnt error", server.GetConnectionCnt())
	}
	addrs := server.GetNeighborAddrs()
	if len(addrs) != 5 {
		t.Error("TestNetServerNbrPeer GetNeighborAddrs error")
	}
	if server.NodeEstablished(0x7533345) == false {
		t.Error("TestNetServerNbrPeer NodeEstablished error")
	}
	if server.GetPeer(0x7533345) == nil {
		t.Error("TestNetServerNbrPeer GetPeer error")
	}
	p, ok := server.DelNbrNode(0x7533345)
	if ok != true || p == nil {
		t.Error("TestNetServerNbrPeer DelNbrNode error")
	}
	if len(server.GetNeighbors()) != 4 {
		t.Error("TestNetServerNbrPeer GetNeighbors error")
	}
	sp := &peer.Peer{}
	cp := &peer.Peer{}
	server.AddPeerSyncAddress("127.0.0.1:10338", sp)
	server.AddPeerConsAddress("127.0.0.1:20338", cp)
	if server.GetPeerFromAddr("127.0.0.1:10338") != sp {
		t.Error("TestNetServerNbrPeer Get/AddPeerConsAddress error")
	}

}
