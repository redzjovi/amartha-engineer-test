package entity

import "time"

type LoanState string

const (
	Proposed  LoanState = "proposed"
	Approved  LoanState = "approved"
	Invested  LoanState = "invested"
	Disbursed LoanState = "disbursed"
)

type Loan struct {
	ID              uint      `gorm:"column:id;primaryKey"`
	BorrowerID      uint      `gorm:"column:borrower_id"`
	PrincipalAmount float64   `gorm:"column:principal_amount"`
	Rate            float64   `gorm:"column:rate"`
	ROI             float64   `gorm:"column:roi"`
	AgreementLetter string    `gorm:"column:agreement_letter"`
	State           LoanState `gorm:"column:state"`
	TotalInvested   float64   `gorm:"column:total_invested"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
}

func (u *Loan) TableName() string {
	return "loans"
}
