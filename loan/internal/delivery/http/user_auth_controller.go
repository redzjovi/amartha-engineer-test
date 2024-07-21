package http

import (
	"loan/internal/delivery/http/middleware"
	"loan/internal/model"
	"loan/internal/usecase"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserAuthController struct {
	AuthUsecase *usecase.AuthUsecase
	Log         *logrus.Logger
}

func NewUserAuthController(
	authUsecase *usecase.AuthUsecase,
	log *logrus.Logger,
) *UserAuthController {
	return &UserAuthController{
		AuthUsecase: authUsecase,
		Log:         log,
	}
}

func (c *UserAuthController) Logout(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.AuthLogoutRequest{
		ID: auth.ID,
	}

	if err := c.AuthUsecase.Logout(ctx.UserContext(), request); err != nil {
		return err
	}

	return ctx.SendStatus(http.StatusNoContent)
}
