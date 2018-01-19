package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/turbo-iot/models"
	"github.com/mojocn/turbo-iot/utils"
)

type cmdJson struct {
	D_no    string
	Payload string
}

const redisDeviceCmdKey = "TBIoTSTDownlinkPrepareMsgList"

func CommandAdd(c *gin.Context) {
	uid := c.GetInt("uid")
	var input cmdJson
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.JsonResponseError(c, err)
		return
	}

	instance := models.Device{}
	db := models.DB.Where("d_userid = ? and d_no = ?", uid, input.D_no).First(&instance)
	if db.RecordNotFound() {
		utils.JsonResponseError(c, "这个设备不属于您!")
		return
	}

	data := gin.H{"method": "disired", "devid": input.D_no, "state": gin.H{"payload": input.Payload}}
	valueByets, err := json.Marshal(data)
	if err != nil {
		utils.JsonResponseError(c, "json序列化错误!")
		return
	}

	if err := models.RedisClient.RPush(redisDeviceCmdKey, valueByets).Err(); err != nil {
		utils.JsonResponseError(c, err)
		return
	}
	utils.JsonResponseSuccess(c, "发送设备命令成功!")
}
