package http

import (
	"loan/internal/model"
	"loan/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type GuestLoanController struct {
	LoanUsecase *usecase.LoanUsecase
	Log         *logrus.Logger
}

func NewGuestLoanController(
	loanUsecase *usecase.LoanUsecase,
	log *logrus.Logger,
) *GuestLoanController {
	return &GuestLoanController{
		LoanUsecase: loanUsecase,
		Log:         log,
	}
}

func (c *GuestLoanController) List(ctx *fiber.Ctx) error {
	response, err := c.LoanUsecase.ListApproved(ctx.UserContext())
	if err != nil {
		return err
	}

	return ctx.JSON(model.WebResponse[[]model.LoanResponse]{Data: response})
}
