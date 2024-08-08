package server

import (
	"github.com/L4B0MB4/Musicfriends/pkg/models"
	"github.com/L4B0MB4/Musicfriends/pkg/utils"
)

type InMemorySessionStore struct {
	sessions map[string]models.UserContext
}

func (s *InMemorySessionStore) AddSession(m models.UserContext) string {
	if s.sessions == nil {
		s.sessions = map[string]models.UserContext{}
	}
	sessionKey := utils.RandomString(16)
	s.sessions[sessionKey] = m
	return sessionKey

}
func (s *InMemorySessionStore) HasSession(sessionKey string) bool {
	_, ok := s.sessions[sessionKey]
	return ok
}
func (s *InMemorySessionStore) GetSession(sessionKey string) (models.UserContext, bool) {
	session, ok := s.sessions[sessionKey]
	return session, ok
}
func (s *InMemorySessionStore) RemoveSession(sessionKey string) {
	delete(s.sessions, sessionKey)
}
