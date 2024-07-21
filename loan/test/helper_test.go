package test

import (
	"loan/internal/entity"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func ClearAll() {
	ClearLoanApprovals()
	ClearLoanDisbursements()
	ClearLoanInvestments()
	ClearLoans()
	ClearPayments()
	ClearUserRoles()
	ClearUsers()
}

func ClearLoanApprovals() {
	err := db.Where("id is not null").Delete(&entity.LoanApproval{}).Error
	if err != nil {
		log.Fatalf("Failed clear loan approval data : %+v", err)
	}
}

func ClearLoanDisbursements() {
	err := db.Where("id is not null").Delete(&entity.LoanDisbursement{}).Error
	if err != nil {
		log.Fatalf("Failed clear loan disbursement data : %+v", err)
	}
}

func ClearLoanInvestments() {
	err := db.Where("id is not null").Delete(&entity.LoanInvestment{}).Error
	if err != nil {
		log.Fatalf("Failed clear loan investment data : %+v", err)
	}
}

func ClearLoans() {
	err := db.Where("id is not null").Delete(&entity.Loan{}).Error
	if err != nil {
		log.Fatalf("Failed clear loan data : %+v", err)
	}
}

func ClearPayments() {
	err := db.Where("id is not null").Delete(&entity.Payment{}).Error
	if err != nil {
		log.Fatalf("Failed clear payment data : %+v", err)
	}
}

func ClearUserRoles() {
	err := db.Where("user_id is not null").Delete(&entity.UserRole{}).Error
	if err != nil {
		log.Fatalf("Failed clear user role data : %+v", err)
	}
}

func ClearUsers() {
	err := db.Where("id is not null").Delete(&entity.User{}).Error
	if err != nil {
		log.Fatalf("Failed clear user data : %+v", err)
	}
}

func CreateLoanApproved(t *testing.T) (res *entity.Loan) {
	user := entity.User{
		Email:    faker.Email(),
		Password: "password",
		Token:    faker.UUIDDigit(),
	}
	err := db.Create(&user).Error
	assert.Nil(t, err)
	loan := entity.Loan{
		BorrowerID:      user.ID,
		PrincipalAmount: 100000000,
		Rate:            0.1,
		ROI:             110000000,
		AgreementLetter: faker.URL(),
		State:           entity.Approved,
		TotalInvested:   0,
	}
	err = db.Create(&loan).Error
	assert.Nil(t, err)
	return &loan
}

func CreateLoanApprovedAndFullfilled(t *testing.T) (res *entity.Loan) {
	user := entity.User{
		Email:    faker.Email(),
		Password: "password",
		Token:    faker.UUIDDigit(),
	}
	err := db.Create(&user).Error
	assert.Nil(t, err)

	loan := entity.Loan{
		BorrowerID:      user.ID,
		PrincipalAmount: 100000000,
		Rate:            0.1,
		ROI:             110000000,
		AgreementLetter: faker.URL(),
		State:           entity.Approved,
		TotalInvested:   0,
	}
	err = db.Create(&loan).Error
	assert.Nil(t, err)

	paidAt := time.Now()
	paidLoanInvestment := entity.LoanInvestment{
		LoanID:     loan.ID,
		InvestorID: user.ID,
		Amount:     90000000,
		PaidAt:     &paidAt,
	}
	err = db.Create(&paidLoanInvestment).Error
	assert.Nil(t, err)

	unpaidLoanInvestment := entity.LoanInvestment{
		LoanID:     loan.ID,
		InvestorID: user.ID,
		Amount:     90000000,
		PaidAt:     nil,
	}
	err = db.Create(&unpaidLoanInvestment).Error
	assert.Nil(t, err)

	unpaidPayment := entity.Payment{
		ProductType: entity.PaymentProductTypeLoanInvestment,
		ProductID:   unpaidLoanInvestment.ID,
		Amount:      10000000,
		ExpiredAt:   time.Now().Add(1 * time.Hour),
		PaidAt:      nil,
	}
	err = db.Create(&unpaidPayment).Error
	assert.Nil(t, err)

	return &loan
}

func CreateLoanInvested(t *testing.T) (res *entity.Loan) {
	user := entity.User{
		Email:    faker.Email(),
		Password: "password",
		Token:    faker.UUIDDigit(),
	}
	err := db.Create(&user).Error
	assert.Nil(t, err)
	loan := entity.Loan{
		BorrowerID:      user.ID,
		PrincipalAmount: 100000000,
		Rate:            0.1,
		ROI:             110000000,
		AgreementLetter: faker.URL(),
		State:           entity.Invested,
		TotalInvested:   0,
	}
	err = db.Create(&loan).Error
	assert.Nil(t, err)
	return &loan
}

func CreateLoanProposed(t *testing.T) (res *entity.Loan) {
	user := entity.User{
		Email:    faker.Email(),
		Password: "password",
		Token:    faker.UUIDDigit(),
	}
	err := db.Create(&user).Error
	assert.Nil(t, err)
	loan := entity.Loan{
		BorrowerID:      user.ID,
		PrincipalAmount: 100000000,
		Rate:            0.1,
		ROI:             110000000,
		AgreementLetter: faker.URL(),
		State:           entity.Proposed,
		TotalInvested:   0,
	}
	err = db.Create(&loan).Error
	assert.Nil(t, err)
	return &loan
}

func CreateUser(t *testing.T) (res *entity.User) {
	user := entity.User{
		Email:    faker.Email(),
		Password: "password",
		Token:    faker.UUIDDigit(),
	}
	err := db.Create(&user).Error
	assert.Nil(t, err)
	return &user
}

func CreateUserAdmin(t *testing.T) (res *entity.User) {
	user := entity.User{
		Email:    faker.Email(),
		Password: "password",
		Token:    faker.UUIDDigit(),
	}
	err := db.Create(&user).Error
	assert.Nil(t, err)
	userRole := entity.UserRole{
		UserID:   user.ID,
		RoleCode: entity.RoleCodeAdmin,
	}
	err = db.Create(&userRole).Error
	assert.Nil(t, err)
	return &user
}
