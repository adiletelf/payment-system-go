package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/adiletelf/payment-system-go/pkg/models"
	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var input models.CreateTransactionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := checkCurrency(input.Currency, c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	transaction := models.NewTransaction(models.User{
		ID:    input.UserID,
		Email: input.Email,
	}, input.Amount, input.Currency)

	DB.Create(&transaction)

	c.JSON(http.StatusCreated, transaction)
}

func checkCurrency(c models.Currency, ctx *gin.Context) error {
	switch c {
	case models.Ruble:
	case models.Dollar:
	case models.Euro:
	default:
		currencies := strings.Join([]string{string(models.Ruble), string(models.Dollar), string(models.Euro)}, ",")
		return fmt.Errorf("supported currencies: (%q)", currencies)
	}
	return nil
}

func GetAllTransactions(c *gin.Context) {
	var transactions []models.Transaction
	DB.Find(&transactions)

	c.JSON(http.StatusOK, transactions)
}
