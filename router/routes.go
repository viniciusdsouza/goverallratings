package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/viniciusdsouza/goverallratings/docs"
	"github.com/viniciusdsouza/goverallratings/handler"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize Handler
	handler.InitializeHandler()
	basePath := "/api/v1/"
	docs.SwaggerInfo.BasePath = basePath

	movie := router.Group(basePath + "/movies")
	{
		movie.GET("/:id", handler.GetMovieHandler)
		movie.POST("", handler.CreateMovieHandler)
		movie.PUT("/:id", handler.UpdateMovieHandler)
		movie.DELETE("/:id", handler.DeleteMovieHandler)
		movie.GET("", handler.ListMoviesHandler)
	}

	show := router.Group(basePath + "/shows")
	{
		show.GET("/:id", handler.GetShowHandler)
		show.POST("", handler.CreateShowHandler)
		show.PUT("/:id", handler.UpdateShowHandler)
		show.DELETE("/:id", handler.DeleteShowHandler)
		show.GET("", handler.ListShowsHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
