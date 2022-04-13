package matchRoute

import (
	"github.com/nmcalinden/footpal/models"
	"time"
)

var MockMatches = []models.Match{
	{
		MatchId:             1,
		BookingId:           1,
		MatchAccessStatusId: 2,
		MatchStatusId:       1,
		Cost:                30.0,
		IsPaid:              false,
		Created:             time.Now().String(),
		LastUpdated:         time.Now().String(),
	},
	{
		MatchId:             2,
		BookingId:           1,
		MatchAccessStatusId: 2,
		MatchStatusId:       2,
		Cost:                30.0,
		IsPaid:              false,
		Created:             time.Now().String(),
		LastUpdated:         time.Now().String(),
	},
	{
		MatchId:             3,
		BookingId:           1,
		MatchAccessStatusId: 2,
		MatchStatusId:       2,
		Cost:                30.0,
		IsPaid:              false,
		Created:             time.Now().String(),
		LastUpdated:         time.Now().String(),
	},
	{
		MatchId:             4,
		BookingId:           1,
		MatchAccessStatusId: 2,
		MatchStatusId:       2,
		Cost:                30.0,
		IsPaid:              false,
		Created:             time.Now().String(),
		LastUpdated:         time.Now().String(),
	},
}

var MockMatchPlayers = []models.MatchPlayer{
	{
		MatchPlayerId: 1,
		MatchId:       1,
		PlayerId:      1,
		PaymentTypeId: 1,
	},
	{
		MatchPlayerId: 2,
		MatchId:       1,
		PlayerId:      2,
		PaymentTypeId: 1,
	},
	{
		MatchPlayerId: 3,
		MatchId:       1,
		PlayerId:      3,
		PaymentTypeId: 1,
	},
}
