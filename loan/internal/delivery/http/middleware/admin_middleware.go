package middleware

import (
	"loan/internal/model"
	"loan/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func NewAdmin(authUsecase *usecase.AuthUsecase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		request := &model.AuthVerifyIsAdminRequest{Token: ctx.Get("Authorization", "NOT_FOUND")}
		authUsecase.Log.Debugf("Authorization : %s", request.Token)

		auth, err := authUsecase.VerifyIsAdmin(ctx.UserContext(), request)
		if err != nil {
			authUsecase.Log.Warnf("Failed find user by token : %+v", err)
			return fiber.ErrUnauthorized
		}

		authUsecase.Log.Debugf("User : %+v", auth.ID)
		ctx.Locals("admin", auth)
		return ctx.Next()
	}
}

func GetAdmin(ctx *fiber.Ctx) *model.Auth {
	return ctx.Locals("admin").(*model.Auth)
}
