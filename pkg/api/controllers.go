package api

import (
	"net/http"

	"github.com/adiletelf/payment-system-go/pkg/models"
	"github.com/gin-gonic/gin"
)

func ReturnTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, models.NewTransaction(models.User{ID: 1, Email: "test@gmail.com"}, 99.99, models.Ruble))
}
