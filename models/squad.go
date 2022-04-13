package models

type SquadRequest struct {
	Name string `json:"name" validate:"required,min=2,max=30"`
	City string `json:"city" validate:"required,min=3,max=50"`
}

type Squad struct {
	SquadId int    `json:"squadId,omitempty"`
	Name    string `json:"name" validate:"required,min=2,max=30"`
	City    string `json:"city" validate:"required,min=3,max=50"`
}

type SquadPlayer struct {
	SquadPlayerId       int    `json:"squadPlayerId,omitempty"`
	SquadId             int    `json:"squadId"`
	PlayerId            int    `json:"playerId"`
	Role                string `json:"role"`
	SquadPlayerStatusId int    `json:"squadPlayerStatusId"`
}
