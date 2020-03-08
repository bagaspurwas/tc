package models

type Block struct {
    Timestamp       int64           `bson:"timestamp" json:"timestamp"`
    Data            []byte          `bson:"data" json:"data"`
    Transactions	[]*Transaction   `bson:"data" json:"data"`
    PrevBlockHash   []byte          `bson:"prevblockhash" json:"prevblockhash"`
    Hash            []byte          `bson:"hash" json:"hash"`
    Nonce           int64           `bson:"nonce" json:"nonce"`
}
