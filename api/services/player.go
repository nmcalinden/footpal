package services

import (
	"log"
	"sort"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/hashicorp/go-multierror"
	errors2 "github.com/nmcalinden/footpal/errors"
	"github.com/nmcalinden/footpal/mappers"
	"github.com/nmcalinden/footpal/models"
	"github.com/nmcalinden/footpal/payloads"
	"github.com/nmcalinden/footpal/repository"
	"github.com/nmcalinden/footpal/utils"
	"github.com/nmcalinden/footpal/views"
	"gopkg.in/guregu/null.v4"
)

const (
	playersQuery = "/players?limit=%d&after_id=%d"
)

type PlayerService struct {
	playerRepo      repository.PlayerRepositoryI
	userRepo        repository.UserRepositoryI
	matchRepo       repository.MatchRepositoryI
	matchPlayerRepo repository.MatchPlayerRepositoryI
	squadRepo       repository.SquadRepositoryI
}

func NewPlayerService(playerRepo repository.PlayerRepositoryI, userRepo repository.UserRepositoryI,
	matchRepo repository.MatchRepositoryI, matchPlayerRepo repository.MatchPlayerRepositoryI,
	squadRepo repository.SquadRepositoryI) *PlayerService {
	return &PlayerService{
		playerRepo:      playerRepo,
		userRepo:        userRepo,
		matchRepo:       matchRepo,
		matchPlayerRepo: matchPlayerRepo,
		squadRepo:       squadRepo,
	}
}

func (s *PlayerService) GetPlayers(limit int, after int) (*views.Players, error) {
	t, err := s.playerRepo.GetTotal()
	if err != nil {
		return nil, err
	}

	players, err := s.playerRepo.FindAll(limit, after)
	if err != nil {
		return nil, err
	}

	ps, err := s.buildPlayers(players)
	if err != nil {
		return nil, err
	}

	response := buildPlayersResponse(limit, after, *ps, t)
	return &response, nil
}

func (s *PlayerService) buildPlayers(players *[]models.Player) (*[]views.Player, error) {
	var ps []views.Player

	numP := len(*players)
	playerChan, errorChan := make(chan views.Player, numP), make(chan error)
	var wg sync.WaitGroup
	wg.Add(numP)
	for _, player := range *players {
		go s.createPlayerView(player, &wg, playerChan, errorChan)
	}

	wg.Wait()
	close(errorChan)
	close(playerChan)

	var mErr error
	for e := range errorChan {
		mErr = multierror.Append(mErr, e)
	}
	if mErr != nil {
		return nil, mErr
	}

	for pC := range playerChan {
		ps = append(ps, pC)
	}
	sort.Slice(ps[:], func(i, j int) bool {
		return *ps[i].PlayerId < *ps[j].PlayerId
	})
	return &ps, nil
}

func (s PlayerService) createPlayerView(player models.Player, wg *sync.WaitGroup, pChan chan views.Player, errorChan chan<- error) {
	defer wg.Done()
	u, err := s.userRepo.FindById(&player.UserId)
	if err != nil {
		pChan <- views.Player{}
		errorChan <- err
		return
	}

	var p views.Player
	err = mappers.MapToPlayerView(&p, player, *u)
	if err != nil {
		pChan <- views.Player{}
		errorChan <- err
		return
	}

	pChan <- p
}

func buildPlayersResponse(limit int, after int, ps []views.Player, t *int) views.Players {
	lastId := ps[len(ps)-1].PlayerId
	pag := utils.BuildPagination(*t, limit, len(ps), after, *lastId, playersQuery)
	response := views.Players{
		Pagination: pag,
		Data:       ps,
	}
	return response
}

func (s *PlayerService) GetPlayerById(playerId *int) (*views.Player, error) {
	p, err := s.playerRepo.FindById(playerId)
	if err != nil {
		log.Println(err)
		return nil, errors2.GetError(errors2.RecordNotFound, fiber.StatusNotFound, "Player does not exist")
	}

	u, err := s.userRepo.FindById(&p.UserId)
	if err != nil {
		log.Println(err)
		return nil, errors2.GetError(errors2.RecordNotFound, fiber.StatusNotFound, "User does not exist")
	}

	var player views.Player
	err = mappers.MapToPlayerView(&player, *p, *u)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &player, nil
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

func (s *PlayerService) EditPlayer(userId *int, playerRequest *payloads.PlayerRequest) (*views.PlayerUser, error) {
	p, err := s.playerRepo.FindByUserId(userId)
	if err != nil {
		return nil, err
	}

	test := null.StringFromPtr(&playerRequest.Nickname).Ptr()
	p.Nickname = test
	p.PhoneNo = playerRequest.PhoneNo
	p.Postcode = playerRequest.Postcode
	p.City = playerRequest.City

	player, err := s.playerRepo.Update(p)
	if err != nil {
		return nil, err
	}

	usr, _ := s.userRepo.FindById(userId)

	var user views.PlayerUser
	err = mappers.MapToUser(&user, *player, *usr)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &user, nil
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
