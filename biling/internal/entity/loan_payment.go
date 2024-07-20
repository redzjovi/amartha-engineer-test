package entity

import "time"

type LoanPayment struct {
	ID      int
	LoanID  int
	StartAt time.Time
	EndAt   time.Time
	Amount  int
	PaidAt  *time.Time
}
