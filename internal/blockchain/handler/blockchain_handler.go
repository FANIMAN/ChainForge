package handler

import (
	"net/http"
	"strconv"

	"github.com/FANIMAN/chainforge/internal/blockchain/domain"
	"github.com/gin-gonic/gin"
)

// BlockchainHandler handles blockchain routes
type BlockchainHandler struct {
	BC *domain.Blockchain
}

// NewBlockchainHandler initializes handler
func NewBlockchainHandler(bc *domain.Blockchain) *BlockchainHandler {
	return &BlockchainHandler{BC: bc}
}

// GetBlockchain handles GET /blockchain
func (h *BlockchainHandler) GetBlockchain(c *gin.Context) {
	c.JSON(http.StatusOK, h.BC.Blocks)
}

// GetMempool handles GET /mempool
func (h *BlockchainHandler) GetMempool(c *gin.Context) {
	c.JSON(http.StatusOK, h.BC.GetPendingTxs())
}

// SendTransaction handles POST /transaction/send
func (h *BlockchainHandler) SendTransaction(c *gin.Context, wallets map[string]*domain.Wallet) {
	var req struct {
		From   string `json:"from"`
		To     string `json:"to"`
		Amount int    `json:"amount"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sender, ok := wallets[req.From]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sender wallet not found"})
		return
	}

	tx := domain.NewTransaction(req.From, req.To, req.Amount)
	if err := tx.Sign(sender.PrivateKey); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to sign transaction"})
		return
	}

	// Add to mempool
	h.BC.AddToMempool(tx)

	c.JSON(http.StatusOK, gin.H{
		"message": "transaction added to mempool",
	})
}

// MineBlock handles POST /blockchain/mine
func (h *BlockchainHandler) MineBlock(c *gin.Context, wallets map[string]*domain.Wallet) {
	miner := c.Query("miner")
	rewardStr := c.Query("reward")

	if miner == "" || wallets[miner] == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid miner address"})
		return
	}

	reward := 50 // default reward
	if rewardStr != "" {
		if r, err := strconv.Atoi(rewardStr); err == nil {
			reward = r
		}
	}

	h.BC.MinePendingTxs(miner, wallets, reward)

	c.JSON(http.StatusOK, gin.H{
		"message": "pending transactions mined",
		"miner":   miner,
		"reward":  reward,
	})
}
