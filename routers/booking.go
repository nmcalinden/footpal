package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/middleware"
)

func ConfigureBookingHandlers(app *fiber.App) {
	group := app.Group("/bookings", middleware.IsAuthenticated)

	bookController := InitializeBookingController()

	roles := []middleware.UserRole{{Role: "player"}, {Role: "venueAdmin"}}
	group.Use(middleware.NewRoles(roles).HasRole)
	group.Get("/", bookController.RetrieveBookings)
	group.Post("/", bookController.CreateBooking)
	group.Get("/:bookingId", bookController.GetBookingById)
	group.Put("/:bookingId", bookController.UpdateBooking)
	group.Delete("/:bookingId", bookController.CancelBooking)
	group.Get("/:bookingId/matches", bookController.GetMatchesByBooking)
}
