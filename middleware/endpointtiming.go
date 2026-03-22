package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func EndpointTimingMiddleware(handler func(c *gin.Context) (int, string)) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		status, body := handler(c);

		elapsed := time.Since(start)
		elapsedMs := elapsed.Milliseconds();

		c.Header("X-Endpoint-Elapsed-Ms", strconv.FormatInt(elapsedMs, 10))
		c.JSON(status, gin.H{
			"result": body,
		})
	}
}
