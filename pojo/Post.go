package pojo

import (
	"log"
	"regexp"
	"strings"
	"test/database"
	"time"

	"gorm.io/gorm"
)

// Post Model
type Post struct {
	Id      int    `json:"Id"`
	Title   string `json:"Title" binding:"required"`
	Content string `json:"Content" binding:"required"`
	User    string `json:"User"`
}

// Get Model
type PostGet struct {
	Id          int             `json:"Id"`
	Title       string          `json:"Title" binding:"required"`
	Content     string          `json:"Content" binding:"required"`
	User        string          `json:"User"`
	Create_Time time.Time       `json:"create_time,omitempty"`
	Update_Time *time.Time      `json:"update_time,omitempty"`
	Delete_Time *gorm.DeletedAt `gorm:"index" json:"delete_time,omitempty"`
	Src         string          `json:"Src,omitempty"`
}

// Update Mode
type PostEdit struct {
	Id          int        `json:"Id"`
	Title       string     `json:"Title" binding:"required"`
	Content     string     `json:"Content" binding:"required"`
	Update_Time *time.Time `json:"update_time,omitempty"`
	Src         string     `json:"Src,omitempty"`
}

type Image struct {
	Id  int    `json:"ImagesId"`
	Img string `json:"ImagesImg"`
}

type Images struct {
	ImageList     []Image `json:"ImageList"`
	ImageListSize int     `json:"ImageListSize"`
}

func DeletePost(postId int) error {
	var postget PostGet

	// 查詢要删除的資料
	if err := database.DBconnect.Table("posts").First(&postget, postId).Error; err != nil {
		return err
	}

	// 進行軟删除操作，不需要傳 postId
	if err := database.DBconnect.Table("posts").Delete(&postget).Error; err != nil {
		return err
	}

	return nil
}

func FindALLPosts(userName string) []PostGet {
	var postgets []PostGet
	database.DBconnect.Table("posts").Where("user = ?", userName).Find(&postgets)
	return postgets
}

func GetALLPosts() []PostGet {
	var postgets []PostGet
	database.DBconnect.Table("posts").Select("Id, Title, Content, User, Create_Time, Update_Time").Find(&postgets)
	for i, postget := range postgets {
		regex := regexp.MustCompile(`<img[^>]*src="([^"]+)"[^>]*>`)
		matchs := regex.FindAllStringSubmatch(postget.Content, -1) //在postget.Content內尋找所有符合字串 -1代表全部
		srcList := []string{}
		for _, match := range matchs { //走訪所有匹配結果
			if len(match) > 1 {
				src := match[1]
				srcList = append(srcList, src)
			}
		}
		if len(srcList) > 0 {
			postgets[i].Src = strings.Join(srcList, ", ")
		} else {
			postgets[i].Src = "" // 如果没有匹配到 src 内容，则设置为空字符串
		}
		if len(postget.Content) > 50 {
			postgets[i].Content = postget.Content[:50]
		}
	}
	return postgets
}

func GetPost(postId string) PostGet {
	var postget PostGet
	database.DBconnect.Table("posts").Where("id = ?", postId).First(&postget) //id的值由userId注入，在使用First(找出第一個ID)去尋找指向user的table
	return postget
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

// EditPost
func EditPost(postId string, postedit PostEdit) PostEdit {
	database.DBconnect.Model(&postedit).Table("posts").Where("id = ?", postId).Updates(postedit) //更改User結構
	return postedit
}
