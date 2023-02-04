package main

import (
	"lambda/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRuter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.RequestID(6))
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "lambda"})
	})
	return r
}
