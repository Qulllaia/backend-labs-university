package main

import (
	"main/config"
	"main/database"
	"main/router"

	_ "main/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example API
// @host      localhost:5000
// @BasePath  /api
func main() {
	config := config.InitConfig()
	db := database.InitDatabase(config)

	app := gin.Default()
	router.RouterStart(app, db)
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := app.Run(":5000")
	if err != nil {
		panic(err.Error())
	}
}
