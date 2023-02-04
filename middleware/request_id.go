package middleware

import (
	"math/rand"

	"github.com/gin-gonic/gin"
)

const dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890-?=!+"

func RequestID(n int) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := make([]byte, n)
		for i := 0; i < n; i++ {
			id[i] = dictionary[rand.Intn(len(dictionary))]
		}
		c.Header("X-Request-ID", string(id))
	}
}
