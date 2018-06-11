package model

import (
	"gbyk/config"
	"gbyk/db"
)

func init()  {
	conf      := config.GetConfig()
	mysql := db.GetMysqlDB()
	mysql.AutoMigrate(&User{}, &Wechat{}, )
	mysql.LogMode(conf.Server.Debug)
}