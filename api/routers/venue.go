package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/enums"
	"github.com/nmcalinden/footpal/middleware"
)

func ConfigureVenueHandlers(app *fiber.App) {
	group := app.Group("/venues")

	venueController := InitializeVenueController()

	group.Get("/", venueController.RetrieveVenues)
	group.Get("/summary", venueController.RetrieveVenueSummaries)
	group.Get("/:venueId", venueController.RetrieveVenueById)
	group.Get("/:venueId/hours", venueController.RetrieveVenueOpeningHours)
	group.Get("/:venueId/pitches", venueController.RetrievePitchesByVenue)
	group.Get("/:venueId/pitches/:pitchId", venueController.RetrievePitch)
	group.Get("/:venueId/pitches/:pitchId/timeslots", venueController.RetrievePitchTimeSlots)
	group.Get("/:venueId/timeslots", venueController.RetrieveVenueTimeSlots)
	group.Get("/:venueId/timeslots/:pitchTimeslotId/pitch", venueController.RetrievePitchByTimeSlot)

	group.Use(middleware.IsAuthenticated)
	group.Use(middleware.NewRoles(enums.VenueAdmin).HasRole)

	group.Post("/", venueController.CreateVenue)
	group.Put("/:venueId", venueController.UpdateVenue)
	group.Put("/:venueId", venueController.DeleteVenue)
	group.Get("/:venueId/admins", venueController.RetrieveVenueAdmins)
	group.Post("/:venueId/admins", venueController.AddAdminToVenue)
	group.Delete("/:venueId/admins/:adminId", venueController.RemoveAdminFromVenue)
	group.Post("/:venueId/pitches", venueController.AddPitchToVenue)
	group.Put("/:venueId/pitches/:pitchId", venueController.UpdatePitchInfo)
	group.Delete("/:venueId/pitches/:pitchId", venueController.RemovePitch)

}
