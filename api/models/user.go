package models

type User struct {
	UserId   *int   `json:"venueId,omitempty" db:"id"`
	Forename string `json:"forename" db:"forename"`
	Surname  string `json:"surname" db:"surname"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
