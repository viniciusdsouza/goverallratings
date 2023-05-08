package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusdsouza/goverallratings/schemas"
)

func UpdateMovieHandler(ctx *gin.Context) {
	cmd := UpdateMovieRequest{}
	
	ctx.BindJSON(&cmd)
	
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParam").Error())
	}

	err := cmd.Validate()

	if err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	movie := schemas.Movie{}
	
	if err := db.First(&movie, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("movie with id: %s not found", id))
		return
	}

	if movie.Title != "" {
		movie.Title = cmd.Title
	}

	if movie.Duration > 0 {
		movie.Duration = cmd.Duration
	}

	if movie.Rating > 0 {
		movie.Rating = cmd.Rating
	}

	if movie.Genre != "" {
		movie.Genre = cmd.Genre
	}

	if err := db.Save(&movie).Error; err != nil {
		logger.Errorf("error updating movie: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating movie")
		return
	}

	sendSuccess(ctx, "Update Movie", movie)
}
