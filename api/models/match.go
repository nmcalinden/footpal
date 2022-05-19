package models

import "time"

type Match struct {
	MatchId             *int      `json:"matchId,omitempty" db:"id"`
	BookingId           int       `json:"bookingId" validate:"required" db:"booking_id"`
	MatchAccessStatusId int       `json:"matchAccessStatusId" validate:"required" db:"match_access_status_id"`
	MatchStatusId       int       `json:"matchStatusId" validate:"required" db:"match_status_id"`
	SquadId             *int      `json:"squadId" validate:"required" db:"squad_id"`
	MatchDate           string    `json:"matchDate" validate:"required" db:"match_date"`
	Cost                float32   `json:"cost" validate:"required" db:"cost"`
	IsPaid              bool      `json:"isPaid" validate:"required" db:"is_paid"`
	Created             time.Time `json:"created" validate:"required" db:"created"`
	LastUpdated         time.Time `json:"lastUpdated" validate:"required" db:"last_updated"`
}

type MatchBookingDetail struct {
	VenueId    int       `json:"venueId" validate:"required" db:"venue_id"`
	VenueName  string    `json:"venueName" validate:"required,min=3,max=100" db:"venue_name"`
	PitchId    int       `json:"pitchId,omitempty" db:"pitch_id"`
	PitchName  string    `json:"pitchName" validate:"required,min=3,max=100" db:"pitch_name"`
	MaxPlayers int       `json:"maxPlayers" validate:"required" db:"max_players"`
	Cost       float32   `json:"cost" validate:"required" db:"cost"`
	MatchDate  string    `json:"matchDate" validate:"required" db:"match_date"`
	StartTime  time.Time `json:"startTime" validate:"required" db:"start_time"`
}

type MatchPlayer struct {
	MatchPlayerId *int    `json:"matchPlayerId,omitempty" db:"id"`
	MatchId       *int    `json:"matchId" validate:"required" db:"match_id"`
	PlayerId      *int    `json:"playerId" validate:"required" db:"player_id"`
	AmountToPay   float32 `json:"amountToPay" validate:"required" db:"amount_to_pay"`
	PaymentTypeId int     `json:"paymentTypeId" validate:"required" db:"payment_type_id"`
}
