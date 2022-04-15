package models

type SquadRequest struct {
	Name string `json:"name" validate:"required,min=2,max=30"`
	City string `json:"city" validate:"required,min=3,max=50"`
}

type Squad struct {
	SquadId int    `json:"squadId,omitempty" db:"id"`
	Name    string `json:"name" validate:"required,min=2,max=30" db:"squad_name"`
	City    string `json:"city" validate:"required,min=3,max=50" db:"city"`
}

type SquadPlayer struct {
	SquadPlayerId       int    `json:"squadPlayerId,omitempty" db:"id"`
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
