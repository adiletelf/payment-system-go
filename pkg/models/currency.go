package models

type Currency int

const (
	Ruble Currency = iota
	Dollar
	Euro
)

func (c Currency) String() string {
	return []string{"ruble", "dollar", "euro"}[c]
}
