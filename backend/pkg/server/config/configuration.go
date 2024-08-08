package config

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Configuration struct {
	ClientId     string
	ClientSecret string
	Host         string
	Port         string
}

func (c *Configuration) SetUp() {
	envFile, _ := godotenv.Read(".env")
	var ok bool
	c.ClientId, ok = envFile["CLIENT_ID"]
	if !ok {
		log.Info().Msg("No value for CLIENT_ID found")
	}
	c.ClientSecret, ok = envFile["CLIENT_SECRET"]
	if !ok {
		log.Info().Msg("No value for CLIENT_SECRET found")
	}
	c.Host, ok = envFile["HOST"]
	if !ok {
		log.Info().Msg("No value for HOST found falling back to localhost")
		c.Host = "localhost"
	}
	c.Port, ok = envFile["PORT"]
	if !ok {
		log.Info().Msg("No value for HOST found falling back to 8080")
		c.Port = "8080"
	}
}
