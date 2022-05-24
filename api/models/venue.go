package models

type Venue struct {
	VenueId  *int   `json:"id,omitempty" db:"id"`
	Name     string `json:"name" db:"venue_name"`
	Address  string `json:"address" db:"venue_address"`
	Postcode string `json:"postcode" db:"postcode"`
	City     string `json:"city" db:"city"`
	PhoneNo  string `json:"phoneNo" db:"phone_no"`
	Email    string `json:"email" db:"email"`
}

type VenueAdmin struct {
	VenueAdminId *int `json:"venueAdminId,omitempty" db:"id"`
	UserId       int  `json:"userId,omitempty" db:"footpal_user_id"`
	VenueId      int  `json:"venueId,omitempty" db:"venue_id"`
}

type OpeningHours struct {
	Day       string `json:"day"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type VenueTimeSlot struct {
	PitchTimeSlotId int    `json:"pitchTimeSlotId" db:"pitch_time_slot_id"`
	DayOfWeek       string `json:"dayOfWeek" db:"day_of_week"`
	MatchDate       string `json:"matchDate" db:"match_date"`
}
