package models

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

type Booking struct {
	BookingId       int    `json:"bookingId,omitempty"`
	BookingStatusId int    `json:"bookingStatusId" validate:"required"`
	CreatedBy       int    `json:"createdBy" validate:"required"`
	Created         string `json:"created" validate:"required"`
	LastUpdated     string `json:"lastUpdated" validate:"required"`
}
