package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/adiletelf/payment-system-go/pkg/models"
	"github.com/gin-gonic/gin"
)

func GetAllTransactions(c *gin.Context) {
	var transactions []models.Transaction

	email := c.Query("email")
	userId, err := strconv.Atoi(c.Query("userId"))
	if err != nil {
		userId = 0
	}

	// query ignores parameter if zero field (0, "", false)
	if err := DB.Where(&models.Transaction{UserID: uint(userId), Email: email}).Find(&transactions).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, transactions)
}

func CreateTransaction(c *gin.Context) {
	var input models.CreateTransactionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := checkCurrency(input.Currency); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	transaction := models.NewTransaction(models.User{
		ID:    input.UserID,
		Email: input.Email,
	}, input.Amount, input.Currency)

	DB.Create(&transaction)

	c.JSON(http.StatusCreated, transaction)
}

func GetTransactionStatus(c *gin.Context) {
	var t models.Transaction

	if err := DB.Where("id = ?", c.Param("id")).First(&t).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactionStatus": t.Status.String()})
}

func checkCurrency(c models.Currency) error {
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
