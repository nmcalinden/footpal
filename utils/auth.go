package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nmcalinden/footpal/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func GetClaims(token interface{}) jwt.MapClaims {
	t := token.(*jwt.Token)
	return t.Claims.(jwt.MapClaims)
}

func GetAccessToken(user *models.User, isAdmin bool, isPlayer bool, secretKey string) (string, error) {
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
	return token.SignedString([]byte(secretKey))
}

func GetRefreshToken(user *models.User, secretKey string) (string, error) {
	rtClaims := jwt.MapClaims{
		"sub": user.UserId,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	return refreshToken.SignedString([]byte(secretKey))
}

func ParseRefreshToken(refreshToken *string) (*jwt.Token, error) {
	return jwt.Parse(*refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("refreshSecret"), nil
	})
}

func HashPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(b), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
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
