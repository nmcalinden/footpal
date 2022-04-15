package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/models"
	"github.com/nmcalinden/footpal/services"
	"github.com/nmcalinden/footpal/utils"
	"strconv"
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
// @Param 		 message body models.Login true "Request"
// @Success      200 {object} models.UserResponse
// @Failure      400 {object} utils.ErrorResponse
// @Router       /login [post]
func (controller UserController) LoginHandler(c *fiber.Ctx) error {
	l := new(models.Login)
	if err := c.BodyParser(&l); err != nil {
		return err
	}

	if errors := utils.ValidateStruct(*l); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	usr, err := controller.userService.Login(l)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "No User found")
	}

	response := models.UserResponse{Id: usr}
	return c.Status(fiber.StatusOK).JSON(response)
}

// RegisterHandler @Summary      Register
// @Description  Register as new player
// @Tags         user
// @Produce      json
// @Param 		 message body models.Register true "Request"
// @Success      200 {object} models.UserResponse
// @Failure      400 {object} utils.ErrorResponse
// @Router       /register [post]
func (controller UserController) RegisterHandler(c *fiber.Ctx) error {
	r := new(models.Register)
	if err := c.BodyParser(&r); err != nil {
		return err
	}

	if errors := utils.ValidateStruct(*r); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	usr, err := controller.userService.Register(r)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to register user")
	}
	response := models.UserResponse{Id: usr}
	return c.Status(fiber.StatusOK).JSON(response)
}

// DeactivateHandler @Summary      Deactivate User
// @Description  Delete user from Footpal
// @Tags         user
// @Produce      json
// @Param        userId   path  int  true  "User ID"
// @Success      200
// @Failure      400 {object} utils.ErrorResponse
// @Router       /deactivate/{userId} [delete]
func (controller UserController) DeactivateHandler(c *fiber.Ctx) error {
	userId, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "UserId supplied is invalid")
	}

	err = controller.userService.Deactivate(&userId)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to delete user")
	}
	return c.SendStatus(fiber.StatusOK)
}
