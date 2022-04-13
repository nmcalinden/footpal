package bookingRoute

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/models"
	matchRoute "github.com/nmcalinden/footpal/routers/match"
	"github.com/nmcalinden/footpal/utils"
	"strconv"
)

func ConfigureBookingHandlers(app *fiber.App) {
	group := app.Group("/bookings")

	group.Get("/", retrieveBookings)
	group.Post("/", createBooking)
	group.Get("/:bookingId", getBookingById)
	group.Put("/:bookingId", updateBooking)
	group.Delete("/:bookingId", cancelBooking)
	group.Get("/:bookingId/matches", getMatchesByBooking)
}

// @Summary      Retrieve bookings
// @Description  Retrieve all bookings by user
// @Tags         booking
// @Produce      json
// @Success      200  {array}  models.Booking
// @Router       /bookings [get]
func retrieveBookings(c *fiber.Ctx) error {
	mockBookings := MockBookings
	return c.Status(fiber.StatusOK).JSON(mockBookings)
}

// @Summary      Create new booking
// @Description  Create new single or recurring booking
// @Tags         booking
// @Produce      json
// @Param 		 message body models.BookingRequest true "Request"
// @Success      202
// @Failure      400 {object} utils.ErrorResponse
// @Router       /bookings [post]
func createBooking(c *fiber.Ctx) error {
	newBooking := new(models.BookingRequest)
	if err := c.BodyParser(&newBooking); err != nil {
		return err
	}

	if errors := utils.ValidateStruct(*newBooking); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return c.SendStatus(fiber.StatusAccepted)
}

// @Summary      Retrieve booking
// @Description  Retrieve booking by bookingId
// @Tags         booking
// @Produce      json
// @Param        bookingId   path  int  true  "Booking ID"
// @Success      200 {object} models.Booking
// @Failure      400 {object} utils.ErrorResponse
// @Router       /bookings/{bookingId} [get]
func getBookingById(c *fiber.Ctx) error {
	bookingId, err := strconv.Atoi(c.Params("bookingId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "BookingId supplied is invalid")
	}

	mockBookings := MockBookings
	result := models.Booking{}
	for _, s := range mockBookings {
		if s.BookingId == bookingId {
			result = s
		}
	}

	if result.BookingId == 0 {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Booking does not exist")
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

// @Summary      Edit booking
// @Description  Edit booking details
// @Tags         booking
// @Produce      json
// @Param        bookingId   path  int  true  "Booking ID"
// @Param 		 message body models.BookingRequest true "Request"
// @Success      200 {object} models.Booking
// @Failure      400 {object} utils.ErrorResponse
// @Router       /bookings/{bookingId} [put]
func updateBooking(c *fiber.Ctx) error {
	_, err := strconv.Atoi(c.Params("bookingId"))
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

	response := MockBookings[0]
	return c.Status(fiber.StatusOK).JSON(response)
}

// @Summary      Cancel booking
// @Description  Cancel active or pending booking
// @Tags         booking
// @Produce      json
// @Param        bookingId   path  int  true  "Booking ID"
// @Success      200 {string} bookingId
// @Failure      400 {object} utils.ErrorResponse
// @Router       /bookings/{bookingId} [delete]
func cancelBooking(c *fiber.Ctx) error {
	bookingId, err := strconv.Atoi(c.Params("bookingId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "BookingId supplied is not a number")
	}

	return c.Status(fiber.StatusOK).JSON(bookingId)
}

// @Summary      Get matches by booking
// @Description  Get all matches linked to booking
// @Tags         booking
// @Produce      json
// @Param        bookingId   path  int  true  "Booking ID"
// @Success      200 {array} models.Match
// @Failure      400 {object} utils.ErrorResponse
// @Router       /bookings/{bookingId}/matches [get]
func getMatchesByBooking(c *fiber.Ctx) error {
	_, err := strconv.Atoi(c.Params("bookingId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "BookingId supplied is not a number")
	}

	mockMatches := matchRoute.MockMatches
	return c.Status(fiber.StatusOK).JSON(mockMatches)
}
