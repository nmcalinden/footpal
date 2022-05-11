package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/nmcalinden/footpal/api/enums"
	"github.com/nmcalinden/footpal/api/utils"
	"github.com/nmcalinden/footpal/config"
	"golang.org/x/exp/slices"
)

var IsAuthenticated = jwtware.New(jwtware.Config{
	SigningKey:   []byte(config.AccessSecret),
	ErrorHandler: jwtError,
})

func jwtError(ctx *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT"})
	}
	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})

}

type Roles struct {
	Roles []enums.Role
}

func NewRoles(roles ...enums.Role) Roles {
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

func isValid(roles []interface{}, expectedRoles []enums.Role) bool {
	if len(roles) == 0 {
		return false
	}

	var isValid bool
	for _, r := range roles {
		i := slices.IndexFunc(expectedRoles, func(rol enums.Role) bool { return rol.String() == r })
		if i != -1 {
			isValid = true
			break
		}
	}

	return isValid
}
