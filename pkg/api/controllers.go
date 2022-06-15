package api

import (
	"net/http"

	"github.com/adiletelf/payment-system-go/pkg/models"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	r.GET("/transaction", returnTransaction)

	return r
}

func returnTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, models.NewTransaction(models.User{ID: 1, Email: "test@gmail.com"}, 99.99, models.Ruble))
}
