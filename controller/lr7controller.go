package controller

import (
	"main/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) (int , string) {
	return http.StatusOK, "pong"
}

func TraceHandler(c *gin.Context) (int , string){
	traceId, exists := c.Get(middleware.TraceIdKey)
	if !exists {
			return http.StatusInternalServerError, "TraceId not found"
	}
	
	return http.StatusOK, traceId.(string)
}

func ErrorHandler(c *gin.Context) {
	panic("test error occurred")
}
