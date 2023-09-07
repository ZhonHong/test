package main

import (
	"io"
	"os"
	"test/database"
	"test/middlewares"
	. "test/src"

	"test/pojo"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func setupLogging() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout) //寫入
}

func main() {
	database.DD() //專案啟動同步啟動資料庫
	// defer database.DBconnect.Statement.ReflectValue.Close() // 在 main 函式結束時關閉資料庫連接
	setupLogging()
	router := gin.Default()

	//register Validator Func
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("userpasd", middlewares.UserPasd)
		v.RegisterStructValidation(middlewares.UserList, pojo.Users{})

	}

	router.Use(gin.Recovery(), middlewares.Logger(), cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // 允许的源，使用通配符 "*" 表示任何来源都允许
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})) //操作驗證，若發生panics則會強制關閉應用程式(代號500)
	api := router.Group("/api")
	AddUserRouter(api)

	router.Run(":8080")

}
