package main

import (
	"fmt"
	"time"

	"github.com/FANIMAN/chainforge/internal/blockchain/domain"
	"github.com/FANIMAN/chainforge/internal/blockchain/storage"
)

func main() {
	start := time.Now()

	// ----------------------------
	// Create wallets
	// ----------------------------
	w1 := domain.NewWallet()
	w2 := domain.NewWallet()

	fmt.Println("Wallet 1:", w1.Address())
	fmt.Println("Wallet 2:", w2.Address())

	wallets := map[string]*domain.Wallet{
		w1.Address(): w1,
		w2.Address(): w2,
	}

	// ----------------------------
	// Initialize blockchain (persistent)
	// ----------------------------
	store := storage.NewBadgerStore()
	defer store.Close()

	bc := domain.NewBlockchain(store)

	// ----------------------------
	// Create transactions and add to mempool
	// ----------------------------
	tx1 := domain.NewTransaction(w1.Address(), w2.Address(), 10)
	_ = tx1.Sign(w1.PrivateKey)
	bc.AddToMempool(tx1)

	tx2 := domain.NewTransaction(w2.Address(), w1.Address(), 5)
	_ = tx2.Sign(w2.PrivateKey)
	bc.AddToMempool(tx2)

	// ----------------------------
	// Mine pending transactions with reward
	// ----------------------------
	bc.MinePendingTxs(w1.Address(), wallets, 50) // Wallet1 gets 50 coin reward

	// ----------------------------
	// Print blockchain
	// ----------------------------
	for i, block := range bc.Blocks {
		fmt.Printf("\nBlock %d\n", i)
		fmt.Printf("  Hash: %s\n", block.HashHex())
		fmt.Printf("  Nonce: %d\n", block.Nonce)
		fmt.Println("  Transactions:")
		for _, tx := range block.Transactions {
			fmt.Printf("    %s\n", tx.String())
		}
	}

	// ----------------------------
	// Print wallet balances
	// ----------------------------
	fmt.Printf("\nWallet balances:\n")
	fmt.Printf("  Wallet 1: %d\n", bc.GetBalance(w1.Address()))
	fmt.Printf("  Wallet 2: %d\n", bc.GetBalance(w2.Address()))

	fmt.Printf("\nExecution took: %s\n", time.Since(start))
}
