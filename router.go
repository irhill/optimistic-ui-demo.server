package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"optimistic-ui-demo/middleware"
	"optimistic-ui-demo/users"
	"time"
)

func SetupRouter(pgxMiddleware *middleware.PgxMiddleware, teapot bool) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	router.Use(middleware.DelayMiddleware(2 * time.Second))
	if teapot {
		// this middleware will fail any request sent to the router
		router.Use(middleware.TeapotMiddleware())
	}
	router.Use(pgxMiddleware.Middleware())

	initialiseRoutes(router)

	return router
}

func initialiseRoutes(router *gin.Engine) {
	users.SetupRoutes(router)
}
