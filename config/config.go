package config

import (
	"log"

	"f.in/v/utils"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Server server
	Mysql  mysql
	Redis  redis
	Mongo  mongo
	File file
}
type server struct {
	Port      int    `json:"port"`
	ServerEnv string `json:"serverEnv"`
	Debug     bool   `json:"debug"`
}

type mysql struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"dbName"`
	Charset  string `json:"charset"`
	Pool     int    `json:"pool"`
}

type redis struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	DbName   int    `json:"dbName"`
	Pool     int    `json:"pool"`
}

type mongo struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	Pool     int    `json:"pool"`
}
type file struct {
	Path string `json:"path"` //文件保存路径
}

func GetConfig() (conf Config) {
	if _, err := toml.DecodeFile(getCOnfigPath(), &conf); err != nil {
		log.Println(err)
		return
	}
	return
}

func getCOnfigPath() string {
	return utils.SelfDir() + "/config/config.toml"
}
