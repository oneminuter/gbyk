package main

import (
	"fmt"
	"gbyk/config"
	"log"

	"github.com/gin-gonic/gin"
)

var conf = config.GetConfig()

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Run(fmt.Sprintf(":%d", conf.Server.Port))
	log.Println(fmt.Sprintf("Listen and Server in 0.0.0.0:%d", conf.Server.Port))
}
