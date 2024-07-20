package test

import (
	"billing/internal/entity"
	"time"
)

func ClearAll() {
	ClearLoans()
}

func ClearLoans() {
	loanRepository.Truncate()
	loanPaymentRepository.Truncate()
}

func CreateLoan() *entity.Loan {
	now := time.Now().Add(-3 * 7 * 24 * time.Hour)
	loan := &entity.Loan{
		StartAt:           now,
		EndAt:             now.Add(time.Duration(50) * 7 * 24 * time.Hour),
		Principal:         5000000,
		InterestRate:      0.1,
		OutstandingAmount: 5500000,
	}
	loanRepository.Create(loan)
	return loan
}

func CreateLoanAndLoanPayments() *entity.Loan {
	loan := CreateLoan()
	for i := 1; i <= 50; i++ {
		endAt := loan.StartAt.Add(time.Duration(i) * 7 * 24 * time.Hour)
		var paidAt *time.Time
		if i <= 2 {
			paidAt = &endAt
		}
		loanPayment := &entity.LoanPayment{
			LoanID:  loan.ID,
			StartAt: loan.StartAt.Add(time.Duration(i-1) * 7 * 24 * time.Hour),
			EndAt:   endAt,
			Amount:  int(float64(loan.Principal) * (1 + loan.InterestRate) / float64(50)),
			PaidAt:  paidAt,
		}
		loanPaymentRepository.Create(loanPayment)
	}
	return loan
}

func CreateLoanAndOutstandingLoanPayments() *entity.Loan {
	loan := CreateLoan()
	for i := 1; i <= 50; i++ {
		loanPayment := &entity.LoanPayment{
			LoanID:  loan.ID,
			StartAt: loan.StartAt.Add(time.Duration(i-1) * 7 * 24 * time.Hour),
			EndAt:   loan.StartAt.Add(time.Duration(i) * 7 * 24 * time.Hour),
			Amount:  int(float64(loan.Principal) * (1 + loan.InterestRate) / float64(50)),
			PaidAt:  nil,
		}
		loanPaymentRepository.Create(loanPayment)
	}
	return loan
}

func CreatePaidLoan() *entity.Loan {
	now := time.Now().Add(-3 * 7 * 24 * time.Hour)
	loan := &entity.Loan{
		StartAt:           now,
		EndAt:             now.Add(time.Duration(50) * 7 * 24 * time.Hour),
		Principal:         5000000,
		InterestRate:      0.1,
		OutstandingAmount: 0,
	}
	loanRepository.Create(loan)
	return loan
}
