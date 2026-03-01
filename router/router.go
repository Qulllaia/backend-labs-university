package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"main/router/lr2"
	"main/router/lr3"

	"main/router/lr4"
)

func RouterStart(router *gin.Engine, db *sqlx.DB) {
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
	}
}
