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
	VenueId      int            `json:"venueId,omitempty"`
	Name         string         `json:"venueName" validate:"required,min=3,max=100"`
	Address      string         `json:"address" validate:"required,min=3,max=100"`
	Postcode     string         `json:"postcode" validate:"required,min=5,max=8"`
	City         string         `json:"city" validate:"required,min=3,max=50"`
	PhoneNo      string         `json:"phoneNo" validate:"required,min=11,max=15"`
	Email        string         `json:"email" validate:"required,min=10,max=100"`
	OpeningHours []OpeningHours `json:"openingHours"`
}

type OpeningHours struct {
	Day       string `json:"day" validate:"required,min=3,max=100"`
	StartTime string `json:"startTime" validate:"required,min=3,max=15"`
	EndTime   string `json:"endTime" validate:"required,min=3,max=15"`
}

type PitchTimeSlots struct {
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
