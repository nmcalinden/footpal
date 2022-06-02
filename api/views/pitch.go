package views

type PitchTimeSlot struct {
	Pitch    Pitch    `json:"pitch"`
	TimeSlot TimeSlot `json:"timeSlot"`
}
type Pitch struct {
	PitchId    *int    `json:"id"`
	VenueId    int     `json:"venueId"`
	Name       string  `json:"name"`
	MaxPlayers int     `json:"maxPlayers"`
	Cost       float32 `json:"cost"`
}

type TimeSlot struct {
	PitchTimeSlotId int    `json:"id,omitempty"`
	DayOfWeek       string `json:"dayOfWeek"`
	StartTime       string `json:"startTime"`
	EndTime         string `json:"endTime"`
}
