package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhaohaijun/blockchain-crypto/keypair"
)

func TestAddressFromBookkeepers(t *testing.T) {
	_, pubKey1, _ := keypair.GenerateKeyPair(keypair.PK_ECDSA, keypair.P256)
	_, pubKey2, _ := keypair.GenerateKeyPair(keypair.PK_ECDSA, keypair.P256)
	_, pubKey3, _ := keypair.GenerateKeyPair(keypair.PK_ECDSA, keypair.P256)
	pubkeys := []keypair.PublicKey{pubKey1, pubKey2, pubKey3}

	addr, _ := AddressFromBookkeepers(pubkeys)
	addr2, _ := AddressFromMultiPubKeys(pubkeys, 3)
	assert.Equal(t, addr, addr2)

	pubkeys = []keypair.PublicKey{pubKey3, pubKey2, pubKey1}
	addr3, _ := AddressFromMultiPubKeys(pubkeys, 3)

	assert.Equal(t, addr2, addr3)
}
