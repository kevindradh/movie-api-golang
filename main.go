package main

import (
	"MovieAPI/config"
	"MovieAPI/routes"

	_ "MovieAPI/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           MovieAPI
// @version         1.0
// @description     A simple API to manage movies use Go (Gin) & MongoDB (Test)

// @host      localhost:8080
// @BasePath  /api

func main() {
	config.ConnectDB()

	r := routes.SetupRouter()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
