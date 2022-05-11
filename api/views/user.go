package views

type PlayerUser struct {
	Name     string  `json:"name"`
	Nickname *string `json:"nickname"`
	Email    string  `json:"email"`
	PhoneNo  string  `json:"phoneNo"`
	Postcode string  `json:"postcode"`
	City     string  `json:"city"`
}
