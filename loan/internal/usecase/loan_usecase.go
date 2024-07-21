package usecase

import (
	"context"
	"loan/internal/entity"
	"loan/internal/model"
	"loan/internal/model/converter"
	"loan/internal/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type LoanUsecase struct {
	DB                         *gorm.DB
	loanApprovalRepository     *repository.LoanApprovalRepository
	loanDisbursementRepository *repository.LoanDisbursementRepository
	loanInvestmentRepository   *repository.LoanInvestmentRepository
	loanRepository             *repository.LoanRepository
	Log                        *logrus.Logger
	paymentRepository          *repository.PaymentRepository
	Validate                   *validator.Validate
}

func NewLoanUsecase(
	db *gorm.DB,
	loanRepository *repository.LoanRepository,
	loanApprovalRepository *repository.LoanApprovalRepository,
	loanDisbursementRepository *repository.LoanDisbursementRepository,
	loanInvestmentRepository *repository.LoanInvestmentRepository,
	log *logrus.Logger,
	paymentRepository *repository.PaymentRepository,
	validate *validator.Validate,
) *LoanUsecase {
	return &LoanUsecase{
		DB:                         db,
		loanRepository:             loanRepository,
		loanApprovalRepository:     loanApprovalRepository,
		loanDisbursementRepository: loanDisbursementRepository,
		loanInvestmentRepository:   loanInvestmentRepository,
		Log:                        log,
		paymentRepository:          paymentRepository,
		Validate:                   validate,
	}
}

func (u *LoanUsecase) Approve(ctx context.Context, loanId uint, validatorId uint, request *model.LoanApproveRequest) error {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	loan := new(entity.Loan)
	if err := u.loanRepository.FindById(tx, loan, loanId); err != nil {
		if err == gorm.ErrRecordNotFound {
			return fiber.ErrNotFound
		}
		u.Log.Warnf("Failed find loan by id : %+v", err)
		return fiber.ErrInternalServerError
	} else if loan.State != entity.Proposed {
		return fiber.NewError(fiber.StatusConflict, "loan is not proposed")
	}

	loan.State = entity.Approved

	if err := u.loanRepository.Update(tx, loan); err != nil {
		u.Log.Warnf("Failed update loan : %+v", err)
		return fiber.ErrInternalServerError
	}

	loanApproval := new(entity.LoanApproval)
	loanApproval.LoanID = loanId
	loanApproval.ValidatorID = validatorId
	loanApproval.PictureProof = request.PictureProof

	if err := u.loanApprovalRepository.Create(tx, loanApproval); err != nil {
		u.Log.Warnf("Failed create loan approval : %+v", err)
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.Warnf("Failed commit transaction : %+v", err)
		return fiber.ErrInternalServerError
	}

	return nil
}

func (u *LoanUsecase) Disburse(ctx context.Context, loanId uint, fieldOfficerID uint, request *model.LoanDisburseRequest) error {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	loan := new(entity.Loan)
	if err := u.loanRepository.FindById(tx, loan, loanId); err != nil {
		if err == gorm.ErrRecordNotFound {
			return fiber.ErrNotFound
		}
		u.Log.Warnf("Failed find loan by id : %+v", err)
		return fiber.ErrInternalServerError
	} else if loan.State != entity.Invested {
		return fiber.NewError(fiber.StatusConflict, "loan is not invested")
	}

	loan.State = entity.Disbursed

	if err := u.loanRepository.Update(tx, loan); err != nil {
		u.Log.Warnf("Failed create loan : %+v", err)
		return fiber.ErrInternalServerError
	}

	loanDisbursement := new(entity.LoanDisbursement)
	loanDisbursement.LoanID = loanId
	loanDisbursement.FieldOfficerID = fieldOfficerID
	loanDisbursement.AgreementLetter = request.AgreementLetter

	if err := u.loanDisbursementRepository.Create(tx, loanDisbursement); err != nil {
		u.Log.Warnf("Failed create loan disbursement : %+v", err)
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.Warnf("Failed commit transaction : %+v", err)
		return fiber.ErrInternalServerError
	}

	return nil
}

