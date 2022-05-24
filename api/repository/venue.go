package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/nmcalinden/footpal/api/models"
)

//go:generate mockgen -destination=./mocks/venue_mock.go -package=mocks github.com/nmcalinden/footpal/api/repository VenueRepositoryI

type VenueRepositoryI interface {
	FindAll() (*[]models.Venue, error)
	FindById(id *int) (*models.Venue, error)
	FindAdminsByVenue(venueId *int) (*[]models.VenueAdmin, error)
	FindAdminByVenue(venueId *int, adminId *int) (*models.VenueAdmin, error)
	FindAdminByUserId(userId *int) (*models.VenueAdmin, error)
	FindPitchesByVenue(venueId *int) (*[]models.Pitch, error)
	FindPitchByVenueIdAndPitchId(venueId *int, pitchId *int) (*models.Pitch, error)
	Save(venue *models.Venue) (*int, error)
	SaveAdmin(admin *models.VenueAdmin) (*int, error)
	SavePitch(pitch *models.Pitch) (*int, error)
	Update(venue *models.Venue) (*models.Venue, error)
	Delete(id *int) error
	DeleteAdmin(venueAdminId *int) error
	UpdatePitch(pitch *models.Pitch) (*models.Pitch, error)
	DeletePitch(pitchId *int) error
	FindTimeslotsByVenueIdAndDateRange(venueId int, from string, to string) (*[]models.VenueTimeSlot, error)
	FindPitchTimeslots(pitchId *int) (*[]models.PitchTimeSlot, error)
}

type VenueRepository struct {
	database *sqlx.DB
}

func NewVenueRepository(database *sqlx.DB) *VenueRepository {
	return &VenueRepository{database: database}
}

func (r VenueRepository) FindAll() (*[]models.Venue, error) {
	var venues []models.Venue
	err := r.database.Select(&venues, "SELECT * FROM footpaldb.public.venue")
	if err != nil || len(venues) == 0 {
		return nil, err
	}
	return &venues, nil
}

