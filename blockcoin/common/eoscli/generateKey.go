package eoscli

import (
	"crypto"
	"fmt"

	"github.com/eoscanada/eos-go/ecc"
)

func NewKeyPair() {
	priKey, _ := ecc.NewRandomPrivateKey()
	pubKey := priKey.PublicKey()
	fmt.Printf("Private Key: %s\nPublic  Key: %s\n", priKey, pubKey)
}

//can return keys
func NewKeysReturn() (*ecc.PrivateKey, crypto.PublicKey, error) {
	priKey, err := ecc.NewRandomPrivateKey()
	pubKey := priKey.PublicKey()
	return priKey, pubKey, err
}
