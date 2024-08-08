package server

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/L4B0MB4/Musicfriends/pkg/server/config"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Server struct {
	router     *gin.Engine
	config     *config.Configuration
	initialzed bool
}

func (s *Server) SetUp(r *gin.Engine, c *config.Configuration) {
	s.router = r
	s.initialzed = true
	s.config = c
}

func (s *Server) Start() {
	if !s.initialzed {
		log.Error().Msg("Server not setup properly. Shutting down...")
		return
	}

	signalCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	ctx, cancel := context.WithCancel(context.Background())

	srv := &http.Server{
		Addr:    s.config.Host + ":" + s.config.Port,
		Handler: s.router,
	}
	go func() {
		log.Info().Str("Address", srv.Addr).Msg("Starting server")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("error starting server")
		}
	}()
	defer cancel()
	<-signalCtx.Done()
	log.Info().Msg("Shutting down ...")
	cancel()
	srv.Shutdown(ctx)

}
