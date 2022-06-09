package views

type BookingsByUser struct {
	Pagination Pagination    `json:"pagination"`
	Bookings   []UserBooking `json:"data"`
}
type UserBooking struct {
	BookingId     int                   `json:"id"`
	BookingStatus string                `json:"status"`
	MatchDate     string                `json:"matchDate"`
	StartTime     string                `json:"startTime"`
	NoOfWeeks     int                   `json:"noOfWeeks"`
	TotalCost     float32               `json:"totalCost"`
	IsBookingPaid bool                  `json:"isPaid"`
	Venue         BookingVenueSummary   `json:"venue"`
	Pitch         BookingPitchSummary   `json:"pitch"`
	Matches       []BookingMatchSummary `json:"matches"`
}

type BookingVenueSummary struct {
	VenueId int    `json:"id"`
	Name    string `json:"name"`
}

type BookingPitchSummary struct {
	PitchId int    `json:"id"`
	Name    string `json:"name"`
}

type BookingMatchSummary struct {
	MatchId   int    `json:"id"`
	MatchDate string `json:"matchDate"`
}

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
