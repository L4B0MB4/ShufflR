package routes

import (
	"github.com/L4B0MB4/Musicfriends/pkg/server"
	"github.com/L4B0MB4/Musicfriends/pkg/server/config"
	"github.com/L4B0MB4/Musicfriends/pkg/server/manager"
	"github.com/gin-gonic/gin"
)

type MeController struct {
	sessionStore server.SessionStore
	config       *config.Configuration
	manager      *manager.Manager
}

func (ctrl *MeController) SetUp(router gin.IRouter, sessionStore server.SessionStore, c *config.Configuration, m *manager.Manager) {
	ctrl.sessionStore = sessionStore
	ctrl.manager = m
}
