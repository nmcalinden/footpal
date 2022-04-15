package models

type Player struct {
	PlayerId int    `json:"playerId,omitempty" db:"id"`
	UserId   int    `json:"userId" db:"footpal_user_id"`
	Nickname string `json:"nickname" validate:"required,min=2,max=16" db:"nickname"`
	PhoneNo  string `json:"address" validate:"required,min=11,max=15" db:"phone_no"`
	Postcode string `json:"postcode" validate:"required,min=5,max=8" db:"postcode"`
	City     string `json:"city" validate:"required,min=3,max=50" db:"city"`
}

type PlayerRequest struct {
	Nickname string `json:"nickname" validate:"required,min=2,max=16"`
	PhoneNo  string `json:"address" validate:"required,min=11,max=15"`
	Postcode string `json:"postcode" validate:"required,min=5,max=8"`
	City     string `json:"city" validate:"required,min=3,max=50"`
}
type JoinMatchResponse struct {
	MatchId *int
}

type MatchPaymentRequest struct {
	AmountToPay float32 `json:"amountToPay" validate:"required"`
}
