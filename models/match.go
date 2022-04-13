package models

type Match struct {
	MatchId             int     `json:"matchId,omitempty"`
	BookingId           int     `json:"bookingId" validate:"required"`
	MatchAccessStatusId int     `json:"matchAccessStatusId" validate:"required"`
	MatchStatusId       int     `json:"matchStatusId" validate:"required"`
	Cost                float32 `json:"cost" validate:"required"`
	IsPaid              bool    `json:"isPaid" validate:"required"`
	Created             string  `json:"created" validate:"required"`
	LastUpdated         string  `json:"lastUpdated" validate:"required"`
}

type MatchPlayer struct {
	MatchPlayerId int     `json:"matchPlayerId,omitempty"`
	MatchId       int     `json:"matchId" validate:"required"`
	PlayerId      int     `json:"playerId" validate:"required"`
	AmountToPay   float32 `json:"amountToPay" validate:"required"`
	PaymentTypeId int     `json:"paymentTypeId" validate:"required"`
}
