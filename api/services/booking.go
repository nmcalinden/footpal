package services

import (
	"errors"
	"github.com/nmcalinden/footpal/api/models"
	"github.com/nmcalinden/footpal/api/payloads"
	"github.com/nmcalinden/footpal/api/repository"
	"time"
)

type BookingService struct {
	bookingRepo repository.BookingRepositoryI
}

func NewBookingService(bookingRepo repository.BookingRepositoryI) *BookingService {
	return &BookingService{bookingRepo: bookingRepo}
}

func (s *BookingService) GetBookings() (*[]models.Booking, error) {
	return s.bookingRepo.FindAll()
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
			fmd, err := time.Parse("2006-01-02", md)
			if err != nil {
				return nil, err
			}

			fmd = fmd.AddDate(0, 0, 7)
			md = fmd.Format("2006-01-02")
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
			fmd, err := time.Parse("2006-01-02", md)
			if err != nil {
				return nil, err
			}

			fmd = fmd.AddDate(0, 0, 7)
			md = fmd.Format("2006-01-02")
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
