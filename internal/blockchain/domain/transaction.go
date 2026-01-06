package domain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)

type Transaction struct {
	ID        string
	From      string
	To        string
	Amount    int
	Signature []byte
}

// NewTransaction creates a transaction and calculates ID
func NewTransaction(from, to string, amount int) *Transaction {
	tx := &Transaction{
		From:   from,
		To:     to,
		Amount: amount,
	}
	tx.ID = tx.calculateID()
	return tx
}

// calculateID generates SHA256 hash of transaction contents
func (tx *Transaction) calculateID() string {
	data := fmt.Sprintf("%s%s%d", tx.From, tx.To, tx.Amount)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// Sign signs the transaction using sender's private key
func (tx *Transaction) Sign(privKey *ecdsa.PrivateKey) error {
	hash := sha256.Sum256([]byte(tx.ID))
	r, s, err := ecdsa.Sign(randReader{}, privKey, hash[:])
	if err != nil {
		return err
	}
	signature := append(r.Bytes(), s.Bytes()...)
	tx.Signature = signature
	return nil
}

// Verify verifies the transaction signature using sender public key
func (tx *Transaction) Verify(pubKey []byte) bool {
	if tx.Signature == nil {
		return false
	}
	hash := sha256.Sum256([]byte(tx.ID))

	keyLen := len(pubKey) / 2
	x := new(big.Int).SetBytes(pubKey[:keyLen])
	y := new(big.Int).SetBytes(pubKey[keyLen:])
	publicKey := &ecdsa.PublicKey{Curve: elliptic.P256(), X: x, Y: y}

	r := new(big.Int).SetBytes(tx.Signature[:len(tx.Signature)/2])
	s := new(big.Int).SetBytes(tx.Signature[len(tx.Signature)/2:])

	return ecdsa.Verify(publicKey, hash[:], r, s)
}

// For demonstration purposes (we’ll fix randReader below)
type randReader struct{}

func (r randReader) Read(b []byte) (int, error) {
	return rand.Read(b)
}

// String prints transaction info
func (tx *Transaction) String() string {
	return fmt.Sprintf("TX %s: %s -> %s | %d coins", tx.ID, tx.From, tx.To, tx.Amount)
}


func (tx *Transaction) IsValid() bool {
	if tx.From == "" || tx.To == "" {
		return false
	}
	if tx.Amount <= 0 {
		return false
	}
	if tx.Signature == nil {
		return false
	}
	return true
}
