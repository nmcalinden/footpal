package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/enums"
	"github.com/nmcalinden/footpal/middleware"
)

func ConfigureUserHandlers(app *fiber.App) {
	group := app.Group("/")

	userController := InitializeUserController()

	group.Post("/login", userController.LoginHandler)
	group.Post("/register", userController.RegisterHandler)
	group.Post("/refresh", userController.RefreshToken)

	group.Get("/me", middleware.IsAuthenticated, middleware.NewRoles(enums.All).HasRole, userController.UserHandler)
	group.Delete("/deactivate", middleware.IsAuthenticated, middleware.NewRoles(enums.Player).HasRole, userController.DeactivateHandler)
}
