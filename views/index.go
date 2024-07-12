package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jasonodonnell/go-opendota"
	"github.com/rs/zerolog/log"
)

func GetIndex(client *opendota.Client) gin.HandlerFunc {
	proMatches, _, err := client.ProMatchService.Matches()
	if err != nil {
		log.Error().Err(err).Msg("failed to get pro matches")
	}
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.gohtml", gin.H{
			"title":   "Dota2 Tracker",
			"matches": proMatches,
		})
	}
}
