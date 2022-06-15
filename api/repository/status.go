package repository

import (
	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -destination=./mocks/status_mock.go -package=mocks github.com/nmcalinden/footpal/repository StatusRepositoryI

type StatusRepositoryI interface {
	FindBookingStatusById(int *int) (*string, error)
	FindMatchStatusById(id *int) (*string, error)
}

type StatusRepository struct {
	database *sqlx.DB
}

func NewStatusRepository(database *sqlx.DB) *StatusRepository {
	return &StatusRepository{database: database}
}

func (r StatusRepository) FindBookingStatusById(id *int) (*string, error) {
	q := "SELECT description FROM footpaldb.public.booking_status_ref WHERE id=$1"
	return r.findStatusById(id, q)
}

func (r StatusRepository) FindMatchStatusById(id *int) (*string, error) {
	q := "SELECT description FROM footpaldb.public.match_status_ref WHERE id=$1"
	return r.findStatusById(id, q)
}

func (r StatusRepository) findStatusById(id *int, q string) (*string, error) {
	var d string
	err := r.database.Get(&d, q, id)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
