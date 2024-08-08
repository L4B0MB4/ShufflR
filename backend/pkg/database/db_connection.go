package database

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
)

type DatabaseConnection struct {
	initialized bool
	db          *sql.DB
}

func (d *DatabaseConnection) SetUp() {
	db, err := sql.Open("sqlite3", "./shufflr.db")
	if err != nil {

		log.Info().Err(err).Msg("Opening sqlite connection")
		return
	}
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS users (id TEXT PRIMARY KEY, data BLOB)")
	if err != nil {

		log.Info().Err(err).Msg("Preparing statement for users table")
		return
	}
	_, err = stmt.Exec()
	if err != nil {

		log.Info().Err(err).Msg("Creating users table")
		return
	}
	d.db = db
	d.initialized = true
}

func (d *DatabaseConnection) GetDbConnection() (*sql.DB, error) {
	if !d.initialized {
		return nil, errors.New("DatabaseConnection not properly initialized")
	}
	return d.db, nil
}
