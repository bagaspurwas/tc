package main

import (
	"./models"
	"crypto/rsa"
	"crypto/rand"
)

type PartySafe models.PartySafe

func GenerateKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	rng := rand.Reader
	PrivKey := rsa.GenerateKey(rng, 4096)
	PubKey := rsa.Public()
	return PrivKey, PubKey
}

func GeneratePartySafe(Name string, Address string) *PartySafe {
	PrivKey, PubKey := GenerateKeyPair()
	PS := PartySafe{Name, Address, PrivKey, PubKey}
	return PS
}

func GenerateParty(Name string, Address string) *Party {
	
}