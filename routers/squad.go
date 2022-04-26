package routers

import (
	"github.com/gofiber/fiber/v2"
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
	roles := []middleware.UserRole{{Role: "player"}}
	group.Use(middleware.NewRoles(roles).HasRole)

	group.Post("/", squadController.CreateSquadGroup)
	group.Put("/:squadId", squadController.UpdateSquadInfo)
	group.Put("/:squadId/players/:playerId", squadController.ApprovePlayerToSquad)

	roles = []middleware.UserRole{{Role: "player"}, {Role: "venueAdmin"}}
	group.Use(middleware.NewRoles(roles).HasRole)

	group.Delete("/:squadId", squadController.RemoveSquad)
	group.Delete("/:squadId/players/:playerId", squadController.RemovePlayerFromSquad)
}
