package db

import (
	"fmt"
	"gbyk/config"
	"log"
	"sync"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"gopkg.in/mgo.v2"
)

const (
	REDIS = "redis"
	MYSQL = "mysql"
	MGO   = "mgo"
)

var (
	onceMysql sync.Once
	once      sync.Once
	instance  *singleton
	conf      = config.GetConfig()
)

type (
	Driver string
)
type singleton struct {
	services *sync.Map
}

//获取mysql
func GetMysqlDB() *gorm.DB {
	return newInstance(MYSQL).(*gorm.DB)
}

//获取redis
func GetRedisDB() *redis.Client {
	return newInstance(REDIS).(*redis.Client)
}

//获取mgo
func GetMgoDB() *mgo.Session {
	return newInstance(MGO).(*mgo.Session)
}

func newInstance(driver Driver) interface{} {
	s := getInstance()
	if val, ok := s.services.Load(driver); ok {
		return val
	}
	var r interface{}
	switch driver {
	case REDIS:
		r, _ = s.getOrSetMap(REDIS, newRedis())
	case MYSQL:
		r, _ = s.getOrSetMap(MYSQL, newMysql())
	case MGO:
		r, _ = s.getOrSetMap(MGO, newMGO())
	default:
	}
	return r
}

func getInstance() *singleton {
	if instance == nil {
		once.Do(func() {
			instance = &singleton{services: &sync.Map{}}
		})
		instance.getOrSetMap(REDIS, newRedis())
	}
	return instance
}

func (s *singleton) getOrSetMap(name Driver, service interface{}) (interface{}, bool) {
	return s.services.LoadOrStore(name, service)
}

func newRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Redis.Host, conf.Redis.Port),
		Password: conf.Redis.Password,
		DB:       conf.Redis.DbName,
		PoolSize: conf.Redis.Pool,
	})
	client.Ping().Result()
	return client
}

func newMysql() *gorm.DB {
	var mysql *gorm.DB
	onceMysql.Do(func() {
		var err error
		dbURL := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
			conf.Mysql.User, conf.Mysql.Password, "tcp", conf.Mysql.Host,
			conf.Mysql.Port, conf.Mysql.DbName, conf.Mysql.Charset)
		mysql, err = gorm.Open("mysql", dbURL)
		if err != nil {
			log.Panic(err)
		}
		mysql.DB().SetMaxIdleConns(10)
		mysql.DB().SetMaxOpenConns(conf.Mysql.Pool)
	})
	return mysql
}

func newMGO() *mgo.Session {
	var session *mgo.Session
	var err error

	if conf.Server.ServerEnv == "inline" {
		session, err = mgo.Dial(fmt.Sprintf("%s:%d", conf.Mongo.Host, conf.Mongo.Port))
	} else {
		session, err = mgo.Dial(fmt.Sprintf("%s:%s@%s:%d", conf.Mongo.User,
			conf.Mongo.Password, conf.Mongo.Host, conf.Mongo.Port))
		if err != nil {
			log.Panic(err)
		}
	}
	session.SetMode(mgo.Monotonic, true)
	session.SetPoolLimit(conf.Mongo.Pool)
	return session.Clone()
}
