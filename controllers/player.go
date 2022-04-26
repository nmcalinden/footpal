package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/errors"
	"github.com/nmcalinden/footpal/payloads"
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
// @Param        limit   query  int  true  "Limit" default(10)
// @Param        after_id   query  int  true  "After ID" default(0)
// @Success      200  {object}  views.Players
// @Failure      500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router       /players [get]
func (con PlayerController) RetrievePlayers(c *fiber.Ctx) error {
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Limit supplied is invalid")
	}

	after, err := strconv.Atoi(c.Query("after_id"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "ID supplied is invalid")
	}

	p, err := con.playerService.GetPlayers(limit, after)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to get Players")
	}
	return c.Status(fiber.StatusOK).JSON(p)
}

// UpdatePlayer @Summary      Edit Player
// @Description  Edit player information
// @Tags         player
// @Produce      json
// @Param 		 message body payloads.PlayerRequest true "Request"
// @Success      200 {object} models.Player
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router       /players [put]
func (con PlayerController) UpdatePlayer(c *fiber.Ctx) error {
	claims := utils.GetClaims(c.Locals("user"))
	userId := claims["id"].(int)

	p := new(payloads.PlayerRequest)
	if err := c.BodyParser(&p); err != nil {
		return err
	}

	if e := utils.ValidateStruct(*p); e != nil {
		return c.Status(fiber.StatusBadRequest).JSON(e)
	}

	res, err := con.playerService.EditPlayer(&userId, p)
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
// @Success      200 {object} views.Player
// @Failure      400 {object} utils.ErrorResponse
// @Failure      404 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router       /players/{playerId} [get]
func (con PlayerController) RetrievePlayerById(c *fiber.Ctx) error {
	playerId, err := strconv.Atoi(c.Params("playerId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Player Id supplied is invalid")
	}

	p, err := con.playerService.GetPlayerById(&playerId)
	if err != nil {
		e, ok := err.(*errors.FpError)
		if ok {
			return utils.BuildErrorResponse(c, e.Status, e.Error())
		}
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to retrieve player")
	}
	return c.Status(fiber.StatusOK).JSON(p)
}

// GetSquadsByUser @Summary      Retrieve Squads
// @Description  Retrieve all squads linked to player
// @Tags         player
// @Produce      json
// @Success      200  {array}  models.Squad
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router       /players/squads [get]
func (con PlayerController) GetSquadsByUser(c *fiber.Ctx) error {
	claims := utils.GetClaims(c.Locals("user"))
	userId := claims["id"].(int)

	p, err := con.playerService.GetAllSquadsByPlayer(&userId)
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
// @Security ApiKeyAuth
// @Router       /players/squads/{squadId} [get]
func (con PlayerController) GetSquadByPlayer(c *fiber.Ctx) error {
	claims := utils.GetClaims(c.Locals("user"))
	userId := claims["id"].(int)

	squadId, err := strconv.Atoi(c.Params("squadId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Squad Id supplied is invalid")
	}

	p, err := con.playerService.GetSquadByPlayer(&userId, &squadId)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to get squad")
	}

	return c.Status(fiber.StatusOK).JSON(p)
}

// GetPlayerMatches @Summary      Retrieve Matches by player
// @Description  Retrieve a players opted in matches
// @Tags         player
// @Produce      json
// @Success      200  {array}  models.Match
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router       /players/matches [get]
func (con PlayerController) GetPlayerMatches(c *fiber.Ctx) error {
	claims := utils.GetClaims(c.Locals("user"))
	userId := claims["id"].(int)

	m, err := con.playerService.GetMatchesByPlayer(&userId)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to get player matches")
	}
	return c.Status(fiber.StatusOK).JSON(m)
}

