package model

import (
	"time"
	_ "github.com/jinzhu/gorm"
)

type User struct {
	ID        uint       `gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`

	Username string `json:"username" grom:"size:20"`   //姓名
	Password string `json:"password" grom:"size:32"`   //密码
	Random   string `json:"random" grom:"size:15"`     //加密字符串
	Mobile string `json:"mobile" gorm:"size:15"`
	Usernick string `string:"usernick" grom:"size:20"` //昵称
	Sex      int    `json:"sex"  grom:"size:4"`        //行别 0：未知，1：男，2：女
}
