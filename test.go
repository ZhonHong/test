package main

import (
	"io"
	"os"
	"test/database"
	"test/middlewares"
	. "test/src"

	"github.com/gin-gonic/gin"
)

func setupLogging() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout) //寫入
}

func main() {

	setupLogging()
	router := gin.Default()
	router.Use(gin.BasicAuth(gin.Accounts{"Tom": "123456"}), middlewares.Logger()) //操作驗證
	v1 := router.Group("/v1")
	AddUserRouter(v1)

	go func() {
		database.DD()
	}() //專案啟動同步啟動資料庫

	router.Run(":8080")
}
