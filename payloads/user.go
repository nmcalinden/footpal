package payloads

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
	IsPlayer bool `json:"isPlayer" validate:"required"`
}

type Refresh struct {
	RefreshToken *string `json:"refresh_token" validate:"required"`
}

type RegisterResponse struct {
	Id *int `json:"id"`
}

type TokenPairResponse struct {
	AccessToken  *string `json:"access_token"`
	RefreshToken *string `json:"refresh_token"`
}
