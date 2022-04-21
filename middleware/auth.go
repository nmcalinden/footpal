package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/nmcalinden/footpal/utils"
	"golang.org/x/exp/slices"
)

var IsAuthenticated = jwtware.New(jwtware.Config{SigningKey: []byte("accessSecret")})

type Roles struct {
	Roles []UserRole
}

type UserRole struct {
	R string
}

func NewRoles(roles []UserRole) Roles {
	return Roles{Roles: roles}
}

func (r Roles) HasRole(c *fiber.Ctx) error {
	claims := utils.GetClaims(c.Locals("user"))
	roles := claims["roles"].([]interface{})
	if ok := isValid(roles, r.Roles); !ok {
		return utils.BuildErrorResponse(c, fiber.StatusForbidden, "User does not have correct permissions")
	}
	return c.Next()
}

func isValid(roles []interface{}, expectedRoles []UserRole) bool {
	if len(roles) == 0 {
		return false
	}

	var isValid bool
	for _, r := range roles {
		i := slices.IndexFunc(expectedRoles, func(uR UserRole) bool { return uR.R == r })
		if i != -1 {
			isValid = true
			break
		}
	}

	return isValid
}
