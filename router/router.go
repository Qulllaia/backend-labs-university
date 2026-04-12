package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"

	"main/database/repos"
	"main/middleware"
	"main/router/lr11"
	"main/router/lr2"
	"main/router/lr3"
	"main/router/lr6"
	"main/router/lr7"
	"main/router/lr8"

	"main/router/lr4"

	"main/router/lr5"
)

func RouterStart(router *gin.Engine, db *sqlx.DB, gormDB *gorm.DB) {
	api := router.Group("/api")
	{

		lr2router := api.Group("/lr2")
		{
			lr2.RegisterRouterForLR2(lr2router)
		}

		lr3router := api.Group("/lr3")
		{
			lr3.RegisterRouterForLR3(lr3router)
		}

		lr4router := api.Group("/lr4")
		{
			lr4.RegisterRouterForLR4(lr4router, db)
		}

		orderRouter := api.Group("/lr5")
		{
			lr5.RegisterRouterForLR5(orderRouter, db)
		}
		orderRouterlr6 := api.Group("/lr6")
		{
			lr6.RegisterRouterForLR6(orderRouterlr6, db)
		}
		lr7router := api.Group("/lr7")
		{
			lr7router.Use(middleware.BlockPathMiddleware())

			lr7router.Use(middleware.RequestTraceMiddleware())

			lr7.RegisterRouterForLR7(lr7router)
		}

		lr8router := api.Group("/lr8")
		{
			lr8.RegisterRouterForLR8(lr8router)
		}

		lr11router := api.Group("/lr11")
		{
			ur := &repos.UserRepo{DB: gormDB}
			br := &repos.BalanceRepo{DB: gormDB}
			tr := &repos.TransactionRepo{DB: gormDB}

			lr11.RegisterRouterForLR11(lr11router, ur, br, tr)
		}
	}
}
