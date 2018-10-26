package program

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhaohaijun/blockchain-crypto/keypair"
)

func TestProgramBuilder_PushBytes(t *testing.T) {
	fmt.Println("aaaa")
	N := 20000
	fmt.Println("hello 222222222")
	fmt.Println("aaaa")
	builder := NewProgramBuilder()
	fmt.Println("hello 1")
	for i := 0; i < N; i++ {
		builder.PushNum(uint16(i))
	}

	parser := newProgramParser(builder.Finish())
	for i := 0; i < N; i++ {
		n, err := parser.ReadNum()
		assert.Nil(t, err)
		assert.Equal(t, n, uint16(i))
	}
}

func TestGetProgramInfo(t *testing.T) {
	N := 10
	M := N / 2
	var pubkeys []keypair.PublicKey
	for i := 0; i < N; i++ {
		_, key, _ := keypair.GenerateKeyPair(keypair.PK_ECDSA, keypair.P256)
		pubkeys = append(pubkeys, key)
	}

	pubkeys = keypair.SortPublicKeys(pubkeys)
	progInfo := ProgramInfo{PubKeys: pubkeys, M: uint16(M)}
	prog, err := ProgramFromMultiPubKey(progInfo.PubKeys, int(progInfo.M))
	assert.Nil(t, err)

	info2, err := GetProgramInfo(prog)
	assert.Nil(t, err)
	assert.Equal(t, progInfo, info2)
}
