package main

import (
	"lambda/controllers"
	"lambda/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRuter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.RequestID(6))
	r.GET("/", controllers.HelloController)
	if os.Getenv("ENVIRONMENT") != "prod" {
		r.GET("/api-docs", func(c *gin.Context) {
			c.File("./docs/index.html")
		})
	}
	return r
}
