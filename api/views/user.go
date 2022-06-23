package views

type UserProfile struct {
	Id       *int    `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Nickname *string `json:"nickname"`
	PhoneNo  string  `json:"phoneNo"`
	Postcode string  `json:"postcode"`
	City     string  `json:"city"`
}

type PlayerUser struct {
	Name     string  `json:"name"`
	Nickname *string `json:"nickname"`
	Email    string  `json:"email"`
	PhoneNo  string  `json:"phoneNo"`
	Postcode string  `json:"postcode"`
	City     string  `json:"city"`
}
