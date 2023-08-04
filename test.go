package main

import (
	"test/database"
	. "test/src"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	AddUserRouter(v1)

	go func() {
		database.DD()
	}() //專案啟動同步啟動資料庫

	router.Run(":8080")
}
