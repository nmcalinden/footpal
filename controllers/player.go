package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/models"
	"github.com/nmcalinden/footpal/services"
	"github.com/nmcalinden/footpal/utils"
	"strconv"
)

type PlayerController struct {
	playerService *services.PlayerService
}

func NewPlayerController(playerService *services.PlayerService) *PlayerController {
	return &PlayerController{playerService: playerService}
}

// RetrievePlayers @Summary      Retrieve Players
// @Description  Retrieve all players
// @Tags         player
// @Produce      json
// @Success      200  {array}  models.Player
// @Failure      500 {object} utils.ErrorResponse
// @Router       /players [get]
func (controller PlayerController) RetrievePlayers(c *fiber.Ctx) error {
	p, err := controller.playerService.GetPlayers()
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to get Players")
	}
	return c.Status(fiber.StatusOK).JSON(p)
}

// UpdatePlayer @Summary      Edit Player
// @Description  Edit player information
// @Tags         player
// @Produce      json
// @Param        playerId   path  int  true  "Player ID"
// @Param 		 message body models.PlayerRequest true "Request"
// @Success      200 {object} models.Player
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Router       /players/{playerId} [put]
func (controller PlayerController) UpdatePlayer(c *fiber.Ctx) error {
	playerId, err := strconv.Atoi(c.Params("playerId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "PlayerId supplied is invalid")
	}

	p := new(models.PlayerRequest)
	if err := c.BodyParser(&p); err != nil {
		return err
	}

	if errors := utils.ValidateStruct(*p); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	res, err := controller.playerService.EditPlayer(&playerId, p)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to edit player")
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

// RetrievePlayerById @Summary      Retrieve Player
// @Description  Get Player information
// @Tags         player
// @Produce      json
// @Param        playerId   path  int  true  "Player ID"
// @Success      200 {object} models.Player
// @Failure      400 {object} utils.ErrorResponse
// @Router       /players/{playerId} [get]
func (controller PlayerController) RetrievePlayerById(c *fiber.Ctx) error {
	playerId, err := strconv.Atoi(c.Params("playerId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Player Id supplied is invalid")
	}

	p, err := controller.playerService.GetPlayerById(&playerId)

	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Player does not exist")
	}
	return c.Status(fiber.StatusOK).JSON(p)
}

// GetSquadsByUser @Summary      Retrieve Squads
// @Description  Retrieve all squads linked to player
// @Tags         player
// @Produce      json
// @Param        playerId   path  int  true  "Player ID"
// @Success      200  {array}  models.Squad
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Router       /players/{playerId}/squads [get]
func (controller PlayerController) GetSquadsByUser(c *fiber.Ctx) error {
	playerId, err := strconv.Atoi(c.Params("playerId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Player Id supplied is invalid")
	}

	p, err := controller.playerService.GetAllSquadsByPlayer(&playerId)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to get squads by player")
	}
	return c.Status(fiber.StatusOK).JSON(p)
}

// GetSquadByPlayer @Summary      Retrieve Squad by squadId
// @Description  Retrieve squad by squad id
// @Tags         player
// @Produce      json
// @Param        playerId   path  int  true  "Player ID"
// @Param        squadId   path  int  true  "Squad ID"
// @Success      200  {array}  models.Player
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Router       /players/{playerId}/squads/{squadId} [get]
func (controller PlayerController) GetSquadByPlayer(c *fiber.Ctx) error {
	playerId, err := strconv.Atoi(c.Params("playerId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Player Id supplied is invalid")
	}

	squadId, err := strconv.Atoi(c.Params("squadId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Squad Id supplied is invalid")
	}

	p, err := controller.playerService.GetSquadByPlayer(&playerId, &squadId)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to get squad")
	}

	return c.Status(fiber.StatusOK).JSON(p)
}

// GetPlayerMatches @Summary      Retrieve Matches by player
// @Description  Retrieve a players opted in matches
// @Tags         player
// @Produce      json
// @Param        playerId   path  int  true  "Player ID"
// @Success      200  {array}  models.Match
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Router       /players/{playerId}/matches [get]
func (controller PlayerController) GetPlayerMatches(c *fiber.Ctx) error {
	playerId, err := strconv.Atoi(c.Params("playerId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Player Id supplied is invalid")
	}

	m, err := controller.playerService.GetMatchesByPlayer(&playerId)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to get player matches")
	}
	return c.Status(fiber.StatusOK).JSON(m)
}

