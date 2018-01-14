package middlewares

import (
	"github.com/gin-gonic/gin"
	"strings"
)

const tokenPrefix = "Bearer "

func JwtTokenCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		//        'Authorization' => 'Bearer '.$accessToken,
		//'Accept' => 'application/json',
		//'Authorization' => 'Bearer '.$accessToken,
		token := strings.Replace(c.GetHeader("Authorization"), tokenPrefix, "", 1)
		claims, err := JwtParseToken(token)
		if err != nil {
			JsonResponseError(c, err)
			return
		}
		if err := claims.Valid(); err != nil {
			JsonResponseError(c, err)
			return
		}

		c.Next()

		// after request

	}
}
