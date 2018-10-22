package account

import (
	"github.com/zhaohaijun/matrixchain/log"

	"github.com/ontio/ontology/core/types"
	"github.com/zhaohaijun/matrixchain/common"

	"github.com/zhaohaijun/blockchain-crypto/keypair"
	s "github.com/zhaohaijun/blockchain-crypto/signature"
)

type Account struct {
	PrivateKey keypair.PrivateKey
	PublicKey  keypair.PublicKey
	Address    common.Address
	SigScheme  s.SignatureScheme
}

func NewAccount(encrypt string) *Account {
	var pkAlgorithm keypair.KeyType
	var params interface{}
	var scheme s.SignatureScheme
	var err error
	if "" != encrypt {
		scheme, err = s.GetScheme(encrypt)
	} else {
		scheme = s.SHA256withECDSA
	}
	if err != nil {
		log.Warn("unknow signature scheme,use SHA256withECDSA as default.")
		scheme = s.SHA256withECDSA
	}
	switch scheme {
	case s.SHA224withECDSA, s.SHA3_224withECDSA:
		pkAlgorithm = keypair.PK_ECDSA
		params = keypair.P224
	case s.SHA256withECDSA, s.SHA3_256withECDSA, s.RIPEMD160withECDSA:
		pkAlgorithm = keypair.PK_ECDSA
		params = keypair.P256
	case s.SHA384withECDSA, s.SHA3_384withECDSA:
		pkAlgorithm = keypair.PK_ECDSA
		params = keypair.P384
	case s.SHA512withECDSA, s.SHA3_512withECDSA:
		pkAlgorithm = keypair.PK_ECDSA
		params = keypair.P521
	case s.SM3withSM2:
		pkAlgorithm = keypair.PK_SM2
		params = keypair.SM2P256V1
	case s.SHA512withEDDSA:
		pkAlgorithm = keypair.PK_EDDSA
		params = keypair.ED25519
	}

	pri, pub, _ := keypair.GenerateKeyPair(pkAlgorithm, params)
	address := types.AddressFromPubKey(pub)
	return &Account{
		PrivateKey: pri,
		PublicKey:  pub,
		Address:    address,
		SigScheme:  scheme,
	}

}
