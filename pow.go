package main

import (
	"bytes"
	"math"
    "encoding/hex"
)

var difficulty, _ = hex.DecodeString("000")

func Validate(block *Block) bool {
	if bytes.HasPrefix(block.Hash, difficulty) {
		return true
	} else {
		return false
	}
}

func Work(block *Block) *Block {
	var nonce int64 = 0

	for nonce < math.MaxInt64 {
		block.Nonce = nonce
		block.HashBlock()
		if Validate(block) {
			break
		} else {
			nonce++	
		}
	}

	return block
}