package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const TraceIdKey = "TraceId"

func RequestTraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := uuid.New().String()
		
		c.Set(TraceIdKey, traceId)
		
		c.Header("X-Trace-Id", traceId)
		
		c.Next()
	}
}
