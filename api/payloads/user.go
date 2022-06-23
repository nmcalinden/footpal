package payloads

import "github.com/nmcalinden/footpal/views"

type Login struct {
	Email    string `json:"email" validate:"required,min=5,max=100"`
	Password string `json:"password" validate:"required,min=5,max=25"`
}

type Register struct {
	Forename string `json:"forename" validate:"required,min=3,max=50"`
	Surname  string `json:"surname" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,min=5,max=100"`
	Password string `json:"password" validate:"required,min=5,max=25"`

	//default: true
	IsPlayer bool `json:"isPlayer"`
}

type Refresh struct {
	RefreshToken *string `json:"refreshToken" validate:"required"`
}

type RegisterResponse struct {
	Id *int `json:"id"`
}

type LoginResponse struct {
	User views.UserProfile `json:"user"`
	JWT  TokenPairResponse `json:"jwt"`
}

type TokenPairResponse struct {
	AccessToken  *string `json:"accessToken"`
	RefreshToken *string `json:"refreshToken"`
}
