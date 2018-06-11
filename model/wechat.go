package model

import (
	"time"

	_ "github.com/jinzhu/gorm"
)

type Wechat struct {
	ID        uint       `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `json:"-"`

	UserId int64 `json:"userId"`

	NickName  string `json:"nickName" gorm:"size:20"`  //昵称
	AvatarUrl string `json:"avatarUrl" gorm:"size:20"` //头像
	Gender    int    `json:"gender" gorm:"size:4"`     //值为1时是男性，值为2时是女性，值为0时是未知
	City      string `json:"city" gorm:"size:100"`     //所在城市
	Country   string `json:"country" gorm:"size:50"`   //所在国家
	Province  string `json:"province" gorm:"size:50"`  //所在省份
	Language  string `json:"language" gorm:"size:20"`  //语言，简体中文为zh_CN
}
