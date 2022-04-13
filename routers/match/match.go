package matchRoute

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/models"
	"github.com/nmcalinden/footpal/utils"
	"strconv"
)

func ConfigureMatchHandlers(app *fiber.App) {
	group := app.Group("/matches")

	group.Get("/", retrieveMatches)
	group.Get("/:matchId", retrieveMatchById)
	group.Delete("/:matchId", cancelMatch)
	group.Get("/:matchId/players", retrieveMatchPlayers)
	group.Delete("/:matchId/players/:playerId", removePlayerFromMatch)
}

// @Summary      Retrieve Matches
// @Description  Retrieve all public matches
// @Tags         match
// @Produce      json
// @Success      200  {array}  models.Match
// @Router       /matches [get]
func retrieveMatches(c *fiber.Ctx) error {
	p := MockMatches
	return c.Status(fiber.StatusOK).JSON(p)
}

// @Summary      Retrieve Match by matchId
// @Description  Retrieve match details by matchId
// @Tags         match
// @Produce      json
// @Param        matchId   path  int  true  "Match ID"
// @Success      200  {object}  models.Match
// @Failure      400 {object} utils.ErrorResponse
// @Router       /matches/{matchId} [get]
func retrieveMatchById(c *fiber.Ctx) error {
	matchId, err := strconv.Atoi(c.Params("matchId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "MatchId supplied is invalid")
	}

	p := MockMatches
	result := models.Match{}
	for _, s := range p {
		if s.MatchId == matchId {
			result = s
		}
	}

	n := 0
	if result.MatchId == n {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Match does not exist")
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

// @Summary      Cancel match
// @Description  Cancel match
// @Tags         match
// @Produce      json
// @Param        matchId   path  int  true  "Match ID"
// @Success      200  {string} string "ok"
// @Router       /matches/{matchId} [delete]
func cancelMatch(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

// @Summary      Retrieve Players by Match
// @Description  Retrieve all players opted into a match
// @Tags         match
// @Produce      json
// @Param        matchId   path  int  true  "Match ID"
// @Success      200  {object}  models.Match
// @Failure      400 {object} utils.ErrorResponse
// @Router       /matches/{matchId}/players [get]
func retrieveMatchPlayers(c *fiber.Ctx) error {
	matchId, err := strconv.Atoi(c.Params("matchId"))

	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "MatchId supplied is invalid")
	}

	p := MockMatchPlayers
	var result []models.MatchPlayer

	for _, s := range p {
		if s.MatchId == matchId {
			result = append(result, s)
		}
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

// @Summary      Remove player from match
// @Description  As a booking/squad admin, remove player from match
// @Tags         match
// @Produce      json
// @Param        matchId   path  int  true  "Match ID"
// @Param        playerId   path  int  true  "Player ID"
// @Success      200  {string} string "ok"
// @Failure      400 {object} utils.ErrorResponse
// @Router       /matches/{matchId}/players/{playerId} [delete]
func removePlayerFromMatch(c *fiber.Ctx) error {
	_, err := strconv.Atoi(c.Params("matchId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "MatchId supplied is invalid")
	}

	_, err = strconv.Atoi(c.Params("playerId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "PlayerId supplied is invalid")
	}

	return c.SendStatus(fiber.StatusOK)
}