func (u *LoanUsecase) Invest(ctx context.Context, loanId uint, investorId uint, request *model.LoanInvestRequest) error {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	loan := new(entity.Loan)
	if err := u.loanRepository.FindById(tx, loan, loanId); err != nil {
		if err == gorm.ErrRecordNotFound {
			return fiber.ErrNotFound
		}
		u.Log.Warnf("Failed find loan by id : %+v", err)
		return fiber.ErrInternalServerError
	} else if loan.State != entity.Approved {
		return fiber.NewError(fiber.StatusConflict, "loan is not approved")
	}

	totalInvested, err := u.loanInvestmentRepository.SumAmountByLoanIDAndPay(tx, loan.ID)
	if err != nil {
		u.Log.Warnf("Failed sum amount loan investment : %+v", err)
		return fiber.ErrInternalServerError
	}
	if totalInvested >= loan.PrincipalAmount {
		return fiber.NewError(fiber.StatusNotAcceptable, "loan already fullfilled")
	} else if totalInvested+request.Amount > loan.PrincipalAmount {
		return fiber.NewError(fiber.StatusNotAcceptable, "invest amount invalid")
	}

	loanInvestment := new(entity.LoanInvestment)
	loanInvestment.LoanID = loanId
	loanInvestment.InvestorID = investorId
	loanInvestment.Amount = request.Amount
	loanInvestment.PaidAt = nil

	if err := u.loanInvestmentRepository.Create(tx, loanInvestment); err != nil {
		u.Log.Warnf("Failed create loan investment : %+v", err)
		return fiber.ErrInternalServerError
	}

	payment := new(entity.Payment)
	payment.ProductType = entity.PaymentProductTypeLoanInvestment
	payment.ProductID = loanInvestment.LoanID
	payment.Amount = loanInvestment.Amount
	payment.ExpiredAt = time.Now().Add(1 * time.Hour)
	payment.PaidAt = nil

	if err := u.paymentRepository.Create(tx, payment); err != nil {
		u.Log.Warnf("Failed create payment : %+v", err)
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.Warnf("Failed commit transaction : %+v", err)
		return fiber.ErrInternalServerError
	}

	return nil
}

func (u *LoanUsecase) Propose(ctx context.Context, borrowerId uint, request *model.LoanProposeRequest) error {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	loan := new(entity.Loan)
	loan.BorrowerID = borrowerId
	loan.PrincipalAmount = request.PrincipalAmount
	loan.Rate = request.Rate
	loan.ROI = request.PrincipalAmount * (1 + request.Rate)
	loan.AgreementLetter = uuid.NewString()
	loan.State = entity.Proposed
	loan.TotalInvested = 0

	if err := u.loanRepository.Create(tx, loan); err != nil {
		u.Log.Warnf("Failed create loan : %+v", err)
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.Warnf("Failed commit transaction : %+v", err)
		return fiber.ErrInternalServerError
	}

	return nil
}

func (u *LoanUsecase) List(ctx context.Context) ([]model.LoanResponse, error) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	loans, err := u.loanRepository.List(tx)
	if err != nil {
		u.Log.Warnf("Failed list loan : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	res := make([]model.LoanResponse, 0)
	for _, loan := range loans {
		res = append(res, *converter.LoanToResponse(&loan))
	}

	return res, nil
}

func (u *LoanUsecase) ListApproved(ctx context.Context) ([]model.LoanResponse, error) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	loans, err := u.loanRepository.ListByState(tx, entity.Approved)
	if err != nil {
		u.Log.Warnf("Failed list loan by state : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	res := make([]model.LoanResponse, 0)
	for _, loan := range loans {
		res = append(res, *converter.LoanToResponse(&loan))
	}

	return res, nil
}
