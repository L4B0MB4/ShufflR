package database

import (
	"database/sql"
	"encoding/json"
	"time"

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
	stmt, err := con.Prepare("Select data FROM topTracks where userId =? where fromDate>?")
	if err != nil {
		return nil
	}
	var topTracksModel models.TopTracksResponse
	row := stmt.QueryRow(userId, time.Now().Add(-(time.Hour * 24)))
	if deserializeRow(row, &topTracksModel) != nil {
		return nil
	}
	return &topTracksModel
}

func SaveTopTracks(db *DatabaseConnection, userId string, topTracks *models.TopTracksResponse) error {
	con, _ := db.GetDbConnection()
	stmt, err := con.Prepare("Insert into topTracks (userId,data,fromDate) values(?,?,?)")
	if err != nil {
		return nil
	}
	body, err := json.Marshal(topTracks)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(userId, body, time.Now().Unix())
	if err != nil {
		return err
	}
	return nil
}
