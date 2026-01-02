package domain

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)

const targetBits = 16 // difficulty (adjust for testing)

// ProofOfWork represents a mining task for a block
type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

// NewProof creates a PoW for a given block
func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, 256-targetBits)
	return &ProofOfWork{b, target}
}

// Run performs mining: finds a nonce so hash < target
func (pow *ProofOfWork) Run() ([]byte, int64) {
	var hashInt big.Int
	var hash [32]byte
	nonce := int64(0)

	for {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.Target) == -1 {
			break
		}
		nonce++
	}

	return hash[:], nonce
}

// prepareData prepares the block fields + nonce for hashing
func (pow *ProofOfWork) prepareData(nonce int64) []byte {
	block := pow.Block
	data := bytes.Join(
		[][]byte{
			block.PrevBlockHash,
			block.Data,
			[]byte(fmt.Sprintf("%d", block.Timestamp)),
			[]byte(fmt.Sprintf("%d", nonce)),
		},
		[]byte{},
	)
	return data
}

// Validate verifies the block satisfies PoW
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int
	data := pow.prepareData(pow.Block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	return hashInt.Cmp(pow.Target) == -1
}

// HashToString converts hash bytes to hex string
func HashToString(hash []byte) string {
	return hex.EncodeToString(hash)
}
