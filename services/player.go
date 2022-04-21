package services

import (
	"database/sql"
	"github.com/hashicorp/go-multierror"
	"github.com/jmoiron/sqlx"
	"github.com/nmcalinden/footpal/models"
	"github.com/nmcalinden/footpal/payloads"
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

func (s *PlayerService) GetPlayers() (*[]models.Player, error) {
	return s.playerRepo.FindAll()
}

func (s *PlayerService) GetPlayerById(playerId *int) (*models.Player, error) {
	return s.playerRepo.FindById(playerId)
}

func (s *PlayerService) GetAllSquadsByPlayer(userId *int) (*[]models.Squad, error) {
	p, err := s.playerRepo.FindByUserId(userId)
	if err != nil {
		return nil, err
	}
	return s.squadRepo.FindAllByPlayerId(p.PlayerId)
}

func (s *PlayerService) GetSquadByPlayer(userId *int, squadId *int) (*models.Squad, error) {
	p, err := s.playerRepo.FindByUserId(userId)
	if err != nil {
		return nil, err
	}
	return s.squadRepo.FindSquadByPlayerId(squadId, p.PlayerId)
}

func (s *PlayerService) GetMatchesByPlayer(userId *int) (*[]models.Match, error) {
	p, err := s.playerRepo.FindByUserId(userId)
	if err != nil {
		return nil, err
	}
	return s.matchPlayerRepo.FindMatchesByPlayer(p.PlayerId)
}

func (s *PlayerService) EditPlayer(userId *int, playerRequest *payloads.PlayerRequest) (*models.Player, error) {
	p, err := s.playerRepo.FindByUserId(userId)
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

	return s.playerRepo.Update(p)
}

func (s *PlayerService) JoinSquad(userId *int, squadId *int) error {
	p, err := s.playerRepo.FindByUserId(userId)
	if err != nil {
		return err
	}

	sq, err := s.squadRepo.FindById(squadId)
	if err != nil {
		return err
	}

	squadPlayer := models.SquadPlayer{
		SquadId:             *sq.SquadId,
		PlayerId:            *p.PlayerId,
		Role:                "player",
		SquadPlayerStatusId: 4,
	}
	err = s.squadRepo.AddPlayer(squadPlayer)
	if err != nil {
		return err
	}

	return nil
}

func (s *PlayerService) JoinMatch(userId *int, matchId *int) (*int, error) {
	var wg sync.WaitGroup
	wg.Add(2)

	playerChan, matchChan, errsChan := make(chan int), make(chan int), make(chan error)
	go getId(s.getPlayerId, userId, &wg, playerChan, errsChan)
	go getId(s.getMatchId, matchId, &wg, matchChan, errsChan)

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
	res, dErr := s.matchPlayerRepo.Save(matchPlayer)
	if dErr != nil {
		return nil, dErr
	}

	return res, nil
}

func (s *PlayerService) LeaveMatch(userId *int, matchId *int) error {
	var wg sync.WaitGroup
	wg.Add(2)

	// Go routines example
	var errs []error

	var pId int
	go func() {
		defer wg.Done()
		player, err := s.playerRepo.FindByUserId(userId)
		if err != nil {
			errs = append(errs, err)
		}
		pId = *player.PlayerId
	}()

	var mId int
	go func() {
		defer wg.Done()
		match, err := s.matchRepo.FindById(matchId)
		if err != nil {
			errs = append(errs, err)
		}
		mId = *match.MatchId
	}()

	wg.Wait()
	if len(errs) != 0 {
		var err error
		return multierror.Append(err, errs...)
	}

	dErr := s.matchPlayerRepo.Delete(&mId, &pId)
	if dErr != nil {
		return dErr
	}
	return nil
}

func (s *PlayerService) Pay(userId *int, matchId *int, amountToPay *float32) error {
	p, err := s.playerRepo.FindByUserId(userId)
	if err != nil {
		return err
	}

	res, dErr := s.matchPlayerRepo.FindByMatchIdAndPlayerId(matchId, p.PlayerId)
	if dErr != nil {
		return dErr
	}

	if res.AmountToPay > 0 && *amountToPay <= res.AmountToPay {
		newAmount := res.AmountToPay - *amountToPay
		res.AmountToPay = newAmount

		_, err := s.matchPlayerRepo.Update(res)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *PlayerService) UpdatePaymentMethod(matchId *int, userId *int, paymentTypeId *int) error {
	p, err := s.playerRepo.FindByUserId(userId)
	if err != nil {
		return err
	}

	res, err := s.matchPlayerRepo.FindByMatchIdAndPlayerId(matchId, p.PlayerId)
	if err != nil {
		return err
	}

	res.PaymentTypeId = *paymentTypeId
	_, dErr := s.matchPlayerRepo.Update(res)
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

func (s PlayerService) getPlayerId(userId *int, ch chan<- int) error {
	p, err := s.playerRepo.FindByUserId(userId)
	if err != nil {
		return err
	}
	ch <- *p.PlayerId
	return nil
}

func (s *PlayerService) getMatchId(matchId *int, ch chan<- int) error {
	m, err := s.matchRepo.FindById(matchId)
	if err != nil {
		return err
	}
	ch <- *m.MatchId
	return nil
}
