package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/nmcalinden/footpal/models"
)

type UserRepository struct {
	database *sqlx.DB
}

func NewUserRepository(database *sqlx.DB) *UserRepository {
	return &UserRepository{database: database}
}

func (repository UserRepository) FindByEmail(e *string) (*models.User, error) {
	var user models.User
	err := repository.database.Get(&user, "SELECT * FROM footpaldb.public.footpal_user WHERE email = $1", e)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository UserRepository) Save(user *models.User) (int, error) {
	_, err := repository.database.NamedExec(`INSERT INTO footpaldb.public.footpal_user(forename, surname, email) 
					VALUES(:forename, :surname, :email)`, user)
	return user.UserId, err
}

func (repository UserRepository) Delete(id *int) error {
	res, err := repository.database.Exec("DELETE FROM footpaldb.public.footpal_user WHERE id=$1", id)

	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err == nil && count == 1 {
		return nil
	}

	return err
}
