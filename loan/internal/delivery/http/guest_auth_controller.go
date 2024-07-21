package http

import (
	"loan/internal/model"
	"loan/internal/usecase"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type GuestAuthController struct {
	AuthUsecase *usecase.AuthUsecase
	Log         *logrus.Logger
}

func NewGuestAuthController(
	authUsecase *usecase.AuthUsecase,
	log *logrus.Logger,
) *GuestAuthController {
	return &GuestAuthController{
		AuthUsecase: authUsecase,
		Log:         log,
	}
}

func (c *GuestAuthController) Login(ctx *fiber.Ctx) error {
	request := new(model.AuthLoginRequest)
	if err := ctx.BodyParser(request); err != nil {
		return fiber.ErrBadRequest
	}

	response, err := c.AuthUsecase.Login(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(model.WebResponse[*model.AuthLoginResponse]{Data: response})
}

func (c *GuestAuthController) SignUp(ctx *fiber.Ctx) error {
	request := new(model.AuthSignUpRequest)
	if err := ctx.BodyParser(request); err != nil {
		return fiber.ErrBadRequest
	}

	if err := c.AuthUsecase.SignUp(ctx.UserContext(), request); err != nil {
		return err
	}

	return ctx.SendStatus(http.StatusNoContent)
}
