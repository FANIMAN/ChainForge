package domain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Wallet represents a blockchain wallet with a key pair
type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  []byte
}

// NewWallet creates a new wallet
func NewWallet() *Wallet {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	pubKey := append(privKey.PublicKey.X.Bytes(), privKey.PublicKey.Y.Bytes()...)
	return &Wallet{
		PrivateKey: privKey,
		PublicKey:  pubKey,
	}
}

// Address returns a hashed representation of the public key
func (w *Wallet) Address() string {
	hash := sha256.Sum256(w.PublicKey)
	return hex.EncodeToString(hash[:])
}

// String prints wallet info (for testing)
func (w *Wallet) String() string {
	return fmt.Sprintf("Address: %s", w.Address())
}
