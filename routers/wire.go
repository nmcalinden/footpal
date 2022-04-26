//go:build wireinject
// +build wireinject

package routers

import (
	"github.com/google/wire"
	"github.com/nmcalinden/footpal/config"
	"github.com/nmcalinden/footpal/controllers"
	"github.com/nmcalinden/footpal/repository"
	"github.com/nmcalinden/footpal/services"
)

func InitializeBookingController() *controllers.BookingController {
	wire.Build(
		controllers.NewBookingController, services.NewBookingService,
		repository.BookingRepoSet, config.GetConnection)
	return nil
}

func InitializeMatchController() *controllers.MatchController {
	wire.Build(
		controllers.NewMatchController, services.NewMatchService,
		repository.MatchRepoSet, repository.MatchPlayerRepoSet,
		config.GetConnection)
	return nil
}

func InitializePlayerController() *controllers.PlayerController {
	wire.Build(
		controllers.NewPlayerController, services.NewPlayerService,
		repository.PlayerRepoSet, repository.UserRepoSet, repository.MatchRepoSet,
		repository.MatchPlayerRepoSet, repository.SquadRepoSet,
		config.GetConnection)
	return nil
}

func InitializeSquadController() *controllers.SquadController {
	wire.Build(
		controllers.NewSquadController, services.NewSquadService,
		repository.PlayerRepoSet, repository.SquadRepoSet,
		config.GetConnection)
	return nil
}

func InitializeUserController() *controllers.UserController {
	wire.Build(
		controllers.NewUserController, services.NewUserService,
		repository.UserRepoSet, repository.PlayerRepoSet, repository.VenueRepoSet,
		config.GetConnection)
	return nil
}

func InitializeVenueController() *controllers.VenueController {
	wire.Build(
		controllers.NewVenueController, services.NewVenueService,
		repository.VenueRepoSet,
		config.GetConnection)
	return nil
}
