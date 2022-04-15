package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/nmcalinden/footpal/models"
	"github.com/nmcalinden/footpal/repository"
)

type PlayerService struct {
	playerRepo      *repository.PlayerRepository
	matchRepo       *repository.MatchRepository
	matchPlayerRepo *repository.MatchPlayerRepository
	squadRepo       *repository.SquadRepository
}

func NewPlayerService(database *sqlx.DB) *PlayerService {
	return &PlayerService{
		playerRepo:      repository.NewPlayerRepository(database),
		matchRepo:       repository.NewMatchRepository(database),
		matchPlayerRepo: repository.NewMatchPlayerRepository(database),
		squadRepo:       repository.NewSquadRepository(database),
	}
}

func (service *PlayerService) GetPlayers() (*[]models.Player, error) {
	return service.playerRepo.FindAll()
}

func (service *PlayerService) GetPlayerById(playerId *int) (*models.Player, error) {
	return service.playerRepo.FindById(playerId)
}

func (service *PlayerService) GetAllSquadsByPlayer(playerId *int) (*[]models.Squad, error) {
	return service.squadRepo.FindAllByPlayerId(playerId)
}

func (service *PlayerService) GetSquadByPlayer(playerId *int, squadId *int) (*models.Squad, error) {
	return service.squadRepo.FindSquadByPlayerId(squadId, playerId)
}

func (service *PlayerService) GetMatchesByPlayer(playerId *int) (*[]models.Match, error) {
	return service.matchPlayerRepo.FindMatchesByPlayer(playerId)
}

func (service *PlayerService) EditPlayer(playerId *int, playerRequest *models.PlayerRequest) (*models.Player, error) {
	p, err := service.playerRepo.FindById(playerId)
	if err != nil {
		return nil, err
	}

	p.Nickname = playerRequest.Nickname
	p.PhoneNo = playerRequest.PhoneNo
	p.Postcode = playerRequest.Postcode
	p.City = playerRequest.City

	return service.playerRepo.Update(p)
}

func (service *PlayerService) JoinSquad(playerId *int, squadId *int) error {
	p, err := service.playerRepo.FindById(playerId)
	if err != nil {
		return err
	}

	s, err := service.squadRepo.FindById(squadId)
	if err != nil {
		return err
	}

	squadPlayer := models.SquadPlayer{
		SquadPlayerId:       0,
		SquadId:             s.SquadId,
		PlayerId:            p.PlayerId,
		Role:                "player",
		SquadPlayerStatusId: 4,
	}
	err = service.squadRepo.AddPlayer(squadPlayer)
	if err != nil {
		return err
	}

	return nil
}

func (service *PlayerService) JoinMatch(playerId *int, matchId *int) (*int, error) {
	p, err := service.playerRepo.FindById(playerId)
	if err != nil {
		return nil, err
	}

	m, err := service.matchRepo.FindById(matchId)
	if err != nil {
		return nil, err
	}

	matchPlayer := models.MatchPlayer{
		MatchPlayerId: 0,
		MatchId:       m.MatchId,
		PlayerId:      p.PlayerId,
		AmountToPay:   0,
		PaymentTypeId: 1,
	}
	res, dErr := service.matchPlayerRepo.Save(matchPlayer)
	if dErr != nil {
		return nil, dErr
	}

	return res, nil
}

func (service *PlayerService) LeaveMatch(playerId *int, matchId *int) error {
	// Go Routines wait groups and goroutines
	p, err := service.playerRepo.FindById(playerId)
	if err != nil {
		return err
	}

	m, err := service.matchRepo.FindById(matchId)
	if err != nil {
		return err
	}

	dErr := service.matchPlayerRepo.Delete(&m.MatchId, &p.PlayerId)
	if dErr != nil {
		return dErr
	}
	return nil
}

func (service *PlayerService) Pay(playerId *int, matchId *int, amountToPay *float32) error {
	res, dErr := service.matchPlayerRepo.FindByMatchIdAndPlayerId(matchId, playerId)
	if dErr != nil {
		return dErr
	}

	if res.AmountToPay > 0 && *amountToPay <= res.AmountToPay {
		newAmount := res.AmountToPay - *amountToPay
		res.AmountToPay = newAmount

		_, err := service.matchPlayerRepo.Update(res)
		if err != nil {
			return err
		}
	}
	return nil
}

func (service *PlayerService) UpdatePaymentMethod(matchId *int, playerId *int, paymentTypeId *int) error {
	p, err := service.matchPlayerRepo.FindByMatchIdAndPlayerId(matchId, playerId)
	if err != nil {
		return err
	}

	p.PaymentTypeId = *paymentTypeId
	_, dErr := service.matchPlayerRepo.Update(p)
	if dErr != nil {
		return dErr
	}
	return nil
}
