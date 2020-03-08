package models

type Party struct {
    Name            string    ``
    Address         string
    PubKey          byte[]
}

// Here's how transaction is going to be tracked
// Signature = sign(TX1.data)
// TX1.ID = hash(sign(TX1.data)); TX1.TXChainID = nill
// TX2.ID = hash(sign(TX2.data)); TX2.TXChainID = TX1.ID
// Transaction authenticity can be verified as follow:
// 1. Sign (TX1.Data) using issuer party Private Key
// 2. Verify sign(TX1.Data) using issuer Public Key

type Transaction struct {
    Signature       []byte    
    Hash            []byte
    Data            TransactionData
    TXChainID       []byte
}

type TransactionData struct {
    Issuer          Party
    Receiver        Party
    ServiceType     string
    Value           int32
    Info            string
}