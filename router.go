package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"optimistic-ui-demo/users"
	"time"
)

func SetupRouter(pgxMiddleware *PgxMiddleware, failure bool) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	router.Use(DelayMiddleware(2 * time.Second))
	if failure {
		// this middleware will fail any request sent to the router
		router.Use(ForceFailureMiddleware())
	}
	router.Use(pgxMiddleware.Middleware())

	initialiseRoutes(router)

	return router
}

func initialiseRoutes(router *gin.Engine) {
	users.SetupRoutes(router)
}
