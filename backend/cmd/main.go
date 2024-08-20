package main

import (
	"net/http"
	"os"

	"github.com/L4B0MB4/Musicfriends/pkg/database"
	"github.com/L4B0MB4/Musicfriends/pkg/server"
	"github.com/L4B0MB4/Musicfriends/pkg/server/config"
	"github.com/L4B0MB4/Musicfriends/pkg/server/manager"
	"github.com/L4B0MB4/Musicfriends/pkg/server/routes"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	server := setup()
	if server != nil {
		server.Start()
	}
}
func setup() *server.Server {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	router := gin.Default()
	router.StaticFile("/", "./static/index.html")
	router.StaticFS("/static", http.Dir("./static"))
	c := config.Configuration{}
	sMw := server.SessionMiddleware{}
	sessionStore := server.InMemorySessionStore{}
	gC := routes.GeneralController{}
	mC := routes.MeController{}
	dbConn := database.DatabaseConnection{}
	mng := manager.PersonalInfoManager{}
	c.SetUp()
	sMw.SetUp(router, &sessionStore)
	dbConn.SetUp()
	mng.SetUp(&dbConn)
	gC.SetUp(router, &sessionStore, &c, &mng)
	mC.SetUp(router, &sessionStore, &mng)
	_, err := dbConn.GetDbConnection()
	if err != nil {
		log.Error().Err(err).Msg("Error setting up db connection. Stoping...")
		return nil
	}
	server := server.Server{}
	server.SetUp(router, &c)
	return &server

}
