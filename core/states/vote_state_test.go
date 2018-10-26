package states

import (
	"testing"

	"bytes"

	"github.com/stretchr/testify/assert"
	"github.com/zhaohaijun/blockchain-crypto/keypair"
)

func TestVoteState_Deserialize_Serialize(t *testing.T) {
	_, pubKey1, _ := keypair.GenerateKeyPair(keypair.PK_ECDSA, keypair.P256)
	_, pubKey2, _ := keypair.GenerateKeyPair(keypair.PK_ECDSA, keypair.P256)

	vs := VoteState{
		StateBase:  StateBase{(byte)(1)},
		PublicKeys: []keypair.PublicKey{pubKey1, pubKey2},
		Count:      10,
	}

	buf := bytes.NewBuffer(nil)
	vs.Serialize(buf)
	bs := buf.Bytes()

	var vs2 VoteState
	vs2.Deserialize(buf)
	assert.Equal(t, vs, vs2)

	buf = bytes.NewBuffer(bs[:len(bs)-1])
	err := vs2.Deserialize(buf)
	assert.NotNil(t, err)
}
