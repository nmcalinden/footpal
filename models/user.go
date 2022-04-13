package models

type User struct {
	UserId   int    `json:"venueId,omitempty"`
	Forename string `json:"forename" validate:"required,min=3,max=50"`
	Surname  string `json:"surname" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,min=5,max=8"`
}

type Login struct {
	Email    string `json:"email" validate:"required,min=5,max=8"`
	Password string `json:"password" validate:"required,min=5,max=30"`
}

type Register struct {
	Forename string `json:"forename" validate:"required,min=3,max=50"`
	Surname  string `json:"surname" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,min=5,max=8"`
	Password string `json:"password" validate:"required,min=5,max=30"`
}
