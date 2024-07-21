package entity

import "time"

type PaymentProductType string

const (
	PaymentProductTypeLoanInvestment PaymentProductType = "loan_investment"
)

type Payment struct {
	ID          uint               `gorm:"column:id;primaryKey"`
	ProductType PaymentProductType `gorm:"column:product_type"`
	ProductID   uint               `gorm:"column:product_id"`
	Amount      float64            `gorm:"column:amount"`
	CreatedAt   time.Time          `gorm:"column:created_at"`
	UpdatedAt   time.Time          `gorm:"column:updated_at"`
	ExpiredAt   time.Time          `gorm:"column:expired_at"`
	PaidAt      *time.Time         `gorm:"column:paid_at"`
}

func (p *Payment) TableName() string {
	return "payments"
}
