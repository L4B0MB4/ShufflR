package interfaces

import "github.com/L4B0MB4/Musicfriends/pkg/models"

type SessionStore interface {
	AddSession(models.UserContext) string
	HasSession(string) bool
	GetSession(string) (models.UserContext, bool)
	RemoveSession(string)
}
