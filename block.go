package main

import (
	"time"
	"crypto/sha256"
	"bytes"
    "encoding/binary"
    "log"
)

// IntToHex converts an int64 to a byte array
func IntToHex(num int64) []byte {
    buff := new(bytes.Buffer)
    err := binary.Write(buff, binary.BigEndian, num)
    if err != nil {
        log.Panic(err)
    }

    return buff.Bytes()
}

func GenerateBlockHeader(block *Block) []byte {
    data := bytes.Join(
            [][]byte{
                block.PrevBlockHash,
                IntToHex(block.Timestamp),
                IntToHex(block.Nonce),
            },
            []byte{},       
    )   
    return data
}

func (block *Block) HashBlock() {
    headers := GenerateBlockHeader(block)
    hash := sha256.Sum256(headers)
    block.Hash = hash[:]
}

func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
    block := &Block{time.Now().Unix(), transactions, prevBlockHash, []byte{}, 0}
    block.HashBlock()
    return block
}
