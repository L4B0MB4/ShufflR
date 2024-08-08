package server

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type SessionMiddleware struct {
	sessionStore SessionStore
}

func (m *SessionMiddleware) SetUp(router gin.IRouter, s SessionStore) {
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
	if strings.Contains(ctx.Request.URL.String(), "/api/") && !ok {
		ctx.Redirect(301, "/forbidden")
		return
	}
	ctx.Next()

}
