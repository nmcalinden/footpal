package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/controllers"
)

func ConfigureHealthHandler(app *fiber.App) {
	healthController := controllers.HealthController{}
	app.Get("/health", healthController.HandleHealthCheck())
}
