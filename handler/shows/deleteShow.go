package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteShowHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "DELETE Show",
	})
}