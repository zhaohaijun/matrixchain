package util

import (
	"crypto/sha256"
	"errors"
	"io"

	"github.com/zhaohaijun/blockchain-crypto/keypair"
	s "github.com/zhaohaijun/blockchain-crypto/signature"
	"github.com/zhaohaijun/matrixchain/common/log"
	ontErrors "github.com/zhaohaijun/matrixchain/errors"
	"golang.org/x/crypto/ripemd160"
)

type ECDsaCrypto struct {
}

func (c *ECDsaCrypto) Hash160(message []byte) []byte {
	temp := sha256.Sum256(message)
	md := ripemd160.New()
	io.WriteString(md, string(temp[:]))
	hash := md.Sum(nil)
	return hash
}

func (c *ECDsaCrypto) Hash256(message []byte) []byte {
	temp := sha256.Sum256(message)
	f := sha256.Sum256(temp[:])
	return f[:]
}

func (c *ECDsaCrypto) VerifySignature(message []byte, signature []byte, pubkey []byte) (bool, error) {

	log.Debugf("message: %x", message)
	log.Debugf("signature: %x", signature)
	log.Debugf("pubkey: %x", pubkey)

	pk, err := keypair.DeserializePublicKey(pubkey)
	if err != nil {
		return false, ontErrors.NewDetailErr(errors.New("[ECDsaCrypto], deserializing public key failed."), ontErrors.ErrNoCode, "")
	}

	sig, err := s.Deserialize(signature)
	ok := s.Verify(pk, message, sig)
	if !ok {
		return false, ontErrors.NewDetailErr(errors.New("[ECDsaCrypto], VerifySignature failed."), ontErrors.ErrNoCode, "")
	}

	return true, nil
}
