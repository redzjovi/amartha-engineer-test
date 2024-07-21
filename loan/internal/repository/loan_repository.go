package repository

import (
	"loan/internal/entity"

	"gorm.io/gorm"
)

type LoanRepository struct {
	Repository[entity.Loan]
}

func NewLoanRepository() *LoanRepository {
	return &LoanRepository{}
}

func (r *LoanRepository) List(db *gorm.DB) (res []entity.Loan, err error) {
	err = db.
		Order("created_at DESC").
		Find(&res).
		Error
	return res, err
}

func (r *LoanRepository) ListByState(db *gorm.DB, state entity.LoanState) (res []entity.Loan, err error) {
	err = db.
		Where("state = ?", state).
		Order("created_at DESC").
		Find(&res).
		Error
	return res, err
}
