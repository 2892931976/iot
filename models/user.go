package models

import (
	"fmt"
	"github.com/mojocn/turbo-iot/config"
)

type User struct {
	Uid      int64 `gorm:"primary_key"` // set AnimalId to be primary key
	Email    string
	Nickname string
	Pwd      string
	Sex      string
}

func (User) TableName() string {
	return "pm_member"
}

func (m *User) RedisKey() string {
	return fmt.Sprint(config.RedisKeyPrefix, "jwt:", m.Uid)
}
