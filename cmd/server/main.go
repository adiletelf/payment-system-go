package main

import (
	"log"

	"github.com/adiletelf/payment-system-go/internal/handler"
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
	r.GET("/transactions", h.GetAllTransactions)
	r.POST("/transaction", h.CreateTransaction)
	r.GET("/transaction/:id", h.GetTransactionStatus)
	r.PUT("/transaction/:id", h.UpdateTransactionStatus)
	r.GET("/transaction/cancel/:id", h.CancelTransaction)

	r.POST("/register", h.Register)
	r.POST("/login", h.Login)

	r.Run(":8080")
}
