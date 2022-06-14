package services

import (
	"errors"
	"sort"
	"time"

	"github.com/gofiber/fiber/v2"
	errors2 "github.com/nmcalinden/footpal/errors"
	"github.com/nmcalinden/footpal/models"
	"github.com/nmcalinden/footpal/payloads"
	"github.com/nmcalinden/footpal/repository"
	"github.com/nmcalinden/footpal/utils"
	"github.com/nmcalinden/footpal/views"
)

type BookingService struct {
	bookingRepo repository.BookingRepositoryI
	matchRepo   repository.MatchRepositoryI
	statusRepo  repository.StatusRepositoryI
}

func NewBookingService(bookingRepo repository.BookingRepositoryI, matchRepo repository.MatchRepositoryI,
	statusRepo repository.StatusRepositoryI) *BookingService {
	return &BookingService{bookingRepo: bookingRepo, matchRepo: matchRepo, statusRepo: statusRepo}
}

func (s *BookingService) GetBookings(userId *int) (*[]views.UserBooking, error) {
	var userBookings []views.UserBooking
	b, err := s.bookingRepo.FindAllByUserId(userId)
	if err != nil {
		return nil, err
	}

	if b == nil {
		return nil, errors2.GetError(errors2.NoResults, fiber.StatusOK, "No Results")
	}
	for _, booking := range *b {
		res, err := s.buildUserBooking(booking)
		if err != nil {
			return nil, err
		}
		userBookings = append(userBookings, *res)
	}

	sort.Slice(userBookings[:], func(i, j int) bool {
		md1, _ := utils.ParseDateFromTimestampString(userBookings[i].MatchDate)
		md2, _ := utils.ParseDateFromTimestampString(userBookings[j].MatchDate)
		return md1.Before(md2)
	})

	return &userBookings, err
}

func (s *BookingService) buildUserBooking(booking models.Booking) (*views.UserBooking, error) {
	m, err := s.bookingRepo.FindMatchesByBookingId(booking.BookingId)
	if err != nil || len(*m) == 0 {
		return nil, err
	}

	mbd, err := s.matchRepo.FindMatchDetailsByBookingId(booking.BookingId)
	if err != nil {
		return nil, err
	}

	var matchSummary []views.BookingMatchSummary
	var bookingPaid = true
	for _, match := range *m {
		ms := views.BookingMatchSummary{
			MatchId:   *match.MatchId,
			MatchDate: match.MatchDate,
		}

		if !match.IsPaid {
			bookingPaid = false
		}
		matchSummary = append(matchSummary, ms)
	}

	bStatus, err := s.statusRepo.FindBookingStatusById(&booking.BookingStatusId)
	if err != nil {
		return nil, err
	}

	matches := *mbd
	totalCost := matches[0].Cost * float32(len(matches))
	res := views.UserBooking{
		BookingId:     *booking.BookingId,
		BookingStatus: *bStatus,
		MatchDate:     matches[0].MatchDate,
		StartTime:     utils.GetFormattedTime(matches[0].StartTime),
		NoOfWeeks:     len(matches),
		TotalCost:     totalCost,
		IsBookingPaid: bookingPaid,
		Venue: views.BookingVenueSummary{
			VenueId: matches[0].VenueId,
			Name:    matches[0].VenueName,
		},
		Pitch: views.BookingPitchSummary{
			PitchId: matches[0].PitchId,
			Name:    matches[0].PitchName,
		},
		Matches: matchSummary,
	}
	return &res, nil
}

func (s *BookingService) GetBookingById(bookingId *int) (*models.Booking, error) {
	return s.bookingRepo.FindById(bookingId)
}

func (s *BookingService) GetMatchesByBooking(bookingId *int) (*[]models.Match, error) {
	return s.bookingRepo.FindMatchesByBookingId(bookingId)
}

