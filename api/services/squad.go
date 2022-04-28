package services

import (
	"github.com/nmcalinden/footpal/api/models"
	"github.com/nmcalinden/footpal/api/payloads"
	"github.com/nmcalinden/footpal/api/repository"
)

type SquadService struct {
	playerRepo repository.PlayerRepositoryI
	squadRepo  repository.SquadRepositoryI
}

func NewSquadService(playerRepo repository.PlayerRepositoryI, squadRepo repository.SquadRepositoryI) *SquadService {
	return &SquadService{
		playerRepo: playerRepo,
		squadRepo:  squadRepo,
	}
}

func (s *SquadService) GetSquads() (*[]models.Squad, error) {
	return s.squadRepo.FindAll()
}

func (s *SquadService) GetSquadById(squadId *int) (*models.Squad, error) {
	return s.squadRepo.FindById(squadId)
}

func (s *SquadService) GetAllPlayersBySquad(squadId *int) (*[]models.SquadPlayerDetails, error) {
	return s.squadRepo.FindPlayersBySquadId(squadId)
}

func (s *SquadService) GetSquadByPlayer(playerId *int, squadId *int) (*models.Squad, error) {
	return s.squadRepo.FindSquadByPlayerId(squadId, playerId)
}

func (s *SquadService) CreateNewSquad(squadRequest *payloads.SquadRequest) (*int, error) {
	newSquad := models.Squad{
		Name: squadRequest.Name,
		City: squadRequest.City,
	}
	return s.squadRepo.Save(&newSquad)
}

func (s *SquadService) EditSquad(squadId *int, squadRequest *payloads.SquadRequest) (*models.Squad, error) {
	p, err := s.squadRepo.FindById(squadId)
	if err != nil {
		return nil, err
	}

	p.Name = squadRequest.Name
	p.City = squadRequest.City
	return s.squadRepo.Update(p)
}

func (s *SquadService) DeleteSquad(squadId *int) error {
	p, err := s.squadRepo.FindById(squadId)
	if err != nil {
		return err
	}
	return s.squadRepo.Delete(p.SquadId)
}

func (s *SquadService) ApprovePlayer(squadId *int, playerId *int) error {
	sq, err := s.squadRepo.FindSquadByPlayerId(squadId, playerId)
	if err != nil {
		return err
	}

	return s.squadRepo.UpdatePlayerStatus(sq.SquadId, playerId, 1)
}

func (s *SquadService) RemovePlayer(squadId *int, playerId *int) error {
	sq, err := s.squadRepo.FindSquadByPlayerId(squadId, playerId)
	if err != nil {
		return err
	}

	return s.squadRepo.UpdatePlayerStatus(sq.SquadId, playerId, 3)

}
