package router

import (
	"github.com/gin-gonic/gin"
	"github.com/viniciusdsouza/goverallratings/handler"
)

func initializeRoutes(router *gin.Engine) {
	movie := router.Group("/api/v1/movies")
	{
		movie.GET("/:id", handler.GetMovieHandler)
		movie.POST("", handler.CreateMovieHandler)
		movie.PUT("/:id", handler.UpdateMovieHandler)
		movie.DELETE("/:id", handler.DeleteMovieHandler)
		movie.GET("", handler.ListMoviesHandler)
	}

	show := router.Group("/api/v1/shows")
	{
		show.GET("/:id", handler.GetShowHandler)
		show.POST("", handler.CreateShowHandler)
		show.PUT("/:id", handler.UpdateShowHandler)
		show.DELETE("/:id", handler.DeleteShowHandler)
		show.GET("", handler.ListShowsHandler)
	}
}
