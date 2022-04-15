package models

type VenueRequest struct {
	Name     string `json:"venueName" validate:"required,min=3,max=100"`
	Address  string `json:"address" validate:"required,min=3,max=100"`
	Postcode string `json:"postcode" validate:"required,min=5,max=8"`
	City     string `json:"city" validate:"required,min=3,max=50"`
	PhoneNo  string `json:"phoneNo" validate:"required,min=11,max=15"`
	Email    string `json:"email" validate:"required,min=10,max=100"`
}

type Venue struct {
	VenueId  int    `json:"venueId,omitempty" db:"id"`
	Name     string `json:"venueName" validate:"required,min=3,max=100" db:"venue_name"`
	Address  string `json:"address" validate:"required,min=3,max=100" db:"venue_address"`
	Postcode string `json:"postcode" validate:"required,min=5,max=8" db:"postcode"`
	City     string `json:"city" validate:"required,min=3,max=50" db:"city"`
	PhoneNo  string `json:"phoneNo" validate:"required,min=11,max=15" db:"phone_no"`
	Email    string `json:"email" validate:"required,min=10,max=100" db:"email"`
}

type VenueAdmin struct {
	VenueAdminId int `json:"venueAdminId,omitempty" db:"id"`
	UserId       int `json:"userId,omitempty" db:"footpal_user_id"`
	VenueId      int `json:"venueId,omitempty" db:"venue_id"`
}

type OpeningHours struct {
	Day       string `json:"day" validate:"required,min=3,max=100"`
	StartTime string `json:"startTime" validate:"required,min=3,max=15"`
	EndTime   string `json:"endTime" validate:"required,min=3,max=15"`
}

type VenueTimeSlot struct {
	PitchId   int         `json:"pitchId,omitempty"`
	TimeSlots []TimeSlots `json:"timeSlots"`
}

type TimeSlots struct {
	DayOfWeek string  `json:"day" validate:"required,min=3,max=100"`
	Slots     []Times `json:"slots"`
}

type Times struct {
	StartTime string `json:"startTime" validate:"required,min=3,max=15"`
	EndTime   string `json:"endTime" validate:"required,min=3,max=15"`
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
