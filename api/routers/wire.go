//go:build wireinject
// +build wireinject

package routers

import (
	"github.com/google/wire"
	"github.com/nmcalinden/footpal/api/controllers"
	"github.com/nmcalinden/footpal/api/repository"
	"github.com/nmcalinden/footpal/api/services"
	"github.com/nmcalinden/footpal/config"
)

var BookingRepoSet = wire.NewSet(repository.NewBookingRepository, wire.Bind(new(repository.BookingRepositoryI), new(*repository.BookingRepository)))
var MatchRepoSet = wire.NewSet(repository.NewMatchRepository, wire.Bind(new(repository.MatchRepositoryI), new(*repository.MatchRepository)))
var MatchPlayerRepoSet = wire.NewSet(repository.NewMatchPlayerRepository, wire.Bind(new(repository.MatchPlayerRepositoryI), new(*repository.MatchPlayerRepository)))
var PitchRepoSet = wire.NewSet(repository.NewPitchRepository, wire.Bind(new(repository.PitchRepositoryI), new(*repository.PitchRepository)))
var PlayerRepoSet = wire.NewSet(repository.NewPlayerRepository, wire.Bind(new(repository.PlayerRepositoryI), new(*repository.PlayerRepository)))
var SquadRepoSet = wire.NewSet(repository.NewSquadRepository, wire.Bind(new(repository.SquadRepositoryI), new(*repository.SquadRepository)))
var StatusRepoSet = wire.NewSet(repository.NewStatusRepository, wire.Bind(new(repository.StatusRepositoryI), new(*repository.StatusRepository)))
var UserRepoSet = wire.NewSet(repository.NewUserRepository, wire.Bind(new(repository.UserRepositoryI), new(*repository.UserRepository)))
var VenueRepoSet = wire.NewSet(repository.NewVenueRepository, wire.Bind(new(repository.VenueRepositoryI), new(*repository.VenueRepository)))

func InitializeBookingController() *controllers.BookingController {
	wire.Build(controllers.NewBookingController, services.NewBookingService, BookingRepoSet, MatchRepoSet, StatusRepoSet, config.GetConnection)
	return nil
}

func InitializeMatchController() *controllers.MatchController {
	wire.Build(controllers.NewMatchController, services.NewMatchService, MatchRepoSet, MatchPlayerRepoSet, SquadRepoSet, config.GetConnection)
	return nil
}

func InitializePlayerController() *controllers.PlayerController {
	wire.Build(
		controllers.NewPlayerController, services.NewPlayerService, PlayerRepoSet, UserRepoSet, MatchRepoSet,
		MatchPlayerRepoSet, SquadRepoSet, config.GetConnection)
	return nil
}

func InitializeSquadController() *controllers.SquadController {
	wire.Build(controllers.NewSquadController, services.NewSquadService, PlayerRepoSet, SquadRepoSet, config.GetConnection)
	return nil
}

func InitializeUserController() *controllers.UserController {
	wire.Build(controllers.NewUserController, services.NewUserService, UserRepoSet, PlayerRepoSet,
		VenueRepoSet, config.GetConnection)
	return nil
}

func InitializeVenueController() *controllers.VenueController {
	wire.Build(controllers.NewVenueController, services.NewVenueService, VenueRepoSet, PitchRepoSet, config.GetConnection)
	return nil
}
