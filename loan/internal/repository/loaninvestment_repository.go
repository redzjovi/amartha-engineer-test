package repository

import (
	"loan/internal/entity"

	"gorm.io/gorm"
)

type LoanInvestmentRepository struct {
	Repository[entity.LoanInvestment]
}

func NewLoanInvestmentRepository() *LoanInvestmentRepository {
	return &LoanInvestmentRepository{}
}

func (r *LoanInvestmentRepository) SumAmountByLoanIDAndPay(db *gorm.DB, loanID uint) (float64, error) {
	var loanInvestment entity.LoanInvestment
	var payment entity.Payment
	result := struct {
		TotalAmount float64
	}{}
	err := db.
		Select("SUM(li.amount) AS total_amount").
		Table(loanInvestment.TableName()+" AS li").
		Joins("LEFT JOIN "+payment.TableName()+" AS p ON p.product_type = ? AND p.product_id = li.id", entity.PaymentProductTypeLoanInvestment).
		Where("li.loan_id = ?", loanID).
		Where("li.paid_at IS NOT NULL OR p.expired_at > NOW()").
		Scan(&result).Error
	return result.TotalAmount, err
}
