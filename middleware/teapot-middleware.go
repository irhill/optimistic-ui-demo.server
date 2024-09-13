package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TeapotMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_ = c.AbortWithError(http.StatusTeapot, errors.New("error: no coffee for you"))
		return
	}
}
