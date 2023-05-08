package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusdsouza/goverallratings/schemas"
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
		"content": content,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateBookResponse struct {
	Message string               `json:"message"`
	Data    schemas.BookResponse `json:"data"`
}

type GetBookResponse struct {
	Message string               `json:"message"`
	Data    schemas.BookResponse `json:"data"`
}

type UpdateBookResponse struct {
	Message string               `json:"message"`
	Data    schemas.BookResponse `json:"data"`
}

type DeleteBookResponse struct {
	Message string               `json:"message"`
	Data    schemas.BookResponse `json:"data"`
}

type ListBookResponse struct {
	Message string                 `json:"message"`
	Data    []schemas.BookResponse `json:"data"`
}

type CreateShowResponse struct {
	Message string               `json:"message"`
	Data    schemas.ShowResponse `json:"data"`
}

type GetShowResponse struct {
	Message string               `json:"message"`
	Data    schemas.ShowResponse `json:"data"`
}

type UpdateShowResponse struct {
	Message string               `json:"message"`
	Data    schemas.ShowResponse `json:"data"`
}

type DeleteShowResponse struct {
	Message string               `json:"message"`
	Data    schemas.ShowResponse `json:"data"`
}

type ListShowResponse struct {
	Message string                 `json:"message"`
	Data    []schemas.ShowResponse `json:"data"`
}
