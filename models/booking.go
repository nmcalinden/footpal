package models

import "time"

type Booking struct {
	BookingId       int       `json:"bookingId,omitempty" db:"id"`
	BookingStatusId int       `json:"bookingStatusId" validate:"required" db:"booking_status_id"`
	CreatedBy       int       `json:"createdBy" validate:"required" db:"created_by"`
	Created         time.Time `json:"created" validate:"required" db:"created"`
	LastUpdated     time.Time `json:"lastUpdated" validate:"required" db:"last_updated"`
}

type BookingRequest struct {
	VenueId             int    `json:"venueId" validate:"required"`
	PitchId             int    `json:"pitchId" validate:"required"`
	Day                 string `json:"day" validate:"required,min=3,max=10"`
	StartTime           string `json:"startTime" validate:"required"`
	EndTime             string `json:"endTime" validate:"required"`
	NoOfWeeks           int    `json:"noOfWeeks" validate:"required"`
	MatchAccessStatusId int    `json:"matchAccessStatusId" validate:"required"`
	SquadId             int    `json:"squadId" validate:"omitempty"`
}
