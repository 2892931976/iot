package models

import (
	"time"
)

type Device struct {
	D_no          string `gorm:"primary_key"` // set AnimalId to be primary key
	D_pno         string
	D_appeui      string
	D_key         string
	D_name        string
	D_userid      int
	D_model       string
	D_version     string
	D_cid         int
	D_regist_time string
}

func (Device) TableName() string {
	return "iot_device"
}

func (d *Device) BeforeSave() (err error) {
	d.D_regist_time = time.Now().Format("2006-01-02 15:04:05")
	return
}
