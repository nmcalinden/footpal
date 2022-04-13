package models

type PlayerRequest struct {
	Nickname string `json:"nickname" validate:"required,min=2,max=16"`
	PhoneNo  string `json:"address" validate:"required,min=11,max=15"`
	Postcode string `json:"postcode" validate:"required,min=5,max=8"`
	City     string `json:"city" validate:"required,min=3,max=50"`
}

type Player struct {
	PlayerId int    `json:"playerId,omitempty"`
	Nickname string `json:"nickname" validate:"required,min=2,max=16"`
	PhoneNo  string `json:"address" validate:"required,min=11,max=15"`
	Postcode string `json:"postcode" validate:"required,min=5,max=8"`
	City     string `json:"city" validate:"required,min=3,max=50"`
}
