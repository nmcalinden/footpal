package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
	"github.com/nmcalinden/footpal/config"
	"github.com/nmcalinden/footpal/router"
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
// @BasePath /
func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(logger.New())

	config.InitializeDatabase()

	utils.ConfigureSwagger(app)

	router.ConfigureUserHandlers(app)
	router.ConfigureVenueHandlers(app)
	router.ConfigurePlayerHandlers(app)
	router.ConfigureSquadPlayers(app)
	router.ConfigureBookingHandlers(app)
	router.ConfigureMatchHandlers(app)

	log.Fatal(app.Listen(":3000"))
}
