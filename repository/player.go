package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/nmcalinden/footpal/models"
)

type PlayerRepositoryI interface {
	FindAll() (*[]models.Player, error)
	FindById(id *int) (*models.Player, error)
	FindByUserId(userId *int) (*models.Player, error)
	Update(player *models.Player) (*models.Player, error)
}

type PlayerRepository struct {
	database *sqlx.DB
}

func NewPlayerRepository(database *sqlx.DB) *PlayerRepository {
	return &PlayerRepository{database: database}
}

func (r PlayerRepository) FindAll() (*[]models.Player, error) {
	var players []models.Player
	err := r.database.Select(&players, "SELECT * FROM footpaldb.public.player")
	if err != nil || len(players) == 0 {
		return nil, err
	}
	return &players, nil
}

func (r PlayerRepository) FindById(id *int) (*models.Player, error) {
	query := "SELECT * FROM footpaldb.public.player WHERE id = $1"
	return r.findByInt(id, query)
}

func (r PlayerRepository) FindByUserId(userId *int) (*models.Player, error) {
	query := "SELECT * FROM footpaldb.public.player WHERE footpal_user_id = $1"
	return r.findByInt(userId, query)
}

func (r PlayerRepository) findByInt(id *int, query string) (*models.Player, error) {
	var player models.Player
	err := r.database.Get(&player, query, id)
	if err != nil {
		return nil, err
	}
	return &player, nil
}

func (r PlayerRepository) Update(player *models.Player) (*models.Player, error) {
	_, err := r.database.NamedExec(`UPDATE footpaldb.public.player SET nickname=:nickname, 
                                    phone_no=:phone_no, postcode=:postcode, city=:city WHERE id=:id`, player)

	if err != nil {
		return nil, err
	}
	return player, nil
}
