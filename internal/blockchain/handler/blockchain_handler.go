package handler

import (
	"net/http"

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

	h.BC.AddBlockWithVerification([]*domain.Transaction{tx}, wallets)
	c.JSON(http.StatusOK, gin.H{"message": "transaction added"})
}
