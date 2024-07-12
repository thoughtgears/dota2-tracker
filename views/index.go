package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetIndex(ctx *gin.Context) {
	proMatches, err := s.dotaClient.GetProMatches()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.HTML(http.StatusOK, "index.gohtml", gin.H{
		"title":   "Dota2 Tracker",
		"matches": proMatches,
	})
}
