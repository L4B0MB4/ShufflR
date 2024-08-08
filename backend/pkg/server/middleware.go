package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type SessionMiddleware struct {
	sessionStore SessionStore
}

func (m *SessionMiddleware) SetUp(s SessionStore, router gin.IRouter) {
	m.sessionStore = s
	router.Use(m.UseSession)
}

func (m *SessionMiddleware) UseSession(ctx *gin.Context) {
	val, err := ctx.Cookie("session")
	if err != nil {
		log.Error().Err(err).Msg("Error during session recovery")
		ctx.Next()
		return
	}
	session, ok := m.sessionStore.GetSession(val)
	if ok {
		ctx.Set("session", session)
	}
	ctx.Next()

}
