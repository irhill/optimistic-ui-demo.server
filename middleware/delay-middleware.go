package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
)

func DelayMiddleware(delay time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		time.Sleep(delay)
		c.Next()
	}
}
