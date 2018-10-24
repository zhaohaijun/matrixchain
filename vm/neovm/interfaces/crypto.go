package interfaces

type Crypto interface {
	Hash160(message []byte) []byte
	Hash256(message []byte) []byte
	VerffySignature(message []byte, signature []byte, pubkey []byte) (bool, error)
}
