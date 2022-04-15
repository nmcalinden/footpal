package services

import (
	"github.com/jmoiron/sqlx"
	"github.com/nmcalinden/footpal/models"
	"github.com/nmcalinden/footpal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(database *sqlx.DB) *UserService {
	return &UserService{userRepo: repository.NewUserRepository(database)}
}

func (service *UserService) Login(login *models.Login) (*int, error) {
	res, err := service.userRepo.FindByEmail(&login.Email)
	if err != nil {
		return nil, err
	}
	return &res.UserId, nil
}

func (service *UserService) Register(register *models.Register) (*int, error) {
	newUser := models.User{
		UserId:   0,
		Forename: register.Forename,
		Surname:  register.Surname,
		Email:    register.Email,
	}
	return service.userRepo.Save(&newUser)
}

func (service *UserService) Deactivate(userId *int) error {
	return service.userRepo.Delete(userId)
}
