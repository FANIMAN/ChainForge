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

// AddBlock appends a mined block
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// NewGenesisBlock creates the first block in the chain
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
