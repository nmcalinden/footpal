package repository

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/nmcalinden/footpal/api/models"
	"strconv"
)

//go:generate mockgen -destination=./mocks/booking_mock.go -package=mocks github.com/nmcalinden/footpal/api/repository BookingRepositoryI

type BookingRepositoryI interface {
	FindAll() (*[]models.Booking, error)
	FindById(id *int) (*models.Booking, error)
	FindMatchesByBookingId(id *int) (*[]models.Match, error)
	Save(booking *models.Booking, matches *[]models.Match, pitchSlots *[]models.PitchSlot) (*int, error)
	Update(booking *models.Booking) (*models.Booking, error)
	FindAvailableVenues(venueId *int, matchDate string, city *string, players *int) (*[]models.Venue, error)
	IsExistingMatchPresent(matchDate *string, pitchTimeslotId *int) (*bool, error)
}

type BookingRepository struct {
	database *sqlx.DB
}

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

func (r BookingRepository) FindAvailableVenues(v *int, m string, c *string, p *int) (*[]models.Venue, error) {
	var vs []models.Venue
	q := "SELECT DISTINCT v.* FROM venue v " +
		"LEFT JOIN pitch p on v.id = p.venue_id JOIN pitch_time_slot pts on p.id = pts.pitch_id " +
		"WHERE pts.id NOT IN (SELECT pitch_time_slot_id FROM pitch_slot where match_date = $1)"

	if v != nil {
		q += " AND v.id = " + strconv.Itoa(*v)
	}

	if c != nil {
		q += " AND v.city = " + *c
	}

	if p != nil {
		q += " AND p.max_players = " + strconv.Itoa(*p)
	}

	err := r.database.Select(&vs, q, m)
	if err != nil {
		return nil, err
	}
	return &vs, nil
}

func (r BookingRepository) IsExistingMatchPresent(matchDate *string, pitchTimeslotId *int) (*bool, error) {
	b := false
	q := "SELECT m.match_status_id, ps.booking_status_id FROM match m JOIN pitch_slot ps on m.match_date = ps.match_date WHERE ps.match_date = $1" +
		"AND ps.pitch_time_slot_id = $2"

	res, err := r.database.Query(q, matchDate, pitchTimeslotId)
	if err != nil {
		return nil, err
	}

	for res.Next() {
		var m int
		var bs int
		err = res.Scan(&m, &bs)
		if err != nil {
			return nil, err
		}
		if m != 2 {
			b = true
		}
	}

	return &b, nil
}

func (r BookingRepository) Save(booking *models.Booking, matches *[]models.Match, pitchSlots *[]models.PitchSlot) (*int, error) {
	tx := r.database.MustBegin()
	var lastInsertedId int
	err := tx.QueryRow(`INSERT INTO footpaldb.public.booking(booking_status_id, created_by, created, last_updated)
							VALUES($1, $2, $3, $4) RETURNING id`,
		booking.BookingStatusId, booking.CreatedBy, booking.Created, booking.LastUpdated).Scan(&lastInsertedId)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, err
		}
		return nil, errors.New("insert booking error")
	}

	err = saveMatches(tx, matches, lastInsertedId)
	if err != nil {
		return nil, err
	}

	err = savePitchSlots(tx, pitchSlots, lastInsertedId)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.New("tx commit error")
	}
	return booking.BookingId, nil
}

func saveMatches(tx *sqlx.Tx, matches *[]models.Match, bId int) error {
	for _, m := range *matches {
		m.BookingId = bId
		_, err := tx.Exec(`INSERT INTO footpaldb.public.match(booking_id, match_access_status_id, match_status_id, 
                                   squad_id, match_date, cost, is_paid, created, last_updated)
							VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
			m.BookingId, m.MatchAccessStatusId, m.MatchStatusId, m.SquadId, m.MatchDate, m.Cost, m.IsPaid, m.Created, m.LastUpdated)
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return err
			}
			return errors.New("insert match error")
		}

	}
	return nil
}

func savePitchSlots(tx *sqlx.Tx, pitchSlots *[]models.PitchSlot, bId int) error {
	for _, ps := range *pitchSlots {
		ps.BookingId = bId
		_, err := tx.Exec(`INSERT INTO footpaldb.public.pitch_slot(booking_id, pitch_time_slot_id, match_date, 
                                   booking_status_id) VALUES($1, $2, $3, $4)`,
			ps.BookingId, ps.PitchTimeSlotId, ps.MatchDate, ps.BookingStatusId)
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return err
			}
			return errors.New("insert pitch slot error")
		}

	}
	return nil
}

func (r BookingRepository) Update(booking *models.Booking) (*models.Booking, error) {
	_, err := r.database.NamedExec(`UPDATE footpaldb.public.booking SET booking_status_id=:booking_status_id, 
                                    created_by=:created_by, created=:created, last_updated=:last_updated WHERE id=:id`, booking)

	if err != nil {
		return nil, err
	}
	return booking, nil
}
