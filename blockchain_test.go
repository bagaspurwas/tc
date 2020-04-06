package main

import (
	"testing"
)


func TestBlockchain(t *testing.T) {
	SafePartyA, PartyA := GenerateParty("Diamond Miner A", "TCA Tower A, Floor 12 Block C1")
	SafePartyB, PartyB := GenerateParty("CNC Precise Cutting Ltd.", "Dorsey Avenue 87")

	tx1data := NewTransactionData(PartyA, PartyB, "Cutting Service", 12000, "must be done by 20/5/2020")
	tx1 := SafePartyA.NewTransaction(tx1data, nil)
	tx1.VerifySignature()

	tx2data := NewTransactionData(PartyB, PartyA, "Returning cut diamond", 0, "")

	transactionA := SafePartyA.NewTransaction(tx1data, nil)
	transactionB := SafePartyB.NewTransaction(tx2data, nil)
	
	TX := []*Transaction {transactionA, transactionB} 

	bc := InitBlockchain()
	bc.AddBlock(TX)
	bc.Print()
	if len(bc.blocks) != 2 {
		t.Errorf("blockchain length must be 2")	
	}
	for _, b := range bc.blocks[1:] {
		if !Validate(b) {
			t.Errorf("Block hash mismatch")	
		}
	}
}
