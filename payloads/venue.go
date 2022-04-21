package payloads

type VenueRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=100"`
	Address  string `json:"address" validate:"required,min=3,max=100"`
	Postcode string `json:"postcode" validate:"required,min=5,max=8"`
	City     string `json:"city" validate:"required,min=3,max=50"`
	PhoneNo  string `json:"phoneNo" validate:"required,min=11,max=15"`
	Email    string `json:"email" validate:"required,min=10,max=100"`
}

type VenueAdminRequest struct {
	VenueId int
	UserId  int
}

type VenueAdminResponse struct {
	VenueAdminId int
	Forename     string
	Surname      string
	Email        string
}
