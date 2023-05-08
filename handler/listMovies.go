package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusdsouza/goverallratings/schemas"
)

func ListMoviesHandler(ctx *gin.Context) {
	movies := []schemas.Movie{}

	if err := db.Find(&movies).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing movies")
		return
	}

	sendSuccess(ctx, "Get List Movies", movies)
}
