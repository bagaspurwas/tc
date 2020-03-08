package main

import (
	"./models"
	"crypto/rsa"
	"crypto/rand"
)

type Party models.Party
type Transaction models.Transaction
type TransactionData models.TransactionData

func GenerateTXHeader(TX *Transaction) []byte {
	TXData := TX.Data
    data := bytes.Join(
            [][]byte{
            	[]byte(TXData.Issuer.PubKey),
            	[]byte(TXData.Receiver.PubKey),
            	[]byte(TXData.ServiceType),
            	[]byte(TXData.Value),
            	[]byte(TXData.Info)
            },
            []byte{},       
    )
    return data
}

func (TX *Transaction) GenerateTXHash() {
    headers := GenerateTXHeader(TX)
    hash := sha256.Sum256(headers)
    TX.Hash = hash[:]
}

func (TX *Transaction) GenerateSignature() {
	rng := rand.Reader
	signature, err := SignPKCS1v15(rng, PrivateKey, crypto.SHA256, TX.Hash[:])
	if err != nil {
	    fmt.Fprintf(os.Stderr, "Error from signing: %s\n", err)
	    return
	}
	TX.Signature = signature
}

func (TX *Transaction) VerifySignature() {
	
}

func NewTransaction(data TransactionData) *Transaction {
	TX := Transaction{Data=data}
	TX.GenerateHash()
	TX.GenerateSignature()
	return TX
}