package services

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/nmcalinden/footpal/middleware"
	"github.com/nmcalinden/footpal/models"
	"github.com/nmcalinden/footpal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(database *sqlx.DB) *UserService {
	return &UserService{userRepo: repository.NewUserRepository(database)}
}

func (service *UserService) Login(login *models.Login) (*string, error) {
	res, err := service.userRepo.FindByEmail(&login.Email)
	if err != nil {
		return nil, err
	}

	name := fmt.Sprintf("%s %s", res.Forename, res.Surname)
	token, err := middleware.CreateToken(&name, &res.Email, false, true)

	if err != nil {
		return nil, err
	}

	bearerToken := fmt.Sprintf("Bearer %s", token)
	return &bearerToken, nil
}

func (service *UserService) Register(register *models.Register) (int, error) {
	newUser := models.User{
		UserId:   0,
		Forename: register.Forename,
		Surname:  register.Surname,
		Email:    register.Email,
	}
	return service.userRepo.Save(&newUser)
}

func (service *UserService) Deactivate(email *string) error {
	res, err := service.userRepo.FindByEmail(email)
	if err != nil {
		return err
	}
	return service.userRepo.Delete(&res.UserId)
}