// JoinSquad @Summary      Join squad
// @Description  Request to join squad
// @Tags         player
// @Produce      json
// @Param        playerId   path  int  true  "Player ID"
// @Param        squadId   path  int  true  "Squad ID"
// @Success      202  {object} models.JoinMatchResponse
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Router       /players/{playerId}/squads/{squadId} [post]
func (controller PlayerController) JoinSquad(c *fiber.Ctx) error {
	playerId, err := strconv.Atoi(c.Params("playerId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Player Id supplied is invalid")
	}

	squadId, err := strconv.Atoi(c.Params("squadId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Match Id supplied is invalid")
	}

	err = controller.playerService.JoinSquad(&playerId, &squadId)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed request to join match")
	}

	return c.SendStatus(fiber.StatusAccepted)
}

// JoinMatch @Summary      Join match
// @Description  Request to join match
// @Tags         player
// @Produce      json
// @Param        playerId   path  int  true  "Player ID"
// @Param        matchId   path  int  true  "Match ID"
// @Success      200  {object} models.JoinMatchResponse
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Router       /players/{playerId}/matches/{matchId} [post]
func (controller PlayerController) JoinMatch(c *fiber.Ctx) error {
	playerId, err := strconv.Atoi(c.Params("playerId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Player Id supplied is invalid")
	}

	matchId, err := strconv.Atoi(c.Params("matchId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Match Id supplied is invalid")
	}

	res, err := controller.playerService.JoinMatch(&playerId, &matchId)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed request to join match")
	}

	response := models.JoinMatchResponse{MatchId: res}
	return c.Status(fiber.StatusOK).JSON(response)
}

// LeaveMatch @Summary      Leave match
// @Description  Leave match
// @Tags         player
// @Produce      json
// @Param        playerId   path  int  true  "Player ID"
// @Param        matchId   path  int  true  "Match ID"
// @Success      200
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Router       /players/{playerId}/matches/{matchId} [delete]
func (controller PlayerController) LeaveMatch(c *fiber.Ctx) error {
	playerId, err := strconv.Atoi(c.Params("playerId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Player Id supplied is invalid")
	}

	matchId, err := strconv.Atoi(c.Params("matchId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Match Id supplied is invalid")
	}

	err = controller.playerService.LeaveMatch(&playerId, &matchId)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed request to leave match")
	}

	return c.SendStatus(fiber.StatusOK)
}

// MakePlayerPayment @Summary      Match player payment
// @Description  For a match, pay amount owed by player
// @Tags         player
// @Produce      json
// @Param        playerId   path  int  true  "Player ID"
// @Param        matchId   path  int  true  "Match ID"
// @Param 		 message body models.MatchPaymentRequest true "Request"
// @Success      200
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Router       /players/{playerId}/matches/{matchId}/pay [post]
func (controller PlayerController) MakePlayerPayment(c *fiber.Ctx) error {
	playerId, err := strconv.Atoi(c.Params("playerId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Player Id supplied is invalid")
	}

	matchId, err := strconv.Atoi(c.Params("matchId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Match Id supplied is invalid")
	}

	pay := new(models.MatchPaymentRequest)
	if err := c.BodyParser(&pay); err != nil {
		return err
	}
	if errors := utils.ValidateStruct(*pay); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	err = controller.playerService.Pay(&playerId, &matchId, &pay.AmountToPay)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to make payment for match")
	}
	return c.SendStatus(fiber.StatusOK)
}

// UpdatePlayerPaymentType @Summary      Update player payment type
// @Description  For a match, pay amount owed by player
// @Tags         player
// @Produce      json
// @Param        playerId   path  int  true  "Player ID"
// @Param        matchId   path  int  true  "Match ID"
// @Param        paymentTypeId   query  int  true  "Payment Type ID"
// @Success      200
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Router       /players/{playerId}/matches/{matchId}/pay [patch]
func (controller PlayerController) UpdatePlayerPaymentType(c *fiber.Ctx) error {
	playerId, err := strconv.Atoi(c.Params("playerId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Player Id supplied is invalid")
	}

	matchId, err := strconv.Atoi(c.Params("matchId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Match Id supplied is invalid")
	}

	paymentTypeId, err := strconv.Atoi(c.Query("paymentTypeId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Payment Type Id supplied is invalid")
	}

	err = controller.playerService.UpdatePaymentMethod(&matchId, &playerId, &paymentTypeId)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to update payment type")
	}
	return c.SendStatus(fiber.StatusOK)
}
