package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/turbo-iot/models"
	"github.com/mojocn/turbo-iot/utils"
	"math"
)

func DeviceAdd(c *gin.Context) {

	var input models.Device

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.JsonResponseError(c, err)
		return
	} else {
		uid := c.GetInt("uid")
		input.D_userid = uid
	}

	if err := models.DB.Create(&input).Error; err != nil {
		utils.JsonResponseError(c, err)
		return
	}
	utils.JsonResponseSuccess(c, input)
}

func DeviceIndex(c *gin.Context) {
	page := utils.GetQueryInt(c, "page", 1)
	pageSize := utils.GetQueryInt(c, "pageSize", 15)
	var list []models.Device
	var totalCount int
	db := models.DB.Model(&models.Device{}).Count(&totalCount)
	totalPage := int(math.Ceil(float64(totalCount) / float64(pageSize)))
	offset := pageSize * page
	db.Limit(pageSize).Offset(offset).Find(&list)
	utils.JsonResponseSuccess(c, gin.H{"totalPage": totalPage, "page": page, "pageSize": pageSize, "list": list,
		"totalCount": totalCount})
}

func DeviceUpdate(c *gin.Context) {
	var input models.Device
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.JsonResponseError(c, err)
		return
	} else {
		uid := c.GetInt("uid")
		input.D_userid = uid
	}
	if input.D_no == "" {
		utils.JsonResponseError(c, "d_no主键不能为空")
		return
	}
	db := models.DB.Model(&input).Updates(input)
	if err := db.Error; err != nil {
		utils.JsonResponseError(c, err)
		return
	}
	if db.RowsAffected == 0 {
		utils.JsonResponseError(c, "主键错误!")
		return
	}
	utils.JsonResponseSuccess(c, input)
}
func DeviceInfo(c *gin.Context) {
	dno := c.Param("dno")
	instance := models.Device{
		D_no: dno,
	}
	if err := models.DB.First(&instance).Error; err != nil {
		utils.JsonResponseError(c, err)
		return
	}
	utils.JsonResponseSuccess(c, instance)
}
func DeviceDelete(c *gin.Context) {
	var input models.Device
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.JsonResponseError(c, err)
		return
	}
	if input.D_no == "" {
		utils.JsonResponseError(c, "d_no主键不能为空")
		return
	}
	if err := models.DB.Delete(&input).Error; err != nil {
		utils.JsonResponseError(c, err)
	} else {
		utils.JsonResponseSuccess(c, nil)
	}
}
