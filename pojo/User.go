package pojo

import (
	"log"
	"test/database"
)

type User struct { //DB : Users
	Id       int    `json:"UserId" binding:"required"`        //required必填項目，Id DB : id, UserId DB: user_id
	Name     string `json:"UserName" binding:"required,gt=5"` // Name DB:name, UserName DB: user_name
	Password string `json:"UserPassword" binding:"min=4,max=20,userpasd"`
	Email    string `json:"UserEmail" binding:"email"`
}

type Users struct {
	UserList     []User `json:"UserList" binding:"required,gt=0,lt=3"`
	UserListSize int    `json:"UserListSize"`
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

// Post
func CreateUser(user User) User {
	database.DBconnect.Create(&user)
	return user
}

// DeleteUser
func DeleteUser(userId string) bool {
	user := User{}
	result := database.DBconnect.Where("id = ?", userId).Delete(&user) //找到id進行Delete user並回傳
	log.Println(result)
	return result.RowsAffected > 0 //依Rows回傳值判斷bool，1則true；0則flase
}

// UpdateUser
func UpdateUser(userId string, user User) User {
	database.DBconnect.Model(&user).Where("id = ?", userId).Updates(user) //更改User結構
	return user
}

// CheckUserPassword
func CheckUserPassword(name string, password string) User {
	user := User{}
	database.DBconnect.Where("name = ? and password = ?", name, password).First(&user)
	return user
}
