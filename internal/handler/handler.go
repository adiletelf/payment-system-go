package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/adiletelf/payment-system-go/internal/model"
	"github.com/gin-gonic/gin"
)

type BaseHandler struct {
	tr model.TransactionRepo
	ar model.AdminRepo
}

func NewBaseHandler(tr model.TransactionRepo, ar model.AdminRepo) *BaseHandler {
	return &BaseHandler{
		tr: tr,
		ar : ar,
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
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": string(t.Status)})
}

func (h *BaseHandler) UpdateTransactionStatus(c *gin.Context) {
	input, err := bindUpdateTransactionInput(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Status == model.Canceled {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Not allowed to cancel transaction using current endpoint."})
		return
	}

	t, err := h.findTransactionById(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := t.Status.IsModifiable(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t.Status = input.Status
	if err := h.tr.UpdateStatusOrCreate(t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, t)
}

func (h *BaseHandler) CancelTransaction(c *gin.Context) {
	t, err := h.findTransactionById(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := t.Status.IsModifiable(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t.Status = model.Canceled
	if err := h.tr.UpdateStatusOrCreate(t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, t)
}

func bindCreateTransactionInput(c *gin.Context) (*model.Transaction, error) {
	var input model.CreateTransactionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	if err := input.Currency.IsSupported(); err != nil {
		return nil, err
	}

	t := model.NewTransaction(model.User{
		ID:    input.UserID,
		Email: input.Email,
	}, input.Amount, input.Currency)

	return &t, nil
}

func bindUpdateTransactionInput(c *gin.Context) (*model.UpdateTransactionInput, error) {
	var input model.UpdateTransactionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		return nil, fmt.Errorf("unable to unmarshall input")
	}

	if err := input.Status.IsSupported(); err != nil {
		return nil, err
	}
	return &input, nil
}

func (h *BaseHandler) findTransactionById(c *gin.Context) (*model.Transaction, error) {
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
