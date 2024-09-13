package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ForceFailureMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_ = c.AbortWithError(http.StatusTeapot, errors.New("error: no coffee for you"))
		return
	}
}
