package model

import (
	"time"
	_ "github.com/jinzhu/gorm"
)

//点赞
type Likes struct {
	ID        uint       `gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`

	UserId   int64  `json:"userId"`
	ByUserId int64  `json:"byUserId"` //点赞人的id
	From     string `json:"from"`     //点赞出处，关联表名
	FromId   int64  `json:"fromId"`   //相关出处id
}
