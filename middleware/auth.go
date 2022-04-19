package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func IsAuthenticated(c *fiber.Ctx) error {
	err := jwtware.New(jwtware.Config{SigningKey: []byte("secret")})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Valid JWT required")
	}
	return nil
}

func CreateToken(name *string, email *string, isAdmin bool, isPlayer bool) (string, error) {
	roles := buildRoles(isAdmin, isPlayer)

	claims := jwt.MapClaims{
		"name":  name,
		"email": email,
		"roles": roles,
		"exp":   time.Now().Add(time.Minute * 60).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
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
