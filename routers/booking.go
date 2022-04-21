package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/config"
	"github.com/nmcalinden/footpal/controllers"
	"github.com/nmcalinden/footpal/middleware"
	"github.com/nmcalinden/footpal/services"
)

func ConfigureBookingHandlers(app *fiber.App) {
	group := app.Group("/bookings", middleware.IsAuthenticated)

	bService := services.NewBookingService(config.GetConnection())
	bookController := controllers.NewBookingController(bService)

	roles := []middleware.UserRole{{R: "player"}, {R: "venueAdmin"}}
	group.Use(middleware.NewRoles(roles).HasRole)
	group.Get("/", bookController.RetrieveBookings)
	group.Post("/", bookController.CreateBooking)
	group.Get("/:bookingId", bookController.GetBookingById)
	group.Put("/:bookingId", bookController.UpdateBooking)
	group.Delete("/:bookingId", bookController.CancelBooking)
	group.Get("/:bookingId/matches", bookController.GetMatchesByBooking)
}
