package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetQueryInt(c *gin.Context, key string, defaultValue int) int {
	pageSize, ok := c.GetQuery(key)
	if !ok {
		return defaultValue
	}
	i, err := strconv.Atoi(pageSize)
	if err == nil {
		return i
	} else {
		return defaultValue
	}

}
