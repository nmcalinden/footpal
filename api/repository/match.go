package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/nmcalinden/footpal/api/models"
)

//go:generate mockgen -destination=./mocks/match_mock.go -package=mocks github.com/nmcalinden/footpal/api/repository MatchRepositoryI

type MatchRepositoryI interface {
	FindAll() (*[]models.Match, error)
	FindById(id *int) (*models.Match, error)
	Update(match *models.Match) (*models.Match, error)
	DeletePlayerByMatch(matchId *int, playerId *int) error
	FindMatchDetailByBookingIdAndMatchDate(bookingId int, matchDate string) (*models.MatchBookingDetail, error)
}

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

func (r MatchRepository) FindMatchDetailByBookingIdAndMatchDate(bookingId int, matchDate string) (*models.MatchBookingDetail, error) {
	q := "SELECT v.id as venue_id, v.venue_name, p.id as pitch_id, p.pitch_name, p.max_players, p.cost, ps.match_date, " +
		"pts.start_time from pitch_slot as ps " +
		"JOIN pitch_time_slot pts on pts.id = ps.pitch_time_slot_id " +
		"JOIN pitch p on pts.pitch_id = p.id " +
		"JOIN venue v on p.venue_id = v.id " +
		"WHERE ps.booking_id = $1 AND ps.match_date = $2"

	var match models.MatchBookingDetail
	err := r.database.Get(&match, q, bookingId, matchDate)
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
