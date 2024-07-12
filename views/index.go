package views

import (
	"net/http"

	"github.com/thoughtgears/dota2-tracker/internal/dota"

	"github.com/gin-gonic/gin"
)

func GetIndex(client *dota.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		proMatches := client.GetProMatches()
		ctx.HTML(http.StatusOK, "index.gohtml", gin.H{
			"title":   "Dota2 Tracker",
			"matches": proMatches,
		})
	}
}
