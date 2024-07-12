package router

import (
	"bytes"
	"fmt"
	"io"

	"github.com/thoughtgears/dota2-tracker/internal/dota"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/thoughtgears/dota2-tracker/views"
)

func NewRouter(client *dota.Client, debug bool) (*gin.Engine, error) {
	router := gin.New()
	router.Use(gin.Recovery(), logger.SetLogger(
		logger.WithLogger(func(_ *gin.Context, l zerolog.Logger) zerolog.Logger {
			return l.Output(gin.DefaultWriter).With().Logger()
		}),
	))

	if err := router.SetTrustedProxies(nil); err != nil {
		return nil, fmt.Errorf("failed to set trusted proxies: %w", err)
	}

	if debug {
		router.Use(debugLogger())
	}

	router.LoadHTMLGlob("templates/**/*.gohtml")
	router.GET("/", views.GetIndex(client))

	return router, nil
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
