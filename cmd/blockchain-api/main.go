package main

import (
	"github.com/FANIMAN/chainforge/internal/blockchain/domain"
	"github.com/FANIMAN/chainforge/internal/blockchain/handler"
	"github.com/FANIMAN/chainforge/internal/blockchain/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	store := storage.NewBadgerStore()
	defer store.Close()

	bc := domain.NewBlockchain(store)

	walletHandler := handler.NewWalletHandler()
	blockchainHandler := handler.NewBlockchainHandler(bc)

	r := gin.Default()

	// Wallet routes
	r.POST("/wallet/create", walletHandler.CreateWallet)
	r.GET("/wallet/balance/:address", func(c *gin.Context) {
		walletHandler.GetBalance(c, bc)
	})

	// Blockchain routes
	r.GET("/blockchain", blockchainHandler.GetBlockchain)
	r.POST("/transaction/send", func(c *gin.Context) {
		blockchainHandler.SendTransaction(c, walletHandler.Wallets)
	})

	r.Run(":8080")
}
