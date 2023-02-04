package main

import (
	"lambda/controllers"
	"lambda/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRuter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.RequestID(6))
	r.GET("/", controllers.HelloController)
	return r
}
