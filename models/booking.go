package models

import "time"

type Booking struct {
	BookingId       *int      `json:"bookingId,omitempty" db:"id"`
	BookingStatusId int       `json:"bookingStatusId" validate:"required" db:"booking_status_id"`
	CreatedBy       int       `json:"createdBy" validate:"required" db:"created_by"`
	Created         time.Time `json:"created" validate:"required" db:"created"`
	LastUpdated     time.Time `json:"lastUpdated" validate:"required" db:"last_updated"`
}
