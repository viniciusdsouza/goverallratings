package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, content interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s successfull", op),
		"content":    content,
	})
}
