package models

type TransactionStatus int

const (
	Created TransactionStatus = iota
	Successed
	Unsuccessed
	Failed
	Canceled
)

func (ts TransactionStatus) String() string {
	return []string{"created", "successed", "unsuccessed", "failed", "canceled"}[ts]
}