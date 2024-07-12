package main

import (
	"fmt"
	"time"

	"github.com/thoughtgears/dota2-tracker/internal/dota"

	"github.com/thoughtgears/dota2-tracker/internal/router"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Config struct to hold the configuration for the application
// Port: port to run the application
// Debug: debug mode to run the application
// host: host to run the application, sets to 127.0.0.1 for local mode, and 0.0.0.0 if local is not set
type Config struct {
	Port  string `envconfig:"PORT" default:"8080"`
	host  string
	Local bool `envconfig:"LOCAL" default:"false"`
	Debug bool `envconfig:"DEBUG" default:"false"`
}

var config Config

// init function to set default config and global config for the application
func init() {
	envconfig.MustProcess("", &config)
	gin.SetMode(gin.ReleaseMode)
	config.host = "0.0.0.0"

	if config.Local {
		gin.SetMode(gin.DebugMode)
		config.host = "127.0.0.1"
	}

	zerolog.LevelFieldName = "severity"
	zerolog.TimestampFieldName = "timestamp"
	zerolog.TimeFieldFormat = time.RFC3339Nano
}

func main() {
	client := dota.NewClient()
	r, err := router.NewRouter(client, config.Debug)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create router")
	}

	log.Info().Msg("server started")
	log.Fatal().Err(r.Run(fmt.Sprintf("%s:%s", config.host, config.Port))).Msg("server stopped")
}
