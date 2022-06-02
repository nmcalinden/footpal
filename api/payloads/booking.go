package payloads

type BookingRequest struct {
	VenueId         int    `json:"venueId" validate:"required"`
	PitchTimeSlotId int    `json:"pitchTimeSlotId" validate:"required"`
	MatchDate       string `json:"matchDate" validate:"required"`
	MatchType       string `json:"matchType" validate:"required"`
	Payment         string `json:"payment" validate:"required"`
	NoOfWeeks       int    `json:"noOfWeeks" validate:"required"`
	SquadId         *int   `json:"squadId" validate:"omitempty"`
}

type BookingSearchRequest struct {
	VenueId    *int    `json:"venueId" validate:"required"`
	Date       string  `json:"date" validate:"required"`
	City       *string `json:"city"`
	MaxPlayers *int    `json:"maxPlayers"`
}
