package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  []byte
	Address    []byte
}

// 生成钱包
func NewWallet() (*Wallet, string, error) {
	c := elliptic.P256()
	private, err := ecdsa.GenerateKey(c, rand.Reader)
	if err != nil {
		return nil, "", err
	}

	return &Wallet{PrivateKey: private, PublicKey: []byte{1}}, "addr", nil
}

//
