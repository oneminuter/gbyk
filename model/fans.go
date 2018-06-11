package model

import (
	"time"
	_ "github.com/jinzhu/gorm"
)

//粉丝
type Fans struct {
	ID        uint       `gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
	
	UserId int64 `json:"userId"`
	FansUserId int64 `json:"fansUserId"`
}
