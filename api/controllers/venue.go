package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/models"
	"github.com/nmcalinden/footpal/payloads"
	"github.com/nmcalinden/footpal/services"
	"github.com/nmcalinden/footpal/utils"
)

type VenueController struct {
	venueService *services.VenueService
}

func NewVenueController(venueService *services.VenueService) *VenueController {
	return &VenueController{venueService: venueService}
}

// RetrieveVenues @Summary      Retrieve Venues
// @Description  Retrieve all venues
// @Tags         venue
// @Produce      json
// @Success      200  {array}  views.Venue
// @Success      500  {object}  utils.ErrorResponse
// @Router       /venues [get]
func (con VenueController) RetrieveVenues(c *fiber.Ctx) error {
	p, err := con.venueService.GetVenues()

	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to retrieve venues")
	}
	return c.Status(fiber.StatusOK).JSON(p)
}

// RetrieveVenueSummaries @Summary  Retrieve List of Venues
// @Description  Retrieve all venues
// @Tags         venue
// @Produce      json
// @Success      200  {array}  views.VenueSummary
// @Success      500  {object}  utils.ErrorResponse
// @Router       /venues/summary [get]
func (con VenueController) RetrieveVenueSummaries(c *fiber.Ctx) error {
	p, err := con.venueService.GetVenueSummaries()

	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to retrieve venues")
	}
	return c.Status(fiber.StatusOK).JSON(p)
}

