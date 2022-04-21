package models

type Squad struct {
	SquadId *int   `json:"squadId,omitempty" db:"id"`
	Name    string `json:"name" db:"squad_name"`
	City    string `json:"city" db:"city"`
}

type SquadPlayer struct {
	SquadPlayerId       *int   `json:"squadPlayerId,omitempty" db:"id"`
	SquadId             int    `json:"squadId" db:"squad_id"`
	PlayerId            int    `json:"playerId" db:"player_id"`
	Role                string `json:"role" db:"user_role"`
	SquadPlayerStatusId int    `json:"squadPlayerStatusId" db:"squad_player_status_id"`
}

type SquadPlayerDetails struct {
	PlayerId int    `json:"playerId,omitempty" db:"id"`
	Nickname string `json:"nickname" db:"nickname"`
	Forename string `json:"forename" db:"forename"`
	Surname  string `json:"surname" db:"surname"`
}
