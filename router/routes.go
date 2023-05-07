package router

import (
	"github.com/gin-gonic/gin"
	moviesHandler "github.com/viniciusdsouza/goverallratings/handler/movies"
	showsHandler "github.com/viniciusdsouza/goverallratings/handler/shows"
)

func initializeRoutes(router *gin.Engine) {
	movie := router.Group("/api/v1/movies")
	{
		movie.GET("/:id", moviesHandler.GetMovieHandler)
		movie.POST("", moviesHandler.CreateMovieHandler)
		movie.PUT("/:id", moviesHandler.UpdateMovieHandler)
		movie.DELETE("/:id", moviesHandler.DeleteMovieHandler)
		movie.GET("", moviesHandler.ListMoviesHandler)
	}

	show := router.Group("/api/v1/shows")
	{
		show.GET("/:id", showsHandler.GetShowHandler)
		show.POST("", showsHandler.CreateShowHandler)
		show.PUT("/:id", showsHandler.UpdateShowHandler)
		show.DELETE("/:id", showsHandler.DeleteShowHandler)
		show.GET("", showsHandler.ListShowsHandler)
	}
}
