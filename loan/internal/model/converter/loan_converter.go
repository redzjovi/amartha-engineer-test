package converter

import (
	"loan/internal/entity"
	"loan/internal/model"
)

func LoanToResponse(loan *entity.Loan) *model.LoanResponse {
	return &model.LoanResponse{
		ID:              loan.ID,
		BorrowerID:      loan.BorrowerID,
		PrincipalAmount: loan.PrincipalAmount,
		Rate:            loan.Rate,
		ROI:             loan.ROI,
		State:           loan.State,
		TotalInvested:   loan.TotalInvested,
		CreatedAt:       loan.CreatedAt,
		UpdatedAt:       loan.UpdatedAt,
	}
}
