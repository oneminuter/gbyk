package model

import (
	"time"
	_ "github.com/jinzhu/gorm"
)

//
type Auth struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time

	UserId int64 `json:"userId"`
	Account int `json:"account" gorm:"default:1 size:4"` //账号状态 0: 封禁，1：正常

}