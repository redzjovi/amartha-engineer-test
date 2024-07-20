package repository

import (
	"billing/internal/entity"
	"sync"
	"time"
)

type LoanPaymentRepository interface {
	Create(loan *entity.LoanPayment)
	ListOutstandingByLoanID(loanID int) (res []entity.LoanPayment)
	Save(loan *entity.LoanPayment)
	Truncate()
}

type InMemoryLoanPaymentRepository struct {
	loanPayments map[int]*entity.LoanPayment
	mu           sync.Mutex
	nextID       int
}

func NewInMemoryLoanPaymentRepository() *InMemoryLoanPaymentRepository {
	return &InMemoryLoanPaymentRepository{
		loanPayments: make(map[int]*entity.LoanPayment),
		nextID:       1,
	}
}

func (r *InMemoryLoanPaymentRepository) Create(loan *entity.LoanPayment) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if loan.ID == 0 {
		loan.ID = r.nextID
		r.nextID++
	}

	r.loanPayments[loan.ID] = loan
}

func (r *InMemoryLoanPaymentRepository) ListOutstandingByLoanID(loanID int) (res []entity.LoanPayment) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, loanPayment := range r.loanPayments {
		if loanPayment.LoanID != loanID {
			continue
		}
		if loanPayment.EndAt.After(time.Now()) {
			continue
		}
		if loanPayment.PaidAt != nil {
			continue
		}
		res = append(res, *loanPayment)
	}

	return res
}

func (r *InMemoryLoanPaymentRepository) Save(loan *entity.LoanPayment) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.loanPayments[loan.ID] = loan
}

func (r *InMemoryLoanPaymentRepository) Truncate() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.loanPayments = make(map[int]*entity.LoanPayment)
	r.nextID = 1
}
