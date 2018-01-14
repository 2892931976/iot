package utils

import "github.com/gin-gonic/gin"

func JsonResponseError(c *gin.Context, err interface{}) {
	switch v := err.(type) {
	case string:
		c.JSON(200, gin.H{"msg": v, "code": 0})
	case error:
		c.JSON(200, gin.H{"msg": v.Error(), "code": 0})
	}
}

func JsonResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{"msg": "success", "data": data, "code": 1})
}
