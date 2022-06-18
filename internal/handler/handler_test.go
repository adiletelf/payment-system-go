package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/adiletelf/payment-system-go/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockTransactionRepo struct {
	tr []model.Transaction
}

func NewMockTransactionRepo() *MockTransactionRepo {
	first := model.User{ID: 1, Email: "first@gmail.com"}
	second := model.User{ID: 2, Email: "second@gmail.com"}
	third := model.User{ID: 3, Email: "third@gmail.com"}

	tr := []model.Transaction{
		{ID: 1, UserID: first.ID, Email: first.Email, Amount: 100, Currency: model.Ruble, Status: model.Created, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 2, UserID: first.ID, Email: first.Email, Amount: 125, Currency: model.Ruble, Status: model.Created, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 3, UserID: first.ID, Email: first.Email, Amount: 126, Currency: model.Ruble, Status: model.Succeed, CreatedAt: time.Now(), UpdatedAt: time.Now()},

		{ID: 4, UserID: second.ID, Email: second.Email, Amount: 251, Currency: model.Ruble, Status: model.Created, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 5, UserID: second.ID, Email: second.Email, Amount: 252, Currency: model.Ruble, Status: model.Created, CreatedAt: time.Now(), UpdatedAt: time.Now()},

		{ID: 6, UserID: third.ID, Email: third.Email, Amount: 300, Currency: model.Ruble, Status: model.Created, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	return &MockTransactionRepo{
		tr: tr,
	}
}

func (m *MockTransactionRepo) Find(userID uint, email string) ([]model.Transaction, error) {
	if userID == 0 && email == "" {
		return m.tr, nil
	}

	if userID == 0 {
		return filterTransactions(m.tr, func(t model.Transaction) bool { return t.Email == email }), nil
	}
	if email == "" {
		return filterTransactions(m.tr, func(t model.Transaction) bool { return t.UserID == userID }), nil
	}

	return filterTransactions(m.tr, func(t model.Transaction) bool {
		return t.Email == email && t.UserID == userID
	}), nil
}

func (m *MockTransactionRepo) FindById(id uint) (*model.Transaction, error) {
	for _, t := range m.tr {
		if t.ID == id {
			return &t, nil
		}
	}

	return nil, fmt.Errorf("record not found")
}

func (m *MockTransactionRepo) UpdateStatusOrCreate(t *model.Transaction) error {
	m.tr = append(m.tr, *t)
	return nil
}

func filterTransactions(tr []model.Transaction, f func(model.Transaction) bool) []model.Transaction {
	var result []model.Transaction
	for _, t := range tr {
		if f(t) {
			result = append(result, t)
		}
	}
	return result
}

var router *gin.Engine

const (
	GetAllTransactionsPath      = "/transactions/"
	CreateTransactionPath       = "/transaction/"
	GetTransactionStatusPath    = "/transaction/"
	UpdateTransactionStatusPath = "/transaction/"
	CancelTransactionPath       = "/transaction/cancel/"
)

func TestMain(m *testing.M) {
	router = gin.Default()
	h := &BaseHandler{
		tr: NewMockTransactionRepo(),
	}
	router.GET(GetAllTransactionsPath, h.GetAllTransactions)
	router.GET(GetTransactionStatusPath+":id", h.GetTransactionStatus)
	router.PUT(UpdateTransactionStatusPath+":id", h.UpdateTransactionStatus)
	router.POST(CreateTransactionPath, h.CreateTransaction)
	router.GET(CancelTransactionPath+":id", h.CancelTransaction)

	m.Run()
}

func performRequest(method, path string, body io.Reader) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, path, body)
	router.ServeHTTP(w, req)

	return w
}

func TestBaseHandler_GetAllTransactions_ReturnAllTransactions(t *testing.T) {
	w := performRequest("GET", GetAllTransactionsPath, nil)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body)
}

