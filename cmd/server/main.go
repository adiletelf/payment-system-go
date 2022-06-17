package main

import (
	"log"

	"github.com/adiletelf/payment-system-go/pkg/api"
	"github.com/adiletelf/payment-system-go/pkg/repositories"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := api.SetupDB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	tr := repositories.NewTransactionRepo(db)
	h := api.NewBaseHandler(tr)

	r := gin.Default()
	r.GET("/transactions", h.GetAllTransactions)

	r.GET("/transaction/:id", h.GetTransactionStatus)
	r.PUT("/transaction/:id", h.UpdateTransactionStatus)
	r.POST("/transaction", h.CreateTransaction)

	r.GET("/transaction/cancel/:id", h.CancelTransaction)

	r.Run(":8080")
}
