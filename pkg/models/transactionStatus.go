package models

type TransactionStatus string

const (
	Created TransactionStatus = "created"
	Succeed TransactionStatus = "succeed"
	Unsucceed TransactionStatus = "unsucceed"
	Failed TransactionStatus = "failed"
	Canceled TransactionStatus = "canceled"
)
