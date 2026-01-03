package main

import (
	"log"
	"os"

	authHandler "github.com/FANIMAN/chainforge/internal/auth/handler"
	authRepo "github.com/FANIMAN/chainforge/internal/auth/repository"
	authService "github.com/FANIMAN/chainforge/internal/auth/service"

	blockDomain "github.com/FANIMAN/chainforge/internal/blockchain/domain"
	blockHandler "github.com/FANIMAN/chainforge/internal/blockchain/handler"
	blockStorage "github.com/FANIMAN/chainforge/internal/blockchain/storage"

	"github.com/FANIMAN/chainforge/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// ----------------------------
	// Initialize blockchain storage
	// ----------------------------
	store := blockStorage.NewBadgerStore()
	defer store.Close()

	// ----------------------------
	// Initialize blockchain
	// ----------------------------
	bc := blockDomain.NewBlockchain(store)

	// ----------------------------
	// Initialize Auth service
	// ----------------------------
	userRepo := authRepo.NewUserRepo()
	authSvc := authService.NewAuthService(userRepo)
	authH := authHandler.NewAuthHandler(authSvc)

	// ----------------------------
	// Initialize Gin & Handlers
	// ----------------------------
	r := gin.Default()
	bcHandler := blockHandler.NewBlockchainHandler(bc)
	walletHandler := blockHandler.NewWalletHandler()

	// ----------------------------
	// Public routes
	// ----------------------------
	r.POST("/auth/register", authH.Register)
	r.POST("/auth/login", authH.Login)

	r.POST("/wallet/create", walletHandler.CreateWallet)
	r.GET("/wallet/balance/:address", func(c *gin.Context) {
		walletHandler.GetBalance(c, bc)
	})

	r.GET("/blockchain", bcHandler.GetBlockchain)
	r.GET("/mempool", bcHandler.GetMempool)

	// ----------------------------
	// Protected routes (JWT)
	// ----------------------------
	authGroup := r.Group("/", middleware.AuthMiddleware())
	authGroup.POST("/transaction/send", func(c *gin.Context) {
		bcHandler.SendTransaction(c, walletHandler.Wallets)
	})
	authGroup.POST("/blockchain/mine", func(c *gin.Context) {
		bcHandler.MineBlock(c, walletHandler.Wallets)
	})

	log.Println("🚀 Blockchain API running on :" + port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
