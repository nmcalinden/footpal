package services

import (
	"github.com/nmcalinden/footpal/models"
	"github.com/nmcalinden/footpal/repository"
)

type MatchService struct {
	matchRepo       repository.MatchRepositoryI
	matchPlayerRepo repository.MatchPlayerRepositoryI
}

func NewMatchService(matchRepo repository.MatchRepositoryI, matchPlayerRepo repository.MatchPlayerRepositoryI) *MatchService {
	return &MatchService{
		matchRepo:       matchRepo,
		matchPlayerRepo: matchPlayerRepo,
	}
}

func (s *MatchService) GetMatches() (*[]models.Match, error) {
	return s.matchRepo.FindAll()
}

func (s *MatchService) GetMatchById(matchId *int) (*models.Match, error) {
	return s.matchRepo.FindById(matchId)
}

func (s *MatchService) GetMatchPlayers(matchId *int) (*[]models.MatchPlayer, error) {
	return s.matchPlayerRepo.FindByMatchId(matchId)
}

func (s *MatchService) CancelMatch(matchId *int) (*int, error) {
	m, err := s.matchRepo.FindById(matchId)
	if err != nil {
		return nil, err
	}

	m.MatchStatusId = 2
	response, dErr := s.matchRepo.Update(m)
	return response.MatchId, dErr
}

func (s *MatchService) RemovePlayer(matchId *int, playerId *int) error {
	return s.matchPlayerRepo.Delete(matchId, playerId)
}
