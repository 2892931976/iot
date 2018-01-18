package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/turbo-iot/models"
	"github.com/mojocn/turbo-iot/utils"
	"math"
)

type jsonInput struct {
	models.Device
	models.Lorawan
}

func DeviceAdd(c *gin.Context) {
	var input jsonInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.JsonResponseError(c, err)
		return
	} else {
		uid := c.GetInt("uid")
		input.D_userid = uid
	}
	if models.IsProductNotValid(input.D_pno) {
		utils.JsonResponseError(c, "d_no为无效产品型号")
		return
	}
	//正常设备
	if models.IsNormalDevice(input.D_pno) {
		input.D_ptype = 1
	}
	//网关设备
	if models.IsGetwayDevice(input.D_pno) {
		input.D_ptype = 2
	} else {
		if input.D_appeui == "" {
			utils.JsonResponseError(c, "d_appeui参数必填!")
			return
		}
	}

	tx := models.DB.Begin()
	//lorawan设备
	if models.IsLoraWanDevice(input.D_pno) {
		if input.Lda_dev_addr == "" || input.Lda_app_key == "" || input.Lda_app_s_key == "" || input.Lda_nwk_s_key == "" {
			tx.Rollback()
			utils.JsonResponseError(c, "loraWan设备4个参数必填")
			return
		}
		input.D_ptype = 1
		input.Lda_dno = input.D_no
		if err := tx.Create(&input.Lorawan).Error; err != nil {
			tx.Rollback()
			utils.JsonResponseError(c, err)
			return
		}
	}
	//网关设备 和 普通设备 lorawan 设备写入设备表
	if err := tx.Create(&input.Device).Error; err != nil {
		tx.Rollback()
		utils.JsonResponseError(c, err)
		return
	}
	tx.Commit()
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
	//检查设备编号
	if models.IsNormalDevice(input.D_pno) {
		utils.JsonResponseError(c, "P_no:型号有误")
		return
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

	dno := c.Query("dno")
	if dno == "" {
		utils.JsonResponseError(c, "d_no主键不能为空")
		return
	}
	if err := models.DB.Delete(&models.Device{D_no: dno}).Error; err != nil {
		utils.JsonResponseError(c, err)
	} else {
		utils.JsonResponseSuccess(c, nil)
	}
}
