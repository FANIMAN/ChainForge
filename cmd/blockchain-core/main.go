package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"time"

	"github.com/FANIMAN/chainforge/internal/blockchain/domain"
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

	// Map wallets for verification
	wallets := map[string]*domain.Wallet{
		w1.Address(): w1,
		w2.Address(): w2,
	}

	// ----------------------------
	// Initialize blockchain
	// ----------------------------
	bc := domain.NewBlockchain()

	// ----------------------------
	// Create transactions
	// ----------------------------
	tx1 := domain.NewTransaction(w1.Address(), w2.Address(), 10)
	err := tx1.Sign(w1.PrivateKey)
	if err != nil {
		fmt.Println("Failed to sign tx1:", err)
		return
	}

	tx2 := domain.NewTransaction(w2.Address(), w1.Address(), 5)
	err = tx2.Sign(w2.PrivateKey)
	if err != nil {
		fmt.Println("Failed to sign tx2:", err)
		return
	}

	// Add block with verification
	bc.AddBlockWithVerification([]*domain.Transaction{tx1, tx2}, wallets)

	// ----------------------------
	// Print blockchain
	// ----------------------------
	for i, block := range bc.Blocks {
		fmt.Printf("\nBlock %d\n", i)
		fmt.Printf("  Hash: %s\n", block.HashHex())
		fmt.Printf("  Nonce: %d\n", block.Nonce)
		fmt.Printf("  Transactions:\n")
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

	fmt.Printf("\nMining + transaction verification took: %s\n", time.Since(start))
}

// ----------------------------
// Optional helper: generate new ECDSA key (used by Wallet struct)
// ----------------------------
func newKey() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}
