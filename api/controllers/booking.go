package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/errors"
	"github.com/nmcalinden/footpal/payloads"
	"github.com/nmcalinden/footpal/services"
	"github.com/nmcalinden/footpal/utils"
)

type BookingController struct {
	bookingService *services.BookingService
}

func NewBookingController(bookingService *services.BookingService) *BookingController {
	return &BookingController{bookingService: bookingService}
}

// RetrieveBookings @Summary      Retrieve bookings
// @Description  Retrieve all bookings by user
// @Tags         booking
// @Produce      json
// @Success      200  {array} views.UserBooking
// @Failure      500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router       /bookings [get]
func (con BookingController) RetrieveBookings(c *fiber.Ctx) error {
	claims := utils.GetClaims(c.Locals("user"))
	userId := int(claims["sub"].(float64))

	bookingRecords, err := con.bookingService.GetBookings(&userId)
	if err != nil {
		e, ok := err.(*errors.FpError)
		if ok && e.ErrorCode == errors.NoResults {
			return c.Status(fiber.StatusOK).JSON(new(interface{}))
		}
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to get bookings")
	}
	return c.Status(fiber.StatusOK).JSON(bookingRecords)
}

// CreateBooking @Summary      Create new booking
// @Description  Create new single or recurring booking testing
// @Tags         booking
// @Produce      json
// @Param 		 message body payloads.BookingRequest true "Request"
// @Success      202 {object} payloads.BookingResponse
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router       /bookings [post]
func (con BookingController) CreateBooking(c *fiber.Ctx) error {
	claims := utils.GetClaims(c.Locals("user"))
	userId := int(claims["sub"].(float64))

	newBooking := new(payloads.BookingRequest)
	if err := c.BodyParser(&newBooking); err != nil {
		return err
	}

	if err := utils.ValidateStruct(*newBooking); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	bookingId, err := con.bookingService.CreateNewBooking(newBooking, userId)

	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to create booking")
	}

	res := payloads.BookingResponse{BookingId: *bookingId}
	return c.Status(fiber.StatusAccepted).JSON(res)
}

// FindAvailableSlotsByVenue @Summary  Find venues with available slots
// @Description  Find Venues with available bookings
// @Tags         booking
// @Produce      json
// @Param 		 message body payloads.BookingSearchRequest true "Request"
// @Success      200 {array} models.Venue
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Router       /bookings/search [post]
func (con BookingController) FindAvailableSlotsByVenue(c *fiber.Ctx) error {
	bs := new(payloads.BookingSearchRequest)
	if err := c.BodyParser(&bs); err != nil {
		return err
	}

	if err := utils.ValidateStruct(*bs); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	res, err := con.bookingService.FindVenuesWithAvailableBookings(bs)

	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to find venues")
	}
	return c.Status(fiber.StatusAccepted).JSON(res)
}

// GetBookingById @Summary      Retrieve booking
// @Description  Retrieve booking by bookingId
// @Tags         booking
// @Produce      json
// @Param        bookingId   path  int  true  "Booking ID"
// @Success      200 {object} models.Booking
// @Failure      400 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router       /bookings/{bookingId} [get]
func (con BookingController) GetBookingById(c *fiber.Ctx) error {
	bookingId, err := strconv.Atoi(c.Params("bookingId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "BookingId supplied is invalid")
	}

	booking, bErr := con.bookingService.GetBookingById(&bookingId)
	if bErr != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Booking does not exist")
	}
	return c.Status(fiber.StatusOK).JSON(booking)
}

// UpdateBooking @Summary      Edit booking
// @Description  Edit booking details
// @Tags         booking
// @Produce      json
// @Param        bookingId   path  int  true  "Booking ID"
// @Param 		 message body payloads.BookingRequest true "Request"
// @Success      200 {object} models.Booking
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router       /bookings/{bookingId} [put]
func (con BookingController) UpdateBooking(c *fiber.Ctx) error {
	bookingId, err := strconv.Atoi(c.Params("bookingId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "BookingId supplied is not a number")
	}

	b := new(payloads.BookingRequest)

	if err := c.BodyParser(&b); err != nil {
		return err
	}

	if err := utils.ValidateStruct(*b); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	response, bErr := con.bookingService.EditBooking(&bookingId, b)
	if bErr != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to edit booking")
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

// CancelBooking @Summary      Cancel booking
// @Description  Cancel active or pending booking
// @Tags         booking
// @Produce      json
// @Param        bookingId   path  int  true  "Booking ID"
// @Success      200 {string} bookingId
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router       /bookings/{bookingId} [delete]
func (con BookingController) CancelBooking(c *fiber.Ctx) error {
	bookingId, err := strconv.Atoi(c.Params("bookingId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "BookingId supplied is not a number")
	}

	response, dErr := con.bookingService.CancelBooking(&bookingId)
	if dErr != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to cancel booking")
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

// GetMatchesByBooking @Summary      Get matches by booking
// @Description  Get all matches linked to booking
// @Tags         booking
// @Produce      json
// @Param        bookingId   path  int  true  "Booking ID"
// @Success      200 {array} models.Match
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router       /bookings/{bookingId}/matches [get]
func (con BookingController) GetMatchesByBooking(c *fiber.Ctx) error {
	bookingId, err := strconv.Atoi(c.Params("bookingId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "BookingId supplied is not a number")
	}

	matches, mErr := con.bookingService.GetMatchesByBooking(&bookingId)
	if mErr != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to retrieve matches by booking")
	}
	return c.Status(fiber.StatusOK).JSON(matches)
}