func (s *BookingService) CreateNewBooking(bookingRequest *payloads.BookingRequest, userId int) (*int, error) {
	m, err := s.bookingRepo.IsExistingMatchPresent(&bookingRequest.MatchDate, &bookingRequest.PitchTimeSlotId)
	if err != nil {
		return nil, err
	}

	if *m {
		return nil, errors.New("bookingAlreadyExists")
	}

	newBooking := models.Booking{
		BookingStatusId: 1,
		CreatedBy:       userId,
		Created:         time.Now(),
		LastUpdated:     time.Now(),
	}

	ms, err := buildMatchesFromBooking(bookingRequest)
	ps, err := buildPitchSlotsFromBooking(bookingRequest)
	if err != nil {
		return nil, err
	}
	return s.bookingRepo.Save(&newBooking, ms, ps)
}

func buildMatchesFromBooking(br *payloads.BookingRequest) (*[]models.Match, error) {
	var ms []models.Match
	match := models.Match{
		BookingId:           0,
		MatchAccessStatusId: 1,
		MatchStatusId:       1,
		SquadId:             br.SquadId,
		MatchDate:           br.MatchDate,
		Cost:                30,
		IsPaid:              false,
		Created:             time.Now(),
		LastUpdated:         time.Now(),
	}

	ms = append(ms, match)

	md := br.MatchDate

	if br.NoOfWeeks > 1 {
		var i = 1

		for i < br.NoOfWeeks {
			fmd, err := utils.ParseDateFromString(md)
			if err != nil {
				return nil, err
			}

			fmd = fmd.AddDate(0, 0, 7)
			md = utils.GetFormattedDate(fmd)
			match = models.Match{
				BookingId:           0,
				MatchAccessStatusId: 1,
				MatchStatusId:       4,
				SquadId:             br.SquadId,
				MatchDate:           md,
				Cost:                30,
				IsPaid:              false,
				Created:             time.Now(),
				LastUpdated:         time.Now(),
			}

			ms = append(ms, match)
			i += 1
		}
	}

	return &ms, nil
}

func buildPitchSlotsFromBooking(br *payloads.BookingRequest) (*[]models.PitchSlot, error) {
	var ps []models.PitchSlot
	pitchSlot := models.PitchSlot{
		PitchTimeSlotId: br.PitchTimeSlotId,
		MatchDate:       br.MatchDate,
		BookingStatusId: 1,
	}

	ps = append(ps, pitchSlot)

	md := br.MatchDate

	if br.NoOfWeeks > 1 {
		var i = 1

		for i < br.NoOfWeeks {
			fmd, err := utils.ParseDateFromString(md)
			if err != nil {
				return nil, err
			}

			fmd = fmd.AddDate(0, 0, 7)
			md = utils.GetFormattedDate(fmd)
			pitchSlot = models.PitchSlot{
				PitchTimeSlotId: br.PitchTimeSlotId,
				MatchDate:       md,
				BookingStatusId: 1,
			}

			ps = append(ps, pitchSlot)
			i++
		}
	}

	return &ps, nil
}

func (s *BookingService) EditBooking(bookingId *int, bookingRequest *payloads.BookingRequest) (*models.Booking, error) {
	b, err := s.bookingRepo.FindById(bookingId)
	if err != nil {
		return nil, err
	}

	b.LastUpdated = time.Now()
	b.CreatedBy = 1 //TODO - Test example
	return s.bookingRepo.Update(b)
}

func (s *BookingService) CancelBooking(bookingId *int) (*int, error) {
	b, err := s.bookingRepo.FindById(bookingId)
	if err != nil {
		return nil, err
	}

	b.BookingStatusId = 2
	response, dErr := s.bookingRepo.Update(b)
	return response.BookingId, dErr
}

func (s *BookingService) FindVenuesWithAvailableBookings(b *payloads.BookingSearchRequest) (*[]models.Venue, error) {
	v, err := s.bookingRepo.FindAvailableVenues(b.VenueId, b.Date, b.City, b.MaxPlayers)

	if err != nil {
		return nil, err
	}

	return v, nil
}
