package venueRoute

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/models"
	"github.com/nmcalinden/footpal/utils"
	"strconv"
)

type VenueAdminRequest struct {
	VenueId int
	UserId  string
}

type VenueAdminResponse struct {
	VenueAdminId int
	Forename     string
	Surname      string
	Email        string
}

func ConfigureVenueHandlers(app *fiber.App) {
	group := app.Group("/venues")

	group.Get("/", retrieveVenues)
	group.Post("/", createVenue)
	group.Get("/:venueId", retrieveVenueById)
	group.Put("/:venueId", updateVenue)
	group.Put("/:venueId", deleteVenue)
	group.Get("/:venueId/admins", retrieveVenueAdmins)
	group.Post("/:venueId/admins", addAdminToVenue)
	group.Delete("/:venueId/admins/:adminId", removeAdminFromVenue)
	group.Get("/:venueId/pitches", retrievePitchesByVenue)
	group.Post("/:venueId/pitches", addPitchToVenue)
	group.Get("/:venueId/pitches/:pitchId", retrievePitch)
	group.Put("/:venueId/pitches/:pitchId", updatePitchInfo)
	group.Delete("/:venueId/pitches/:pitchId", removePitch)
	group.Get("/:venueId/pitches/:pitchId/timeslots", retrievePitchTimeSlots)
	group.Get("/:venueId/timeslots", retrieveVenueTimeSlots)

}

// @Summary      Retrieve Venues
// @Description  Retrieve all venues
// @Tags         venue
// @Produce      json
// @Success      200  {array}  models.Venue
// @Router       /venues [get]
func retrieveVenues(c *fiber.Ctx) error {
	p := MockVenues
	return c.Status(fiber.StatusOK).JSON(p)
}

