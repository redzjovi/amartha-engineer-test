package repository

import (
	"loan/internal/entity"

	"gorm.io/gorm"
)

type LoanApprovalRepository struct {
	Repository[entity.LoanApproval]
}

func NewLoanApprovalRepository() *LoanApprovalRepository {
	return &LoanApprovalRepository{}
}

func (r *LoanApprovalRepository) CountByLoanId(db *gorm.DB, loanId uint) (total int64, err error) {
	err = db.Model(entity.LoanApproval{}).Where("loan_id = ?", loanId).Count(&total).Error
	return total, err
}
