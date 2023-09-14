package pojo

import (
	"log"
	"test/database"
)

type Post struct {
	Id         int    `json:"PostId"`
	Title      string `json:"PostTitle"`
	Content    string `json:"PostContent"`
	Created_at string `json:"Postcreated_at"`
	User       string `json:"PostUser"`
}

type Image struct {
	Id  int    `json:"ImagesId"`
	Img string `json:"ImagesImg"`
}

type Images struct {
	ImageList     []Image `json:"ImageList"`
	ImageListSize int     `json:"ImageListSize"`
}

// 多筆上傳Model
func FindALLImages() []Image {
	var images []Image
	database.DBconnect.Find(&images) //若未指定尋找內容，則全部列印出來
	return images
}

func CreatePost(post Post) *Post {

	db := database.DBconnect
	tx := db.Create(&post)
	if tx.Error != nil {
		log.Printf("error:%s", tx.Error.Error())
		return nil
	} else {
		return &post
	}
}

func CreatePostImg(image Image) *Image {
	db := database.DBconnect
	tx := db.Create(&image)
	if tx.Error != nil {
		log.Printf("error:%s", tx.Error.Error())
		return nil
	} else {
		return &image
	}
}
