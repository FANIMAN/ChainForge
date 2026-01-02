package handler

import (
	"net/http"

	"github.com/FANIMAN/chainforge/internal/blockchain/domain"
	"github.com/gin-gonic/gin"
)

// WalletHandler holds wallets map
type WalletHandler struct {
	Wallets map[string]*domain.Wallet
}

// NewWalletHandler creates WalletHandler
func NewWalletHandler() *WalletHandler {
	return &WalletHandler{
		Wallets: make(map[string]*domain.Wallet),
	}
}

// CreateWallet handles POST /wallet/create
func (h *WalletHandler) CreateWallet(c *gin.Context) {
	w := domain.NewWallet()
	h.Wallets[w.Address()] = w
	c.JSON(http.StatusOK, gin.H{
		"address": w.Address(),
	})
}

// GetBalance handles GET /wallet/balance/:address
func (h *WalletHandler) GetBalance(c *gin.Context, bc *domain.Blockchain) {
	address := c.Param("address")
	if _, ok := h.Wallets[address]; !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "wallet not found"})
		return
	}
	balance := bc.GetBalance(address)
	c.JSON(http.StatusOK, gin.H{"balance": balance})
}
