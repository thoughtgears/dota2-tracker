package views

import (
	"bytes"
	"fmt"
	"io"

	"github.com/thoughtgears/dota2-tracker/internal/dota"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Server struct {
	Router     *gin.Engine
	dotaClient *dota.Client
}

func NewServer(client *dota.Client, debug bool) (*Server, error) {
	server := &Server{
		Router:     gin.New(),
		dotaClient: client,
	}

	server.Router.Use(gin.Recovery(), logger.SetLogger(
		logger.WithLogger(func(_ *gin.Context, l zerolog.Logger) zerolog.Logger {
			return l.Output(gin.DefaultWriter).With().Logger()
		}),
	))

	if err := server.Router.SetTrustedProxies(nil); err != nil {
		return nil, fmt.Errorf("failed to set trusted proxies: %w", err)
	}

	if debug {
		server.Router.Use(debugLogger())
	}

	server.Router.LoadHTMLGlob("templates/**/*.gohtml")
	server.Router.GET("/", server.GetIndex)

	return server, nil
}

func debugLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyBytes, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		c.Next()
		log.Info().
			Int("status", c.Writer.Status()).
			RawJSON("request_body", bodyBytes).
			Msg("incoming request")
	}
}
