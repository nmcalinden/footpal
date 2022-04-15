package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/models"
	"github.com/nmcalinden/footpal/service"
	"github.com/nmcalinden/footpal/utils"
	"strconv"
)

type SquadResponse struct {
	id *int
}

type SquadController struct {
	squadService *service.SquadService
}

func NewSquadController(squadService *service.SquadService) *SquadController {
	return &SquadController{squadService: squadService}
}

// RetrieveSquads @Summary      Retrieve Squads
// @Description  Retrieve all squads
// @Tags         squad
// @Produce      json
// @Success      200  {array}  models.Squad
// @Failure      500  {object}  utils.ErrorResponse
// @Router       /squads [get]
func (controller SquadController) RetrieveSquads(c *fiber.Ctx) error {
	s, err := controller.squadService.GetSquads()
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to get squads")
	}
	return c.Status(fiber.StatusOK).JSON(s)
}

// CreateSquadGroup @Summary      Create new squad
// @Description  Create new squad for building up a team of players
// @Tags         squad
// @Produce      json
// @Param 		 message body models.SquadRequest true "Request"
// @Success      201 {object} SquadResponse
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Router       /squads [post]
func (controller SquadController) CreateSquadGroup(c *fiber.Ctx) error {
	s := new(models.SquadRequest)
	if err := c.BodyParser(&s); err != nil {
		return err
	}

	if errors := utils.ValidateStruct(*s); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	res, err := controller.squadService.CreateNewSquad(s)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to get squads")
	}

	response := SquadResponse{id: res}
	return c.Status(fiber.StatusCreated).JSON(response)
}

// RetrieveSquadById @Summary      Retrieve squad by id
// @Description  Get squad by squadId
// @Tags         squad
// @Produce      json
// @Param        squadId   path  int  true  "Squad ID"
// @Success      200 	{object} models.Squad
// @Failure      400 {object} utils.ErrorResponse
// @Router       /squads/{squadId} [get]
func (controller SquadController) RetrieveSquadById(c *fiber.Ctx) error {
	squadId, err := strconv.Atoi(c.Params("squadId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Squad Id supplied is invalid")
	}

	s, err := controller.squadService.GetSquadById(&squadId)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Squad does not exist")
	}

	return c.Status(fiber.StatusOK).JSON(s)
}

// UpdateSquadInfo @Summary      Update Squad Info
// @Description  Update Squad Info
// @Tags         squad
// @Produce      json
// @Param 		 message body models.SquadRequest true "Request"
// @Param        squadId   path  int  true  "Squad ID"
// @Success      200 	{object} models.Squad
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Router       /squads/{squadId} [put]
func (controller SquadController) UpdateSquadInfo(c *fiber.Ctx) error {
	squadId, err := strconv.Atoi(c.Params("squadId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "SquadId supplied is invalid")
	}

	s := new(models.SquadRequest)
	if err := c.BodyParser(&s); err != nil {
		return err
	}

	res, err := controller.squadService.EditSquad(&squadId, s)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to update squad info")
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

// RemoveSquad @Summary      Remove squad
// @Description  Delete squad from footpal
// @Tags         squad
// @Produce      json
// @Param        squadId   path  int  true  "Squad ID"
// @Success      202 {string} string "accepted"
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Router       /squads/{squadId} [delete]
func (controller SquadController) RemoveSquad(c *fiber.Ctx) error {
	squadId, err := strconv.Atoi(c.Params("squadId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "SquadId supplied is invalid")
	}

	err = controller.squadService.DeleteSquad(&squadId)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to delete squad")
	}
	return c.SendStatus(fiber.StatusAccepted)
}

// RetrieveSquadPlayers @Summary      Retrieve players by squad
// @Description  List of players associated with squad
// @Tags         squad
// @Produce      json
// @Param        squadId   path  int  true  "Squad ID"
// @Success      200 {array} models.SquadPlayerDetails
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Router       /squads/{squadId}/players [get]
func (controller SquadController) RetrieveSquadPlayers(c *fiber.Ctx) error {
	squadId, err := strconv.Atoi(c.Params("squadId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "SquadId supplied is invalid")
	}

	p, err := controller.squadService.GetAllPlayersBySquad(&squadId)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to delete squad")
	}
	return c.Status(fiber.StatusOK).JSON(p)
}

// ApprovePlayerToSquad @Summary      Approve squad request
// @Description  Approve player to join squad
// @Tags         squad
// @Produce      json
// @Param        squadId   path  int  true  "Squad ID"
// @Param        playerId   path  int  true  "Player ID"
// @Success      200 {string} string "ok"
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Router       /squads/{squadId}/players/{playerId} [put]
func (controller SquadController) ApprovePlayerToSquad(c *fiber.Ctx) error {
	squadId, err := strconv.Atoi(c.Params("squadId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "SquadId supplied is invalid")
	}

	playerId, err := strconv.Atoi(c.Params("playerId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "PlayerId supplied is invalid")
	}

	err = controller.squadService.ApprovePlayer(&squadId, &playerId)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to approve player to squad")
	}
	return c.SendStatus(fiber.StatusOK)
}

// RemovePlayerFromSquad @Summary      Remove squad player
// @Description  Remove player from squad
// @Tags         squad
// @Produce      json
// @Param        squadId   path  int  true  "Squad ID"
// @Param        playerId   path  int  true  "Player ID"
// @Success      204
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Router       /squads/{squadId}/players/{playerId} [delete]
func (controller SquadController) RemovePlayerFromSquad(c *fiber.Ctx) error {
	squadId, err := strconv.Atoi(c.Params("squadId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "SquadId supplied is invalid")
	}

	playerId, err := strconv.Atoi(c.Params("playerId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "PlayerId supplied is invalid")
	}

	err = controller.squadService.RemovePlayer(&squadId, &playerId)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to approve player to squad")
	}
	return c.SendStatus(fiber.StatusNoContent)
}
