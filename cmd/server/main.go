package main

import (
	"net/http"
	"time"

	"github.com/adiletelf/payment-system-go/pkg/models"
	"github.com/gin-gonic/gin"
)

func returnTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"transaction": models.Transaction{
			ID:        1,
			UserID:    1,
			Email:     "test@gmail.com",
			Amount:    100,
			Currency:  models.Ruble,
			Status:    models.Created,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	})
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	r.GET("/transaction", returnTransaction)

	r.Run(":8080")
}
