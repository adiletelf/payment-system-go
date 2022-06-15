package models

type CreateTransactionInput struct {
	UserID   uint     `json:"userId" binding:"required"`
	Email    string   `json:"email" binding:"required"`
	Amount   float64  `json:"amount" binding:"required"`
	Currency Currency `json:"currency" binding:"required"`
}
