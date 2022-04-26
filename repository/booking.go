package repository

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/nmcalinden/footpal/models"
)

type BookingRepositoryI interface {
	FindAll() (*[]models.Booking, error)
	FindById(id *int) (*models.Booking, error)
	FindMatchesByBookingId(id *int) (*[]models.Match, error)
	Save(booking *models.Booking) (*int, error)
	Update(booking *models.Booking) (*models.Booking, error)
}

type BookingRepository struct {
	database *sqlx.DB
}

var BookingRepoSet = wire.NewSet(NewBookingRepository, wire.Bind(new(BookingRepositoryI), new(*BookingRepository)))

func NewBookingRepository(database *sqlx.DB) *BookingRepository {
	return &BookingRepository{database: database}
}

func (r BookingRepository) FindAll() (*[]models.Booking, error) {
	var bookingRecords []models.Booking
	err := r.database.Select(&bookingRecords, "SELECT * FROM footpaldb.public.booking")
	if err != nil || len(bookingRecords) == 0 {
		return nil, err
	}
	return &bookingRecords, nil
}

func (r BookingRepository) FindById(id *int) (*models.Booking, error) {
	var booking models.Booking
	err := r.database.Get(&booking, "SELECT * FROM footpaldb.public.booking WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &booking, nil
}

func (r BookingRepository) FindMatchesByBookingId(id *int) (*[]models.Match, error) {
	var matches []models.Match
	err := r.database.Select(&matches, "SELECT * FROM footpaldb.public.match WHERE booking_id = $1", id)
	if err != nil {
		return nil, err
	}
	return &matches, nil
}

func (r BookingRepository) Save(booking *models.Booking) (*int, error) {
	_, err := r.database.NamedExec(`INSERT INTO footpaldb.public.booking(booking_status_id, created_by, created, last_updated)
							VALUES(:booking_status_id, :created_by, :created, :last_updated)`, booking)

	if err != nil {
		return nil, err
	}
	return booking.BookingId, nil
}

func (r BookingRepository) Update(booking *models.Booking) (*models.Booking, error) {
	_, err := r.database.NamedExec(`UPDATE footpaldb.public.booking SET booking_status_id=:booking_status_id, 
                                    created_by=:created_by, created=:created, last_updated=:last_updated WHERE id=:id`, booking)

	if err != nil {
		return nil, err
	}
	return booking, nil
}
