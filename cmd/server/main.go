package main

import (
	"log"

	"github.com/adiletelf/payment-system-go/internal/handlers"
	"github.com/adiletelf/payment-system-go/internal/repositories"
	"github.com/adiletelf/payment-system-go/internal/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := utils.SetupDB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	tr := repositories.NewTransactionRepo(db)
	h := handlers.NewBaseHandler(tr)

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	r.GET("/transactions", h.GetAllTransactions)

	r.GET("/transaction/:id", h.GetTransactionStatus)
	r.PUT("/transaction/:id", h.UpdateTransactionStatus)
	r.POST("/transaction", h.CreateTransaction)

	r.GET("/transaction/cancel/:id", h.CancelTransaction)

	r.Run(":8080")
}
