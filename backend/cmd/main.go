package main

import (
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
	c := config.Configuration{}
	sMw := server.SessionMiddleware{}
	sessionStore := server.InMemorySessionStore{}
	gC := routes.GeneralController{}
	dbConn := database.DatabaseConnection{}
	mng := manager.Manager{}
	c.SetUp()
	sMw.SetUp(router, &sessionStore)
	dbConn.SetUp()
	mng.SetUp(&dbConn)
	gC.SetUp(router, &sessionStore, &c, &mng)
	_, err := dbConn.GetDbConnection()
	if err != nil {
		log.Error().Err(err).Msg("Error setting up db connection. Stoping...")
		return nil
	}
	server := server.Server{}
	server.SetUp(router, &c)
	return &server

}
