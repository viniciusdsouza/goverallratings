package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusdsouza/goverallratings/schemas"
)

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
		Title: cmd.Title,
		Duration: cmd.Duration,
		Rating: cmd.Rating,
		Genre: cmd.Genre,
	}

	err = db.Create(&movie).Error
	if err != nil{
		logger.Errorf("error creating movie: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating movie")
		return
	}

	sendSuccess(ctx, "Create Movie", movie)
}
