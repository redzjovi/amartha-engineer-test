package entity

import "time"

type LoanDisbursement struct {
	ID              uint      `gorm:"column:id;primaryKey"`
	LoanID          uint      `gorm:"column:loan_id"`
	FieldOfficerID  uint      `gorm:"column:field_officer_id"`
	AgreementLetter string    `gorm:"column:agreement_letter"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
}

func (li *LoanDisbursement) TableName() string {
	return "loan_disbursements"
}
