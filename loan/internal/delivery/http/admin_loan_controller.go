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

type AdminLoanController struct {
	LoanUsecase *usecase.LoanUsecase
	Log         *logrus.Logger
}

func NewAdminLoanController(
	loanUsecase *usecase.LoanUsecase,
	log *logrus.Logger,
) *AdminLoanController {
	return &AdminLoanController{
		LoanUsecase: loanUsecase,
		Log:         log,
	}
}

func (c *AdminLoanController) Approve(ctx *fiber.Ctx) error {
	auth := middleware.GetAdmin(ctx)

	loanIdUint64, err := strconv.ParseUint(ctx.Params("loanId"), 10, 64)
	if err != nil {
		return fiber.ErrBadRequest
	}

	request := new(model.LoanApproveRequest)
	if err := ctx.BodyParser(request); err != nil {
		return fiber.ErrBadRequest
	}

	if err := c.LoanUsecase.Approve(ctx.UserContext(), uint(loanIdUint64), auth.ID, request); err != nil {
		return err
	}

	return ctx.SendStatus(http.StatusNoContent)
}

func (c *AdminLoanController) Disburse(ctx *fiber.Ctx) error {
	auth := middleware.GetAdmin(ctx)

	loanIdUint64, err := strconv.ParseUint(ctx.Params("loanId"), 10, 64)
	if err != nil {
		return fiber.ErrBadRequest
	}

	request := new(model.LoanDisburseRequest)
	if err := ctx.BodyParser(request); err != nil {
		return fiber.ErrBadRequest
	}

	if err := c.LoanUsecase.Disburse(ctx.UserContext(), uint(loanIdUint64), auth.ID, request); err != nil {
		return err
	}

	return ctx.SendStatus(http.StatusNoContent)
}

func (c *AdminLoanController) List(ctx *fiber.Ctx) error {
	response, err := c.LoanUsecase.List(ctx.UserContext())
	if err != nil {
		return err
	}

	return ctx.JSON(model.WebResponse[[]model.LoanResponse]{Data: response})
}
