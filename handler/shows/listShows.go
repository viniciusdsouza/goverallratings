package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListShowsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "LIST Shows",
	})
}
