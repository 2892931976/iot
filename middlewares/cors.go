package middlewares

import (
	"github.com/gin-gonic/gin"
)

func EnableCors() gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.GetHeader("Request Method") == "OPTIONS" {

		} else {
			c.Next()
		}
	}
}
