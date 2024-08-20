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
	if createUsersTable(db) != nil {
		return
	}
	if createTopTracksTable(db) != nil {
		return
	}
	d.db = db
	d.initialized = true
}

func createUsersTable(db *sql.DB) error {
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS users (id TEXT PRIMARY KEY, data BLOB)")
	if err != nil {

		log.Info().Err(err).Msg("Preparing statement for users table")
		return err
	}
	_, err = stmt.Exec()
	if err != nil {

		log.Info().Err(err).Msg("Creating users table")
		return err
	}
	return nil
}

func createTopTracksTable(db *sql.DB) error {
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS topTracks (userId TEXTNOT NULL,data BLOB, fromDate INTEGER,UNIQUE(userId, fromDate) ON CONFLICT FAIL )")
	if err != nil {

		log.Info().Err(err).Msg("Preparing statement for users table")
		return err
	}
	_, err = stmt.Exec()
	if err != nil {

		log.Info().Err(err).Msg("Creating users table")
		return err
	}
	return nil
}

func (d *DatabaseConnection) GetDbConnection() (*sql.DB, error) {
	if !d.initialized {
		return nil, errors.New("DatabaseConnection not properly initialized")
	}
	return d.db, nil
}
