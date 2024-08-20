package routes

import (
	"net/url"

	"github.com/L4B0MB4/Musicfriends/pkg/models"
	"github.com/L4B0MB4/Musicfriends/pkg/server/interfaces"
	"github.com/L4B0MB4/Musicfriends/pkg/server/manager"
	"github.com/L4B0MB4/Musicfriends/pkg/utils"
	"github.com/gin-gonic/gin"
)

type MeController struct {
	sessionStore interfaces.SessionStore
	manager      *manager.PersonalInfoManager
}

func (ctrl *MeController) SetUp(router gin.IRouter, sessionStore interfaces.SessionStore, m *manager.PersonalInfoManager) {
	ctrl.sessionStore = sessionStore
	ctrl.manager = m
	router.GET("/api/me", ctrl.getRoute)
	router.GET("/api/me/top/tracks", ctrl.getTopTracksRoute)
}

func (ctrl *MeController) getRoute(ctx *gin.Context) {
	userContext := utils.GetUserContextFromCtx(ctx, ctrl.sessionStore)
	profile := ctrl.manager.GetUserProfile(userContext.ID)
	ctx.JSON(200, profile)
}

func (ctrl *MeController) getTopTracksRoute(ctx *gin.Context) {
	userContext := utils.GetUserContextFromCtx(ctx, ctrl.sessionStore)
	query := url.Values{"time_range": []string{"short_term"}, "limit": []string{"50"}}
	topTracks := utils.SpotifyApiCall[models.TopTracksResponse]("/v1/me/top/tracks", userContext.AccessToken, "GET", query, nil)
	ctx.JSON(200, topTracks.Items)
}
