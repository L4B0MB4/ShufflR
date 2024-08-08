package manager

import (
	"github.com/L4B0MB4/Musicfriends/pkg/database"
	"github.com/L4B0MB4/Musicfriends/pkg/models"
)

type Manager struct {
	db *database.DatabaseConnection
}

func (m *Manager) SetUp(db *database.DatabaseConnection) {
	m.db = db
}

func (m Manager) UpsertProfile(user *models.CurrentUserProfile) {
	res := database.GetUserProfile(m.db, user.ID)
	if res == nil {
		database.InsertUserProfile(m.db, user)
	}
}
