package middleware

import (
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

func BlockPathMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		slicedPath := strings.Split(c.Request.URL.Path, "/");
		if slices.Contains(slicedPath, "blocked") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}