func TestBaseHandler_GetAllTransactions_ReturnEmpty(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", GetAllTransactionsPath, nil)
	q := req.URL.Query()
	q.Add("email", "nonexisting@gmail.com")
	q.Add("userId", "0")
	req.URL.RawQuery = q.Encode()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "null", w.Body.String())
}

func TestBaseHandler_CreateTransaction_EmptyRequest(t *testing.T) {
	w := performRequest("POST", CreateTransactionPath, nil)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestBaseHandler_CreateTransaction_NonEmptyRequest(t *testing.T) {
	input := model.CreateTransactionInput{
		UserID:   4,
		Email:    "fourth@gmail.com",
		Amount:   444,
		Currency: model.Ruble,
	}
	json_data, _ := json.Marshal(input)
	w := performRequest("POST", CreateTransactionPath, bytes.NewBuffer(json_data))

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestBaseHandler_CreateTransaction_AlreadyExist(t *testing.T) {
	input := model.CreateTransactionInput{
		UserID:   1,
		Email:    "",
		Amount:   444,
		Currency: model.Ruble,
	}
	json_data, _ := json.Marshal(input)
	w := performRequest("POST", CreateTransactionPath, bytes.NewBuffer(json_data))

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestBaseHandler_GetTransactionStatus_NotFound(t *testing.T) {
	id := "100"
	w := performRequest("GET", GetTransactionStatusPath+id, nil)
	var response map[string]any
	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestBaseHandler_GetTransactionStatus_ReturnsStatus(t *testing.T) {
	// id := "1"
	// w := performRequest("GET", GetTransactionStatusPath+id, nil)
	w := performRequest("GET", "/transaction/1", nil)

	var response map[string]any
	err := json.Unmarshal(w.Body.Bytes(), &response)
	status, exists := response["status"]

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, err)
	assert.True(t, exists)
	switch status {
	case string(model.Created), string(model.Failed):
	default:
		t.Fatalf("unexpected status: %v", status)
	}
}

func TestBaseHandler_UpdateTransactionStatus_CancelNotAllowed(t *testing.T) {
	input := model.UpdateTransactionInput{
		Status: model.Canceled,
	}
	json_data, _ := json.Marshal(input)
	id := "1"
	w := performRequest("PUT", UpdateTransactionStatusPath+id, bytes.NewBuffer(json_data))

	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
}

func TestBaseHandler_UpdateTransactionStatus_NotFound(t *testing.T) {
	input := model.UpdateTransactionInput{
		Status: model.Succeed,
	}
	json_data, _ := json.Marshal(input)
	id := "100"
	w := performRequest("PUT", UpdateTransactionStatusPath+id, bytes.NewBuffer(json_data))

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestBaseHandler_UpdateTransactionStatus_NotModifiable(t *testing.T) {
	input := model.UpdateTransactionInput{
		Status: model.Unsucceed,
	}
	json_data, _ := json.Marshal(input)
	// unmodifiable transaction
	id := "3"
	w := performRequest("PUT", UpdateTransactionStatusPath+id, bytes.NewBuffer(json_data))

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestBaseHandler_UpdateTransactionStatus_ReturnsOK(t *testing.T) {
	input := model.UpdateTransactionInput{
		Status: model.Succeed,
	}
	json_data, _ := json.Marshal(input)
	// unmodifiable transaction
	id := "4"
	w := performRequest("PUT", UpdateTransactionStatusPath+id, bytes.NewBuffer(json_data))

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestBaseHandler_CancelTransaction_NotFound(t *testing.T) {
	id := "100"
	w := performRequest("GET", CancelTransactionPath+id, nil)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestBaseHandler_CancelTransaction_NotModifiable(t *testing.T) {
	id := "3"
	w := performRequest("GET", CancelTransactionPath+id, nil)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestBaseHandler_CancelTransaction_ReturnsOK(t *testing.T) {
	id := "1"
	w := performRequest("GET", CancelTransactionPath+id, nil)

	var response map[string]any
	err := json.Unmarshal(w.Body.Bytes(), &response)
	status, exists := response["status"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, status, string(model.Canceled))
}
