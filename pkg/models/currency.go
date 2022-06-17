package models

import (
	"fmt"
	"strings"
)

type Currency string

const (
	Ruble  Currency = "ruble"
	Dollar Currency = "dollar"
	Euro   Currency = "euro"
)

func (c Currency) IsSupported() error {
	switch c {
	case Ruble, Dollar, Euro:
		return nil
	default:
		currencies := strings.Join([]string{string(Ruble), string(Dollar), string(Euro)}, ", ")
		return fmt.Errorf("supported currencies: (%v)", currencies)
	}
}
