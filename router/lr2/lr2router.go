package lr2

import (
	"main/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRouterForLR2(router *gin.RouterGroup) {
	router.GET("/get", controller.GeyQuerryController)

	router.POST("/post/body", controller.PostBodyController)

	router.POST("/post", controller.PostBodyQueryController)

	router.PUT("/put", controller.PutBodyQueryController)

	router.PATCH("/patch", controller.PatchBodyQueryController)
}
