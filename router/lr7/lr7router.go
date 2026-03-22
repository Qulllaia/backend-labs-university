package lr7

import (
	"main/controller"
	"main/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRouterForLR7(router *gin.RouterGroup) {
	router.GET("/ping", middleware.EndpointTimingMiddleware(controller.PingHandler))
	router.GET("/blocked/ping", middleware.EndpointTimingMiddleware(controller.PingHandler))
	router.GET("/trace", middleware.EndpointTimingMiddleware(controller.TraceHandler))
	router.GET("/error", controller.ErrorHandler)

}
