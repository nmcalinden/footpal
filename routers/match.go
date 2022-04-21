package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/config"
	"github.com/nmcalinden/footpal/controllers"
	"github.com/nmcalinden/footpal/middleware"
	"github.com/nmcalinden/footpal/services"
)

func ConfigureMatchHandlers(app *fiber.App) {
	group := app.Group("/matches", middleware.IsAuthenticated)

	mService := services.NewMatchService(config.GetConnection())
	matchController := controllers.NewMatchController(mService)

	roles := []middleware.UserRole{{R: "player"}, {R: "venueAdmin"}}
	group.Use(middleware.NewRoles(roles).HasRole)

	group.Get("/", matchController.RetrieveMatches)
	group.Get("/:matchId", matchController.RetrieveMatchById)
	group.Delete("/:matchId", matchController.CancelMatch)
	group.Get("/:matchId/players", matchController.RetrieveMatchPlayers)
	group.Delete("/:matchId/players/:playerId", matchController.RemovePlayerFromMatch)
}
