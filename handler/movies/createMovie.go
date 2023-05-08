package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusdsouza/goverallratings/schemas"
)

// @BasePath /api/v1
// @Summary Create movie
// @Tags Movies
// @Accept json
// @Produce json
// @Param request body CreateMovieRequest true "Request Body"
// @Success 200 {object} CreateMovieResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /movies [post]
func CreateMovieHandler(ctx *gin.Context) {
	cmd := CreateMovieRequest{}

	ctx.BindJSON(&cmd)

	err := cmd.Validate()
	if err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	movie := schemas.Movie{
		Title:    cmd.Title,
		Duration: cmd.Duration,
		Rating:   cmd.Rating,
		Genre:    cmd.Genre,
	}

	err = db.Create(&movie).Error
	if err != nil {
		logger.Errorf("error creating movie: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating movie")
		return
	}

	sendSuccess(ctx, "Create Movie", movie)
}
