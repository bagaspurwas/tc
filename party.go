package main

import (
	"crypto/rsa"
	"crypto/rand"
)

func GenerateKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	rng := rand.Reader
	PrivKey, _ := rsa.GenerateKey(rng, 4096)
	PubKey := &PrivKey.PublicKey
	return PrivKey, PubKey
}

func GenerateParty(Name string, Address string) (*PartySafe, *Party) {
	PrivKey, PubKey := GenerateKeyPair()
	PS := PartySafe{Name, Address, PubKey, PrivKey}
	Party := Party{Name, Address, PubKey}
	return &PS, &Party
}
