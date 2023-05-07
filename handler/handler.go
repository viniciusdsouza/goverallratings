package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMovieHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET Movie",
	})
}
func CreateMovieHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "CREATE Movie",
	})
}
func UpdateMovieHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "UPDATE Movie",
	})
}
func DeleteMovieHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "DELETE Movie",
	})
}
func ListMoviesHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "LIST Movies",
	})
}

func GetShowHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET Show",
	})
}
func CreateShowHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "CREATE Show",
	})
}
func UpdateShowHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "UPDATE Show",
	})
}
func DeleteShowHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "DELETE Show",
	})
}
func ListShowsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "LIST Shows",
	})
}