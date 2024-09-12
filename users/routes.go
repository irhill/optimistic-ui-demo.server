package users

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	router.GET("/users", GetUsers)
	router.GET("/users/:id", GetUser)

	router.POST("/users", PostUser)

	router.PUT("/users/:id", PutUser)

	router.DELETE("/users/:id", DeleteUser)
}
