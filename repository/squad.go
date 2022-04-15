package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/nmcalinden/footpal/models"
)

type SquadRepository struct {
	database *sqlx.DB
}

func NewSquadRepository(database *sqlx.DB) *SquadRepository {
	return &SquadRepository{database: database}
}

func (repository SquadRepository) FindAll() (*[]models.Squad, error) {
	var squads []models.Squad
	err := repository.database.Select(&squads, "SELECT * FROM footpaldb.public.squad")
	if err != nil || len(squads) == 0 {
		return nil, err
	}
	return &squads, nil
}

func (repository SquadRepository) FindById(id *int) (*models.Squad, error) {
	var squad models.Squad
	err := repository.database.Get(&squad, "SELECT * FROM footpaldb.public.squad WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &squad, nil
}

func (repository SquadRepository) FindAllByPlayerId(playerId *int) (*[]models.Squad, error) {
	var squads []models.Squad
	err := repository.database.Select(&squads, getFindAllByPlayerIdQuery(), playerId)
	if err != nil {
		return nil, err
	}
	return &squads, nil
}

func (repository SquadRepository) FindSquadByPlayerId(squadId *int, playerId *int) (*models.Squad, error) {
	var squads models.Squad
	err := repository.database.Get(&squads, getFindBySquadIdAndPlayerIdQuery(), squadId, playerId)
	if err != nil {
		return nil, err
	}
	return &squads, nil
}

func (repository SquadRepository) FindPlayersBySquadId(squadId *int) (*[]models.SquadPlayerDetails, error) {
	var players []models.SquadPlayerDetails
	err := repository.database.Select(&players, getFindPlayersBySquadIdQuery(), squadId)
	if err != nil {
		return nil, err
	}
	return &players, nil
}

func (repository SquadRepository) Update(squad *models.Squad) (*models.Squad, error) {
	_, err := repository.database.NamedExec(`UPDATE footpaldb.public.squad SET squad_name=:squad_name, 
                                    city=:city WHERE id=:id`, squad)

	if err != nil {
		return nil, err
	}
	return squad, nil
}

func (repository SquadRepository) Save(squad *models.Squad) (*int, error) {
	_, err := repository.database.NamedExec(`INSERT INTO footpaldb.public.squad(squad_name, city) VALUES(:squad_name, :city)`, squad)

	if err != nil {
		return nil, err
	}
	return &squad.SquadId, nil
}

func (repository SquadRepository) AddPlayer(squadPlayer models.SquadPlayer) error {
	_, err := repository.database.NamedExec(`INSERT INTO footpaldb.public.squad_player(squad_id, player_id, user_role, squad_player_status_id)
					VALUES(:squad_id, :player_id, :user_role, :squad_player_status_id)`, squadPlayer)

	if err != nil {
		return err
	}
	return nil
}

func getFindAllByPlayerIdQuery() string {
	return fmt.Sprintf("SELECT sq.* FROM squad sq JOIN squad_player sp on sq.id = sp.squad_id WHERE sp.player_id = $1")
}

func getFindPlayersBySquadIdQuery() string {
	return fmt.Sprintf("SELECT p.id, p.nickname, fu.forename, fu.surname FROM squad_player sp " +
		"JOIN player p on sp.player_id = p.id " +
		"JOIN footpal_user fu on p.footpal_user_id = fu.id " +
		"WHERE sp.squad_id = $1")
}

func getFindBySquadIdAndPlayerIdQuery() string {
	return fmt.Sprintf("SELECT sq.* FROM squad sq JOIN squad_player sp on sq.id = sp.squad_id WHERE sp.squad_id = $1 AND sp.player_id = $2")
}

func (repository SquadRepository) Delete(id *int) error {
	res, err := repository.database.Exec("DELETE FROM footpaldb.public.squad WHERE id=$1", id)

	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err == nil && count == 1 {
		return nil
	}

	return err
}

func (repository SquadRepository) UpdatePlayerStatus(squadId *int, playerId *int, status int) error {
	_, err := repository.database.Query("UPDATE footpaldb.public.squad_player SET squad_player_status_id=$1 WHERE squad_id=$2 AND player_id=$3", status, squadId, playerId)
	if err != nil {
		return err
	}

	return nil
}
