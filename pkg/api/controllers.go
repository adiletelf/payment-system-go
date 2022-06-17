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
	input, err := bindCreateTransactionInput(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.tr.FindById(input.ID)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record already exists."})
		return
	}

	err = h.tr.UpdateStatusOrCreate(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, input)
}

func (h *BaseHandler) GetTransactionStatus(c *gin.Context) {
	t, err := h.findTransactionById(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactionStatus": t.Status})
}

func (h *BaseHandler) UpdateTransactionStatus(c *gin.Context) {
	var input models.UpdateTransactionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t, err := h.findTransactionById(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	t.Status = input.Status
	if err := h.tr.UpdateStatusOrCreate(t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, t)
}


func bindCreateTransactionInput(c *gin.Context) (*models.Transaction, error) {
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

func (h *BaseHandler) findTransactionById(c *gin.Context) (*models.Transaction, error) {
	id, err := strconv.Atoi(c.Param("id"))	
	if err != nil {
		return nil, fmt.Errorf("invalid id")
	}

	t, err := h.tr.FindById(uint(id))
	if err != nil {
		return nil, err
	}

	return t, nil
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
