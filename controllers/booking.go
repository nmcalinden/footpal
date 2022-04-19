package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/models"
	"github.com/nmcalinden/footpal/services"
	"github.com/nmcalinden/footpal/utils"
	"strconv"
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
// @Success      200  {array} models.Booking
// @Failure      500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router       /bookings [get]
func (controller BookingController) RetrieveBookings(c *fiber.Ctx) error {
	bookingRecords, err := controller.bookingService.GetBookings()
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to get bookings")
	}
	return c.Status(fiber.StatusOK).JSON(bookingRecords)
}

// CreateBooking @Summary      Create new booking
// @Description  Create new single or recurring booking testing
// @Tags         booking
// @Produce      json
// @Param 		 message body models.BookingRequest true "Request"
// @Success      202
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router       /bookings [post]
func (controller BookingController) CreateBooking(c *fiber.Ctx) error {
	newBooking := new(models.BookingRequest)
	if err := c.BodyParser(&newBooking); err != nil {
		return err
	}

	if errors := utils.ValidateStruct(*newBooking); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	bookingId, err := controller.bookingService.CreateNewBooking(newBooking)

	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to create booking")
	}
	return c.Status(fiber.StatusAccepted).JSON(bookingId)
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
func (controller BookingController) GetBookingById(c *fiber.Ctx) error {
	bookingId, err := strconv.Atoi(c.Params("bookingId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "BookingId supplied is invalid")
	}

	booking, bErr := controller.bookingService.GetBookingById(&bookingId)
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
// @Param 		 message body models.BookingRequest true "Request"
// @Success      200 {object} models.Booking
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router       /bookings/{bookingId} [put]
func (controller BookingController) UpdateBooking(c *fiber.Ctx) error {
	bookingId, err := strconv.Atoi(c.Params("bookingId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "BookingId supplied is not a number")
	}

	b := new(models.BookingRequest)

	if err := c.BodyParser(&b); err != nil {
		return err
	}

	if errors := utils.ValidateStruct(*b); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	response, bErr := controller.bookingService.EditBooking(&bookingId, b)
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
func (controller BookingController) CancelBooking(c *fiber.Ctx) error {
	bookingId, err := strconv.Atoi(c.Params("bookingId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "BookingId supplied is not a number")
	}

	response, dErr := controller.bookingService.CancelBooking(&bookingId)
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
func (controller BookingController) GetMatchesByBooking(c *fiber.Ctx) error {
	bookingId, err := strconv.Atoi(c.Params("bookingId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "BookingId supplied is not a number")
	}

	matches, mErr := controller.bookingService.GetMatchesByBooking(&bookingId)
	if mErr != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to retrieve matches by booking")
	}
	return c.Status(fiber.StatusOK).JSON(matches)
}
