package common

import (
	"crypto/sha256"
	"errors"
	"io"
	"math/big"

	base58 "github.com/itchyny/base58-go"
	"golang.org/x/crypto/ripemd160"
)

const ADDR_LEN = 20

type Address [ADDR_LEN]byte

var ADDRESS_EMPTY = Address{}

func (self *Address) Serialize(w io.Writer) error {
	_, err := w.Write(self[:])
	return err
}
func (self *Address) Deserialize(r io.Reader) error {
	_, err := io.ReadFull(r, self[:])
	if err != nil {
		return errors.New("deserialize address error")
	}
	return nil
}
func (f *Address) ToBase58() string {
	data := append([]byte{23}, f[:]...)
	temp := sha256.Sum256(data)
	temps := sha256.Sum256(temp[:])
	data = append(data, temps[0:4]...)
	bi := new(big.Int).SetBytes(data).String()
	encoded, _ := base58.BitcoinEncoding.Encode([]byte(bi))
	return string(encoded)
}
func AddressParseFromBytes(f []byte) (Address, error) {
	if len(f) != ADDR_LEN {
		return ADDRESS_EMPTY, errors.New("[Common]:AddressParseFromBytes err,len!=20")
	}
	var addr Address
	copy(addr[:], f)
	return addr, nil
}
func AddressFromHexString(s string) (Address, error) {
	hx, err := HexToBytes(s)
	if err != nil {
		return ADDRESS_EMPTY, err
	}
	return AddressParseFromBytes(ToArrayReverse(hx))
}
func AddressFromBase58(encoded string) (Address, error) {
	if encoded == "" {
		return ADDRESS_EMPTY, errors.New("invalid address")
	}
	decoded, err := base58.BitcoinEncoding.Decode([]byte(encoded))
	if err != nil {
		return ADDRESS_EMPTY, err
	}
	x, ok := new(big.Int).SetString(string(decoded), 10)
	if !ok {
		return ADDRESS_EMPTY, errors.New("invalid address")
	}
	buf := x.Bytes()
	if len(buf) != 1+ADDR_LEN+4 || buf[0] != byte(23) {
		return ADDRESS_EMPTY, errors.New("wrong encoded address")
	}

	ph, err := AddressParseFromBytes(buf[1:21])
	if err != nil {
		return ADDRESS_EMPTY, err
	}

	addr := ph.ToBase58()

	if addr != encoded {
		return ADDRESS_EMPTY, errors.New("[AddressFromBase58]: decode encoded verify failed.")
	}

	return ph, nil

}
func AddressFromVmCode(code []byte) Address {
	var addr Address
	temp := sha256.Sum256(code)
	md := ripemd160.New()
	md.Write(temp[:])
	md.Sum(addr[:0])
	return addr
}
