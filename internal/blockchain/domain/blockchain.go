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

// AddBlock appends a block with given transactions (without verification)
func (bc *Blockchain) AddBlock(transactions []*Transaction) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(transactions, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// AddBlockWithVerification validates transactions and adds only valid ones
func (bc *Blockchain) AddBlockWithVerification(transactions []*Transaction, wallets map[string]*Wallet) {
	validTxs := []*Transaction{}
	for _, tx := range transactions {
		// Ensure wallet exists
		wallet, ok := wallets[tx.From]
		if !ok {
			continue
		}
		// Check balance
		if bc.GetBalance(tx.From) < tx.Amount {
			continue
		}
		// Verify signature
		if !tx.Verify(wallet.PublicKey) {
			continue
		}
		validTxs = append(validTxs, tx)
	}

	if len(validTxs) == 0 {
		return
	}

	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(validTxs, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// GetBalance scans all blocks to calculate wallet balance
func (bc *Blockchain) GetBalance(address string) int {
	balance := 0
	for _, block := range bc.Blocks {
		for _, tx := range block.Transactions {
			if tx.From == address {
				balance -= tx.Amount
			}
			if tx.To == address {
				balance += tx.Amount
			}
		}
	}
	return balance
}

// NewGenesisBlock creates the first block with a dummy transaction
func NewGenesisBlock() *Block {
	genesisTx := NewTransaction("network", "genesis", 0)
	return NewBlock([]*Transaction{genesisTx}, []byte{})
}
