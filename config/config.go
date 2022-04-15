package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "password"
	dbname   = "footpaldb"
)

var conn *sqlx.DB

func InitializeDatabase() {
	db, err := sqlx.Open("postgres", getDSN())
	if err != nil {
		panic("Failed to connect to database")
	}

	// https://github.com/go-sql-driver/mysql#important-settings
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(time.Hour)

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)

	conn = db
}

func GetConnection() *sqlx.DB {
	return conn
}

func getDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}
