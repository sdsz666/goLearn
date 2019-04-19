package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/users", Users)

	router.POST("/user", Store)

	router.PUT("/user/:id", Update)

	router.DELETE("/user/:id", Destroy)

	return router
}