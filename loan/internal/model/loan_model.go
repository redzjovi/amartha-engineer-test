package model

import (
	"loan/internal/entity"
	"time"
)

type LoanApproveRequest struct {
	PictureProof string `json:"picture_proof" validate:"required"`
}

type LoanDisburseRequest struct {
	AgreementLetter string `json:"column:agreement_letter" validate:"required"`
}

type LoanInvestRequest struct {
	Amount float64 `json:"amount" validate:"required"`
}

type LoanProposeRequest struct {
	PrincipalAmount float64 `json:"principal_amount" validate:"required"`
	Rate            float64 `json:"rate" validate:"required"`
}

type LoanResponse struct {
	ID              uint             `json:"column:id"`
	BorrowerID      uint             `json:"column:borrower_id"`
	PrincipalAmount float64          `json:"column:principal_amount"`
	Rate            float64          `json:"column:rate"`
	ROI             float64          `json:"column:roi"`
	AgreementLetter string           `json:"column:agreement_letter"`
	State           entity.LoanState `json:"column:state"`
	TotalInvested   float64          `json:"column:total_invested"`
	CreatedAt       time.Time        `json:"column:created_at"`
	UpdatedAt       time.Time        `json:"column:updated_at"`
}
