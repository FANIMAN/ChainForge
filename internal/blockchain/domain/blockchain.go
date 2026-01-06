package domain

import "errors"

type Blockchain struct {
	Blocks  []*Block
	Storage Storage
	Mempool []*Transaction
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


// Add transaction to mempool
// func (bc *Blockchain) AddToMempool(tx *Transaction) {
//     bc.Mempool = append(bc.Mempool, tx)
// }


// AddToMempool adds a transaction to mempool after full validation
func (bc *Blockchain) AddToMempool(tx *Transaction, wallets map[string]*Wallet) error {
	if tx == nil {
		return errors.New("transaction is nil")
	}

	// Validate transaction structure
	if !tx.IsValid() {
		return errors.New("invalid transaction structure")
	}

	// Check sender exists
	senderWallet, ok := wallets[tx.From]
	if !ok {
		return errors.New("sender wallet not found")
	}

	// Check balance
	if bc.GetBalance(tx.From) < tx.Amount {
		return errors.New("insufficient balance")
	}

	// Verify signature
	if !tx.Verify(senderWallet.PublicKey) {
		return errors.New("invalid transaction signature")
	}

	// Passed all checks, add to mempool
	bc.Mempool = append(bc.Mempool, tx)
	return nil
}





// Get pending transactions
func (bc *Blockchain) GetPendingTxs() []*Transaction {
    return bc.Mempool
}

// Mine pending transactions (with coinbase reward)
func (bc *Blockchain) MinePendingTxs(minerAddress string, wallets map[string]*Wallet, reward int) {
    if len(bc.Mempool) == 0 {
        return
    }

    // Coinbase tx
    coinbase := NewTransaction("network", minerAddress, reward)
    
    // Combine coinbase + pending txs
    txs := append([]*Transaction{coinbase}, bc.Mempool...)

    prevBlock := bc.Blocks[len(bc.Blocks)-1]
    newBlock := NewBlock(txs, prevBlock.Hash)

    bc.Blocks = append(bc.Blocks, newBlock)
    bc.Storage.SaveBlock(newBlock)

    // Clear mempool
    bc.Mempool = []*Transaction{}
}