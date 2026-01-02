package domain

type Blockchain struct {
	Blocks []*Block
}

// NewBlockchain initializes blockchain with genesis block
func NewBlockchain() *Blockchain {
	genesis := NewGenesisBlock()
	return &Blockchain{
		Blocks: []*Block{genesis},
	}
}

// AddBlock appends a mined block with transactions
func (bc *Blockchain) AddBlock(transactions []*Transaction) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(transactions, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// NewGenesisBlock creates the first block in the chain
func NewGenesisBlock() *Block {
	genesisTx := NewTransaction("network", "genesis", 0) // dummy genesis transaction
	return NewBlock([]*Transaction{genesisTx}, []byte{})
}
