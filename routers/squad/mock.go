package squadRoute

import (
	"github.com/nmcalinden/footpal/models"
)

var MockSquads = []models.Squad{
	{
		SquadId: 1,
		Name:    "Tester Team",
		City:    "Belfast",
	},
	{
		SquadId: 2,
		Name:    "The B Team",
		City:    "Lisburn",
	},
}

var MockSquadPlayers = []models.SquadPlayer{
	{
		SquadPlayerId:       1,
		SquadId:             1,
		PlayerId:            1,
		Role:                "admin",
		SquadPlayerStatusId: 1,
	},
	{
		SquadPlayerId:       2,
		SquadId:             1,
		PlayerId:            2,
		Role:                "player",
		SquadPlayerStatusId: 1,
	},
	{
		SquadPlayerId:       3,
		SquadId:             2,
		PlayerId:            2,
		Role:                "player",
		SquadPlayerStatusId: 1,
	},
	{
		SquadPlayerId:       4,
		SquadId:             1,
		PlayerId:            3,
		Role:                "player",
		SquadPlayerStatusId: 1,
	},
	{
		SquadPlayerId:       5,
		SquadId:             2,
		PlayerId:            3,
		Role:                "admin",
		SquadPlayerStatusId: 1,
	},
}
