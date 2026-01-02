package main

import (
	"fmt"

	"github.com/FANIMAN/chainforge/internal/blockchain/domain"
)

func main() {
	bc := domain.NewBlockchain()

	bc.AddBlock("First block after genesis")
	bc.AddBlock("Second block after genesis")

	for i, block := range bc.Blocks {
		fmt.Printf("Block %d\n", i)
		fmt.Printf("  Data: %s\n", block.Data)
		fmt.Printf("  Hash: %s\n", block.HashHex())
		fmt.Printf("  Prev: %x\n\n", block.PrevBlockHash)
	}
}
