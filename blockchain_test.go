package main

import (
	"testing"
	"fmt"
)

func TestBlockchain(t *testing.T) {
	transactionA := `{
		"src": "Alice",
		"dst": "Bob",
		"amount": "1"
	}`
	transactionB := `{
		"src": "Bob",
		"dst": "Alice",
		"amount": "2"
	}`

	bc := InitBlockchain()
	bc.AddBlock(transactionA)
	bc.AddBlock(transactionB)
	bc.Print()
	if len(bc.blocks) != 3 {
		t.Errorf("blockchain length must be 3")	
	}
	for _, b := range bc.blocks {
		if !Validate(b) {
			t.Errorf("Block hash mismatch")	
		}
	}
}
