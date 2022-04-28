package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/nmcalinden/footpal/api/models"
)

//go:generate mockgen -destination=./mocks/user_mock.go -package=mocks github.com/nmcalinden/footpal/api/repository UserRepositoryI

type UserRepositoryI interface {
	FindById(id *int) (*models.User, error)
	FindByEmail(e *string) (*models.User, error)
	Save(user *models.User) (*int, error)
	Delete(id *int) error
}

type UserRepository struct {
	database *sqlx.DB
}

func NewUserRepository(database *sqlx.DB) *UserRepository {
	return &UserRepository{database: database}
}

func (r UserRepository) FindById(id *int) (*models.User, error) {
	var user models.User
	err := r.database.Get(&user, "SELECT * FROM footpaldb.public.footpal_user WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r UserRepository) FindByEmail(e *string) (*models.User, error) {
	var user models.User
	err := r.database.Get(&user, "SELECT * FROM footpaldb.public.footpal_user WHERE email = $1", e)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r UserRepository) Save(user *models.User) (*int, error) {
	stmt, err := r.database.PrepareNamed("INSERT INTO footpaldb.public.footpal_user(forename, surname, email, password)" +
		" VALUES(:forename, :surname, :email, :password) RETURNING id")

	var id int
	err = stmt.Get(&id, user)
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (r UserRepository) Delete(id *int) error {
	res, err := r.database.Exec("DELETE FROM footpaldb.public.footpal_user WHERE id=$1", id)

	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err == nil && count == 1 {
		return nil
	}

	return err
}
