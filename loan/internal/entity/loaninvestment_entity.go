package entity

import "time"

type LoanInvestment struct {
	ID         uint       `gorm:"column:id;primaryKey"`
	LoanID     uint       `gorm:"column:loan_id"`
	InvestorID uint       `gorm:"investor_id"`
	Amount     float64    `gorm:"amount"`
	CreatedAt  time.Time  `gorm:"column:created_at"`
	UpdatedAt  time.Time  `gorm:"column:updated_at"`
	PaidAt     *time.Time `gorm:"column:paid_at"`
}

func (li *LoanInvestment) TableName() string {
	return "loan_investments"
}
