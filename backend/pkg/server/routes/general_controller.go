package routes

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/L4B0MB4/Musicfriends/pkg/models"
	"github.com/L4B0MB4/Musicfriends/pkg/server"
	"github.com/L4B0MB4/Musicfriends/pkg/server/config"
	"github.com/L4B0MB4/Musicfriends/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type GeneralController struct {
	sessionStore server.SessionStore
	config       *config.Configuration
	redirect_uri string
}

func (g *GeneralController) SetUp(sessionStore server.SessionStore, router gin.IRouter, c *config.Configuration) {
	g.sessionStore = sessionStore
	g.config = c
	g.redirect_uri = "http://" + g.config.Host + ":" + g.config.Port + "/callback"
	router.GET("/", g.defaultRoute)
	router.GET("/login", g.loginRoute)
	router.GET("/callback", g.callbackRoute)
}

func (g *GeneralController) defaultRoute(ctx *gin.Context) {
	re, _ := ctx.Get("session")
	session, ok := re.(models.CurrentUserProfile)
	if ok {
		ctx.JSON(200, session)
	} else {
		ctx.Writer.Write([]byte("helloooo"))
	}
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
	req := http.Request{}
	spotifyUrl, _ := url.Parse("https://accounts.spotify.com/api/token")
	req.URL = spotifyUrl
	form := url.Values{}
	form.Add("code", code)
	form.Add("redirect_uri", g.redirect_uri)
	form.Add("grant_type", "authorization_code")
	req.Body = io.NopCloser(bytes.NewBufferString(form.Encode()))
	req.Header = http.Header{}
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(g.config.ClientId+":"+g.config.ClientSecret)))
	req.Method = "POST"
	client := &http.Client{}
	res, err := client.Do(&req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to do request to spotify")
		return
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error().Err(err).Msg("Failed to do read spotify response body")
		return
	}

	tokenRes := models.TokenResponse{}
	json.Unmarshal(b, &tokenRes)
	log.Info().Interface("o", tokenRes).Msg("boddyyy")
	profile := RequestMe(tokenRes)
	if profile != nil {
		sessionKey := g.sessionStore.AddSession(*profile)
		ctx.SetCookie("session", sessionKey, 3600, "/", g.config.Host, true, true)
		ctx.Writer.Write([]byte("<html><body><script>window.location.href='/'</script></body></html>"))
	}
}

func RequestMe(tokenRes models.TokenResponse) *models.CurrentUserProfile {
	req := http.Request{}
	spotifyUrl, _ := url.Parse("https://api.spotify.com/v1/me")
	req.URL = spotifyUrl
	req.Header = http.Header{}
	req.Header.Add("Authorization", "Bearer "+tokenRes.AccessToken)
	req.Method = "GET"
	client := &http.Client{}

	res, err := client.Do(&req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to do request to spotify")
		return nil
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error().Err(err).Msg("Failed to do read spotify response body")
		return nil
	}
	profile := models.CurrentUserProfile{}
	json.Unmarshal(b, &profile)
	log.Info().Str("b", string(b)).Msg("boddyyy")
	return &profile
}
