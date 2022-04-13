package playerRoute

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/models"
	"github.com/nmcalinden/footpal/routers/match"
	"github.com/nmcalinden/footpal/routers/squad"
	"github.com/nmcalinden/footpal/utils"
	"strconv"
)

func ConfigurePlayerHandlers(app *fiber.App) {
	group := app.Group("/players")

	group.Get("/", retrievePlayers)
	group.Put("/", updatePlayer)
	group.Get("/:playerId", retrievePlayerById)
	group.Get("/squads", getSquadsByUser)
	group.Get("/squads/:squadId", getSquadById)
	group.Get("/matches", getPlayerMatches)
	group.Get("/matches/:matchId", retrieveMatchById)
	group.Post("/matches/:matchId", joinMatch)
	group.Delete("/matches/:matchId", leaveMatch)
	group.Post("/matches/:matchId/pay", makePlayerPayment)
	group.Patch("/matches/:matchId/pay", updatePlayerPaymentType)
}

// @Summary      Retrieve Players
// @Description  Retrieve all players
// @Tags         player
// @Produce      json
// @Success      200  {array}  models.Player
// @Router       /players [get]
func retrievePlayers(c *fiber.Ctx) error {
	p := MockPlayers
	return c.Status(fiber.StatusOK).JSON(p)
}

// @Summary      Edit Player
// @Description  Edit player information
// @Tags         player
// @Produce      json
// @Param 		 message body models.PlayerRequest true "Request"
// @Success      200 {object} models.Player
// @Failure      400 {object} utils.ErrorResponse
// @Router       /players [put]
func updatePlayer(c *fiber.Ctx) error {
	p := new(models.Player)
	p.PlayerId = 9999

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	if errors := utils.ValidateStruct(*p); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return c.Status(fiber.StatusOK).JSON(p)
}

// @Summary      Retrieve Player
// @Description  Get Player information
// @Tags         player
// @Produce      json
// @Param        playerId   path  int  true  "Player ID"
// @Success      200 {object} models.Player
// @Failure      400 {object} utils.ErrorResponse
// @Router       /players/{playerId} [get]
func retrievePlayerById(c *fiber.Ctx) error {
	playerId, err := strconv.Atoi(c.Params("playerId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Player Id supplied is invalid")
	}

	p := MockPlayers
	result := models.Player{}
	for _, s := range p {
		if s.PlayerId == playerId {
			result = s
		}
	}

	n := 0
	if result.PlayerId == n {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Player does not exist")
	}

	result.PlayerId = 0
	return c.Status(fiber.StatusOK).JSON(result)
}

// @Summary      Retrieve Squads
// @Description  Retrieve all squads linked to player
// @Tags         player
// @Produce      json
// @Success      200  {array}  models.Squad
// @Failure      400 {object} utils.ErrorResponse
// @Router       /players/squads [get]
func getSquadsByUser(c *fiber.Ctx) error {
	p := squadRoute.MockSquads
	return c.Status(fiber.StatusOK).JSON(p)
}

// @Summary      Retrieve Squad by squadId
// @Description  Retrieve squad by squad id
// @Tags         player
// @Produce      json
// @Param        squadId   path  int  true  "Squad ID"
// @Success      200  {array}  models.Player
// @Failure      400 {object} utils.ErrorResponse
// @Router       /players/squads/{squadId} [get]
func getSquadById(c *fiber.Ctx) error {
	_, err := strconv.Atoi(c.Params("squadId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Squad Id supplied is invalid")
	}
	p := squadRoute.MockSquads[0]
	return c.Status(fiber.StatusOK).JSON(p)
}

// @Summary      Retrieve Matches by player
// @Description  Retrieve a players opted in matches
// @Tags         player
// @Produce      json
// @Success      200  {array}  models.Match
// @Router       /players/matches [get]
func getPlayerMatches(c *fiber.Ctx) error {
	p := matchRoute.MockMatches
	return c.Status(fiber.StatusOK).JSON(p)
}

// @Summary      Retrieve Match by match id
// @Description  Retrieve a match linked to a player
// @Tags         player
// @Produce      json
// @Param        matchId   path  int  true  "Match ID"
// @Success      200  {array}  models.Match
// @Failure      400 {object} utils.ErrorResponse
// @Router       /players/matches/{matchId} [get]
func retrieveMatchById(c *fiber.Ctx) error {
	matchId, err := strconv.Atoi(c.Params("matchId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "MatchId supplied is invalid")
	}

	p := matchRoute.MockMatches
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

// @Summary      Join match
// @Description  Request to join match
// @Tags         player
// @Produce      json
// @Param        matchId   path  int  true  "Match ID"
// @Success      202  {string} string "accepted"
// @Failure      400 {object} utils.ErrorResponse
// @Router       /players/matches/{matchId} [post]
func joinMatch(c *fiber.Ctx) error {
	_, err := strconv.Atoi(c.Params("matchId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Match Id supplied is invalid")
	}

	return c.SendStatus(fiber.StatusAccepted)
}

// @Summary      Leave match
// @Description  Leave match
// @Tags         player
// @Produce      json
// @Param        matchId   path  int  true  "Match ID"
// @Success      200  {string} string "ok"
// @Failure      400 {object} utils.ErrorResponse
// @Router       /players/matches/{matchId} [delete]
func leaveMatch(c *fiber.Ctx) error {
	_, err := strconv.Atoi(c.Params("matchId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Match Id supplied is invalid")
	}

	return c.SendStatus(fiber.StatusOK)
}

// @Summary      Match player payment
// @Description  For a match, pay amount owed by player
// @Tags         player
// @Produce      json
// @Param        matchId   path  int  true  "Match ID"
// @Success      200  {string} string "ok"
// @Router       /players/matches/{matchId}/pay [post]
func makePlayerPayment(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

// @Summary      Update player payment type
// @Description  For a match, pay amount owed by player
// @Tags         player
// @Produce      json
// @Param        matchId   path  int  true  "Match ID"
// @Param        paymentType   query  string  true  "Payment Type"
// @Success      200  {string} string "ok"
// @Router       /players/matches/{matchId}/pay [patch]
func updatePlayerPaymentType(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
