package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/nmcalinden/footpal/api/models"
)

//go:generate mockgen -destination=./mocks/player_mock.go -package=mocks github.com/nmcalinden/footpal/api/repository PlayerRepositoryI

type PlayerRepositoryI interface {
	FindAll(limit int, after int) (*[]models.Player, error)
	FindById(id *int) (*models.Player, error)
	FindByUserId(userId *int) (*models.Player, error)
	Update(player *models.Player) (*models.Player, error)
	Save(player *models.Player) (*int, error)
	GetTotal() (*int, error)
}

type PlayerRepository struct {
	database *sqlx.DB
}

func NewPlayerRepository(database *sqlx.DB) *PlayerRepository {
	return &PlayerRepository{database: database}
}

func (r PlayerRepository) FindAll(limit int, after int) (*[]models.Player, error) {
	var players []models.Player
	err := r.database.Select(&players, "SELECT * FROM footpaldb.public.player WHERE id > $1 LIMIT $2", after, limit)
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

func (r PlayerRepository) Save(player *models.Player) (*int, error) {
	stmt, err := r.database.PrepareNamed("INSERT INTO footpaldb.public.player(footpal_user_id, nickname, " +
		"phone_no, postcode, city) VALUES(:footpal_user_id, :nickname, :phone_no, :postcode, :city) RETURNING id")

	if err != nil {
		return nil, err
	}
	var id int
	err = stmt.Get(&id, player)
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (r PlayerRepository) GetTotal() (*int, error) {
	var count int
	err := r.database.Get(&count, "SELECT count(*) from footpaldb.public.player")
	if err != nil {
		return nil, err
	}
	return &count, nil
}
