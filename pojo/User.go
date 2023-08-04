package pojo

import (
	"test/database"
)

type User struct {
	Id       int    `json:"UserId"`
	Name     string `json:"UserName"`
	Password string `json:"UserPassword"`
	Email    string `json:"UserEmail"`
}

func FindALLUsers() []User {
	var users []User
	database.DBconnect.Find(&users) //若未指定尋找內容，則全部列印出來
	return users
}

func FindByUserId(userId string) User {
	var user User
	database.DBconnect.Where("id = ?", userId).First(&user) //id的值由userId注入，在使用First(找出第一個ID)去尋找指向user的table
	return user
}
