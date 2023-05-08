package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusdsouza/goverallratings/schemas"
)

// @BasePath /api/v1
// @Summary List movies
// @Tags Movies
// @Accept json
// @Produce json
// @Success 200 {object} ListMovieResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /movies [get]
func ListMoviesHandler(ctx *gin.Context) {
	movies := []schemas.Movie{}

	if err := db.Find(&movies).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing movies")
		return
	}

	sendSuccess(ctx, "Get List Movies", movies)
}
