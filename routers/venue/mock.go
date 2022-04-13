package venueRoute

import (
	"github.com/nmcalinden/footpal/models"
)

var MockVenues = []models.Venue{
	{
		VenueId:      1,
		Name:         "Soccer Land",
		Address:      "123 Skill Street",
		Postcode:     "BT65 8RE",
		City:         "Belfast",
		PhoneNo:      "07912345678",
		Email:        "tester@test.com",
		OpeningHours: *mockOpeningHours,
	},
	{
		VenueId:      2,
		Name:         "Big Dome Games",
		Address:      "99 Domer Ave",
		Postcode:     "BT99 12W",
		City:         "Belfast",
		PhoneNo:      "07952525355",
		Email:        "bigdome@uno.com",
		OpeningHours: *mockOpeningHours,
	},
	{
		VenueId:      3,
		Name:         "Football Fever",
		Address:      "85 Premier Street",
		Postcode:     "BT12 9RE",
		City:         "Belfast",
		PhoneNo:      "0798238128",
		Email:        "ff@test.com",
		OpeningHours: *mockOpeningHours,
	},
}

var MockVenuePitches = []models.Pitch{
	{
		PitchId:    1,
		VenueId:    1,
		Name:       "Pitch A",
		MaxPlayers: 10,
		Cost:       float32(30),
	},
	{
		PitchId:    2,
		VenueId:    1,
		Name:       "Pitch B",
		MaxPlayers: 12,
		Cost:       float32(50),
	},
}

var mockOpeningHours = &[]models.OpeningHours{
	{Day: "Monday", StartTime: "09:00", EndTime: "22:00"},
	{Day: "Tuesday", StartTime: "09:00", EndTime: "22:00"},
	{Day: "Wednesday", StartTime: "09:00", EndTime: "22:00"},
	{Day: "Thursday", StartTime: "09:00", EndTime: "22:00"},
	{Day: "Friday", StartTime: "09:00", EndTime: "22:00"},
	{Day: "Saturday", StartTime: "09:00", EndTime: "17:00"},
	{Day: "Sunday", StartTime: "09:00", EndTime: "17:00"},
}
