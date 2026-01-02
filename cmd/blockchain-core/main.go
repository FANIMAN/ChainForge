package main

import (
	"fmt"
	"time"

	"github.com/FANIMAN/chainforge/internal/blockchain/domain"
)

func main() {
	start := time.Now()

	// Initialize blockchain with genesis block
	bc := domain.NewBlockchain()

	// Add blocks
	bc.AddBlock("First block after genesis")
	bc.AddBlock("Second block after genesis")

	// Print blocks
	for i, block := range bc.Blocks {
		fmt.Printf("Block %d\n", i)
		fmt.Printf("  Data: %s\n", block.Data)
		fmt.Printf("  Hash: %s\n", block.HashHex())
		fmt.Printf("  Nonce: %d\n", block.Nonce)
		fmt.Printf("  Prev: %x\n\n", block.PrevBlockHash)
	}

	fmt.Printf("Mining took: %s\n", time.Since(start))
}
