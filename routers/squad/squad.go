package squadRoute

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/models"
	"github.com/nmcalinden/footpal/utils"
	"strconv"
)

func ConfigureSquadPlayers(app *fiber.App) {
	group := app.Group("/squads")

	group.Get("/", retrieveSquads)
	group.Post("/", createSquadGroup)
	group.Get("/:squadId", retrieveSquadById)
	group.Put("/:squadId", updateSquadInfo)
	group.Delete("/:squadId", removeSquad)
	group.Get("/:squadId/players", retrieveSquadPlayers)
	group.Put("/:squadId/players/:playerId", approvePlayerToSquad)
	group.Delete("/:squadId/players/:playerId", removePlayerFromSquad)
}

// @Summary      Retrieve Squads
// @Description  Retrieve all squads
// @Tags         squad
// @Produce      json
// @Success      200  {array}  models.Squad
// @Router       /squads [get]
func retrieveSquads(c *fiber.Ctx) error {
	p := MockSquads
	return c.Status(fiber.StatusOK).JSON(p)
}

// @Summary      Create new squad
// @Description  Create new squad for building up a team of players
// @Tags         squad
// @Produce      json
// @Param 		 message body models.SquadRequest true "Request"
// @Success      201 	{string} string "ok"
// @Failure      400 {object} utils.ErrorResponse
// @Router       /squads [post]
func createSquadGroup(c *fiber.Ctx) error {
	p := new(models.SquadRequest)
	if err := c.BodyParser(&p); err != nil {
		return err
	}

	if errors := utils.ValidateStruct(*p); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	squadId := 999
	return c.Status(fiber.StatusCreated).SendString(string(squadId))
}

// @Summary      Retrieve squad by id
// @Description  Get squad by squadId
// @Tags         squad
// @Produce      json
// @Param        squadId   path  int  true  "Squad ID"
// @Success      200 	{object} models.Squad
// @Failure      400 {object} utils.ErrorResponse
// @Router       /squads/{squadId} [get]
func retrieveSquadById(c *fiber.Ctx) error {
	squadId, err := strconv.Atoi(c.Params("squadId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Squad Id supplied is invalid")
	}

	p := MockSquads
	result := models.Squad{}
	for _, s := range p {
		if s.SquadId == squadId {
			result = s
		}
	}

	if result.SquadId == 0 {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Squad does not exist")
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

// @Summary      Retrieve squad by id
// @Description  Get squad by squadId
// @Tags         squad
// @Produce      json
// @Param 		 message body models.SquadRequest true "Request"
// @Param        squadId   path  int  true  "Squad ID"
// @Success      200 	{object} models.Squad
// @Failure      400 {object} utils.ErrorResponse
// @Router       /squads/{squadId} [put]
func updateSquadInfo(c *fiber.Ctx) error {
	_, err := strconv.Atoi(c.Params("squadId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "SquadId supplied is invalid")
	}

	s := new(models.SquadRequest)

	if err := c.BodyParser(&s); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(s)
}

// @Summary      Remove squad
// @Description  Delete squad from footpal
// @Tags         squad
// @Produce      json
// @Param        squadId   path  int  true  "Squad ID"
// @Success      202 	{string} string "accepted"
// @Failure      400 {object} utils.ErrorResponse
// @Router       /squads/{squadId} [delete]
func removeSquad(c *fiber.Ctx) error {
	_, err := strconv.Atoi(c.Params("squadId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "SquadId supplied is invalid")
	}

	return c.SendStatus(fiber.StatusAccepted)
}

// @Summary      Retrieve players by squad
// @Description  List of players associated with squad
// @Tags         squad
// @Produce      json
// @Param        squadId   path  int  true  "Squad ID"
// @Success      200 {array} models.SquadPlayer
// @Failure      400 {object} utils.ErrorResponse
// @Router       /squads/{squadId}/players [get]
func retrieveSquadPlayers(c *fiber.Ctx) error {
	_, err := strconv.Atoi(c.Params("squadId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "SquadId supplied is invalid")
	}

	p := MockSquadPlayers
	return c.Status(fiber.StatusOK).JSON(p)
}

// @Summary      Approve squad request
// @Description  Approve player to join squad
// @Tags         squad
// @Produce      json
// @Param        squadId   path  int  true  "Squad ID"
// @Param        playerId   path  int  true  "Player ID"
// @Success      200 {string} string "ok"
// @Failure      400 {object} utils.ErrorResponse
// @Router       /squads/{squadId}/players/{playerId} [put]
func approvePlayerToSquad(c *fiber.Ctx) error {
	_, err := strconv.Atoi(c.Params("squadId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "SquadId supplied is invalid")
	}

	_, err = strconv.Atoi(c.Params("playerId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "PlayerId supplied is invalid")
	}

	return c.SendStatus(fiber.StatusOK)
}

// @Summary      Remove squad player
// @Description  Remove player from squad
// @Tags         squad
// @Produce      json
// @Param        squadId   path  int  true  "Squad ID"
// @Param        playerId   path  int  true  "Player ID"
// @Success      204
// @Failure      400 {object} utils.ErrorResponse
// @Router       /squads/{squadId}/players/{playerId} [delete]
func removePlayerFromSquad(c *fiber.Ctx) error {
	_, err := strconv.Atoi(c.Params("squadId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "SquadId supplied is invalid")
	}

	_, err = strconv.Atoi(c.Params("playerId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "PlayerId supplied is invalid")
	}

	return c.SendStatus(fiber.StatusNoContent)
}
