package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoanCreate(t *testing.T) {
	ClearAll()

	loan := loanService.Create(5000000, 0.1, 50)

	assert.Equal(t, 5000000, loan.Principal)
	assert.Equal(t, 0.1, loan.InterestRate)
	assert.Equal(t, 5500000, loan.OutstandingAmount)
}

func TestLoanGetOutstandingNotFound(t *testing.T) {
	ClearAll()

	_, err := loanService.GetOutstanding(1)

	assert.EqualError(t, err, "loan not found")
}

func TestLoanGetOutstanding(t *testing.T) {
	ClearAll()
	loan := CreateLoan()

	outstanding, err := loanService.GetOutstanding(loan.ID)

	assert.Equal(t, loan.OutstandingAmount, outstanding)
	assert.Nil(t, err)
}

func TestLoanIsNotDelinquent(t *testing.T) {
	ClearAll()
	loan := CreateLoanAndLoanPayments()

	dilinquent := loanService.IsDelinquent(loan.ID)

	assert.Equal(t, false, dilinquent)
}

func TestLoanIsDelinquent(t *testing.T) {
	ClearAll()
	loan := CreateLoanAndOutstandingLoanPayments()

	dilinquent := loanService.IsDelinquent(loan.ID)

	assert.Equal(t, true, dilinquent)
}

func TestLoanMakePaymentNotFound(t *testing.T) {
	ClearAll()

	err := loanService.MakePayment(1, 550000)

	assert.EqualError(t, err, "loan not found")
}

func TestLoanMakePaymentAlreadyPaid(t *testing.T) {
	ClearAll()
	load := CreatePaidLoan()

	err := loanService.MakePayment(load.ID, 550000)

	assert.EqualError(t, err, "loan already paid")
}

func TestLoanMakePaymentInvalidPaymentAmount(t *testing.T) {
	ClearAll()
	load := CreateLoanAndOutstandingLoanPayments()

	err := loanService.MakePayment(load.ID, 100000)

	assert.EqualError(t, err, "invalid payment amount")
}

func TestLoanMakePayment(t *testing.T) {
	ClearAll()
	load := CreateLoanAndOutstandingLoanPayments()

	err := loanService.MakePayment(load.ID, 110000)

	assert.Nil(t, err)
}