// JoinSquad @Summary      Join squad
// @Description  Request to join squad
// @Tags         player
// @Produce      json
// @Param        squadId   path  int  true  "Squad ID"
// @Success      202
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router       /players/squads/{squadId} [post]
func (con PlayerController) JoinSquad(c *fiber.Ctx) error {
	claims := utils.GetClaims(c.Locals("user"))
	userId := claims["id"].(int)

	squadId, err := strconv.Atoi(c.Params("squadId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Match Id supplied is invalid")
	}

	err = con.playerService.JoinSquad(&userId, &squadId)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed request to join match")
	}

	return c.SendStatus(fiber.StatusAccepted)
}

// JoinMatch @Summary      Join match
// @Description  Request to join match
// @Tags         player
// @Produce      json
// @Param        matchId   path  int  true  "Match ID"
// @Success      200  {object} payloads.JoinMatchResponse
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router       /players/matches/{matchId} [post]
func (con PlayerController) JoinMatch(c *fiber.Ctx) error {
	claims := utils.GetClaims(c.Locals("user"))
	userId := claims["id"].(int)

	matchId, err := strconv.Atoi(c.Params("matchId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Match Id supplied is invalid")
	}

	res, err := con.playerService.JoinMatch(&userId, &matchId)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed request to join match")
	}

	response := payloads.JoinMatchResponse{MatchId: res}
	return c.Status(fiber.StatusOK).JSON(response)
}

// LeaveMatch @Summary      Leave match
// @Description  Leave match
// @Tags         player
// @Produce      json
// @Param        matchId   path  int  true  "Match ID"
// @Success      200
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router       /players/matches/{matchId} [delete]
func (con PlayerController) LeaveMatch(c *fiber.Ctx) error {
	claims := utils.GetClaims(c.Locals("user"))
	userId := claims["id"].(int)

	matchId, err := strconv.Atoi(c.Params("matchId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Match Id supplied is invalid")
	}

	err = con.playerService.LeaveMatch(&userId, &matchId)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed request to leave match")
	}

	return c.SendStatus(fiber.StatusOK)
}

// MakePlayerPayment @Summary      Match player payment
// @Description  For a match, pay amount owed by player
// @Tags         player
// @Produce      json
// @Param        matchId   path  int  true  "Match ID"
// @Param 		 message body payloads.PlayerPaymentRequest true "Request"
// @Success      200
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router       /players/matches/{matchId}/pay [post]
func (con PlayerController) MakePlayerPayment(c *fiber.Ctx) error {
	claims := utils.GetClaims(c.Locals("user"))
	userId := claims["id"].(int)

	matchId, err := strconv.Atoi(c.Params("matchId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Match Id supplied is invalid")
	}

	pay := new(payloads.PlayerPaymentRequest)
	if err := c.BodyParser(&pay); err != nil {
		return err
	}
	if e := utils.ValidateStruct(*pay); e != nil {
		return c.Status(fiber.StatusBadRequest).JSON(e)
	}

	err = con.playerService.Pay(&userId, &matchId, &pay.AmountToPay)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to make payment for match")
	}
	return c.SendStatus(fiber.StatusOK)
}

// UpdatePlayerPaymentType @Summary      Update player payment type
// @Description  For a match, pay amount owed by player
// @Tags         player
// @Produce      json
// @Param        matchId   path  int  true  "Match ID"
// @Param        paymentTypeId   query  int  true  "Payment Type ID"
// @Success      200
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router       /players/matches/{matchId}/pay [put]
func (con PlayerController) UpdatePlayerPaymentType(c *fiber.Ctx) error {
	claims := utils.GetClaims(c.Locals("user"))
	userId := claims["id"].(int)

	matchId, err := strconv.Atoi(c.Params("matchId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Match Id supplied is invalid")
	}

	paymentTypeId, err := strconv.Atoi(c.Query("paymentTypeId"))
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusBadRequest, "Payment Type Id supplied is invalid")
	}

	err = con.playerService.UpdatePaymentMethod(&matchId, &userId, &paymentTypeId)
	if err != nil {
		return utils.BuildErrorResponse(c, fiber.StatusInternalServerError, "Failed to update payment type")
	}
	return c.SendStatus(fiber.StatusOK)
}
