package entity

import "time"

type Loan struct {
	ID                int
	StartAt           time.Time
	EndAt             time.Time
	Principal         int
	InterestRate      float64
	OutstandingAmount int
}
