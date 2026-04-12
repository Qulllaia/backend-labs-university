package lr8

import (
	"main/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRouterForLR8(router *gin.RouterGroup) {
	router.GET("/html", controller.HTMLStatic)
	router.GET("/css", controller.CSSStatic)
	router.GET("/js", controller.JSStatic)
	router.GET("/image", controller.ImageStatic)
}
