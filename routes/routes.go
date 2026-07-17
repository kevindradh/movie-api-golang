package routes

import (
	"MovieAPI/handlers"
	"MovieAPI/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", handlers.Login)
			auth.POST("/register", handlers.Register)
		}

		movies := api.Group("/movies")
		{
			movies.GET("", handlers.GetMovies)
			movies.GET("/:id", handlers.GetMovieByID)

			movies.POST("", middleware.AuthRequired(), handlers.CreateMovie)
			movies.PUT("/:id", middleware.AuthRequired(), handlers.UpdateMovie)
			movies.DELETE("/:id", middleware.AuthRequired(), middleware.AdminOnly(), handlers.DeleteMovie)
			movies.POST("/:id/poster", middleware.AuthRequired(), handlers.UploadPoster)
		}
	}

	return r
}
