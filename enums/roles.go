package enums

type Role string

const (
	All        Role = "everyone"
	Player     Role = "player"
	VenueAdmin Role = "venueAdmin"
)

func (e Role) String() string {
	return string(e)
}
