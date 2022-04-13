package models

import "time"

type Pitch struct {
	PitchId    int     `json:"pitchId,omitempty"`
	VenueId    int     `json:"venueId" validate:"required"`
	Name       string  `json:"name" validate:"required,min=3,max=100"`
	MaxPlayers int     `json:"maxPlayers" validate:"required"`
	Cost       float32 `json:"cost" validate:"required"`
}

type PitchRequest struct {
	Name       string  `json:"name" validate:"required,min=3,max=100"`
	MaxPlayers int     `json:"maxPlayers" validate:"required"`
	Cost       float32 `json:"cost" validate:"required"`
}

type PitchSlot struct {
	PitchSlotId     string `json:"pitchSlotId" validate:"required"`
	BookingId       string `json:"bookingId" validate:"required"`
	PitchTimeSlotId string `json:"pitchTimeSlotId" validate:"required"`
	MatchDate       string `json:"matchDate" validate:"required"`
	BookingStatusId int    `json:"bookingStatusId" validate:"required"`
}

type PitchTimeSlot struct {
	PitchTimeSlotId string    `json:"pitchTimeSlotId" validate:"required"`
	DayOfWeek       string    `json:"dayOfWeek" validate:"required,min=3,max=10"`
	StartTime       time.Time `json:"startTime" validate:"required"`
	EndTime         time.Time `json:"endTime" validate:"required"`
}
