package http

import (
	"loan/internal/delivery/http/middleware"
	"loan/internal/model"
	"loan/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserLoanController struct {
	LoanUsecase *usecase.LoanUsecase
	Log         *logrus.Logger
}

func NewUserLoanController(
	loanUsecase *usecase.LoanUsecase,
	log *logrus.Logger,
) *UserLoanController {
	return &UserLoanController{
		LoanUsecase: loanUsecase,
		Log:         log,
	}
}

func (c *UserLoanController) Invest(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	loanIdUint64, err := strconv.ParseUint(ctx.Params("loanId"), 10, 64)
	if err != nil {
		return fiber.ErrBadRequest
	}

	request := new(model.LoanInvestRequest)
	if err := ctx.BodyParser(request); err != nil {
		return fiber.ErrBadRequest
	}

	if err := c.LoanUsecase.Invest(ctx.UserContext(), uint(loanIdUint64), auth.ID, request); err != nil {
		return err
	}

	return ctx.SendStatus(http.StatusNoContent)
}

func (c *UserLoanController) Propose(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.LoanProposeRequest)
	if err := ctx.BodyParser(request); err != nil {
		return fiber.ErrBadRequest
	}

	if err := c.LoanUsecase.Propose(ctx.UserContext(), auth.ID, request); err != nil {
		return err
	}

	return ctx.SendStatus(http.StatusNoContent)
}
