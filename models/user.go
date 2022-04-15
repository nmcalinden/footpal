package models

type User struct {
	UserId   int    `json:"venueId,omitempty" db:"id"`
	Forename string `json:"forename" validate:"required,min=3,max=50" db:"forename"`
	Surname  string `json:"surname" validate:"required,min=3,max=50" db:"surname"`
	Email    string `json:"email" validate:"required,email,min=5,max=100" db:"email"`
}

type Login struct {
	Email string `json:"email" validate:"required,min=5,max=100"`
}

type Register struct {
	Forename string `json:"forename" validate:"required,min=3,max=50"`
	Surname  string `json:"surname" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,min=5,max=100"`
}

type UserResponse struct {
	Id *int
}
