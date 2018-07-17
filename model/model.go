package model

import (
	"gbyk/config"
	"gbyk/db"
)

type DbStorage interface {
	Add() error
	Save() error
	Update(update interface{}, where interface{}, args ...interface{}) error
	QueryOne(where interface{}) error
	QueryAll() ([]interface{}, error)
}

func init() {
	conf := config.GetConfig()
	mysql := db.GetMysqlDB()
	mysql.AutoMigrate(&User{}, &Wechat{})
	mysql.LogMode(conf.Server.Debug)
}
