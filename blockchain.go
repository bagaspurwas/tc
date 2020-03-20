package main
import (
    "fmt"
    "encoding/json"
)

type Blockchain struct {
    blocks      []*Block
}

func InitGenesis() *Block {
	SafePartyA, PartyA := GenerateParty("Genesis", "Genesis")
	_ , PartyB := GenerateParty("Genesis2", "Genesis2")
	TX1data := NewTransactionData(PartyA, PartyB, "Genesis", 0, "")
	TX1 := SafePartyA.NewTransaction(TX1data, nil)
	var TXs = []*Transaction {TX1}
    return NewBlock(TXs, []byte{})
}

func InitBlockchain() *Blockchain {
    return &Blockchain{ []*Block{ InitGenesis() } }
}

func (bc *Blockchain) Print() {
	for _, block := range bc.blocks {
		block_JSON, err := json.Marshal(block)
		if err != nil {
			return
		}
		fmt.Printf("%s\n", block_JSON)
	}
}

func (bc *Blockchain) AddBlock(data []*Transaction) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	minedBlock := Work(newBlock)	
	bc.blocks = append(bc.blocks, minedBlock)
}
