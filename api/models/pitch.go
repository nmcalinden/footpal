package models

import "time"

type Pitch struct {
	PitchId    *int    `json:"pitchId,omitempty" db:"id"`
	VenueId    int     `json:"venueId" validate:"required" db:"venue_id"`
	Name       string  `json:"name" validate:"required,min=3,max=100" db:"pitch_name"`
	MaxPlayers int     `json:"maxPlayers" validate:"required" db:"max_players"`
	Cost       float32 `json:"cost" validate:"required" db:"cost"`
}

type PitchSlot struct {
	PitchSlotId     int    `json:"pitchSlotId" validate:"required" db:"id"`
	BookingId       string `json:"bookingId" validate:"required" db:"booking_id"`
	PitchTimeSlotId string `json:"pitchTimeSlotId" validate:"required" db:"pitch_time_slot_id"`
	MatchDate       string `json:"matchDate" validate:"required" db:"match_date"`
	BookingStatusId int    `json:"bookingStatusId" validate:"required" db:"booking_status_id"`
}

type PitchTimeSlot struct {
	PitchTimeSlotId int       `json:"pitchTimeSlotId" validate:"required" db:"id"`
	DayOfWeek       string    `json:"dayOfWeek" validate:"required,min=3,max=10" db:"day_of_week"`
	StartTime       time.Time `json:"startTime" validate:"required" db:"start_time"`
	EndTime         time.Time `json:"endTime" validate:"required" db:"end_time"`
}
