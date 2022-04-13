package bookingRoute

import (
	"github.com/nmcalinden/footpal/models"
	"time"
)

var MockBookings = []models.Booking{
	{
		BookingId:       1,
		BookingStatusId: 2,
		CreatedBy:       1,
		Created:         time.Now().String(),
		LastUpdated:     time.Now().String(),
	},
	{
		BookingId:       2,
		BookingStatusId: 2,
		CreatedBy:       1,
		Created:         time.Now().String(),
		LastUpdated:     time.Now().String(),
	},
}
