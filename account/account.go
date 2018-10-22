package account
type Account struct{
	PrivateKey keypair.PrivateKey
	PublicKey keypair.PublicKey
	Address common.Address
	SigScheme s.SignatureScheme
}