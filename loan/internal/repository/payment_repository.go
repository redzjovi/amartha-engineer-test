package repository

import (
	"loan/internal/entity"
)

type PaymentRepository struct {
	Repository[entity.Payment]
}

func NewPaymentRepository() *PaymentRepository {
	return &PaymentRepository{}
}
