package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusdsouza/goverallratings/schemas"
)

func GetMovieHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParam").Error())
	}

	movie := schemas.Movie{}
	if err := db.First(&movie, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("movie with id: %s not found", id))
		return
	}

	sendSuccess(ctx, "Get Movie", movie)
}
