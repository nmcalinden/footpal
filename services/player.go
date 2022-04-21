package services

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/nmcalinden/footpal/models"
	"github.com/nmcalinden/footpal/repository"
	"sync"
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

func (service *PlayerService) GetAllSquadsByPlayer(userId *int) (*[]models.Squad, error) {
	p, err := service.playerRepo.FindByUserId(userId)
	if err != nil {
		return nil, err
	}
	return service.squadRepo.FindAllByPlayerId(&p.PlayerId)
}

func (service *PlayerService) GetSquadByPlayer(userId *int, squadId *int) (*models.Squad, error) {
	p, err := service.playerRepo.FindByUserId(userId)
	if err != nil {
		return nil, err
	}
	return service.squadRepo.FindSquadByPlayerId(squadId, &p.PlayerId)
}

func (service *PlayerService) GetMatchesByPlayer(userId *int) (*[]models.Match, error) {
	p, err := service.playerRepo.FindByUserId(userId)
	if err != nil {
		return nil, err
	}
	return service.matchPlayerRepo.FindMatchesByPlayer(&p.PlayerId)
}

func (service *PlayerService) EditPlayer(userId *int, playerRequest *models.PlayerRequest) (*models.Player, error) {
	p, err := service.playerRepo.FindByUserId(userId)
	if err != nil {
		return nil, err
	}

	p.Nickname = sql.NullString{
		String: playerRequest.Nickname,
		Valid:  true,
	}
	p.PhoneNo = playerRequest.PhoneNo
	p.Postcode = playerRequest.Postcode
	p.City = playerRequest.City

	return service.playerRepo.Update(p)
}

func (service *PlayerService) JoinSquad(userId *int, squadId *int) error {
	p, err := service.playerRepo.FindByUserId(userId)
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

func (service *PlayerService) JoinMatch(userId *int, matchId *int) (*int, error) {
	var wg sync.WaitGroup
	wg.Add(2)

	playerChan, matchChan, errsChan := make(chan int), make(chan int), make(chan error)
	go getId(service.getPlayerId, userId, &wg, playerChan, errsChan)
	go getId(service.getMatchId, matchId, &wg, matchChan, errsChan)

	go func() {
		wg.Wait()
		close(errsChan)
	}()

	pId, mId, gErrs := <-playerChan, <-matchChan, <-errsChan
	if gErrs != nil {
		return nil, gErrs
	}

	matchPlayer := models.MatchPlayer{
		MatchId:       &mId,
		PlayerId:      &pId,
		AmountToPay:   0,
		PaymentTypeId: 1,
	}
	res, dErr := service.matchPlayerRepo.Save(matchPlayer)
	if dErr != nil {
		return nil, dErr
	}

	return res, nil
}

func (service *PlayerService) LeaveMatch(userId *int, matchId *int) error {
	var wg sync.WaitGroup
	wg.Add(2)

	// Go routines example
	var errs []error

	var pId int
	go func() {
		defer wg.Done()
		player, err := service.playerRepo.FindByUserId(userId)
		if err != nil {
			errs = append(errs, err)
		}
		pId = player.PlayerId
	}()

	var mId int
	go func() {
		defer wg.Done()
		match, err := service.matchRepo.FindById(matchId)
		if err != nil {
			errs = append(errs, err)
		}
		mId = match.MatchId
	}()

	wg.Wait()
	if len(errs) != 0 {
		return errs[0]
	}

	dErr := service.matchPlayerRepo.Delete(&mId, &pId)
	if dErr != nil {
		return dErr
	}
	return nil
}

func (service *PlayerService) Pay(userId *int, matchId *int, amountToPay *float32) error {
	p, err := service.playerRepo.FindByUserId(userId)
	if err != nil {
		return err
	}

	res, dErr := service.matchPlayerRepo.FindByMatchIdAndPlayerId(matchId, &p.PlayerId)
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

func (service *PlayerService) UpdatePaymentMethod(matchId *int, userId *int, paymentTypeId *int) error {
	p, err := service.playerRepo.FindByUserId(userId)
	if err != nil {
		return err
	}

	res, err := service.matchPlayerRepo.FindByMatchIdAndPlayerId(matchId, &p.PlayerId)
	if err != nil {
		return err
	}

	res.PaymentTypeId = *paymentTypeId
	_, dErr := service.matchPlayerRepo.Update(res)
	if dErr != nil {
		return dErr
	}
	return nil
}

func getId(f func(*int, chan<- int) error, id *int, wg *sync.WaitGroup, ch chan<- int, errs chan error) {
	defer wg.Done()
	defer close(ch)
	err := f(id, ch)
	if err != nil {
		ch <- 0
		errs <- err
	}
}

func (service PlayerService) getPlayerId(userId *int, ch chan<- int) error {
	p, err := service.playerRepo.FindByUserId(userId)
	if err != nil {
		return err
	}
	ch <- p.PlayerId
	return nil
}

func (service *PlayerService) getMatchId(matchId *int, ch chan<- int) error {
	m, err := service.matchRepo.FindById(matchId)
	if err != nil {
		return err
	}
	ch <- m.MatchId
	return nil
}
