package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/config"
	"github.com/nmcalinden/footpal/controllers"
	"github.com/nmcalinden/footpal/services"
)

func ConfigureVenueHandlers(app *fiber.App) {
	group := app.Group("/venues")

	vService := services.NewVenueService(config.GetConnection())
	venueController := controllers.NewVenueController(vService)

	group.Get("/", venueController.RetrieveVenues)
	group.Post("/", venueController.CreateVenue)
	group.Get("/:venueId", venueController.RetrieveVenueById)
	group.Put("/:venueId", venueController.UpdateVenue)
	group.Put("/:venueId", venueController.DeleteVenue)
	group.Get("/:venueId/admins", venueController.RetrieveVenueAdmins)
	group.Post("/:venueId/admins", venueController.AddAdminToVenue)
	group.Delete("/:venueId/admins/:adminId", venueController.RemoveAdminFromVenue)
	group.Get("/:venueId/pitches", venueController.RetrievePitchesByVenue)
	group.Post("/:venueId/pitches", venueController.AddPitchToVenue)
	group.Get("/:venueId/pitches/:pitchId", venueController.RetrievePitch)
	group.Put("/:venueId/pitches/:pitchId", venueController.UpdatePitchInfo)
	group.Delete("/:venueId/pitches/:pitchId", venueController.RemovePitch)
	group.Get("/:venueId/pitches/:pitchId/timeslots", venueController.RetrievePitchTimeSlots)
	group.Get("/:venueId/timeslots", venueController.RetrieveVenueTimeSlots)

}
