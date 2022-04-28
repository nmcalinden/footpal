package models

type Player struct {
	PlayerId *int    `json:"playerId,omitempty" db:"id"`
	UserId   int     `json:"userId" db:"footpal_user_id"`
	Nickname *string `json:"nickname" validate:"required,min=2,max=16" db:"nickname" swaggertype:"string"`
	PhoneNo  string  `json:"address" validate:"required,min=11,max=15" db:"phone_no"`
	Postcode string  `json:"postcode" validate:"required,min=5,max=8" db:"postcode"`
	City     string  `json:"city" validate:"required,min=3,max=50" db:"city"`
}
