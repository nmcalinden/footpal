package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/config"
	"github.com/nmcalinden/footpal/controllers"
	"github.com/nmcalinden/footpal/service"
)

func ConfigureBookingHandlers(app *fiber.App) {
	group := app.Group("/bookings")

	bService := service.NewBookingService(config.GetConnection())
	bookController := controllers.NewBookingController(bService)

	group.Get("/", bookController.RetrieveBookings)
	group.Post("/", bookController.CreateBooking)
	group.Get("/:bookingId", bookController.GetBookingById)
	group.Put("/:bookingId", bookController.UpdateBooking)
	group.Delete("/:bookingId", bookController.CancelBooking)
	group.Get("/:bookingId/matches", bookController.GetMatchesByBooking)
}
