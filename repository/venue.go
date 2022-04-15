package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/nmcalinden/footpal/models"
)

type VenueRepository struct {
	database *sqlx.DB
}

func NewVenueRepository(database *sqlx.DB) *VenueRepository {
	return &VenueRepository{database: database}
}

func (repository VenueRepository) FindAll() (*[]models.Venue, error) {
	var venues []models.Venue
	err := repository.database.Select(&venues, "SELECT * FROM footpaldb.public.venue")
	if err != nil || len(venues) == 0 {
		return nil, err
	}
	return &venues, nil
}

func (repository VenueRepository) FindById(id *int) (*models.Venue, error) {
	var venue models.Venue
	err := repository.database.Get(&venue, "SELECT * FROM footpaldb.public.venue WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &venue, nil
}

func (repository VenueRepository) FindAdminsByVenue(venueId *int) (*[]models.VenueAdmin, error) {
	var admins []models.VenueAdmin
	err := repository.database.Select(&admins, "SELECT * FROM footpaldb.public.venue_admin WHERE venue_id = $1", venueId)
	if err != nil || len(admins) == 0 {
		return nil, err
	}
	return &admins, nil
}

func (repository VenueRepository) FindAdminByVenue(venueId *int, adminId *int) (*models.VenueAdmin, error) {
	var admin models.VenueAdmin
	err := repository.database.Get(&admin, "SELECT * FROM footpaldb.public.venue_admin WHERE venue_id = $1 AND id = $2", venueId, adminId)
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (repository VenueRepository) FindPitchesByVenue(venueId *int) (*[]models.Pitch, error) {
	var pitches []models.Pitch
	err := repository.database.Select(&pitches, "SELECT * FROM footpaldb.public.pitch WHERE venue_id = $1", venueId)
	if err != nil || len(pitches) == 0 {
		return nil, err
	}
	return &pitches, nil
}

func (repository VenueRepository) FindPitchByVenueIdAndPitchId(venueId *int, pitchId *int) (*models.Pitch, error) {
	var pitch models.Pitch
	err := repository.database.Get(&pitch, "SELECT * FROM footpaldb.public.pitch WHERE id=$1 AND venue_id = $2", pitchId, venueId)
	if err != nil {
		return nil, err
	}
	return &pitch, nil
}

func (repository VenueRepository) Save(venue *models.Venue) (*int, error) {
	_, err := repository.database.NamedExec(`INSERT INTO footpaldb.public.venue(venue_name, venue_address, postcode, 
                                   city, phone_no, email) VALUES(:venue_name, :venue_address, :postcode, :city, :phone_no, :email)`, venue)

	if err != nil {
		return nil, err
	}
	return &venue.VenueId, nil
}

func (repository VenueRepository) SaveAdmin(admin *models.VenueAdmin) (*int, error) {
	_, err := repository.database.NamedExec(`INSERT INTO footpaldb.public.venue_admin(footpal_user_id, venue_id) 
							VALUES(:footpal_user_id, :venue_id)`, admin)
	if err != nil {
		return nil, err
	}
	return &admin.VenueAdminId, nil
}

func (repository VenueRepository) SavePitch(pitch *models.Pitch) (*int, error) {
	_, err := repository.database.NamedExec(`INSERT INTO footpaldb.public.pitch(venue_id, pitch_name, max_players, cost)
 								VALUES(:venue_id, :pitch_name,:max_players, :cost)`, pitch)
	if err != nil {
		return nil, err
	}
	return &pitch.PitchId, nil
}

func (repository VenueRepository) Update(venue *models.Venue) (*models.Venue, error) {
	_, err := repository.database.NamedExec(`UPDATE footpaldb.public.venue SET venue_name=:venue_name, 
                                    venue_address=:venue_address, postcode=:postcode, city=:city, phone_no=:phone_no, 
                                  email=:email WHERE id=:id`, venue)

	if err != nil {
		return nil, err
	}
	return venue, nil
}

func (repository VenueRepository) Delete(id *int) error {
	res, err := repository.database.Exec("DELETE FROM footpaldb.public.venue WHERE id=$1", id)
	return validateDeletion(res, err)
}

func (repository VenueRepository) DeleteAdmin(venueAdminId *int) error {
	res, err := repository.database.Exec("DELETE FROM footpaldb.public.venue_admin WHERE id=$1", venueAdminId)
	return validateDeletion(res, err)
}

func (repository VenueRepository) UpdatePitch(pitch *models.Pitch) (*models.Pitch, error) {
	_, err := repository.database.NamedExec(`UPDATE footpaldb.public.pitch SET pitch_name=:pitch_name, 
                                    max_players=:max_players, cost=:cost WHERE id=:id`, pitch)

	if err != nil {
		return nil, err
	}
	return pitch, nil
}

func (repository VenueRepository) DeletePitch(pitchId *int) error {
	res, err := repository.database.Exec("DELETE FROM footpaldb.public.pitch WHERE id=$1", pitchId)

	return validateDeletion(res, err)
}

func (repository VenueRepository) FindTimeslotsByVenueId(venueId *int) (*[]models.VenueTimeSlot, error) {
	var timeSlots []models.VenueTimeSlot
	return &timeSlots, nil
}

func (repository VenueRepository) FindPitchTimeslots(pitchId *int) (*[]models.PitchTimeSlot, error) {
	var timeSlots []models.PitchTimeSlot
	err := repository.database.Select(&timeSlots, "SELECT * FROM footpaldb.public.pitch WHERE id = $1", pitchId)
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
