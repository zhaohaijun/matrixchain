package signature

import (
	"errors"

	"github.com/zhaohaijun/blockchain-crypto/keypair"
	s "github.com/zhaohaijun/blockchain-crypto/signature"
)

func Sign(signer Signer, data []byte) ([]byte, error) {
	signature, err := s.Sign(signer.Scheme(), signer.PrivKey(), data, nil)
	if err != nil {
		return nil, err
	}
	return s.Serialize(signature)
}
func Verify(pubkey keypair.PublicKey, data, signature []byte) error {
	sigObj, err := s.Deserialize(signature)
	if err != nil {
		return errors.New("invalid signature data:" + err.Error())

	}
	if !s.Verify(pubkey, data, sigObj) {
		return errors.New("signature verification failed")
	}
	return nil
}
func VerifyMultiSignature(data []byte, keys []keypair.PublicKey, m int, sigs [][]byte) error {
	n := len(keys)

	if len(sigs) < m {
		return errors.New("not enough signatures in multi-signature")
	}

	mask := make([]bool, n)
	for i := 0; i < m; i++ {
		valid := false

		sig, err := s.Deserialize(sigs[i])
		if err != nil {
			return errors.New("invalid signature data")
		}
		for j := 0; j < n; j++ {
			if mask[j] {
				continue
			}
			if s.Verify(keys[j], data, sig) {
				mask[j] = true
				valid = true
				break
			}
		}

		if valid == false {
			return errors.New("multi-signature verification failed")
		}
	}

	return nil
}
