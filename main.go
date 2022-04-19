package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
	"github.com/nmcalinden/footpal/config"
	"github.com/nmcalinden/footpal/routers"
	"github.com/nmcalinden/footpal/utils"
	"log"
)

// @title Footpal API
// @version 1.0
// @description This is swagger for Footpal
// @termsOfService http://swagger.io/terms/
// @contact.name Footpal Support
// @contact.email nathan.mcalinden@unosquare.com
// @host 127.0.0.1:3000
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(logger.New())

	config.InitializeDatabase()

	utils.ConfigureSwagger(app)

	routers.ConfigureUserHandlers(app)
	routers.ConfigureVenueHandlers(app)
	routers.ConfigurePlayerHandlers(app)
	routers.ConfigureSquadPlayers(app)
	routers.ConfigureBookingHandlers(app)
	routers.ConfigureMatchHandlers(app)

	log.Fatal(app.Listen(":3000"))
}
