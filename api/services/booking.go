package services

import (
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

func (s *BookingService) CreateNewBooking(bookingRequest *payloads.BookingRequest) (*int, error) {
	newBooking := models.Booking{
		BookingStatusId: 4,
		CreatedBy:       bookingRequest.VenueId,
		Created:         time.Now(),
		LastUpdated:     time.Now(),
	}
	return s.bookingRepo.Save(&newBooking)
}

func (s *BookingService) EditBooking(bookingId *int, bookingRequest *payloads.BookingRequest) (*models.Booking, error) {
	b, err := s.bookingRepo.FindById(bookingId)
	if err != nil {
		return nil, err
	}

	b.LastUpdated = time.Now()
	b.CreatedBy = bookingRequest.SquadId //TODO - Test example
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
