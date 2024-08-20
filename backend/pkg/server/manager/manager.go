package manager

import (
	"github.com/L4B0MB4/Musicfriends/pkg/database"
	"github.com/L4B0MB4/Musicfriends/pkg/models"
)

type PersonalInfoManager struct {
	db *database.DatabaseConnection
}

func (m *PersonalInfoManager) SetUp(db *database.DatabaseConnection) {
	m.db = db
}

func (m PersonalInfoManager) UpsertProfile(user *models.CurrentUserProfile) {
	res := database.GetUserProfile(m.db, user.ID)
	if res == nil {
		database.InsertUserProfile(m.db, user)
	}
}

func (m PersonalInfoManager) GetUserProfile(userId string) *models.CurrentUserProfile {
	res := database.GetUserProfile(m.db, userId)
	if res == nil {
		return &models.CurrentUserProfile{}
	}
	return res
}
