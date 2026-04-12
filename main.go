package main

import (
	"main/config"
	"main/database"
	"main/logger"
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

	gormDB := database.InitGORM_DB(config)

	logger.InitLogger("app.log")

	app := gin.Default()
	router.RouterStart(app, db, gormDB)

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	app.Static("/static", "./static")

	err := app.Run(":5000")
	if err != nil {
		panic(err.Error())
	}
}
