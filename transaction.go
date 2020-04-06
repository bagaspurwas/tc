package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"crypto/rand"
	"bytes"
)

func GenerateTXHeader(TX *Transaction) []byte {
	TXData := TX.Data

	IssuerPK := TXData.Issuer.PubKey
	RecvPK := TXData.Receiver.PubKey

    data := bytes.Join(
            [][]byte{
            	[]byte(IssuerPK),
            	[]byte(RecvPK),
            	[]byte(TXData.ServiceType),
            	IntToHex(int64(TXData.Value)),
            	[]byte(TXData.Info),
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

func (TX *Transaction) Sign(PrivateKey *rsa.PrivateKey) {
	rng := rand.Reader
	signature, err := rsa.SignPKCS1v15(rng, PrivateKey, crypto.SHA256, TX.Hash[:])
	if err != nil {
	    return
	}
	TX.Signature = signature
}

func (TX *Transaction) VerifySignature() {
	PK, err := x509.ParsePKIXPublicKey(TX.Data.Issuer.PubKey)
	PubKey := PK.(*rsa.PublicKey)
	if err != nil {
	    return
	}
	err = rsa.VerifyPKCS1v15(PubKey, crypto.SHA256, TX.Hash[:], TX.Signature)
	if err != nil {
	    return
	}
}

func NewTransactionData(Issuer *Party, Receiver *Party, ServiceType string,
						Value int32, Info string) *TransactionData {
	data := TransactionData{*Issuer, *Receiver, ServiceType, Value, Info}
	return &data
}

func (Safe *PartySafe) NewTransaction(data *TransactionData,
									  TXChainID []byte) *Transaction {
	TX := Transaction{}
	TX.Data = *data
	TX.TXChainID = TXChainID
	TX.GenerateTXHash()
	TX.Sign(Safe.PrivKey)
	return &TX
}
