package server

import (
	"strings"

	"github.com/L4B0MB4/Musicfriends/pkg/server/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type SessionMiddleware struct {
	sessionStore interfaces.SessionStore
}

func (m *SessionMiddleware) SetUp(router gin.IRouter, s interfaces.SessionStore) {
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
	ok := m.sessionStore.HasSession(val)
	if ok {
		ctx.Set("session", val)
	}
	if strings.Contains(ctx.Request.URL.String(), "/api/") && !ok {
		ctx.Redirect(301, "/forbidden")
		return
	}
	ctx.Next()

}
