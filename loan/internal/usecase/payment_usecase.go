package usecase

import (
	"context"
	"loan/internal/entity"
	"loan/internal/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PaymentUsecase struct {
	DB                       *gorm.DB
	loanInvestmentRepository *repository.LoanInvestmentRepository
	loanRepository           *repository.LoanRepository
	Log                      *logrus.Logger
	paymentRepository        *repository.PaymentRepository
	Validate                 *validator.Validate
}

func NewPaymentUsecase(
	db *gorm.DB,
	loanInvestmentRepository *repository.LoanInvestmentRepository,
	log *logrus.Logger,
	paymentRepository *repository.PaymentRepository,
	validate *validator.Validate,
) *PaymentUsecase {
	return &PaymentUsecase{
		DB:                       db,
		loanInvestmentRepository: loanInvestmentRepository,
		Log:                      log,
		paymentRepository:        paymentRepository,
		Validate:                 validate,
	}
}

func (u *PaymentUsecase) Pay(ctx context.Context, id uint) error {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	payment := new(entity.Payment)
	if err := u.paymentRepository.FindById(tx, payment, id); err != nil {
		if err == gorm.ErrRecordNotFound {
			return fiber.ErrNotFound
		}
		u.Log.Warnf("Failed find payment by id : %+v", err)
		return fiber.ErrInternalServerError
	} else if payment.PaidAt != nil {
		return fiber.NewError(fiber.StatusConflict, "payment already paid")
	}

	paidAt := time.Now()
	payment.PaidAt = &paidAt

	if err := u.paymentRepository.Update(tx, payment); err != nil {
		u.Log.Warnf("Failed update payment : %+v", err)
		return fiber.ErrInternalServerError
	}

	if payment.ProductType == entity.PaymentProductTypeLoanInvestment {
		if err := u.payLoanInvestment(tx, payment.ProductID); err != nil {
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.Warnf("Failed commit transaction : %+v", err)
		return fiber.ErrInternalServerError
	}

	return nil
}

func (u *PaymentUsecase) payLoanInvestment(tx *gorm.DB, loanInvestmentId uint) error {
	loanInvestment := new(entity.LoanInvestment)
	if err := u.loanInvestmentRepository.FindById(tx, loanInvestment, loanInvestmentId); err != nil {
		if err == gorm.ErrRecordNotFound {
			return fiber.ErrNotFound
		}
		u.Log.Warnf("Failed find loan investment by id : %+v", err)
		return fiber.ErrInternalServerError
	}
	paidAt := time.Now()
	loanInvestment.PaidAt = &paidAt
	if err := u.loanInvestmentRepository.Update(tx, loanInvestment); err != nil {
		u.Log.Warnf("Failed update loanInvestment : %+v", err)
		return fiber.ErrInternalServerError
	}

	loan := new(entity.Loan)
	if err := u.loanRepository.FindById(tx, loan, loanInvestment); err != nil {
		if err == gorm.ErrRecordNotFound {
			return fiber.ErrNotFound
		}
		u.Log.Warnf("Failed find loan by id : %+v", err)
		return fiber.ErrInternalServerError
	}
	loan.TotalInvested += loanInvestment.Amount
	if loan.TotalInvested >= loan.PrincipalAmount {
		loan.State = entity.Invested
	}
	if err := u.loanRepository.Update(tx, loan); err != nil {
		u.Log.Warnf("Failed update loan : %+v", err)
		return fiber.ErrInternalServerError
	}

	return nil
}
