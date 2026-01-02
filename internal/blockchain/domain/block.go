package domain

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

// Block represents a single block in the blockchain
type Block struct {
	Timestamp     int64
	Transactions []*Transaction
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int64
}

// NewBlock creates a new block and mines it
func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Transactions:  transactions,
		PrevBlockHash: prevBlockHash,
		Nonce:         0,
	}

	// Mine the block
	pow := NewProof(block)
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce

	return block
}



// calculateHash generates SHA-256 hash (used internally, optional)
func (b *Block) calculateHash() []byte {
	var txData []byte
	for _, tx := range b.Transactions {
		txData = append(txData, []byte(tx.ID)...)
	}
	headers := bytes.Join(
		[][]byte{
			b.PrevBlockHash,
			txData,
			[]byte(strconv.FormatInt(b.Timestamp, 10)),
			[]byte(strconv.FormatInt(b.Nonce, 10)),
		},
		[]byte{},
	)
	hash := sha256.Sum256(headers)
	return hash[:]
}


// HashHex returns block hash as hex string for logging/API
func (b *Block) HashHex() string {
	return hex.EncodeToString(b.Hash)
}
