package accounts

import "crypto/ecdsa"

type Account struct {
	Address           string
	PrivateKey        *ecdsa.PrivateKey
	PrivateKeyEncoded string
	PublicKeyEncoded  string
}
