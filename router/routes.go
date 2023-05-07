package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	movie := router.Group("/api/v1/movies")
	{
		movie.GET("/:id", func(ctx *gin.Context){
			ctx.JSON(http.StatusOK, gin.H{
				"message": "GET Movie",
			})
		})
		movie.POST("", func(ctx *gin.Context){
			ctx.JSON(http.StatusOK, gin.H{
				"message": "CREATE Movie",
			})
		})
		movie.PUT("/:id", func(ctx *gin.Context){
			ctx.JSON(http.StatusOK, gin.H{
				"message": "UPDATE Movie",
			})
		})
		movie.DELETE("/:id", func(ctx *gin.Context){
			ctx.JSON(http.StatusOK, gin.H{
				"message": "DELETE Movie",
			})
		})
		movie.GET("", func(ctx *gin.Context){
			ctx.JSON(http.StatusOK, gin.H{
				"message": "LIST Movies",
			})
		})
	}

	show := router.Group("/api/v1/shows")
	{
		show.GET("/:id", func(ctx *gin.Context){
			ctx.JSON(http.StatusOK, gin.H{
				"message": "GET Show",
			})
		})
		show.POST("", func(ctx *gin.Context){
			ctx.JSON(http.StatusOK, gin.H{
				"message": "CREATE Show",
			})
		})
		show.PUT("/:id", func(ctx *gin.Context){
			ctx.JSON(http.StatusOK, gin.H{
				"message": "UPDATE Show",
			})
		})
		show.DELETE("/:id", func(ctx *gin.Context){
			ctx.JSON(http.StatusOK, gin.H{
				"message": "DELETE Show",
			})
		})
		show.GET("", func(ctx *gin.Context){
			ctx.JSON(http.StatusOK, gin.H{
				"message": "LIST Shows",
			})
		})
	}
}
