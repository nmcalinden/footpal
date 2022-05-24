// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package routers

import (
	"github.com/google/wire"
	"github.com/nmcalinden/footpal/api/controllers"
	"github.com/nmcalinden/footpal/api/repository"
	"github.com/nmcalinden/footpal/api/services"
	"github.com/nmcalinden/footpal/config"
)

// Injectors from wire.go:

func InitializeBookingController() *controllers.BookingController {
	db := config.GetConnection()
	bookingRepository := repository.NewBookingRepository(db)
	bookingService := services.NewBookingService(bookingRepository)
	bookingController := controllers.NewBookingController(bookingService)
	return bookingController
}

func InitializeMatchController() *controllers.MatchController {
	db := config.GetConnection()
	matchRepository := repository.NewMatchRepository(db)
	matchPlayerRepository := repository.NewMatchPlayerRepository(db)
	squadRepository := repository.NewSquadRepository(db)
	matchService := services.NewMatchService(matchRepository, matchPlayerRepository, squadRepository)
	matchController := controllers.NewMatchController(matchService)
	return matchController
}

func InitializePlayerController() *controllers.PlayerController {
	db := config.GetConnection()
	playerRepository := repository.NewPlayerRepository(db)
	userRepository := repository.NewUserRepository(db)
	matchRepository := repository.NewMatchRepository(db)
	matchPlayerRepository := repository.NewMatchPlayerRepository(db)
	squadRepository := repository.NewSquadRepository(db)
	playerService := services.NewPlayerService(playerRepository, userRepository, matchRepository, matchPlayerRepository, squadRepository)
	playerController := controllers.NewPlayerController(playerService)
	return playerController
}

func InitializeSquadController() *controllers.SquadController {
	db := config.GetConnection()
	playerRepository := repository.NewPlayerRepository(db)
	squadRepository := repository.NewSquadRepository(db)
	squadService := services.NewSquadService(playerRepository, squadRepository)
	squadController := controllers.NewSquadController(squadService)
	return squadController
}

func InitializeUserController() *controllers.UserController {
	db := config.GetConnection()
	userRepository := repository.NewUserRepository(db)
	playerRepository := repository.NewPlayerRepository(db)
	venueRepository := repository.NewVenueRepository(db)
	userService := services.NewUserService(userRepository, playerRepository, venueRepository)
	userController := controllers.NewUserController(userService)
	return userController
}

func InitializeVenueController() *controllers.VenueController {
	db := config.GetConnection()
	venueRepository := repository.NewVenueRepository(db)
	pitchRepository := repository.NewPitchRepository(db)
	venueService := services.NewVenueService(venueRepository, pitchRepository)
	venueController := controllers.NewVenueController(venueService)
	return venueController
}

// wire.go:

var BookingRepoSet = wire.NewSet(repository.NewBookingRepository, wire.Bind(new(repository.BookingRepositoryI), new(*repository.BookingRepository)))

var MatchRepoSet = wire.NewSet(repository.NewMatchRepository, wire.Bind(new(repository.MatchRepositoryI), new(*repository.MatchRepository)))

var MatchPlayerRepoSet = wire.NewSet(repository.NewMatchPlayerRepository, wire.Bind(new(repository.MatchPlayerRepositoryI), new(*repository.MatchPlayerRepository)))

var PitchRepoSet = wire.NewSet(repository.NewPitchRepository, wire.Bind(new(repository.PitchRepositoryI), new(*repository.PitchRepository)))

var PlayerRepoSet = wire.NewSet(repository.NewPlayerRepository, wire.Bind(new(repository.PlayerRepositoryI), new(*repository.PlayerRepository)))

var SquadRepoSet = wire.NewSet(repository.NewSquadRepository, wire.Bind(new(repository.SquadRepositoryI), new(*repository.SquadRepository)))

var UserRepoSet = wire.NewSet(repository.NewUserRepository, wire.Bind(new(repository.UserRepositoryI), new(*repository.UserRepository)))

var VenueRepoSet = wire.NewSet(repository.NewVenueRepository, wire.Bind(new(repository.VenueRepositoryI), new(*repository.VenueRepository)))
