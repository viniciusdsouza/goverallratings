package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusdsouza/goverallratings/schemas"
)

func DeleteMovieHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParam").Error())
	}
	
	movie := schemas.Movie{}

	if err := db.First(&movie, id).Error; err != nil{
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("movie with id: %s not found", id))
		return
	}

	if err := db.Delete(&movie).Error; err != nil{
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting movie with id: %s", id))
		return
	}

	sendSuccess(ctx, "Delete Movie", movie)

}
