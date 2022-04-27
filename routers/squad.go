package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/enums"
	"github.com/nmcalinden/footpal/middleware"
)

type SquadResponse struct {
	id *int
}

func ConfigureSquadPlayers(app *fiber.App) {
	group := app.Group("/squads")

	squadController := InitializeSquadController()

	group.Get("/", squadController.RetrieveSquads)
	group.Get("/:squadId", squadController.RetrieveSquadById)
	group.Get("/:squadId/players", squadController.RetrieveSquadPlayers)

	group.Use(middleware.IsAuthenticated)
	group.Use(middleware.NewRoles(enums.Player).HasRole)

	group.Post("/", squadController.CreateSquadGroup)
	group.Put("/:squadId", squadController.UpdateSquadInfo)
	group.Put("/:squadId/players/:playerId", squadController.ApprovePlayerToSquad)

	group.Use(middleware.NewRoles(enums.Player, enums.VenueAdmin).HasRole)

	group.Delete("/:squadId", squadController.RemoveSquad)
	group.Delete("/:squadId/players/:playerId", squadController.RemovePlayerFromSquad)
}
