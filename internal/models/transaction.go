package models

import (
	"math/rand"
	"time"
)

const oneInNFailureRate = 10

type Transaction struct {
	ID        uint              `json:"id" gorm:"primary_key"`
	UserID    uint              `json:"userId"`
	Email     string            `json:"email"`
	Amount    float64           `json:"amount"`
	Currency  Currency          `json:"currency"`
	Status    TransactionStatus `json:"status"`
	CreatedAt time.Time         `json:"createdAt"`
	UpdatedAt time.Time         `json:"updatedAt"`
}

type TransactionRepo interface {
	Find(userId uint, email string) ([]Transaction, error)
	FindById(id uint) (*Transaction, error)
	UpdateStatusOrCreate(t *Transaction) error
}

func NewTransaction(u User, amount float64, currency Currency) Transaction {
	t := Transaction{
		UserID:    u.ID,
		Email:     u.Email,
		Amount:    amount,
		Currency:  currency,
		Status:    Created,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if isFailedTransaction(oneInNFailureRate) {
		t.Status = Failed
	}

	return t
}

func isFailedTransaction(p int) bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(p) == 0
}
