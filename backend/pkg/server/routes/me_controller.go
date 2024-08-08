package routes

import (
	"github.com/L4B0MB4/Musicfriends/pkg/server/interfaces"
	"github.com/L4B0MB4/Musicfriends/pkg/server/manager"
	"github.com/L4B0MB4/Musicfriends/pkg/utils"
	"github.com/gin-gonic/gin"
)

type MeController struct {
	sessionStore interfaces.SessionStore
	manager      *manager.Manager
}

func (ctrl *MeController) SetUp(router gin.IRouter, sessionStore interfaces.SessionStore, m *manager.Manager) {
	ctrl.sessionStore = sessionStore
	ctrl.manager = m
	router.GET("/api/me", ctrl.getRoute)
}

func (ctrl *MeController) getRoute(ctx *gin.Context) {
	userContext := utils.GetUserContextFromCtx(ctx, ctrl.sessionStore)
	profile := ctrl.manager.GetUserProfile(userContext.ID)
	ctx.JSON(200, profile)
}
