package main

import (
	"crypto/rsa"
)

type Block struct {
    Timestamp       int64           `bson:"timestamp" json:"timestamp"`
    Transactions	[]*Transaction  `bson:"transactions" json:"transactions"`
    PrevBlockHash   []byte          `bson:"prevblockhash" json:"prevblockhash"`
    Hash            []byte          `bson:"hash" json:"hash"`
    Nonce           int64           `bson:"nonce" json:"nonce"`
}

type PartySafe struct {
    Name        	string
    Address     	string    
	PubKey			*rsa.PublicKey
	PrivKey			*rsa.PrivateKey
}

type Party struct {
    Name            string    		`bson:"name" json:"name"` 
    Address         string    		`bson:"address" json:"address"`
    PubKey          *rsa.PublicKey  `bson:"publickey" json:"publickey"`
}

type Transaction struct {
    Signature       []byte          `bson:"signature" json:"signature"`
    Hash            []byte          `bson:"hash" json:"hash"`
    Data            TransactionData `bson:"data" json:"data"`
    TXChainID       []byte          `bson:"txchainID" json:"txchainID"`
}

type TransactionData struct {
    Issuer          Party           `bson:"issuer" json:"issuer"`
    Receiver        Party           `bson:"receiver" json:"receiver"`
    ServiceType     string          `bson:"servicetype" json:"servicetype"`
    Value           int32           `bson:"value" json:"value"`
    Info            string          `bson:"info" json:"info"`
}