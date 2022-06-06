package main

import (
	"douyin/common"
	"douyin/model"
	"douyin/server"
	"github.com/gin-gonic/gin"
)

func main() {

	model.InitMysql()
	r := gin.Default()

	r.Use(gin.Logger())

	server.InitRouter(r)

	model.Db.AutoMigrate(&common.User{})

	model.Db.AutoMigrate(&common.Video{})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
