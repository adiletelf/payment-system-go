package models

type UpdateTransactionInput struct {
	Status TransactionStatus `json:"status" binding:"required"`
}
