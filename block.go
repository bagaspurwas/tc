package main

import (
	"./models"
	"time"
	"crypto/sha256"
	"bytes"
    "encoding/binary"
    "log"
)

type Block models.Block

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
                block.Data,
                IntToHex(block.Timestamp),
                IntToHex(block.Nonce),
            },
            []byte{},       
    )   
    return data
}

func (b *Block) HashBlock() {
    headers := GenerateBlockHeader(b)
    hash := sha256.Sum256(headers)
    b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
    block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
    block.HashBlock()
    return block
}
