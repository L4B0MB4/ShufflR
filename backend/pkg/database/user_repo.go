package database

import (
	"encoding/json"

	"github.com/L4B0MB4/Musicfriends/pkg/models"
	"github.com/rs/zerolog/log"
)

func GetUserProfile(db *DatabaseConnection, userId string) *models.CurrentUserProfile {
	con, _ := db.GetDbConnection()
	stmt, err := con.Prepare("Select data FROM users where id =?")
	if err != nil {
		return nil
	}
	var profile []byte
	err = stmt.QueryRow(userId).Scan(&profile)
	if err != nil {
		log.Debug().Err(err).Msg("err")
		return nil
	}
	var profileModel models.CurrentUserProfile
	err = json.Unmarshal(profile, &profileModel)
	if err != nil {
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
