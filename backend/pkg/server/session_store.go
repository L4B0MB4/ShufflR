package server

import (
	"github.com/L4B0MB4/Musicfriends/pkg/models"
	"github.com/L4B0MB4/Musicfriends/pkg/utils"
)

type SessionStore interface {
	AddSession(models.CurrentUserProfile) string
	GetSession(string) (models.CurrentUserProfile, bool)
	RemoveSession(string)
}

type InMemorySessionStore struct {
	sessions map[string]models.CurrentUserProfile
}

func (s *InMemorySessionStore) AddSession(m models.CurrentUserProfile) string {
	if s.sessions == nil {
		s.sessions = map[string]models.CurrentUserProfile{}
	}
	sessionKey := utils.RandomString(16)
	s.sessions[sessionKey] = m
	return sessionKey

}
func (s *InMemorySessionStore) GetSession(sessionKey string) (models.CurrentUserProfile, bool) {
	session, ok := s.sessions[sessionKey]
	return session, ok
}
func (s *InMemorySessionStore) RemoveSession(sessionKey string) {
	delete(s.sessions, sessionKey)
}
