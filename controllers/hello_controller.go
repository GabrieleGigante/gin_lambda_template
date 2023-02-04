package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": "lambda"})
}
