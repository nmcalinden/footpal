package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/services"
	"github.com/nmcalinden/footpal/utils"
	"strconv"
)

type MatchController struct {
	matchService *services.MatchService
}

func NewMatchController(matchService *services.MatchService) *MatchController {
	return &MatchController{matchService: matchService}
}

// RetrieveMatches @Summary      Retrieve Matches
// @Description  Retrieve all public matches
// @Tags         match
// @Produce      json
// @Success      200  {array}  models.Match
// @Failure      500  {object}  utils.ErrorResponse
// @Security ApiKeyAuth
// @Router       /matches [get]
func (con MatchController) RetrieveMatches(c *fiber.Ctx) error {
	m, err := con.matchService.GetMatches()
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
// @Security ApiKeyAuth
// @Router       /matches/{matchId} [get]
func (con MatchController) RetrieveMatchById(c *fiber.Ctx) error {
	matchId, err := strconv.Atoi(c.Params("matchId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "MatchId supplied is invalid")
	}

	m, err := con.matchService.GetMatchById(&matchId)

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
// @Security ApiKeyAuth
// @Router       /matches/{matchId} [delete]
func (con MatchController) CancelMatch(c *fiber.Ctx) error {
	matchId, err := strconv.Atoi(c.Params("matchId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "MatchId supplied is invalid")
	}

	_, err = con.matchService.CancelMatch(&matchId)

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
// @Security ApiKeyAuth
// @Router       /matches/{matchId}/players [get]
func (con MatchController) RetrieveMatchPlayers(c *fiber.Ctx) error {
	matchId, err := strconv.Atoi(c.Params("matchId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "MatchId supplied is invalid")
	}

	m, err := con.matchService.GetMatchPlayers(&matchId)
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
// @Security ApiKeyAuth
// @Router       /matches/{matchId}/players/{playerId} [delete]
func (con MatchController) RemovePlayerFromMatch(c *fiber.Ctx) error {
	matchId, err := strconv.Atoi(c.Params("matchId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "MatchId supplied is invalid")
	}

	playerId, err := strconv.Atoi(c.Params("playerId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "PlayerId supplied is invalid")
	}

	err = con.matchService.RemovePlayer(&matchId, &playerId)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to remove player from match")
	}
	return c.SendStatus(fiber.StatusOK)
}
