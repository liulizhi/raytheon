package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
)

// RequestIDMiddleware add requestID
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get("X-Request-Id")
		if requestID == "" {
			requestID = ksuid.New().String()
		}
		// Expose it for use in the application
		c.Set("RequestID", requestID)
		// Set X-Request-Id header
		c.Writer.Header().Set("X-Request-Id", requestID)
		c.Next()
	}
}