func (r VenueRepository) FindById(id *int) (*models.Venue, error) {
	var venue models.Venue
	err := r.database.Get(&venue, "SELECT * FROM footpaldb.public.venue WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &venue, nil
}

func (r VenueRepository) FindAdminsByVenue(venueId *int) (*[]models.VenueAdmin, error) {
	var admins []models.VenueAdmin
	err := r.database.Select(&admins, "SELECT * FROM footpaldb.public.venue_admin WHERE venue_id = $1", venueId)
	if err != nil || len(admins) == 0 {
		return nil, err
	}
	return &admins, nil
}

func (r VenueRepository) FindAdminByVenue(venueId *int, adminId *int) (*models.VenueAdmin, error) {
	var admin models.VenueAdmin
	err := r.database.Get(&admin, "SELECT * FROM footpaldb.public.venue_admin WHERE venue_id = $1 AND id = $2", venueId, adminId)
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r VenueRepository) FindAdminByUserId(userId *int) (*models.VenueAdmin, error) {
	var admin models.VenueAdmin
	err := r.database.Get(&admin, "SELECT * FROM footpaldb.public.venue_admin WHERE footpal_user_id = $1", userId)
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r VenueRepository) FindPitchesByVenue(venueId *int) (*[]models.Pitch, error) {
	var pitches []models.Pitch
	err := r.database.Select(&pitches, "SELECT * FROM footpaldb.public.pitch WHERE venue_id = $1", venueId)
	if err != nil || len(pitches) == 0 {
		return nil, err
	}
	return &pitches, nil
}

func (r VenueRepository) FindPitchByVenueIdAndPitchId(venueId *int, pitchId *int) (*models.Pitch, error) {
	var pitch models.Pitch
	err := r.database.Get(&pitch, "SELECT * FROM footpaldb.public.pitch WHERE id=$1 AND venue_id = $2", pitchId, venueId)
	if err != nil {
		return nil, err
	}
	return &pitch, nil
}

func (r VenueRepository) Save(venue *models.Venue) (*int, error) {
	_, err := r.database.NamedExec(`INSERT INTO footpaldb.public.venue(venue_name, venue_address, postcode, 
                                   city, phone_no, email) VALUES(:venue_name, :venue_address, :postcode, :city, :phone_no, :email)`, venue)

	if err != nil {
		return nil, err
	}
	return venue.VenueId, nil
}

func (r VenueRepository) SaveAdmin(admin *models.VenueAdmin) (*int, error) {
	_, err := r.database.NamedExec(`INSERT INTO footpaldb.public.venue_admin(footpal_user_id, venue_id) 
							VALUES(:footpal_user_id, :venue_id)`, admin)
	if err != nil {
		return nil, err
	}
	return admin.VenueAdminId, nil
}

func (r VenueRepository) SavePitch(pitch *models.Pitch) (*int, error) {
	_, err := r.database.NamedExec(`INSERT INTO footpaldb.public.pitch(venue_id, pitch_name, max_players, cost)
 								VALUES(:venue_id, :pitch_name,:max_players, :cost)`, pitch)
	if err != nil {
		return nil, err
	}
	return pitch.PitchId, nil
}

func (r VenueRepository) Update(venue *models.Venue) (*models.Venue, error) {
	_, err := r.database.NamedExec(`UPDATE footpaldb.public.venue SET venue_name=:venue_name, 
                                    venue_address=:venue_address, postcode=:postcode, city=:city, phone_no=:phone_no, 
                                  email=:email WHERE id=:id`, venue)

	if err != nil {
		return nil, err
	}
	return venue, nil
}

func (r VenueRepository) Delete(id *int) error {
	res, err := r.database.Exec("DELETE FROM footpaldb.public.venue WHERE id=$1", id)
	return validateDeletion(res, err)
}

func (r VenueRepository) DeleteAdmin(venueAdminId *int) error {
	res, err := r.database.Exec("DELETE FROM footpaldb.public.venue_admin WHERE id=$1", venueAdminId)
	return validateDeletion(res, err)
}

func (r VenueRepository) UpdatePitch(pitch *models.Pitch) (*models.Pitch, error) {
	_, err := r.database.NamedExec(`UPDATE footpaldb.public.pitch SET pitch_name=:pitch_name, 
                                    max_players=:max_players, cost=:cost WHERE id=:id`, pitch)

	if err != nil {
		return nil, err
	}
	return pitch, nil
}

func (r VenueRepository) DeletePitch(pitchId *int) error {
	res, err := r.database.Exec("DELETE FROM footpaldb.public.pitch WHERE id=$1", pitchId)

	return validateDeletion(res, err)
}

func (r VenueRepository) FindTimeslotsByVenueIdAndDateRange(v int, f string, t string) (*[]models.VenueTimeSlot, error) {
	q := "SELECT DISTINCT ps.pitch_time_slot_id, ps.match_date, pts.day_of_week from pitch_slot as ps " +
		"JOIN pitch_time_slot pts on ps.pitch_time_slot_id = pts.id " +
		"JOIN pitch p on pts.pitch_id = p.id " +
		"WHERE p.venue_id = $1 AND ps.booking_status_id= 1 AND ps.match_date between $2 AND $3"

	var bs []models.VenueTimeSlot

	err := r.database.Select(&bs, q, v, f, t)
	if err != nil {
		return nil, err
	}
	return &bs, nil
}
func (r VenueRepository) FindTimeslotsByVenueId(venueId *int) (*[]models.VenueTimeSlot, error) {
	var timeSlots []models.VenueTimeSlot
	return &timeSlots, nil
}

func (r VenueRepository) FindPitchTimeslots(pitchId *int) (*[]models.PitchTimeSlot, error) {
	var timeSlots []models.PitchTimeSlot
	err := r.database.Select(&timeSlots, "SELECT * FROM footpaldb.public.pitch WHERE id = $1", pitchId)
	if err != nil || len(timeSlots) == 0 {
		return nil, err
	}
	return &timeSlots, nil
}

func validateDeletion(res sql.Result, err error) error {
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err == nil && count == 1 {
		return nil
	}

	return err
}
