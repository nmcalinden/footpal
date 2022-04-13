package userRoute

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/models"
)

type UserResponse struct {
	Id int
}

func ConfigureUserHandlers(app *fiber.App) {
	group := app.Group("/")

	group.Post("/login", loginHandler)
	group.Post("/register", registerHandler)
	group.Delete("/deactivate", deactivateHandler)
}

// @Summary      Login
// @Description  Login to Footpal
// @Tags         user
// @Produce      json
// @Param 		 message body models.Login true "Request"
// @Success      200 {object} models.User
// @Failure      400 {object} utils.ErrorResponse
// @Router       /login [post]
func loginHandler(c *fiber.Ctx) error {
	l := new(models.Login)

	if err := c.BodyParser(&l); err != nil {
		return err
	}
	usrEx := models.User{
		UserId:   0,
		Forename: "Tester",
		Surname:  "Test",
		Email:    "tester@test.com",
	}
	return c.Status(fiber.StatusOK).JSON(usrEx)
}

// @Summary      Register
// @Description  Register as new player
// @Tags         user
// @Produce      json
// @Param 		 message body models.Register true "Request"
// @Success      200 {object} UserResponse
// @Failure      400 {object} utils.ErrorResponse
// @Router       /register [post]
func registerHandler(c *fiber.Ctx) error {
	r := new(models.Register)

	if err := c.BodyParser(&r); err != nil {
		return err
	}

	var usr = UserResponse{Id: 1}
	return c.Status(fiber.StatusOK).JSON(usr)
}

// @Summary      Deactivate User
// @Description  Delete user from Footpal
// @Tags         user
// @Produce      json
// @Success      200 {object} UserResponse
// @Router       /deactivate [delete]
func deactivateHandler(c *fiber.Ctx) error {
	var usr = UserResponse{Id: 1}
	return c.Status(fiber.StatusOK).JSON(usr)
}
