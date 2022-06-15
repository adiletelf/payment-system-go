package models

import "time"

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
