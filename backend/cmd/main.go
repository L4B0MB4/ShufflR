package main

import (
	"os"

	"github.com/L4B0MB4/Musicfriends/pkg/server"
	"github.com/L4B0MB4/Musicfriends/pkg/server/config"
	"github.com/L4B0MB4/Musicfriends/pkg/server/routes"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	server := setup()
	server.Start()
}
func setup() *server.Server {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	router := gin.Default()
	c := config.Configuration{}
	sMw := server.SessionMiddleware{}
	sessionStore := server.InMemorySessionStore{}
	gC := routes.GeneralController{}
	c.SetUp()
	sMw.SetUp(&sessionStore, router)
	gC.SetUp(&sessionStore, router, &c)

	server := server.Server{}
	server.SetUp(router, &c)
	return &server

}
