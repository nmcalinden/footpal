package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
	"github.com/nmcalinden/footpal/routers/booking"
	"github.com/nmcalinden/footpal/routers/match"
	"github.com/nmcalinden/footpal/routers/player"
	"github.com/nmcalinden/footpal/routers/squad"
	"github.com/nmcalinden/footpal/routers/user"
	"github.com/nmcalinden/footpal/routers/venue"
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

	utils.ConfigureSwagger(app)

	userRoute.ConfigureUserHandlers(app)
	venueRoute.ConfigureVenueHandlers(app)
	playerRoute.ConfigurePlayerHandlers(app)
	squadRoute.ConfigureSquadPlayers(app)
	bookingRoute.ConfigureBookingHandlers(app)
	matchRoute.ConfigureMatchHandlers(app)

	log.Fatal(app.Listen(":3000"))
}
