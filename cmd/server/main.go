package main

import (
	"net/http"

	"github.com/adiletelf/payment-system-go/pkg/api"
	"github.com/gin-gonic/gin"
)


func handleRequests(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	r.GET("/transaction", api.ReturnTransaction)
}

func main() {
	// db, err := api.SetupDB()
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }

	r := gin.Default()
	handleRequests(r)
	r.Run(":8080")
}
