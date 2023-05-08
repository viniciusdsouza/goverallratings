package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/viniciusdsouza/goverallratings/docs"
	moviesHandler "github.com/viniciusdsouza/goverallratings/handler/movies"
	showsHandler "github.com/viniciusdsouza/goverallratings/handler/shows"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize Handler
	moviesHandler.InitializeHandler()
	showsHandler.InitializeHandler()
	basePath := "/api/v1/"
	docs.SwaggerInfo.BasePath = basePath

	movie := router.Group(basePath + "/movies")
	{
		movie.GET("/:id", moviesHandler.GetMovieHandler)
		movie.POST("", moviesHandler.CreateMovieHandler)
		movie.PUT("/:id", moviesHandler.UpdateMovieHandler)
		movie.DELETE("/:id", moviesHandler.DeleteMovieHandler)
		movie.GET("", moviesHandler.ListMoviesHandler)
	}

	show := router.Group(basePath + "/shows")
	{
		show.GET("/:id", showsHandler.GetShowHandler)
		show.POST("", showsHandler.CreateShowHandler)
		show.PUT("/:id", showsHandler.UpdateShowHandler)
		show.DELETE("/:id", showsHandler.DeleteShowHandler)
		show.GET("", showsHandler.ListShowsHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
