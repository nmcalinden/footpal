package main

import (
	"github.com/nmcalinden/footpal/docs"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
	"github.com/nmcalinden/footpal/config"
	"github.com/nmcalinden/footpal/routers"
	"github.com/nmcalinden/footpal/utils"
)

// @title Footpal API
// @version 1.0
// @description This is swagger for Footpal
// @termsOfService http://swagger.io/terms/
// @contact.name Footpal Support
// @contact.email nathan.mcalinden@unosquare.com
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	docs.SwaggerInfo.Host = getHost()

	fiberConfig := fiber.Config{
		ErrorHandler: defaultErrorHandler,
	}
	app := fiber.New(fiberConfig)

	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(logger.New())
	app.Use(recover.New(recover.Config{EnableStackTrace: true}))

	config.InitializeDatabase()

	utils.ConfigureSwagger(app)

	routers.ConfigureHealthHandler(app)
	routers.ConfigureUserHandlers(app)
	routers.ConfigureVenueHandlers(app)
	routers.ConfigurePlayerHandlers(app)
	routers.ConfigureSquadPlayers(app)
	routers.ConfigureBookingHandlers(app)
	routers.ConfigureMatchHandlers(app)

	log.Fatal(app.Listen(":8081"))
}

func getHost() string {
	appHost := os.Getenv("API_HOST")
	if appHost != "" {
		return appHost
	}

	return "127.0.0.1:8081"
}

func defaultErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := err.Error()

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	errors := map[string]string{"reason": message}
	return utils.BuildMultipleErrorResponse(ctx, code, "Something has gone wrong", errors)
}
