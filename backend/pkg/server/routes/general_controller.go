package routes

import (
	"net/url"

	"github.com/L4B0MB4/Musicfriends/pkg/models"
	"github.com/L4B0MB4/Musicfriends/pkg/server/config"
	"github.com/L4B0MB4/Musicfriends/pkg/server/interfaces"
	"github.com/L4B0MB4/Musicfriends/pkg/server/manager"
	"github.com/L4B0MB4/Musicfriends/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type GeneralController struct {
	sessionStore interfaces.SessionStore
	config       *config.Configuration
	manager      *manager.Manager
	redirect_uri string
}

func (g *GeneralController) SetUp(router gin.IRouter, sessionStore interfaces.SessionStore, c *config.Configuration, m *manager.Manager) {

	//todo: thread safety could be an issue as all of these are singletons - ignoring for now. Won't cause any problems im sure :P
	g.sessionStore = sessionStore
	g.config = c
	g.redirect_uri = "http://" + g.config.Host + ":" + g.config.Port + "/callback"
	g.manager = m
	router.GET("/login", g.loginRoute)
	router.GET("/callback", g.callbackRoute)
	router.GET("/forbidden", g.forbiddenRoute)
}

func (g *GeneralController) forbiddenRoute(ctx *gin.Context) {
	ctx.Writer.Write([]byte("<b>forbidden</b>"))
}

func (g *GeneralController) loginRoute(ctx *gin.Context) {

	q := url.Values{}
	scope := "user-read-private user-read-email"

	state := utils.RandomString(16)
	q.Set("response_type", "code")
	q.Set("client_id", g.config.ClientId)
	q.Set("scope", scope)
	q.Set("redirect_uri", g.redirect_uri)
	q.Set("state", state)

	ctx.Redirect(301, "https://accounts.spotify.com/authorize?"+q.Encode())

}

func (g *GeneralController) callbackRoute(ctx *gin.Context) {
	q := ctx.Request.URL.Query()
	log.Info().Interface("queryobj", q).Str("querystr", q.Encode()).Msg("Return values")
	code := q.Get("code")

	tokenRes := utils.GetAccessToken(code, g.redirect_uri, g.config.ClientId, g.config.ClientSecret)
	if tokenRes == nil {
		ctx.Redirect(301, "/forbidden")
		return
	}

	profile := utils.SpotifyApiCall[models.CurrentUserProfile]("/v1/me", tokenRes.AccessToken, "GET", nil)
	g.manager.UpsertProfile(profile)
	if profile != nil {
		userAuth := models.UserContext{
			ID:           profile.ID,
			AccessToken:  tokenRes.AccessToken,
			RefreshToken: tokenRes.RefreshToken,
		}
		sessionKey := g.sessionStore.AddSession(userAuth)
		ctx.SetCookie("session", sessionKey, 3600, "/", g.config.Host, true, true)
		ctx.Writer.Write([]byte("<html><body><script>window.location.href='/api/me'</script></body></html>"))
	}
}
