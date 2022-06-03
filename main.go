package main

import (
	"douyin/repository"
	"douyin/server"
	"github.com/gin-gonic/gin"
)

func main() {

	repository.InitMysql()

	r := gin.Default()

	r.Use(gin.Logger())

	server.InitRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
