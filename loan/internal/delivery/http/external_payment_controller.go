package http

import (
	"loan/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ExternalPaymentController struct {
	Log            *logrus.Logger
	PaymentUsecase *usecase.PaymentUsecase
}

func NewExternalPaymentController(
	log *logrus.Logger,
	paymentUsecase *usecase.PaymentUsecase,
) *ExternalPaymentController {
	return &ExternalPaymentController{
		Log:            log,
		PaymentUsecase: paymentUsecase,
	}
}

func (c *ExternalPaymentController) Pay(ctx *fiber.Ctx) error {
	paymentIdUint64, err := strconv.ParseUint(ctx.Params("loanId"), 10, 64)
	if err != nil {
		return fiber.ErrBadRequest
	}

	if err := c.PaymentUsecase.Pay(ctx.UserContext(), uint(paymentIdUint64)); err != nil {
		return err
	}

	return ctx.SendStatus(http.StatusNoContent)
}
