package routes

import (
	"MovieAPI/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/movies", handlers.GetMovies)
		api.GET("/movies/:id", handlers.GetMovieByID)
		api.POST("/movies", handlers.CreateMovie)
		api.PUT("/movies/:id", handlers.UpdateMovie)
		api.DELETE("/movies/:id", handlers.DeleteMovie)
		api.POST("/movies/:id/poster", handlers.UploadPoster)
	}

	return r
}
