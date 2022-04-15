package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/nmcalinden/footpal/models"
	"github.com/nmcalinden/footpal/repository"
)

type MatchService struct {
	matchRepo       *repository.MatchRepository
	matchPlayerRepo *repository.MatchPlayerRepository
}

func NewMatchService(database *sqlx.DB) *MatchService {
	return &MatchService{
		matchRepo:       repository.NewMatchRepository(database),
		matchPlayerRepo: repository.NewMatchPlayerRepository(database),
	}
}

func (service *MatchService) GetMatches() (*[]models.Match, error) {
	return service.matchRepo.FindAll()
}

func (service *MatchService) GetMatchById(matchId *int) (*models.Match, error) {
	return service.matchRepo.FindById(matchId)
}

func (service *MatchService) GetMatchPlayers(matchId *int) (*[]models.MatchPlayer, error) {
	return service.matchPlayerRepo.FindByMatchId(matchId)
}

func (service *MatchService) CancelMatch(matchId *int) (*int, error) {
	m, err := service.matchRepo.FindById(matchId)
	if err != nil {
		return nil, err
	}

	m.MatchStatusId = 2
	response, dErr := service.matchRepo.Update(m)
	return &response.MatchId, dErr
}

func (service *MatchService) RemovePlayer(matchId *int, playerId *int) error {
	return service.matchPlayerRepo.Delete(matchId, playerId)
}
