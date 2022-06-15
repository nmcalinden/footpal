package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/nmcalinden/footpal/models"
)

//go:generate mockgen -destination=./mocks/pitch_mock.go -package=mocks github.com/nmcalinden/footpal/repository PitchRepositoryI

type PitchRepositoryI interface {
	FindAllByVenueId(venueId int) (*[]models.PitchTimeSlot, error)
	FindAllByVenueIdAndDay(venueId int, dayOfWeek string) (*[]models.PitchTimeSlot, error)
}

type PitchRepository struct {
	database *sqlx.DB
}

func NewPitchRepository(database *sqlx.DB) *PitchRepository {
	return &PitchRepository{database: database}
}

func (r PitchRepository) FindAllByVenueId(v int) (*[]models.PitchTimeSlot, error) {
	var pts []models.PitchTimeSlot

	q := getFindPitchTimeSlotsByVenueIdQuery()
	err := r.database.Select(&pts, q, v)
	if err != nil || len(pts) == 0 {
		return nil, err
	}
	return &pts, nil
}

func (r PitchRepository) FindAllByVenueIdAndDay(v int, d string) (*[]models.PitchTimeSlot, error) {
	var pts []models.PitchTimeSlot

	q := getFindPitchTimeSlotsByVenueIdQuery()
	q = q + " AND pts.day_of_week = $2"
	err := r.database.Select(&pts, q, v, d)
	if err != nil || len(pts) == 0 {
		return nil, err
	}
	return &pts, nil
}

func getFindPitchTimeSlotsByVenueIdQuery() string {
	return "SELECT pts.id, pts.day_of_week, pts.start_time, pts.end_time FROM pitch_time_slot pts " +
		"LEFT JOIN pitch p on pts.pitch_id = p.id WHERE p.venue_id = $1"
}
