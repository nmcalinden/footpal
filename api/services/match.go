package services

import (
	"log"
	"sort"
	"sync"

	"github.com/hashicorp/go-multierror"
	"github.com/nmcalinden/footpal/mappers"
	"github.com/nmcalinden/footpal/models"
	"github.com/nmcalinden/footpal/repository"
	"github.com/nmcalinden/footpal/views"
)

type MatchService struct {
	matchRepo       repository.MatchRepositoryI
	matchPlayerRepo repository.MatchPlayerRepositoryI
	squadRepo       repository.SquadRepositoryI
}

func NewMatchService(matchRepo repository.MatchRepositoryI, matchPlayerRepo repository.MatchPlayerRepositoryI,
	squadRepo repository.SquadRepositoryI) *MatchService {
	return &MatchService{
		matchRepo:       matchRepo,
		matchPlayerRepo: matchPlayerRepo,
		squadRepo:       squadRepo,
	}
}

func (s *MatchService) GetMatches() (*[]views.MatchSummary, error) {
	m, err := s.matchRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var res []views.MatchSummary
	nM := len(*m)
	mC, eC := make(chan views.MatchSummary, nM), make(chan error)
	var wg sync.WaitGroup
	wg.Add(nM)

	for _, match := range *m {
		match := match
		go func() {
			defer wg.Done()
			msr, err := s.retrieveMatchDetail(match)
			if err != nil {
				mC <- views.MatchSummary{}
				eC <- err
				return
			}

			mC <- *msr
		}()
	}

	wg.Wait()
	close(eC)
	close(mC)

	var mErr error
	for e := range eC {
		mErr = multierror.Append(mErr, e)
	}
	if mErr != nil {
		return nil, mErr
	}

	for ms := range mC {
		res = append(res, ms)
	}

	sort.Slice(res[:], func(i, j int) bool {
		return *res[i].MatchId < *res[j].MatchId
	})
	return &res, nil
}

func (s *MatchService) retrieveMatchDetail(match models.Match) (*views.MatchSummary, error) {
	b, err := s.matchRepo.FindMatchDetailByBookingIdAndMatchDate(match.BookingId, match.MatchDate)
	if err != nil {
		return nil, err
	}
	var sq = new(models.Squad)
	if match.SquadId != nil {
		sq, _ = s.squadRepo.FindById(match.SquadId)
	}

	var ms views.MatchSummary
	err = mappers.MapToMatchSummaryView(&ms, *match.MatchId, *b, *sq)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &ms, nil
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
