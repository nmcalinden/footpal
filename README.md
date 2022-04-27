# Footpal API

Golang POC for a 5 a side booking app

- Book pitches 
- Join Football groups
- Opt in to games
- Connect with local football community

### Local Setup
1. Go `go1.16+`
2. Docker `latest`
3. IDE `GoLand 2021.3.4`
4. Environment Variables: `ACCESS_SECRET` and `REFRESH_SECRET` _Note_: Add environment variables either via terminal or IDE Configuration

## Run/Debug Configuration
1. Clone Repo
2. Install Dependencies: `go install`
3. Initialize DB: `docker-compose up`
4. Run: `go run main.go` or `air` (live reloads)
5. Visit: `http://localhost:3000`
   1. If this does not work, use url from start up in terminal
6. Swagger: `http://localhost:3000/swagger`


 ## Dependencies
|                                                             | Description                                                                                                                                                                                                                                                                                                                                                                                   |
|-------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [air](https://github.com/cosmtrek/air)                      | Live reload for Go apps                                                                                                                                                                                                                                                                                                                                                                       |
| [gofiber](https://github.com/gofiber/fiber)                 | Fiber is an Express inspired web framework built on top of Fasthttp, the fastest HTTP engine for Go. Designed to ease things up for fast development with zero memory allocation and performance in mind.                                                                                                                                                                                     |
| [golang-jwt](https://github.com/golang-jwt/jwt)             | A golang implementation of JSON Web Tokens.                                                                                                                                                                                                                                                                                                                                                   |
| [golang/mock](https://github.com/golang/mock)               | gomock is a mocking framework for the Go programming language. It integrates well with Go's built-in testing package, but can be used in other contexts too.                                                                                                                                                                                                                                  |
| [go-model](https://github.com/jeevatkm/go-model)            | Robust & Easy to use model mapper and utility methods for Go struct. Typical methods increase productivity and make Go development more fun                                                                                                                                                                                                                                                   |
| [go-multierror](https://github.com/hashicorp/go-multierror) | go-multierror is a package for Go that provides a mechanism for representing a list of error values as a single error. <br/>This allows a function in Go to return an error that might actually be a list of errors. If the caller knows this, they can unwrap the list and access the errors. If the caller doesn't know, the error formats to a nice human-readable format.                 |
| [google/wire](https://github.com/google/wire)               | Wire is a code generation tool that automates connecting components using dependency injection. Dependencies between components are represented in Wire as function parameters, encouraging explicit initialization instead of global variables. Because Wire operates without runtime state or reflection, code written to be used with Wire is useful even for hand-written initialization. |
| [guregu/null](https://github.com/guregu/null)               | null is a library with reasonable options for dealing with nullable SQL and JSON values                                                                                                                                                                                                                                                                                                       |
| [sqlx](https://github.com/jmoiron/sqlx)                     | sqlx is a library which provides a set of extensions on go's standard database/sql library. The sqlx versions of sql.DB, sql.TX, sql.Stmt, et al. all leave the underlying interfaces untouched, so that their interfaces are a superset on the standard ones. This makes it relatively painless to integrate existing codebases using database/sql with sqlx.                                |
| [swaggo/swag](https://github.com/swaggo/swag)               | Swag converts Go annotations to Swagger Documentation 2.0.                                                                                                                                                                                                                                                                                                                                    |
