package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateMovieHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "UPDATE Movie",
	})
}
