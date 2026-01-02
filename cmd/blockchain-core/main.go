package main

import (
	"fmt"
	"time"

	"github.com/FANIMAN/chainforge/internal/blockchain/domain"
)

func main() {
	start := time.Now()

	// Create wallets
	w1 := domain.NewWallet()
	w2 := domain.NewWallet()

	fmt.Println("Wallet 1:", w1.Address())
	fmt.Println("Wallet 2:", w2.Address())

	// Initialize blockchain
	bc := domain.NewBlockchain()

	// Create transactions
	tx1 := domain.NewTransaction(w1.Address(), w2.Address(), 10)
	tx2 := domain.NewTransaction(w2.Address(), w1.Address(), 5)

	// Add a block with transactions
	bc.AddBlock([]*domain.Transaction{tx1, tx2})

	// Print blockchain
	for i, block := range bc.Blocks {
		fmt.Printf("\nBlock %d\n", i)
		fmt.Printf("  Hash: %s\n", block.HashHex())
		fmt.Printf("  Nonce: %d\n", block.Nonce)
		fmt.Printf("  Transactions:\n")
		for _, tx := range block.Transactions {
			fmt.Printf("    %s\n", tx.String())
		}
	}

	fmt.Printf("\nMining + transactions took: %s\n", time.Since(start))
}
