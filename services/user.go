package services

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jmoiron/sqlx"
	"github.com/nmcalinden/footpal/models"
	"github.com/nmcalinden/footpal/repository"
	"github.com/nmcalinden/footpal/utils"
	"sync"
	"time"
)

const (
	access  = "accessSecret"
	refresh = "refreshSecret"
)

type UserService struct {
	userRepo   *repository.UserRepository
	playerRepo *repository.PlayerRepository
	venueRepo  *repository.VenueRepository
}

func NewUserService(database *sqlx.DB) *UserService {
	return &UserService{
		userRepo:   repository.NewUserRepository(database),
		playerRepo: repository.NewPlayerRepository(database),
		venueRepo:  repository.NewVenueRepository(database),
	}
}

func (s *UserService) Login(login *models.Login) (*models.TokenPairResponse, error) {
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

	response, err2 := s.getTokenPair(res)
	if err2 != nil {
		return nil, err2
	}
	return response, nil
}

func (s *UserService) Register(register *models.Register) (*int, error) {
	pw, err := utils.HashPassword(register.Password)
	if err != nil {
		return nil, err
	}
	newUser := models.User{
		UserId:   0,
		Forename: register.Forename,
		Surname:  register.Surname,
		Email:    register.Email,
		Password: pw,
	}
	return s.userRepo.Save(&newUser)
}

func (s *UserService) Refresh(refreshToken *string) (*models.TokenPairResponse, error) {
	t, err := parseRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		usr := int(claims["sub"].(float64))
		res, err := s.userRepo.FindById(&usr)
		if err != nil {
			return nil, err
		}

		return s.getTokenPair(res)
	}

	return nil, err
}

func (s *UserService) Deactivate(userId *int) error {
	res, err := s.userRepo.FindById(userId)
	if err != nil {
		return err
	}
	return s.userRepo.Delete(&res.UserId)
}

func (s *UserService) getTokenPair(res *models.User) (*models.TokenPairResponse, error) {
	var wg sync.WaitGroup
	wg.Add(2)

	var isAdmin bool
	go func() {
		defer wg.Done()
		_, err := s.venueRepo.FindAdminByUserId(&res.UserId)
		isAdmin = err == nil
	}()

	var isPlayer bool
	go func() {
		defer wg.Done()
		_, err := s.playerRepo.FindByUserId(&res.UserId)
		isPlayer = err == nil
	}()

	wg.Wait()

	access, err := getAccessToken(res, isAdmin, isPlayer)
	if err != nil {
		return nil, err
	}

	refresh, err := getRefreshToken(res)
	if err != nil {
		return nil, err
	}

	bearerToken := fmt.Sprintf("Bearer %s", access)
	response := models.TokenPairResponse{AccessToken: &bearerToken, RefreshToken: &refresh}
	return &response, nil
}

func getAccessToken(user *models.User, isAdmin bool, isPlayer bool) (string, error) {
	roles := buildRoles(isAdmin, isPlayer)

	claims := jwt.MapClaims{
		"sub":   user.UserId,
		"name":  fmt.Sprintf("%s %s", user.Forename, user.Surname),
		"email": user.Email,
		"roles": roles,
		"exp":   time.Now().Add(time.Minute * 60).Unix(),
		"iat":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(access))
}

func getRefreshToken(user *models.User) (string, error) {
	rtClaims := jwt.MapClaims{
		"sub": user.UserId,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	return refreshToken.SignedString([]byte(refresh))
}

func buildRoles(isAdmin bool, isPlayer bool) []string {
	var roles []string
	if isAdmin {
		roles = append(roles, "venueAdmin")
	}
	if isPlayer {
		roles = append(roles, "player")
	}
	return roles
}

func parseRefreshToken(refreshToken *string) (*jwt.Token, error) {
	return jwt.Parse(*refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("refreshSecret"), nil
	})
}
