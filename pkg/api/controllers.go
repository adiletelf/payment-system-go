package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/adiletelf/payment-system-go/pkg/models"
	"github.com/gin-gonic/gin"
)

type BaseHandler struct {
	tr models.TransactionRepo
}

func NewBaseHandler(tr models.TransactionRepo) *BaseHandler {
	return &BaseHandler{
		tr: tr,
	}
}

func (h *BaseHandler) GetAllTransactions(c *gin.Context) {
	email := c.Query("email")
	userId, err := strconv.Atoi(c.Query("userId"))
	if err != nil {
		userId = 0
	}

	transactions, err := h.tr.Find(uint(userId), email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func (h *BaseHandler) CreateTransaction(c *gin.Context) {
	t, err := bindTransactionInput(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.tr.Save(t)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, t)
}

func (h *BaseHandler) GetTransactionStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id."})
		return
	}

	t, err := h.tr.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactionStatus": t.Status})
}

func bindTransactionInput(c *gin.Context) (*models.Transaction, error) {
	var input models.CreateTransactionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	if err := checkCurrency(input.Currency); err != nil {
		return nil, err
	}

	t := models.NewTransaction(models.User{
		ID:    input.UserID,
		Email: input.Email,
	}, input.Amount, input.Currency)

	return &t, nil
}

func checkCurrency(c models.Currency) error {
	switch c {
	case models.Ruble:
	case models.Dollar:
	case models.Euro:
	default:
		currencies := strings.Join([]string{string(models.Ruble), string(models.Dollar), string(models.Euro)}, ", ")
		return fmt.Errorf("supported currencies: (%v)", currencies)
	}
	return nil
}
