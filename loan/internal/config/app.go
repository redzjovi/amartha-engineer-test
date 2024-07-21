package config

import (
	"loan/internal/delivery/http"
	"loan/internal/delivery/http/middleware"
	"loan/internal/delivery/http/route"
	"loan/internal/repository"
	"loan/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	App      *fiber.App
	DB       *gorm.DB
	Log      *logrus.Logger
	Validate *validator.Validate
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	loanApprovalRepository := repository.NewLoanApprovalRepository()
	loanDisbursementRepository := repository.NewLoanDisbursementRepository()
	loanInvestmentRepository := repository.NewLoanInvestmentRepository()
	loanRepository := repository.NewLoanRepository()
	paymentRepository := repository.NewPaymentRepository()
	userRepository := repository.NewUserRepository()
	userRoleRepository := repository.NewUserRoleRepository()

	// setup use cases
	authUsecase := usecase.NewAuthUsecase(config.DB, config.Log, userRepository, userRoleRepository, config.Validate)
	loanUsecase := usecase.NewLoanUsecase(config.DB, loanRepository, loanApprovalRepository, loanDisbursementRepository, loanInvestmentRepository, config.Log, paymentRepository, config.Validate)
	paymentUsecase := usecase.NewPaymentUsecase(config.DB, loanInvestmentRepository, config.Log, paymentRepository, config.Validate)

	// setup controller
	adminLoanController := http.NewAdminLoanController(loanUsecase, config.Log)
	externalPaymentController := http.NewExternalPaymentController(config.Log, paymentUsecase)
	guestAuthController := http.NewGuestAuthController(authUsecase, config.Log)
	guestLoanController := http.NewGuestLoanController(loanUsecase, config.Log)
	userAuthController := http.NewUserAuthController(authUsecase, config.Log)
	userLoanController := http.NewUserLoanController(loanUsecase, config.Log)

	// setup middleware
	adminMiddleware := middleware.NewAdmin(authUsecase)
	userMiddleware := middleware.NewUser(authUsecase)

	routeConfig := route.RouteConfig{
		AdminLoanController:       adminLoanController,
		AdminMiddleware:           adminMiddleware,
		App:                       config.App,
		ExternalPaymentController: externalPaymentController,
		GuestAuthController:       guestAuthController,
		GuestLoanController:       guestLoanController,
		UserAuthController:        userAuthController,
		UserLoanController:        userLoanController,
		UserMiddleware:            userMiddleware,
	}
	routeConfig.Setup()
}
