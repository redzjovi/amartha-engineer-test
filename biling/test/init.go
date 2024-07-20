package test

import (
	"billing/internal/repository"
	"billing/internal/service"
)

var loanRepository repository.LoanRepository
var loanPaymentRepository repository.LoanPaymentRepository

var loanService service.LoanService

func init() {
	loanRepository = repository.NewInMemoryLoanRepository()
	loanPaymentRepository = repository.NewInMemoryLoanPaymentRepository()

	loanService = service.NewLoanService(loanRepository, loanPaymentRepository)
}
