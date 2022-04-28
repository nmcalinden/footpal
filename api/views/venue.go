package views

type Venue struct {
	VenueId  *int                `json:"id,omitempty"`
	Name     string              `json:"name"`
	Address  string              `json:"address"`
	Postcode string              `json:"postcode"`
	City     string              `json:"city"`
	PhoneNo  string              `json:"phoneNo"`
	Email    string              `json:"email"`
	Pitches  []VenuePitchSummary `json:"pitches"`
}

type VenuePitchSummary struct {
	PitchId *int   `json:"id,omitempty"`
	Name    string `json:"name"`
	Href    string `json:"href"`
}
