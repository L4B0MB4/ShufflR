package database

import (
	"database/sql"
	"encoding/json"

	"github.com/L4B0MB4/Musicfriends/pkg/models"
	"github.com/rs/zerolog/log"
)

func deserializeRow(row *sql.Row, out interface{}) error {
	var bytes []byte
	err := row.Scan(&bytes)
	if err != nil {
		log.Debug().Err(err).Msg("err")
		return err
	}
	err = json.Unmarshal(bytes, out)
	if err != nil {
		log.Debug().Err(err).Msg("err")
		return err
	}
	return nil
}

func GetUserProfile(db *DatabaseConnection, userId string) *models.CurrentUserProfile {
	con, _ := db.GetDbConnection()
	stmt, err := con.Prepare("Select data FROM users where id =?")
	if err != nil {
		return nil
	}
	var profileModel models.CurrentUserProfile
	row := stmt.QueryRow(userId)
	if deserializeRow(row, &profileModel) != nil {
		return nil
	}
	return &profileModel
}

func InsertUserProfile(db *DatabaseConnection, user *models.CurrentUserProfile) error {
	con, _ := db.GetDbConnection()
	stmt, err := con.Prepare("Insert into users (id,data) values(?,?)")
	if err != nil {
		return nil
	}
	body, err := json.Marshal(user)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.ID, body)
	if err != nil {
		return err
	}
	return nil
}

func GetTopTracks(db *DatabaseConnection, userId string) *models.TopTracksResponse {

	con, _ := db.GetDbConnection()
	stmt, err := con.Prepare("Select data FROM topTracks where userId =?")
	if err != nil {
		return nil
	}
	var topTracksModel models.TopTracksResponse
	row := stmt.QueryRow(userId)
	if deserializeRow(row, &topTracksModel) != nil {
		return nil
	}
	return &topTracksModel
}
