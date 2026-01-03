package main

import (
	"log"

	"github.com/FANIMAN/chainforge/internal/blockchain/domain"
	"github.com/FANIMAN/chainforge/internal/blockchain/handler"
	"github.com/FANIMAN/chainforge/internal/blockchain/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	// ----------------------------
	// Initialize persistent storage
	// ----------------------------
	store := storage.NewBadgerStore()
	defer store.Close()

	// ----------------------------
	// Initialize blockchain
	// ----------------------------
	bc := domain.NewBlockchain(store)

	// ----------------------------
	// Initialize Gin & Handlers
	// ----------------------------
	r := gin.Default()
	blockHandler := handler.NewBlockchainHandler(bc)
	walletHandler := handler.NewWalletHandler()

	// ----------------------------
	// Wallet routes
	// ----------------------------
	r.POST("/wallet/create", walletHandler.CreateWallet)
	r.GET("/wallet/balance/:address", func(c *gin.Context) {
		walletHandler.GetBalance(c, bc)
	})

	// ----------------------------
	// Blockchain routes
	// ----------------------------
	r.GET("/blockchain", blockHandler.GetBlockchain)
	r.GET("/mempool", blockHandler.GetMempool)
	r.POST("/transaction/send", func(c *gin.Context) {
		blockHandler.SendTransaction(c, walletHandler.Wallets)
	})
	r.POST("/blockchain/mine", func(c *gin.Context) {
		blockHandler.MineBlock(c, walletHandler.Wallets)
	})

	log.Println("🚀 Blockchain API running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
