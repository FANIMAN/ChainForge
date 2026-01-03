package domain

type Blockchain struct {
	Blocks  []*Block
	Storage Storage
}

type Storage interface {
	SaveBlock(*Block)
	LoadBlockchain() []*Block
}

// NewBlockchain initializes blockchain with genesis block and persistent storage
func NewBlockchain(store Storage) *Blockchain {
	blocks := store.LoadBlockchain()

	if len(blocks) == 0 {
		genesis := NewGenesisBlock()
		store.SaveBlock(genesis)
		blocks = []*Block{genesis}
	}

	return &Blockchain{
		Blocks:  blocks,
		Storage: store,
	}
}

// AddBlock validates transactions, mines block, and persists it
func (bc *Blockchain) AddBlock(transactions []*Transaction, wallets map[string]*Wallet) {
	validTxs := []*Transaction{}

	for _, tx := range transactions {
		wallet, ok := wallets[tx.From]
		if !ok {
			continue
		}
		if bc.GetBalance(tx.From) < tx.Amount {
			continue
		}
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
	bc.Storage.SaveBlock(newBlock)
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
