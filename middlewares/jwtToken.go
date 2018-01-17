package middlewares

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/turbo-iot/models"
	. "github.com/mojocn/turbo-iot/utils"
	"strings"
)

const tokenPrefix = "Bearer "

func JwtTokenCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.Replace(c.GetHeader("Authorization"), tokenPrefix, "", 1)
		claims, err := JwtParseToken(token)
		if err != nil {
			JsonResponseError(c, "Authorization请求头Bearer错误:jwt-token 错误")
			return
		}
		if err := claims.Valid(); err != nil {
			JsonResponseError(c, err)
			return
		}
		//更具jwt的id,读取redis的用户信息
		buffer, err := models.RedisClient.Get(claims.Id).Bytes()
		if err != nil {
			JsonResponseError(c, err)
			return
		}
		var authedUser models.User
		err = json.Unmarshal(buffer, &authedUser)
		if err != nil {
			JsonResponseError(c, err)
			return
		} else {
			c.Set("uid", authedUser.Uid)
			c.Set("user", authedUser)
		}
		c.Next()
		// after request
	}
}
