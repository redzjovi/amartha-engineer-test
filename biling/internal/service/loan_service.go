package service

import (
	"billing/internal/entity"
	"billing/internal/repository"
	"errors"
	"time"
)

type LoanService interface {
	Create(principal int, interestRate float64, week int) *entity.Loan
	GetOutstanding(loanID int) (int, error)
	IsDelinquent(loanID int) bool
	MakePayment(loanID int, amount int) error
}

type loanService struct {
	loanRepository        repository.LoanRepository
	loanPaymentRepository repository.LoanPaymentRepository
}

func NewLoanService(
	loanRepository repository.LoanRepository,
	loanPaymentRepository repository.LoanPaymentRepository,
) LoanService {
	return &loanService{
		loanRepository:        loanRepository,
		loanPaymentRepository: loanPaymentRepository,
	}
}

func (s *loanService) Create(principal int, interestRate float64, week int) *entity.Loan {
	now := time.Now()

	loan := &entity.Loan{
		StartAt:           now,
		EndAt:             now.Add(time.Duration(week) * 7 * 24 * time.Hour),
		Principal:         principal,
		InterestRate:      interestRate,
		OutstandingAmount: int(float64(principal) * (1 + interestRate)),
	}
	s.loanRepository.Create(loan)

	for i := 1; i <= week; i++ {
		loanPayment := &entity.LoanPayment{
			LoanID:  loan.ID,
			StartAt: now.Add(time.Duration(i-1) * 7 * 24 * time.Hour),
			EndAt:   now.Add(time.Duration(i) * 7 * 24 * time.Hour),
			Amount:  int(float64(principal) * (1 + interestRate) / float64(week)),
			PaidAt:  nil,
		}
		s.loanPaymentRepository.Create(loanPayment)
	}

	return loan
}

func (s *loanService) GetOutstanding(loanID int) (int, error) {
	loan, err := s.loanRepository.FindByID(loanID)
	if err != nil {
		return 0, err
	} else if loan == nil {
		return 0, errors.New("loan not found")
	}
	return loan.OutstandingAmount, nil
}

func (s *loanService) IsDelinquent(loanID int) bool {
	loanPayments := s.loanPaymentRepository.ListOutstandingByLoanID(loanID)
	return len(loanPayments) >= 2
}

func (s *loanService) MakePayment(loanID int, amount int) error {
	loan, err := s.loanRepository.FindByID(loanID)
	if err != nil {
		return err
	} else if loan == nil {
		return errors.New("loan not found")
	} else if loan.OutstandingAmount <= 0 {
		return errors.New("loan already paid")
	}

	loanPayments := s.loanPaymentRepository.ListOutstandingByLoanID(loanID)
	for _, loanPayment := range loanPayments {
		if amount < loanPayment.Amount {
			return errors.New("invalid payment amount")
		}
		paidAt := time.Now()
		loanPayment.PaidAt = &paidAt
		s.loanPaymentRepository.Save(&loanPayment)
	}

	loan.OutstandingAmount -= amount
	s.loanRepository.Save(loan)

	return nil
}
