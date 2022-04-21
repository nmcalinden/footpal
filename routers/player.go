package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/config"
	"github.com/nmcalinden/footpal/controllers"
	"github.com/nmcalinden/footpal/middleware"
	"github.com/nmcalinden/footpal/services"
)

func ConfigurePlayerHandlers(app *fiber.App) {
	group := app.Group("/players", middleware.IsAuthenticated)

	pService := services.NewPlayerService(config.GetConnection())
	playerController := controllers.NewPlayerController(pService)

	roles := []middleware.UserRole{{R: "player"}, {R: "venueAdmin"}}
	group.Use(middleware.NewRoles(roles).HasRole)

	group.Get("/", playerController.RetrievePlayers)
	group.Get("/:playerId", playerController.RetrievePlayerById)
	group.Get("/:playerId/matches", playerController.GetPlayerMatches)

	roles = []middleware.UserRole{{R: "player"}}
	group.Use(middleware.NewRoles(roles).HasRole)

	group.Put("/:playerId", playerController.UpdatePlayer)
	group.Get("/:playerId/squads", playerController.GetSquadsByUser)
	group.Get("/:playerId/squads/:squadId", playerController.GetSquadByPlayer)
	group.Post("/:playerId/squads/:squadId", playerController.JoinSquad)
	group.Post("/:playerId/matches/:matchId", playerController.JoinMatch)
	group.Delete("/:playerId/matches/:matchId", playerController.LeaveMatch)
	group.Post("/:playerId/matches/:matchId/pay", playerController.MakePlayerPayment)
	group.Put("/:playerId/matches/:matchId/pay", playerController.UpdatePlayerPaymentType)
}
