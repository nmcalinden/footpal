package views

type PitchBookingDetails struct {
	MatchDate string                 `json:"matchDate"`
	DayOfWeek string                 `json:"dayOfWeek"`
	TimeSlots []PitchTimeSlotBooking `json:"timeSlots"`
}

type PitchTimeSlotBooking struct {
	PitchTimeSlotId int    `json:"id,omitempty"`
	StartTime       string `json:"startTime"`
	EndTime         string `json:"endTime"`
	IsBooked        bool   `json:"isBooked"`
}
