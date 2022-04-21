package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/payloads"
	"github.com/nmcalinden/footpal/services"
	"github.com/nmcalinden/footpal/utils"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

// LoginHandler @Summary      Login
// @Description  Login to Footpal
// @Tags         user
// @Produce      json
// @Param 		 message body payloads.Login true "Request"
// @Success      200 {object} payloads.TokenPairResponse
// @Failure      400 {object} utils.ErrorResponse
// @Router       /login [post]
func (con UserController) LoginHandler(c *fiber.Ctx) error {
	l := new(payloads.Login)
	if err := c.BodyParser(&l); err != nil {
		return err
	}

	if errors := utils.ValidateStruct(*l); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	token, err := con.userService.Login(l)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusUnauthorized, "Failed to login - Check username and password")
	}
	return c.Status(fiber.StatusOK).JSON(token)
}

// RegisterHandler @Summary      Register
// @Description  Register as new player
// @Tags         user
// @Produce      json
// @Param 		 message body payloads.Register true "Request"
// @Success      200 {object} payloads.RegisterResponse
// @Failure      400 {object} utils.ErrorResponse
// @Router       /register [post]
func (con UserController) RegisterHandler(c *fiber.Ctx) error {
	r := new(payloads.Register)
	if err := c.BodyParser(&r); err != nil {
		return err
	}

	if errors := utils.ValidateStruct(*r); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	usr, err := con.userService.Register(r)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to register user")
	}
	response := payloads.RegisterResponse{Id: usr}
	return c.Status(fiber.StatusOK).JSON(response)
}

// RefreshToken @Summary      Refresh Token
// @Description  Refresh token
// @Tags         user
// @Produce      json
// @Param 		 message body payloads.Refresh true "Request"
// @Success      200 {object} payloads.TokenPairResponse
// @Failure      400 {object} utils.ErrorResponse
// @Router       /refresh [post]
func (con UserController) RefreshToken(c *fiber.Ctx) error {
	r := new(payloads.Refresh)
	if err := c.BodyParser(&r); err != nil {
		return err
	}

	token, err := con.userService.Refresh(r.RefreshToken)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to refresh token")
	}
	return c.Status(fiber.StatusOK).JSON(token)
}

// DeactivateHandler @Summary      Deactivate User
// @Description  Delete user from Footpal
// @Tags         user
// @Security ApiKeyAuth
// @Produce      json
// @Success      200
// @Failure      400 {object} utils.ErrorResponse
// @Router       /deactivate [delete]
func (con UserController) DeactivateHandler(c *fiber.Ctx) error {
	claims := utils.GetClaims(c.Locals("user"))
	userId := claims["id"].(int)

	err := con.userService.Deactivate(&userId)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to delete user")
	}
	return c.SendStatus(fiber.StatusOK)
}
