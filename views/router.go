package views

import (
	"bytes"
	"fmt"
	"io"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/jasonodonnell/go-opendota"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewRouter(client *opendota.Client, debug bool) (*gin.Engine, error) {
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
	router.GET("/", GetIndex(client))

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
