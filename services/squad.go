package services

import (
	"github.com/jmoiron/sqlx"
	"github.com/nmcalinden/footpal/models"
	"github.com/nmcalinden/footpal/repository"
)

type SquadService struct {
	playerRepo *repository.PlayerRepository
	squadRepo  *repository.SquadRepository
}

func NewSquadService(database *sqlx.DB) *SquadService {
	return &SquadService{
		playerRepo: repository.NewPlayerRepository(database),
		squadRepo:  repository.NewSquadRepository(database),
	}
}

func (service *SquadService) GetSquads() (*[]models.Squad, error) {
	return service.squadRepo.FindAll()
}

func (service *SquadService) GetSquadById(squadId *int) (*models.Squad, error) {
	return service.squadRepo.FindById(squadId)
}

func (service *SquadService) GetAllPlayersBySquad(squadId *int) (*[]models.SquadPlayerDetails, error) {
	return service.squadRepo.FindPlayersBySquadId(squadId)
}

func (service *SquadService) GetSquadByPlayer(playerId *int, squadId *int) (*models.Squad, error) {
	return service.squadRepo.FindSquadByPlayerId(squadId, playerId)
}

func (service *SquadService) CreateNewSquad(squadRequest *models.SquadRequest) (*int, error) {
	newSquad := models.Squad{
		SquadId: 0,
		Name:    squadRequest.Name,
		City:    squadRequest.City,
	}
	return service.squadRepo.Save(&newSquad)
}

func (service *SquadService) EditSquad(squadId *int, squadRequest *models.SquadRequest) (*models.Squad, error) {
	p, err := service.squadRepo.FindById(squadId)
	if err != nil {
		return nil, err
	}

	p.Name = squadRequest.Name
	p.City = squadRequest.City
	return service.squadRepo.Update(p)
}

func (service *SquadService) DeleteSquad(squadId *int) error {
	p, err := service.squadRepo.FindById(squadId)
	if err != nil {
		return err
	}
	return service.squadRepo.Delete(&p.SquadId)
}

func (service *SquadService) ApprovePlayer(squadId *int, playerId *int) error {
	s, err := service.squadRepo.FindSquadByPlayerId(squadId, playerId)
	if err != nil {
		return err
	}

	return service.squadRepo.UpdatePlayerStatus(&s.SquadId, playerId, 1)
}

func (service *SquadService) RemovePlayer(squadId *int, playerId *int) error {
	s, err := service.squadRepo.FindSquadByPlayerId(squadId, playerId)
	if err != nil {
		return err
	}

	return service.squadRepo.UpdatePlayerStatus(&s.SquadId, playerId, 3)

}