// @Summary      Retrieve Venues by id
// @Description  Retrieve venue by venueId
// @Tags         venue
// @Produce      json
// @Success      200  {object}  models.Venue
// @Router       /venues/{venueId} [get]
func retrieveVenueById(c *fiber.Ctx) error {
	venueId, err := strconv.Atoi(c.Params("venueId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "VenueId supplied is invalid")
	}

	p := MockVenues
	result := models.Venue{}
	for _, s := range p {
		if s.VenueId == venueId {
			result = s
		}
	}

	if result.VenueId == 0 {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Venue does not exist")
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

// @Summary      Create new venue
// @Description  Create new football venue
// @Tags         venue
// @Produce      json
// @Param 		 message body models.VenueRequest true "Request"
// @Success      201 {string} string venueId
// @Failure      400 {object} utils.ErrorResponse
// @Router       /venues [post]
func createVenue(c *fiber.Ctx) error {
	newVenue := new(models.Venue)
	newVenue.VenueId = 4

	if err := c.BodyParser(&newVenue); err != nil {
		return err
	}

	if errors := utils.ValidateStruct(*newVenue); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return c.Status(fiber.StatusCreated).JSON(newVenue.VenueId)
}

// @Summary      Edit venue
// @Description  Edit venue details
// @Tags         venue
// @Produce      json
// @Param 		 message body models.VenueRequest true "Request"
// @Param        venueId   path  int  true  "Venue ID"
// @Success      200 {object} models.Venue
// @Failure      400 {object} utils.ErrorResponse
// @Router       /venues [put]
func updateVenue(c *fiber.Ctx) error {
	venueId, err := strconv.Atoi(c.Params("venueId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "VenueId supplied is not a number")
	}

	venue := new(models.Venue)
	venue.VenueId = venueId

	if err := c.BodyParser(&venue); err != nil {
		return err
	}

	if errors := utils.ValidateStruct(*venue); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return c.Status(fiber.StatusOK).JSON(venue)
}

// @Summary      Delete venue
// @Description  Remove venue from Footpal
// @Tags         venue
// @Produce      json
// @Param        venueId   path  int  true  "Venue ID"
// @Success      200
// @Failure      400 {object} utils.ErrorResponse
// @Router       /venues/{venueId} [delete]
func deleteVenue(c *fiber.Ctx) error {
	venueId, err := strconv.Atoi(c.Params("venueId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "VenueId supplied is not a number")
	}

	return c.Status(fiber.StatusOK).JSON(venueId)
}

// @Summary      Get Venue Admins
// @Description  Retrieve venue administrators
// @Tags         venue
// @Produce      json
// @Param        venueId   path  int  true  "Venue ID"
// @Success      200 {array} VenueAdminResponse
// @Failure      400 {object} utils.ErrorResponse
// @Router       /venues/{venueId}/admins [get]
func retrieveVenueAdmins(c *fiber.Ctx) error {
	var venueAdmins = []VenueAdminResponse{
		{
			VenueAdminId: 1,
			Forename:     "Test",
			Surname:      "Test",
			Email:        "test@test.com",
		},
		{
			VenueAdminId: 2,
			Forename:     "Richard",
			Surname:      "Test",
			Email:        "richard@test.com",
		},
	}
	return c.Status(fiber.StatusOK).JSON(venueAdmins)
}

// @Summary      Add Venue Admin
// @Description  Add new administrator to venue
// @Tags         venue
// @Produce      json
// @Param 		 message body VenueAdminRequest true "Request"
// @Param        venueId   path  int  true  "Venue ID"
// @Success      200 {array} VenueAdminResponse
// @Failure      400 {object} utils.ErrorResponse
// @Router       /venues/{venueId}/admins [post]
func addAdminToVenue(c *fiber.Ctx) error {
	_, err := strconv.Atoi(c.Params("venueId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "VenueId supplied is invalid")
	}

	request := new(VenueAdminRequest)
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	if errors := utils.ValidateStruct(*request); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	admin := new(VenueAdminResponse)
	admin.VenueAdminId = 999
	return c.Status(fiber.StatusCreated).JSON(admin)
}

// @Summary      Remove Venue Admin
// @Description  Remove administrator from venue
// @Tags         venue
// @Produce      json
// @Param        venueId   path  int  true  "Venue ID"
// @Success      204
// @Failure      400 {object} utils.ErrorResponse
// @Router       /venues/{venueId}/admins [delete]
func removeAdminFromVenue(c *fiber.Ctx) error {
	_, err := strconv.Atoi(c.Params("venueId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "VenueId supplied is invalid")
	}

	_, err = strconv.Atoi(c.Params("adminId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Admin Id supplied is invalid")
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// @Summary      Get Pitches by Venue
// @Description  Get Pitches by Venue
// @Tags         venue
// @Produce      json
// @Param        venueId   path  int  true  "Venue ID"
// @Success      200 {array} models.Pitch
// @Failure      400 {object} utils.ErrorResponse
// @Router       /venues/{venueId}/pitches [get]
func retrievePitchesByVenue(c *fiber.Ctx) error {
	_, err := strconv.Atoi(c.Params("venueId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "VenueId supplied is invalid")
	}
	p := MockVenuePitches
	return c.Status(fiber.StatusOK).JSON(p)
}

// @Summary      Add Venue Pitch
// @Description  Add new pitch to existing venue
// @Tags         venue
// @Produce      json
// @Param 		 message body models.PitchRequest true "Request"
// @Param        venueId   path  int  true  "Venue ID"
// @Success      200 {array} models.Pitch
// @Failure      400 {object} utils.ErrorResponse
// @Router       /venues/{venueId}/pitches [post]
func addPitchToVenue(c *fiber.Ctx) error {
	venueId, err := strconv.Atoi(c.Params("venueId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "VenueId supplied is invalid")
	}

	pitch := new(models.PitchRequest)
	if err := c.BodyParser(&pitch); err != nil {
		return err
	}

	if errors := utils.ValidateStruct(*pitch); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	response := models.Pitch{
		PitchId:    999,
		VenueId:    venueId,
		Name:       pitch.Name,
		MaxPlayers: pitch.MaxPlayers,
		Cost:       pitch.Cost,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// @Summary      Get Venue Pitch
// @Description  Get pitch info by venue
// @Tags         venue
// @Produce      json
// @Param        venueId   path  int  true  "Venue ID"
// @Success      200 {object} models.Pitch
// @Failure      400 {object} utils.ErrorResponse
// @Router       /venues/{venueId}/pitches/{pitchId} [get]
func retrievePitch(c *fiber.Ctx) error {
	venueId, err := strconv.Atoi(c.Params("venueId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "VenueId supplied is invalid")
	}

	pitchId, err := strconv.Atoi(c.Params("pitchId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Pitch Id supplied is invalid")
	}

	p := MockVenuePitches

	result := models.Pitch{}
	for _, s := range p {
		if s.VenueId == venueId && s.PitchId == pitchId {
			result = s
			break
		}
	}
	if result.PitchId == 0 {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Pitch does not exist")
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

// @Summary      Edit Pitch
// @Description  Edit pitch details
// @Tags         venue
// @Produce      json
// @Param 		 message body models.PitchRequest true "Request"
// @Param        venueId   path  int  true  "Venue ID"
// @Param        pitchId   path  int  true  "Pitch ID"
// @Success      200 {object} models.Pitch
// @Failure      400 {object} utils.ErrorResponse
// @Router       /venues/{venueId}/pitches/{pitchId} [put]
func updatePitchInfo(c *fiber.Ctx) error {
	venueId, err := strconv.Atoi(c.Params("venueId"))
	pitchId, err2 := strconv.Atoi(c.Params("pitchId"))

	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "VenueId supplied is not a number")
	}

	if err2 != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "PitchId supplied is not a number")
	}

	pitch := new(models.PitchRequest)
	if err := c.BodyParser(&pitch); err != nil {
		return err
	}

	if errors := utils.ValidateStruct(*pitch); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	response := models.Pitch{
		PitchId:    pitchId,
		VenueId:    venueId,
		Name:       pitch.Name,
		MaxPlayers: pitch.MaxPlayers,
		Cost:       pitch.Cost,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

// @Summary      Delete Pitch
// @Description  Remove pitch from venue
// @Tags         venue
// @Produce      json
// @Param        venueId   path  int  true  "Venue ID"
// @Param        pitchId   path  int  true  "Pitch ID"
// @Success      204
// @Failure      400 {object} utils.ErrorResponse
// @Router       /venues/{venueId}/pitches/{pitchId} [delete]
func removePitch(c *fiber.Ctx) error {
	_, err := strconv.Atoi(c.Params("venueId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "VenueId supplied is invalid")
	}

	_, err = strconv.Atoi(c.Params("pitchId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "PitchId supplied is invalid")
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// @Summary      Retrieve Pitch time slots
// @Description  Retrieve all time slots by Pitch
// @Tags         venue
// @Produce      json
// @Param        venueId   path  int  true  "Venue ID"
// @Param        pitchId   path  int  true  "Pitch ID"
// @Success      200
// @Failure      400 {object} utils.ErrorResponse
// @Router       /venues/{venueId}/pitches/{pitchId}/timeslots [get]
func retrievePitchTimeSlots(c *fiber.Ctx) error {
	_, err := strconv.Atoi(c.Params("venueId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "VenueId supplied is invalid")
	}
	return c.SendStatus(fiber.StatusOK)
}

// @Summary      Retrieve Venue time slots
// @Description  Retrieve all time slots by Venue
// @Tags         venue
// @Produce      json
// @Param        venueId   path  int  true  "Venue ID"
// @Success      200
// @Failure      400 {object} utils.ErrorResponse
// @Router       /venues/{venueId}/timeslots [get]
func retrieveVenueTimeSlots(c *fiber.Ctx) error {
	_, err := strconv.Atoi(c.Params("venueId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "VenueId supplied is invalid")
	}
	return c.SendStatus(fiber.StatusOK)
}