// RetrieveVenueById @Summary      Retrieve Venues by id
// @Description  Retrieve venue by venueId
// @Tags         venue
// @Produce      json
// @Param        venueId   path  int  true  "Venue ID"
// @Success      200  {object}  views.Venue
// @Success      400  {object}  utils.ErrorResponse
// @Router       /venues/{venueId} [get]
func (con VenueController) RetrieveVenueById(c *fiber.Ctx) error {
	venueId, err := strconv.Atoi(c.Params("venueId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "VenueId supplied is invalid")
	}

	result, err := con.venueService.GetVenueById(&venueId)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Venue does not exist")
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

// CreateVenue @Summary      Create new venue
// @Description  Create new football venue
// @Tags         venue
// @Produce      json
// @Param 		 message body payloads.VenueRequest true "Request"
// @Success      201 {string} string venueId
// @Failure      400 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router       /venues [post]
func (con VenueController) CreateVenue(c *fiber.Ctx) error {
	newVenue := new(payloads.VenueRequest)
	if err := c.BodyParser(&newVenue); err != nil {
		return err
	}

	if errors := utils.ValidateStruct(*newVenue); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	v, _ := con.venueService.CreateNewVenue(newVenue)
	return c.Status(fiber.StatusCreated).JSON(v)
}

// UpdateVenue @Summary      Edit venue
// @Description  Edit venue details
// @Tags         venue
// @Produce      json
// @Param 		 message body payloads.VenueRequest true "Request"
// @Param        venueId   path  int  true  "Venue ID"
// @Success      200 {object} models.Venue
// @Failure      400 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router       /venues [put]
func (con VenueController) UpdateVenue(c *fiber.Ctx) error {
	venueId, err := strconv.Atoi(c.Params("venueId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "VenueId supplied is not a number")
	}

	venue := new(models.Venue)
	venue.VenueId = &venueId

	if err := c.BodyParser(&venue); err != nil {
		return err
	}

	if errors := utils.ValidateStruct(*venue); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return c.Status(fiber.StatusOK).JSON(venue)
}

// DeleteVenue @Summary      Delete venue
// @Description  Remove venue from Footpal
// @Tags         venue
// @Produce      json
// @Param        venueId   path  int  true  "Venue ID"
// @Success      200
// @Failure      400 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router       /venues/{venueId} [delete]
func (con VenueController) DeleteVenue(c *fiber.Ctx) error {
	venueId, err := strconv.Atoi(c.Params("venueId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "VenueId supplied is not a number")
	}

	return c.Status(fiber.StatusOK).JSON(venueId)
}

// RetrieveVenueAdmins @Summary      Get Venue Admins
// @Description  Retrieve venue administrators
// @Tags         venue
// @Produce      json
// @Param        venueId   path  int  true  "Venue ID"
// @Success      200 {array} payloads.VenueAdminResponse
// @Failure      400 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router       /venues/{venueId}/admins [get]
func (con VenueController) RetrieveVenueAdmins(c *fiber.Ctx) error {
	var venueAdmins = []payloads.VenueAdminResponse{
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

// AddAdminToVenue @Summary      Add Venue Admin
// @Description  Add new administrator to venue
// @Tags         venue
// @Produce      json
// @Param 		 message body payloads.VenueAdminRequest true "Request"
// @Param        venueId   path  int  true  "Venue ID"
// @Success      200 {array} payloads.VenueAdminResponse
// @Failure      400 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router       /venues/{venueId}/admins [post]
func (con VenueController) AddAdminToVenue(c *fiber.Ctx) error {
	_, err := strconv.Atoi(c.Params("venueId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "VenueId supplied is invalid")
	}

	request := new(payloads.VenueAdminRequest)
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	if errors := utils.ValidateStruct(*request); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	admin := new(payloads.VenueAdminResponse)
	admin.VenueAdminId = 999
	return c.Status(fiber.StatusCreated).JSON(admin)
}

// RemoveAdminFromVenue @Summary      Remove Venue Admin
// @Description  Remove administrator from venue
// @Tags         venue
// @Produce      json
// @Param        venueId   path  int  true  "Venue ID"
// @Success      204
// @Failure      400 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router       /venues/{venueId}/admins [delete]
func (con VenueController) RemoveAdminFromVenue(c *fiber.Ctx) error {
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

// RetrievePitchesByVenue @Summary      Get Pitches by Venue
// @Description  Get Pitches by Venue
// @Tags         venue
// @Produce      json
// @Param        venueId   path  int  true  "Venue ID"
// @Success      200 {array} models.Pitch
// @Failure      400 {object} utils.ErrorResponse
// @Router       /venues/{venueId}/pitches [get]
func (con VenueController) RetrievePitchesByVenue(c *fiber.Ctx) error {
	_, err := strconv.Atoi(c.Params("venueId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "VenueId supplied is invalid")
	}
	var p []models.Pitch
	return c.Status(fiber.StatusOK).JSON(p)
}

// AddPitchToVenue @Summary      Add Venue Pitch
// @Description  Add new pitch to existing venue
// @Tags         venue
// @Produce      json
// @Param 		 message body payloads.PitchRequest true "Request"
// @Param        venueId   path  int  true  "Venue ID"
// @Success      200 {array} models.Pitch
// @Failure      400 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router       /venues/{venueId}/pitches [post]
func (con VenueController) AddPitchToVenue(c *fiber.Ctx) error {
	venueId, err := strconv.Atoi(c.Params("venueId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "VenueId supplied is invalid")
	}

	pitch := new(payloads.PitchRequest)
	if err := c.BodyParser(&pitch); err != nil {
		return err
	}

	if errors := utils.ValidateStruct(*pitch); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	response := models.Pitch{
		VenueId:    venueId,
		Name:       pitch.Name,
		MaxPlayers: pitch.MaxPlayers,
		Cost:       pitch.Cost,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// RetrievePitch @Summary      Get Venue Pitch
// @Description  Get pitch info by venue
// @Tags         venue
// @Produce      json
// @Param        venueId   path  int  true  "Venue ID"
// @Param        pitchId   path  int  true  "Pitch ID"
// @Success      200 {object} models.Pitch
// @Failure      400 {object} utils.ErrorResponse
// @Router       /venues/{venueId}/pitches/{pitchId} [get]
func (con VenueController) RetrievePitch(c *fiber.Ctx) error {
	venueId, err := strconv.Atoi(c.Params("venueId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "VenueId supplied is invalid")
	}

	pitchId, err := strconv.Atoi(c.Params("pitchId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Pitch Id supplied is invalid")
	}

	res, err := con.venueService.GetVenuePitch(&venueId, &pitchId)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Pitch does not exist")
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

// RetrievePitchByTimeSlot @Summary      Get Venue Pitch
// @Description  Get pitch info by venue and time slot
// @Tags         venue
// @Produce      json
// @Param        venueId   path  int  true  "Venue ID"
// @Param        pitchTimeslotId   path  int  true  "Pitch Time slot ID"
// @Success      200 {object} models.Pitch
// @Failure      400 {object} utils.ErrorResponse
// @Router       /venues/{venueId}/timeslots/{pitchTimeslotId}/pitch [get]
func (con VenueController) RetrievePitchByTimeSlot(c *fiber.Ctx) error {
	venueId, err := strconv.Atoi(c.Params("venueId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "VenueId supplied is invalid")
	}

	timeSlotId, err := strconv.Atoi(c.Params("pitchTimeslotId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Pitch Id supplied is invalid")
	}

	res, err := con.venueService.GetVenuePitchByTimeslot(&venueId, &timeSlotId)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Pitch does not exist")
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

// UpdatePitchInfo @Summary      Edit Pitch
// @Description  Edit pitch details
// @Tags         venue
// @Produce      json
// @Param 		 message body payloads.PitchRequest true "Request"
// @Param        venueId   path  int  true  "Venue ID"
// @Param        pitchId   path  int  true  "Pitch ID"
// @Success      200 {object} models.Pitch
// @Failure      400 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router       /venues/{venueId}/pitches/{pitchId} [put]
func (con VenueController) UpdatePitchInfo(c *fiber.Ctx) error {
	venueId, err := strconv.Atoi(c.Params("venueId"))
	pitchId, err2 := strconv.Atoi(c.Params("pitchId"))

	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "VenueId supplied is not a number")
	}

	if err2 != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "PitchId supplied is not a number")
	}

	pitch := new(payloads.PitchRequest)
	if err := c.BodyParser(&pitch); err != nil {
		return err
	}

	if errors := utils.ValidateStruct(*pitch); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	response := models.Pitch{
		PitchId:    &pitchId,
		VenueId:    venueId,
		Name:       pitch.Name,
		MaxPlayers: pitch.MaxPlayers,
		Cost:       pitch.Cost,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

// RemovePitch @Summary      Delete Pitch
// @Description  Remove pitch from venue
// @Tags         venue
// @Produce      json
// @Param        venueId   path  int  true  "Venue ID"
// @Param        pitchId   path  int  true  "Pitch ID"
// @Success      204
// @Failure      400 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router       /venues/{venueId}/pitches/{pitchId} [delete]
func (con VenueController) RemovePitch(c *fiber.Ctx) error {
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

// RetrievePitchTimeSlots @Summary      Retrieve Pitch time slots
// @Description  Retrieve all time slots by Pitch
// @Tags         venue
// @Produce      json
// @Param        venueId   path  int  true  "Venue ID"
// @Param        pitchId   path  int  true  "Pitch ID"
// @Success      200
// @Failure      400 {object} utils.ErrorResponse
// @Router       /venues/{venueId}/pitches/{pitchId}/timeslots [get]
func (con VenueController) RetrievePitchTimeSlots(c *fiber.Ctx) error {
	_, err := strconv.Atoi(c.Params("venueId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "VenueId supplied is invalid")
	}
	return c.SendStatus(fiber.StatusOK)
}

// RetrieveVenueTimeSlots @Summary      Retrieve Venue time slots
// @Description  Retrieve pitch slots by venue and date range
// @Tags         booking
// @Produce      json
// @Param        venueId path  int  true  "Venue ID"
// @Param        from query  string  true  "Date from - Format: YYYY-MM-DD"
// @Param        to   query  string  false  "Date to - Format: YYYY-MM-DD"
// @Success      200  {array} views.PitchBookingDetails
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Router       /venues/{venueId}/timeslots [get]
func (con VenueController) RetrieveVenueTimeSlots(c *fiber.Ctx) error {
	v, err := strconv.Atoi(c.Params("venueId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Venue ID supplied is invalid")
	}

	f := c.Query("from")
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Date From supplied is invalid")
	}

	t := c.Query("to")
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Date To supplied is invalid")
	}

	res, err := con.venueService.GetVenueTimeslots(&v, f, t)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to get bookings")
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

// RetrieveVenueOpeningHours @Summary      Retrieve Venue Opening Hours
// @Description  Retrieve opening hours by Venue
// @Tags         venue
// @Produce      json
// @Param        venueId   path  int  true  "Venue ID"
// @Success      200 {array} views.VenueOpeningHour
// @Failure      400 {object} utils.ErrorResponse
// @Router       /venues/{venueId}/hours [get]
func (con VenueController) RetrieveVenueOpeningHours(c *fiber.Ctx) error {
	v, err := strconv.Atoi(c.Params("venueId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "VenueId supplied is invalid")
	}

	o, err := con.venueService.GetVenueOpeningHours(&v)
	return c.Status(fiber.StatusOK).JSON(o)
}
