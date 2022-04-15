package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nmcalinden/footpal/config"
	"github.com/nmcalinden/footpal/controllers"
	"github.com/nmcalinden/footpal/services"
)

type SquadResponse struct {
	id *int
}

func ConfigureSquadPlayers(app *fiber.App) {
	group := app.Group("/squads")

	sService := services.NewSquadService(config.GetConnection())
	squadController := controllers.NewSquadController(sService)

	group.Get("/", squadController.RetrieveSquads)
	group.Post("/", squadController.CreateSquadGroup)
	group.Get("/:squadId", squadController.RetrieveSquadById)
	group.Put("/:squadId", squadController.UpdateSquadInfo)
	group.Delete("/:squadId", squadController.RemoveSquad)
	group.Get("/:squadId/players", squadController.RetrieveSquadPlayers)
	group.Put("/:squadId/players/:playerId", squadController.ApprovePlayerToSquad)
	group.Delete("/:squadId/players/:playerId", squadController.RemovePlayerFromSquad)
}
