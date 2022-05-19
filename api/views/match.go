package views

type MatchSummary struct {
	MatchId    *int    `json:"id,omitempty"`
	VenueId    int     `json:"venueId"`
	VenueName  string  `json:"venueName"`
	PitchId    int     `json:"pitchId"`
	PitchName  string  `json:"pitchName"`
	MatchDate  string  `json:"matchDate"`
	Time       string  `json:"time"`
	MaxPlayers int     `json:"maxPlayers"`
	Cost       float32 `json:"cost"`
	SquadId    *int    `json:"squadId,omitempty"`
	SquadName  *string `json:"squadName,omitempty"`
}
