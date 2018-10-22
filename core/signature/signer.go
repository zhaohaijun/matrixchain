package signature

import (
	"github.com/zhaohaijun/blockchain-crypto/keypair"
	"github.com/zhaohaijun/blockchain-crypto/signature"
)

type Signer interface {
	PrivKey() keypair.PrivateKey
	PubKey() keypair.PublicKey
	Scheme() signature.SignatureScheme
}
