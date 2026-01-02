package domain

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int64
}

// NewBlock creates a new block and calculates its hash
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
		Nonce:         0,
	}

	block.Hash = block.calculateHash()
	return block
}

// calculateHash generates the SHA-256 hash of the block
func (b *Block) calculateHash() []byte {
	headers := bytes.Join(
		[][]byte{
			b.PrevBlockHash,
			b.Data,
			[]byte(strconv.FormatInt(b.Timestamp, 10)),
			[]byte(strconv.FormatInt(b.Nonce, 10)),
		},
		[]byte{},
	)

	hash := sha256.Sum256(headers)
	return hash[:]
}

// HashHex returns the block hash as a hex string (useful for logs & APIs)
func (b *Block) HashHex() string {
	return hex.EncodeToString(b.Hash)
}
