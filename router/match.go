package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/config"
	"github.com/nmcalinden/footpal/controllers"
	"github.com/nmcalinden/footpal/service"
)

func ConfigureMatchHandlers(app *fiber.App) {
	group := app.Group("/matches")

	mService := service.NewMatchService(config.GetConnection())
	matchController := controllers.NewMatchController(mService)

	group.Get("/", matchController.RetrieveMatches)
	group.Get("/:matchId", matchController.RetrieveMatchById)
	group.Delete("/:matchId", matchController.CancelMatch)
	group.Get("/:matchId/players", matchController.RetrieveMatchPlayers)
	group.Delete("/:matchId/players/:playerId", matchController.RemovePlayerFromMatch)
}
