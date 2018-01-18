package models

import (
	"time"
)

type Device struct {
	D_no          string `gorm:"primary_key"` // set AnimalId to be primary key
	D_pno         string
	D_appeui      string
	D_ptype       int
	D_name        string
	D_userid      int
	D_regist_time string
	D_enabled     int
}

func (Device) TableName() string {
	return "iot_device"
}

func (d *Device) BeforeSave() (err error) {
	d.D_regist_time = time.Now().Format("2006-01-02 15:04:05")
	return
}
