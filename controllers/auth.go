package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/turbo-iot/models"
	. "github.com/mojocn/turbo-iot/utils"

	"github.com/gin-gonic/gin/json"
	"log"
	"time"
)

func AuthPost(c *gin.Context) {
	var input models.User
	err := c.ShouldBindJSON(&input)
	if err != nil {
		JsonResponseError(c, err)
		return
	}
	var dbUser models.User
	err = models.DB.Where("email = ?", input.Email).First(&dbUser).Error
	if err != nil {
		JsonResponseError(c, err)
		return
	}
	if input.Pwd != dbUser.Pwd {
		JsonResponseError(c, "密码错误")
		return
	}
	//创建token
	expiration := 24 * 365 * time.Hour
	redisKey := dbUser.RedisKey()
	token, err := JwtGenerateToken(redisKey, expiration)
	if err != nil {
		JsonResponseError(c, err)
		return
	}

	value, err := json.Marshal(dbUser)
	if err != nil {
		log.Print(err)
	}
	err = models.RedisClient.Set(redisKey, value, expiration).Err()
	if err != nil {
		JsonResponseError(c, err)
		return
	}
	expirationSecond := expiration / time.Second
	data := map[string]interface{}{"token": token, "expire": expirationSecond, "user": dbUser}
	JsonResponseSuccess(c, data)
}
