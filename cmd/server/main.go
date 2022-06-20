package main

import (
	"log"
	"os"

	"github.com/adiletelf/payment-system-go/internal/handler"
	"github.com/adiletelf/payment-system-go/internal/middleware"
	"github.com/adiletelf/payment-system-go/internal/repository"
	"github.com/adiletelf/payment-system-go/internal/util"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := util.SetupDB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	tr := repository.NewTransactionRepo(db)
	ar := repository.NewAdminRepo(db)
	h := handler.NewBaseHandler(tr, ar)

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	public := r.Group("/api/admin")
	public.POST("/register", h.Register)
	public.POST("/login", h.Login)

	protected := r.Group("/api")

	useAuthentication := os.Getenv("AUTHENTICATION_ENABLED")
	if useAuthentication == "true" {
		protected.Use(middleware.JwtAuthMiddleware())
	}

	protected.GET("/admin", h.CurrentAdmin)
	protected.GET("/transactions", h.GetAllTransactions)
	protected.POST("/transaction", h.CreateTransaction)
	protected.GET("/transaction/:id", h.GetTransactionStatus)
	protected.PUT("/transaction/:id", h.UpdateTransactionStatus)
	protected.GET("/transaction/cancel/:id", h.CancelTransaction)

	r.Run(":8080")
}
