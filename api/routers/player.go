package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/api/enums"
	"github.com/nmcalinden/footpal/api/middleware"
)

func ConfigurePlayerHandlers(app *fiber.App) {
	group := app.Group("/players", middleware.IsAuthenticated)

	playerController := InitializePlayerController()

	group.Use(middleware.NewRoles(enums.All).HasRole)

	group.Get("/", playerController.RetrievePlayers)
	group.Get("/:playerId", playerController.RetrievePlayerById)

	group.Use(middleware.NewRoles(enums.Player).HasRole)

	group.Get("/:playerId/matches", playerController.GetPlayerMatches)
	group.Put("/", playerController.UpdatePlayer)
	group.Get("/:playerId/squads", playerController.GetSquadsByUser)
	group.Get("/:playerId/squads/:squadId", playerController.GetSquadByPlayer)
	group.Post("/:playerId/squads/:squadId", playerController.JoinSquad)
	group.Post("/:playerId/matches/:matchId", playerController.JoinMatch)
	group.Delete("/:playerId/matches/:matchId", playerController.LeaveMatch)
	group.Post("/:playerId/matches/:matchId/pay", playerController.MakePlayerPayment)
	group.Put("/:playerId/matches/:matchId/pay", playerController.UpdatePlayerPaymentType)
}
