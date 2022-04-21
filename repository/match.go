package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/nmcalinden/footpal/models"
)

type MatchRepository struct {
	database *sqlx.DB
}

func NewMatchRepository(database *sqlx.DB) *MatchRepository {
	return &MatchRepository{database: database}
}

func (r MatchRepository) FindAll() (*[]models.Match, error) {
	var matchRecords []models.Match
	err := r.database.Select(&matchRecords, "SELECT * FROM footpaldb.public.match")
	if err != nil || len(matchRecords) == 0 {
		return nil, err
	}
	return &matchRecords, nil
}

func (r MatchRepository) FindById(id *int) (*models.Match, error) {
	var match models.Match
	err := r.database.Get(&match, "SELECT * FROM footpaldb.public.match WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &match, nil
}

func (r MatchRepository) Update(match *models.Match) (*models.Match, error) {
	_, err := r.database.NamedExec(`UPDATE footpaldb.public.match SET match_access_status_id=:match_access_status_id, 
                                    match_status_id=:match_status_id, is_paid=:is_paid, last_updated=:last_updated WHERE id=:id`, match)

	if err != nil {
		return nil, err
	}
	return match, nil
}

func (r MatchRepository) DeletePlayerByMatch(matchId *int, playerId *int) error {
	res, err := r.database.Exec("DELETE FROM footpaldb.public.match_player WHERE match_id=$1 AND player_id=$2", matchId, playerId)

	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err == nil && count == 1 {
		return nil
	}

	return err
}
