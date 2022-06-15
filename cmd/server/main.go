package main

import (
	"log"

	"github.com/adiletelf/payment-system-go/pkg/api"
	"github.com/gin-gonic/gin"
)

func main() {
	err := api.SetupDB()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.GET("/transactions", api.GetAllTransactions)
	r.POST("/transaction", api.CreateTransaction)

	r.Run(":8080")
}
