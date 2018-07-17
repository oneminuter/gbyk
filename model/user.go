package model

import (
	"time"

	"gbyk/db"

	"gbyk/logs"

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
	Mobile   string `json:"mobile" gorm:"size:15"`     //电话
	Usernick string `string:"usernick" grom:"size:20"` //昵称
	Sex      int    `json:"sex"  grom:"size:4"`        //行别 0：未知，1：男，2：女
}

func (this *User) Add() error {
	mdb := db.GetMysqlDB()
	tx := mdb.Begin()
	err := tx.Model(this).Create(this).Error
	if err != nil {
		logs.Log(err)
		tx.Rollback()
	}
	tx.Commit()
	return err
}

func (this *User) Save() error {
	mdb := db.GetMysqlDB()
	tx := mdb.Begin()
	err := tx.Model(this).Save(this).Error
	if err != nil {
		logs.Log(err)
		tx.Rollback()
	}
	tx.Commit()
	return err
}

func (this *User) Update(update interface{}, where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()
	tx := mdb.Begin()
	err := tx.Model(this).Where(where, args...).Update(update).Error
	if err != nil {
		logs.Log(err)
		tx.Rollback()
	}
	tx.Commit()
	return err
}

func (this *User) QueryOne(where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()
	return mdb.Where(where, args...).Last(this).Error
}

func (*User) QueryAll() ([]User, error) {
	mdb := db.GetMysqlDB()
	var userList []User
	err := mdb.Find(&userList).Error
	return userList, err
}
