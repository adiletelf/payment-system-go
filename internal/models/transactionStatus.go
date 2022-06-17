package models

import (
	"fmt"
	"strings"
)

type TransactionStatus string

const (
	Created   TransactionStatus = "created"
	Succeed   TransactionStatus = "succeed"
	Unsucceed TransactionStatus = "unsucceed"
	Failed    TransactionStatus = "failed"
	Canceled  TransactionStatus = "canceled"
)

func (ts TransactionStatus) IsSupported() error {
	switch ts {
	case Created, Succeed, Unsucceed, Failed, Canceled:
		return nil
	default:
		statuses := strings.Join([]string{string(Created), string(Succeed), string(Unsucceed), string(Failed), string(Canceled)}, ", ")
		return fmt.Errorf("supported statuses: (%v)", statuses)
	}
}

func (ts TransactionStatus) IsModifiable() error {
	switch ts {
	case Created, Failed, Canceled:
		return nil
	default:
		return fmt.Errorf("transaction is in not modifiable status (%v)", ts)
	}
}
