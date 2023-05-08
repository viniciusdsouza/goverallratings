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

type CreateMovieResponse struct {
	Message string                `json:"message"`
	Data    schemas.MovieResponse `json:"data"`
}

type GetMovieResponse struct {
	Message string                `json:"message"`
	Data    schemas.MovieResponse `json:"data"`
}

type UpdateMovieResponse struct {
	Message string                `json:"message"`
	Data    schemas.MovieResponse `json:"data"`
}

type DeleteMovieResponse struct {
	Message string                `json:"message"`
	Data    schemas.MovieResponse `json:"data"`
}

type ListMovieResponse struct {
	Message string                  `json:"message"`
	Data    []schemas.MovieResponse `json:"data"`
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
