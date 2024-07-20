package entity

import "time"

type BankStatement struct {
	UniqueIdentifier string
	Amount           float64
	Date             time.Time
}
