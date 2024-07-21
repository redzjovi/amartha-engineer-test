package entity

import "time"

type LoanApproval struct {
	ID           uint      `gorm:"column:id;primaryKey"`
	LoanID       uint      `gorm:"column:loan_id"`
	ValidatorID  uint      `gorm:"validator_id"`
	PictureProof string    `gorm:"picture_proof"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

func (la *LoanApproval) TableName() string {
	return "loan_approvals"
}
