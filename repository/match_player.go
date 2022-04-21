package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/nmcalinden/footpal/models"
)

type MatchPlayerRepository struct {
	database *sqlx.DB
}

func NewMatchPlayerRepository(database *sqlx.DB) *MatchPlayerRepository {
	return &MatchPlayerRepository{database: database}
}

func (r MatchPlayerRepository) FindByMatchId(matchId *int) (*[]models.MatchPlayer, error) {
	var matchPlayers []models.MatchPlayer
	err := r.database.Select(&matchPlayers, "SELECT * FROM footpaldb.public.match_player WHERE match_id = $1", matchId)
	if err != nil {
		return nil, err
	}
	return &matchPlayers, nil
}

func (r MatchPlayerRepository) FindMatchesByPlayer(playerId *int) (*[]models.Match, error) {
	var matches []models.Match
	err := r.database.Select(&matches, getFindMatchesByPlayerQuery(), playerId)
	if err != nil {
		return nil, err
	}
	return &matches, nil
}

func (r MatchPlayerRepository) FindByMatchIdAndPlayerId(matchId *int, playerId *int) (*models.MatchPlayer, error) {
	var matchPlayer models.MatchPlayer
	err := r.database.Get(&matchPlayer, "SELECT * FROM footpaldb.public.match_player WHERE match_id = $1 AND player_id = $2", matchId, playerId)
	if err != nil {
		return nil, err
	}

	return &matchPlayer, nil
}

func (r MatchPlayerRepository) Update(matchPlayer *models.MatchPlayer) (*models.MatchPlayer, error) {
	_, err := r.database.NamedExec(`UPDATE footpaldb.public.match_player SET amount_to_pay=:amount_to_pay, 
                                   payment_type_id=:payment_type_id WHERE id=:id`, matchPlayer)

	if err != nil {
		return nil, err
	}
	return matchPlayer, nil
}

func (r MatchPlayerRepository) Delete(matchId *int, playerId *int) error {
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

func (r MatchPlayerRepository) Save(player models.MatchPlayer) (*int, error) {
	_, err := r.database.NamedExec(`INSERT INTO footpaldb.public.match_player(match_id, player_id, amount_to_pay, payment_type_id)
 						VALUES(:match_id, :player_id, :amount_to_pay, :payment_type_id)`, player)
	if err != nil {
		return nil, err
	}
	return player.MatchId, nil
}

func getFindMatchesByPlayerQuery() string {
	return fmt.Sprintf("SELECT ma.* FROM match ma JOIN match_player mp on ma.id = mp.match_id WHERE mp.player_id = $1")
}
