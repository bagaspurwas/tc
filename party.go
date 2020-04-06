package main

import (
	"crypto/rsa"
	"crypto/rand"
	"crypto/x509"
)

func GenerateKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	rng := rand.Reader
	PrivKey, _ := rsa.GenerateKey(rng, 4096)
	PubKey := &PrivKey.PublicKey
	return PrivKey, PubKey
}

func GenerateParty(Name string, Address string) (*PartySafe, *Party) {
	PrivKey, PubKey := GenerateKeyPair()
	PubKeyByte, err := x509.MarshalPKIXPublicKey(PubKey)
	if err != nil {
		panic(err)
	}
	PS := PartySafe{Name, Address, PubKey, PrivKey}
	Party := Party{Name, Address, []byte(PubKeyByte)}
	return &PS, &Party
}
