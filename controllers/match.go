package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/service"
	"github.com/nmcalinden/footpal/utils"
	"strconv"
)

type MatchRequest struct {
	MatchAccessStatusId int     `json:"matchAccessStatusId" validate:"required"`
	MatchStatusId       int     `json:"matchStatusId" validate:"required"`
	Cost                float32 `json:"cost" validate:"required"`
	IsPaid              bool    `json:"isPaid" validate:"required"`
}

type MatchController struct {
	matchService *service.MatchService
}

func NewMatchController(matchService *service.MatchService) *MatchController {
	return &MatchController{matchService: matchService}
}

// RetrieveMatches @Summary      Retrieve Matches
// @Description  Retrieve all public matches
// @Tags         match
// @Produce      json
// @Success      200  {array}  models.Match
// @Failure      500  {object}  utils.ErrorResponse
// @Router       /matches [get]
func (controller MatchController) RetrieveMatches(c *fiber.Ctx) error {
	m, err := controller.matchService.GetMatches()
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to retrieve matches")
	}
	return c.Status(fiber.StatusOK).JSON(m)
}

// RetrieveMatchById @Summary      Retrieve Match by matchId
// @Description  Retrieve match details by matchId
// @Tags         match
// @Produce      json
// @Param        matchId   path  int  true  "Match ID"
// @Success      200  {object}  models.Match
// @Failure      400 {object} utils.ErrorResponse
// @Router       /matches/{matchId} [get]
func (controller MatchController) RetrieveMatchById(c *fiber.Ctx) error {
	matchId, err := strconv.Atoi(c.Params("matchId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "MatchId supplied is invalid")
	}

	m, err := controller.matchService.GetMatchById(&matchId)

	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Match does not exist")
	}
	return c.Status(fiber.StatusOK).JSON(m)
}

// CancelMatch @Summary      Cancel match
// @Description  Cancel match
// @Tags         match
// @Produce      json
// @Param        matchId   path  int  true  "Match ID"
// @Success      200  {string} string "ok"
// @Failure      500  {object}  utils.ErrorResponse
// @Router       /matches/{matchId} [delete]
func (controller MatchController) CancelMatch(c *fiber.Ctx) error {
	matchId, err := strconv.Atoi(c.Params("matchId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "MatchId supplied is invalid")
	}

	_, err = controller.matchService.CancelMatch(&matchId)

	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to cancel match")
	}
	return c.SendStatus(fiber.StatusOK)
}

// RetrieveMatchPlayers @Summary      Retrieve Players by Match
// @Description  Retrieve all players opted into a match
// @Tags         match
// @Produce      json
// @Param        matchId   path  int  true  "Match ID"
// @Success      200  {array}  models.MatchPlayer
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Router       /matches/{matchId}/players [get]
func (controller MatchController) RetrieveMatchPlayers(c *fiber.Ctx) error {
	matchId, err := strconv.Atoi(c.Params("matchId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "MatchId supplied is invalid")
	}

	m, err := controller.matchService.GetMatchPlayers(&matchId)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to get match players")
	}

	return c.Status(fiber.StatusOK).JSON(m)
}

// RemovePlayerFromMatch @Summary      Remove player from match
// @Description  As a booking/squad admin, remove player from match
// @Tags         match
// @Produce      json
// @Param        matchId   path  int  true  "Match ID"
// @Param        playerId   path  int  true  "Player ID"
// @Success      200  {string} string "ok"
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Router       /matches/{matchId}/players/{playerId} [delete]
func (controller MatchController) RemovePlayerFromMatch(c *fiber.Ctx) error {
	matchId, err := strconv.Atoi(c.Params("matchId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "MatchId supplied is invalid")
	}

	playerId, err := strconv.Atoi(c.Params("playerId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "PlayerId supplied is invalid")
	}

	err = controller.matchService.RemovePlayer(&matchId, &playerId)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to remove player from match")
	}
	return c.SendStatus(fiber.StatusOK)
}
