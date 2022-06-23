package services

import (
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/golang-jwt/jwt/v4"
	"github.com/nmcalinden/footpal/config"
	"github.com/nmcalinden/footpal/mappers"
	"github.com/nmcalinden/footpal/models"
	"github.com/nmcalinden/footpal/payloads"
	"github.com/nmcalinden/footpal/repository"
	"github.com/nmcalinden/footpal/utils"
	"github.com/nmcalinden/footpal/views"
)

type UserService struct {
	userRepo   repository.UserRepositoryI
	playerRepo repository.PlayerRepositoryI
	venueRepo  repository.VenueRepositoryI
}

func NewUserService(usrRepo repository.UserRepositoryI, pRepo repository.PlayerRepositoryI,
	vRepo repository.VenueRepositoryI) *UserService {
	return &UserService{
		userRepo:   usrRepo,
		playerRepo: pRepo,
		venueRepo:  vRepo,
	}
}

func (s *UserService) GetUser(id int) (*views.UserProfile, error) {
	res, err := s.userRepo.FindById(&id)
	if err != nil {
		return nil, err
	}

	player, _ := s.playerRepo.FindByUserId(&id)

	var user views.UserProfile
	err = mappers.MapToUserProfile(&user, *player, *res)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &user, nil
}

func (s *UserService) GetProfile(id int) (*views.PlayerUser, error) {
	res, err := s.userRepo.FindById(&id)
	if err != nil {
		return nil, err
	}

	player, _ := s.playerRepo.FindByUserId(&id)

	var user views.PlayerUser
	err = mappers.MapToPlayer(&user, *player, *res)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &user, nil
}

func (s *UserService) Login(login *payloads.Login) (*payloads.LoginResponse, error) {
	res, err := s.userRepo.FindByEmail(&login.Email)
	if err != nil {
		return nil, err
	}

	usrPw := res.Password
	pwErr := utils.CheckPasswordHash(login.Password, usrPw)
	if !pwErr {
		invalidPw := errors.New("invalidPassword")
		return nil, invalidPw
	}

	player, _ := s.playerRepo.FindByUserId(res.UserId)
	tp, err := s.getTokenPair(res)
	if err != nil {
		return nil, err
	}
	var user views.UserProfile
	err = mappers.MapToUserProfile(&user, *player, *res)
	if err != nil {
		return nil, err
	}
	return &payloads.LoginResponse{
		User: user,
		JWT:  *tp,
	}, nil
}

func (s *UserService) Register(register *payloads.Register) (*int, error) {
	pw, err := utils.HashPassword(register.Password)
	if err != nil {
		return nil, err
	}
	newUser := models.User{
		Forename: register.Forename,
		Surname:  register.Surname,
		Email:    register.Email,
		Password: pw,
	}

	id, err := s.userRepo.Save(&newUser)
	if err != nil {
		return nil, err
	}

	if register.IsPlayer {
		_, err := s.playerRepo.Save(&models.Player{UserId: *id})
		if err != nil {
			return nil, err
		}
	}

	return id, nil

}

func (s *UserService) Refresh(refreshToken *string) (*payloads.LoginResponse, error) {
	t, err := utils.ParseRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		usr := int(claims["sub"].(float64))
		res, err := s.userRepo.FindById(&usr)
		if err != nil {
			return nil, err
		}

		player, _ := s.playerRepo.FindByUserId(res.UserId)
		tp, err := s.getTokenPair(res)
		if err != nil {
			return nil, err
		}
		var user views.UserProfile
		err = mappers.MapToUserProfile(&user, *player, *res)
		if err != nil {
			return nil, err
		}

		return &payloads.LoginResponse{
			User: user,
			JWT:  *tp,
		}, nil
	}

	return nil, err
}

func (s *UserService) Deactivate(userId *int) error {
	res, err := s.userRepo.FindById(userId)
	if err != nil {
		return err
	}
	return s.userRepo.Delete(res.UserId)
}

func (s *UserService) getTokenPair(res *models.User) (*payloads.TokenPairResponse, error) {
	var wg sync.WaitGroup
	wg.Add(2)

	adminChan, playerChan := make(chan bool, 1), make(chan bool, 1)
	go s.doesAdminExist(*res.UserId, &wg, adminChan)
	go s.doesPlayerExist(*res.UserId, &wg, playerChan)

	wg.Wait()

	isAdmin, isPlayer := <-adminChan, <-playerChan

	at, err := utils.GetAccessToken(res, isAdmin, isPlayer, config.AccessSecret)
	if err != nil {
		return nil, err
	}

	rt, err := utils.GetRefreshToken(res, config.RefreshSecret)
	if err != nil {
		return nil, err
	}

	bearerToken := fmt.Sprintf("Bearer %s", at)
	response := payloads.TokenPairResponse{AccessToken: &bearerToken, RefreshToken: &rt}
	return &response, nil
}

func (s *UserService) doesAdminExist(userId int, wg *sync.WaitGroup, ch chan<- bool) {
	defer wg.Done()
	defer close(ch)
	_, err := s.venueRepo.FindAdminByUserId(&userId)
	res := err == nil
	ch <- res
}

func (s *UserService) doesPlayerExist(userId int, wg *sync.WaitGroup, ch chan<- bool) {
	defer wg.Done()
	defer close(ch)
	_, err := s.playerRepo.FindByUserId(&userId)
	res := err == nil
	ch <- res
}
