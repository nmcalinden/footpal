package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/config"
	"github.com/nmcalinden/footpal/controllers"
	"github.com/nmcalinden/footpal/services"
)

func ConfigureUserHandlers(app *fiber.App) {
	group := app.Group("/")

	uService := services.NewUserService(config.GetConnection())
	userController := controllers.NewUserController(uService)

	group.Post("/login", userController.LoginHandler)
	group.Post("/register", userController.RegisterHandler)
	group.Delete("/deactivate/:userId", userController.DeactivateHandler)
}
