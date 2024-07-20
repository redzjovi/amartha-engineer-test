package entity

import "time"

type TransactionType string

const (
	DEBIT  TransactionType = "DEBIT"
	CREDIT TransactionType = "CREDIT"
)

type Transaction struct {
	TrxID           string
	Amount          float64
	Type            TransactionType
	TransactionTime time.Time
}

func (t Transaction) GetRealAmount() float64 {
	if t.Type == DEBIT {
		return -t.Amount
	}
	return t.Amount
}
