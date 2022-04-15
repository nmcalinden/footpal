package services

import (
	"github.com/jmoiron/sqlx"
	"github.com/nmcalinden/footpal/models"
	"github.com/nmcalinden/footpal/repository"
	"time"
)

type BookingService struct {
	bookingRepo *repository.BookingRepository
}

func NewBookingService(database *sqlx.DB) *BookingService {
	return &BookingService{bookingRepo: repository.NewBookingRepository(database)}
}

func (service *BookingService) GetBookings() (*[]models.Booking, error) {
	return service.bookingRepo.FindAll()
}

func (service *BookingService) GetBookingById(bookingId *int) (*models.Booking, error) {
	return service.bookingRepo.FindById(bookingId)
}

func (service *BookingService) GetMatchesByBooking(bookingId *int) (*[]models.Match, error) {
	return service.bookingRepo.FindMatchesByBookingId(bookingId)
}

func (service *BookingService) CreateNewBooking(bookingRequest *models.BookingRequest) (*int, error) {
	newBooking := models.Booking{
		BookingId:       0,
		BookingStatusId: 4,
		CreatedBy:       bookingRequest.VenueId,
		Created:         time.Now(),
		LastUpdated:     time.Now(),
	}
	return service.bookingRepo.Save(&newBooking)
}

func (service *BookingService) EditBooking(bookingId *int, bookingRequest *models.BookingRequest) (*models.Booking, error) {
	b, err := service.bookingRepo.FindById(bookingId)
	if err != nil {
		return nil, err
	}

	b.LastUpdated = time.Now()
	b.CreatedBy = bookingRequest.SquadId //TODO - Test example
	return service.bookingRepo.Update(b)
}

func (service *BookingService) CancelBooking(bookingId *int) (*int, error) {
	b, err := service.bookingRepo.FindById(bookingId)
	if err != nil {
		return nil, err
	}

	b.BookingStatusId = 2
	response, dErr := service.bookingRepo.Update(b)
	return &response.BookingId, dErr
}
