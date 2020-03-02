package main
import (
    "fmt"
)

type Blockchain struct {
    blocks      []*Block
}

func InitGenesis() *Block {
    data := "This is genesis block"
    return NewBlock(data,[]byte{})
}

func InitBlockchain() *Blockchain {
    return &Blockchain{ []*Block{ InitGenesis() } }
}

func (bc *Blockchain) Print() {
	for _, block := range bc.blocks {
		fmt.Printf("------------- block --------------\n")
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Nonce: %x\n", block.Nonce)
	}
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	minedBlock := Work(newBlock)	
	bc.blocks = append(bc.blocks, minedBlock)
}
