//go:build wireinject
// +build wireinject

package routers

import (
	"github.com/google/wire"
	controllers2 "github.com/nmcalinden/footpal/api/controllers"
	"github.com/nmcalinden/footpal/api/repository"
	services2 "github.com/nmcalinden/footpal/api/services"
	"github.com/nmcalinden/footpal/config"
)

var BookingRepoSet = wire.NewSet(repository.NewBookingRepository, wire.Bind(new(repository.BookingRepositoryI), new(*repository.BookingRepository)))
var MatchRepoSet = wire.NewSet(repository.NewMatchRepository, wire.Bind(new(repository.MatchRepositoryI), new(*repository.MatchRepository)))
var MatchPlayerRepoSet = wire.NewSet(repository.NewMatchPlayerRepository, wire.Bind(new(repository.MatchPlayerRepositoryI), new(*repository.MatchPlayerRepository)))
var PlayerRepoSet = wire.NewSet(repository.NewPlayerRepository, wire.Bind(new(repository.PlayerRepositoryI), new(*repository.PlayerRepository)))
var SquadRepoSet = wire.NewSet(repository.NewSquadRepository, wire.Bind(new(repository.SquadRepositoryI), new(*repository.SquadRepository)))
var UserRepoSet = wire.NewSet(repository.NewUserRepository, wire.Bind(new(repository.UserRepositoryI), new(*repository.UserRepository)))
var VenueRepoSet = wire.NewSet(repository.NewVenueRepository, wire.Bind(new(repository.VenueRepositoryI), new(*repository.VenueRepository)))

func InitializeBookingController() *controllers2.BookingController {
	wire.Build(controllers2.NewBookingController, services2.NewBookingService, BookingRepoSet, config.GetConnection)
	return nil
}

func InitializeMatchController() *controllers2.MatchController {
	wire.Build(controllers2.NewMatchController, services2.NewMatchService, MatchRepoSet, MatchPlayerRepoSet, config.GetConnection)
	return nil
}

func InitializePlayerController() *controllers2.PlayerController {
	wire.Build(
		controllers2.NewPlayerController, services2.NewPlayerService, PlayerRepoSet, UserRepoSet, MatchRepoSet,
		MatchPlayerRepoSet, SquadRepoSet, config.GetConnection)
	return nil
}

func InitializeSquadController() *controllers2.SquadController {
	wire.Build(controllers2.NewSquadController, services2.NewSquadService, PlayerRepoSet, SquadRepoSet, config.GetConnection)
	return nil
}

func InitializeUserController() *controllers2.UserController {
	wire.Build(controllers2.NewUserController, services2.NewUserService, UserRepoSet, PlayerRepoSet,
		VenueRepoSet, config.GetConnection)
	return nil
}

func InitializeVenueController() *controllers2.VenueController {
	wire.Build(controllers2.NewVenueController, services2.NewVenueService, VenueRepoSet, config.GetConnection)
	return nil
}
