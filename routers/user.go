package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/config"
	"github.com/nmcalinden/footpal/controllers"
	"github.com/nmcalinden/footpal/middleware"
	"github.com/nmcalinden/footpal/services"
)

func ConfigureUserHandlers(app *fiber.App) {
	group := app.Group("/")

	uService := services.NewUserService(config.GetConnection())
	userController := controllers.NewUserController(uService)

	group.Post("/login", userController.LoginHandler)
	group.Post("/register", userController.RegisterHandler)
	group.Post("/refresh", userController.RefreshToken)

	roles := []middleware.UserRole{{R: "player"}, {R: "venueAdmin"}}
	group.Delete("/deactivate", middleware.IsAuthenticated, middleware.NewRoles(roles).HasRole, userController.DeactivateHandler)
}
