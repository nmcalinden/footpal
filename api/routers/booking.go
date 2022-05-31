package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/api/enums"
	"github.com/nmcalinden/footpal/api/middleware"
)

func ConfigureBookingHandlers(app *fiber.App) {
	group := app.Group("/bookings")

	bookController := InitializeBookingController()

	group.Use(middleware.IsAuthenticated)
	group.Use(middleware.NewRoles(enums.Player, enums.VenueAdmin).HasRole)
	group.Get("/", bookController.RetrieveBookings)
	group.Post("/", bookController.CreateBooking)
	group.Post("/search", bookController.FindAvailableSlotsByVenue)
	group.Get("/:bookingId", bookController.GetBookingById)
	group.Put("/:bookingId", bookController.UpdateBooking)
	group.Delete("/:bookingId", bookController.CancelBooking)
	group.Get("/:bookingId/matches", bookController.GetMatchesByBooking)
}
