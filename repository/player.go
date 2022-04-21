package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/nmcalinden/footpal/models"
)

type PlayerRepository struct {
	database *sqlx.DB
}

func NewPlayerRepository(database *sqlx.DB) *PlayerRepository {
	return &PlayerRepository{database: database}
}

func (repository PlayerRepository) FindAll() (*[]models.Player, error) {
	var players []models.Player
	err := repository.database.Select(&players, "SELECT * FROM footpaldb.public.player")
	if err != nil || len(players) == 0 {
		return nil, err
	}
	return &players, nil
}

func (repository PlayerRepository) FindById(id *int) (*models.Player, error) {
	query := "SELECT * FROM footpaldb.public.player WHERE id = $1"
	return repository.findByInt(id, query)
}

func (repository PlayerRepository) FindByUserId(userId *int) (*models.Player, error) {
	query := "SELECT * FROM footpaldb.public.player WHERE footpal_user_id = $1"
	return repository.findByInt(userId, query)
}

func (repository PlayerRepository) findByInt(id *int, query string) (*models.Player, error) {
	var player models.Player
	err := repository.database.Get(&player, query, id)
	if err != nil {
		return nil, err
	}
	return &player, nil
}

func (repository PlayerRepository) Update(player *models.Player) (*models.Player, error) {
	_, err := repository.database.NamedExec(`UPDATE footpaldb.public.player SET nickname=:nickname, 
                                    phone_no=:phone_no, postcode=:postcode, city=:city WHERE id=:id`, player)

	if err != nil {
		return nil, err
	}
	return player, nil
}
