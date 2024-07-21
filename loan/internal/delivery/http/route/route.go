package route

import (
	"loan/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	AdminLoanController       *http.AdminLoanController
	AdminMiddleware           fiber.Handler
	App                       *fiber.App
	ExternalPaymentController *http.ExternalPaymentController
	GuestAuthController       *http.GuestAuthController
	GuestLoanController       *http.GuestLoanController
	UserAuthController        *http.UserAuthController
	UserLoanController        *http.UserLoanController
	UserMiddleware            fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupAdminRoute()
	c.SetupExternalRoute()
	c.SetupGuestRoute()
	c.SetupUserRoute()
}

func (c *RouteConfig) SetupAdminRoute() {
	admin := c.App.Group("/api/admin", c.AdminMiddleware)
	admin.Get("/loan", c.AdminLoanController.List)
	admin.Post("/loan/:loanId/approve", c.AdminLoanController.Approve)
	admin.Post("/loan/:loanId/disburse", c.AdminLoanController.Disburse)
}

func (c *RouteConfig) SetupExternalRoute() {
	external := c.App.Group("/api/external")
	external.Post("/payment/:paymentId", c.ExternalPaymentController.Pay)
}

func (c *RouteConfig) SetupGuestRoute() {
	guest := c.App.Group("/api")
	guest.Post("/auth/login", c.GuestAuthController.Login)
	guest.Post("/auth/sign-up", c.GuestAuthController.SignUp)
	guest.Get("/loan", c.GuestLoanController.List)
}

func (c *RouteConfig) SetupUserRoute() {
	user := c.App.Group("/api/user", c.UserMiddleware)
	user.Delete("/auth", c.UserAuthController.Logout)
	user.Post("/loan/propose", c.UserLoanController.Propose)
	user.Post("/loan/:loanId/invest", c.UserLoanController.Invest)
}
