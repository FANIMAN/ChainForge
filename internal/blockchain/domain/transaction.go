package domain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Transaction struct {
	ID     string
	From   string
	To     string
	Amount int
}

// NewTransaction creates a new transaction
func NewTransaction(from, to string, amount int) *Transaction {
	tx := &Transaction{
		From:   from,
		To:     to,
		Amount: amount,
	}
	tx.ID = tx.calculateID()
	return tx
}

// calculateID generates a simple hash of transaction contents
func (tx *Transaction) calculateID() string {
	data := fmt.Sprintf("%s%s%d", tx.From, tx.To, tx.Amount)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// String prints transaction info
func (tx *Transaction) String() string {
	return fmt.Sprintf("TX %s: %s -> %s | %d coins", tx.ID, tx.From, tx.To, tx.Amount)
}
