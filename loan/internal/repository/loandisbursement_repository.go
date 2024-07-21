package repository

import (
	"loan/internal/entity"

	"gorm.io/gorm"
)

type LoanDisbursementRepository struct {
	Repository[entity.LoanDisbursement]
}

func NewLoanDisbursementRepository() *LoanDisbursementRepository {
	return &LoanDisbursementRepository{}
}

func (r *LoanDisbursementRepository) CountByLoanId(db *gorm.DB, loanId uint) (total int64, err error) {
	err = db.Model(entity.LoanDisbursement{}).Where("loan_id = ?", loanId).Count(&total).Error
	return total, err
}
