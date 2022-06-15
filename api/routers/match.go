package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/enums"
	"github.com/nmcalinden/footpal/middleware"
)

func ConfigureMatchHandlers(app *fiber.App) {
	group := app.Group("/matches", middleware.IsAuthenticated)

	matchController := InitializeMatchController()

	group.Use(middleware.NewRoles(enums.Player, enums.VenueAdmin).HasRole)

	group.Get("/", matchController.RetrieveMatches)
	group.Get("/:matchId", matchController.RetrieveMatchById)
	group.Delete("/:matchId", matchController.CancelMatch)
	group.Get("/:matchId/players", matchController.RetrieveMatchPlayers)
	group.Delete("/:matchId/players/:playerId", matchController.RemovePlayerFromMatch)
}
