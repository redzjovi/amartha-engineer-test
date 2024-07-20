package repository

import (
	"billing/internal/entity"
	"sync"
)

type LoanRepository interface {
	Create(loan *entity.Loan)
	FindByID(id int) (loan *entity.Loan, err error)
	Save(loan *entity.Loan)
	Truncate()
}

type InMemoryLoanRepository struct {
	loans  map[int]*entity.Loan
	mu     sync.Mutex
	nextID int
}

func NewInMemoryLoanRepository() *InMemoryLoanRepository {
	return &InMemoryLoanRepository{
		loans:  make(map[int]*entity.Loan),
		nextID: 1,
	}
}

func (r *InMemoryLoanRepository) Create(loan *entity.Loan) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if loan.ID == 0 {
		loan.ID = r.nextID
		r.nextID++
	}

	r.loans[loan.ID] = loan
}

func (r *InMemoryLoanRepository) FindByID(id int) (*entity.Loan, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	loan, exists := r.loans[id]
	if !exists {
		return nil, nil
	}

	return loan, nil
}

func (r *InMemoryLoanRepository) Save(loan *entity.Loan) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.loans[loan.ID] = loan
}

func (r *InMemoryLoanRepository) Truncate() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.loans = make(map[int]*entity.Loan)
	r.nextID = 1
}
