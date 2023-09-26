package service

//定義尋找及新增的服務
import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"test/middlewares"
	"time"

	//"os/user"

	"test/pojo"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//var userList = []pojo.User{}

// Get User，當使用FindALLUser函數時，將userList以JSON格式返回客戶端
func FindALLUser(c *gin.Context) {
	//c.JSON(http.StatusOK, userList) //接受gin.Context類型的參數c，並使用c.JSON方法將userList以JSON格式返回客戶端
	users := pojo.FindALLUsers()
	c.JSON(http.StatusOK, users)
}

// 觀看個人貼文
func FindALLPosts(c *gin.Context) {

	//c.JSON(http.StatusOK, userList) //接受gin.Context類型的參數c，並使用c.JSON方法將userList以JSON格式返回客戶端
	errStr, claims := middlewares.ParseToken(c) //對應ParseToken(string, *jwt.MapClaims)
	if claims != nil {
		// Token Cliams Data
		var test jwt.MapClaims //宣告test為jwt.MapClaims型態
		test = *claims         //使test可以使用claims資料

		userId := int(test["userId"].(float64)) //UserId
		userName := test["userName"].(string)
		userEmail := test["userEmail"].(string)
		fmt.Println("User NameId:", userId)
		fmt.Println("User Name:", userName)
		fmt.Println("User Email:", userEmail)
		posts := pojo.FindALLPosts(userName)
		c.JSON(http.StatusOK, posts)
	} else {
		fmt.Println("error" + errStr)
	}

}

// 觀看全部公開貼文
func GetALLPosts(c *gin.Context) {
	posts := pojo.GetALLPosts()
	c.JSON(http.StatusOK, posts)
}

// 多筆上傳
func FindALLImage(c *gin.Context) {
	images := pojo.FindALLImages()
	c.JSON(http.StatusOK, images)
}

// Get Post by Id
func GetPost(c *gin.Context) {

	post := pojo.GetPost(c.Param("id"))
	if post.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	log.Println("Post ->", post)
	c.JSON(http.StatusOK, post)
}

// Get User by Id
func FindByUserId(c *gin.Context) {
	user := pojo.FindByUserId(c.Param("id"))
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	log.Println("User ->", user)
	c.JSON(http.StatusOK, user)
}

// Delete user Post
func DeletePost(c *gin.Context) {
	postId := c.Param("id")
	postIdInt, err := strconv.Atoi(postId) //將URL中的id字串轉為int
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid ID")
		return
	}

	err = pojo.DeletePost(postIdInt) //使用DeletePost方法進行軟刪除，如果刪除成功會回傳nil值
	if err != nil {
		c.JSON(http.StatusNotFound, "Error: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, "Successfully")
}

// func GetUserData(c *gin.Context){
// 	user := pojo.GetUserData()
// }

// Post User
func PostUser(c *gin.Context) {
	user := pojo.User{}      //user的pojo.User結構變量體
	err := c.BindJSON(&user) //錯誤判斷,用指標指向user變量傳遞給BindJSON，以便在方法內部修改user變量
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error :"+err.Error()) //判斷是否空值，則使用c.JSON方法返回一个HTTP狀態码
		return
	}
	//userList = append(userList, user)            //將uesr變量加入userList切片中，並傳回客戶端
	newUser := pojo.CreateUser(user)
	c.JSON(http.StatusOK, newUser) //使用c.JSON方法返回一个HTTP狀態码及成功字串
}

// DeleteUser
// func DeleteUser(c *gin.Context) {
// 	userId, _ := strconv.Atoi(c.Param("id")) //將回傳值轉為string並取得URL內的參數id
// 	for _, user := range userList {
// 		log.Println(user) //golang中的陣列無法讀取資料後確認資料相同後刪除，只能用index刪除

// 		userList = append(userList[:userId], userList[userId+1:]...)
// 		c.JSON(http.StatusOK, "Successfully deleted")
// 		return

// 	}
// 	c.JSON(http.StatusNotFound, "Error")
// }

// delete User
func DeleteUser(c *gin.Context) {
	user := pojo.DeleteUser(c.Param("id"))
	if !user {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	c.JSON(http.StatusOK, "Successfully")
}

// Put
// func PutUser(c *gin.Context) {
// 	beforeUser := pojo.User{}
// 	err := c.BindJSON(&beforeUser)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, "Error")
// 	}
// 	userId, _ := strconv.Atoi(c.Param("id")) //將回傳值轉為string並取得URL內的參數id
// 	for key, user := range userList {
// 		if userId == user.Id {
// 			userList[key] = beforeUser
// 			log.Println(userList[key])
// 			c.JSON(http.StatusOK, "Successfully")
// 			return
// 		}
// 	}
// 	c.JSON(http.StatusNotFound, "Error")
// }

// Edit Post
func EditPost(c *gin.Context) {
	post := pojo.PostEdit{}
	err := c.BindJSON(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, "err:"+err.Error())
		return
	}
	post = pojo.EditPost(c.Param("id"), post)
	if post.Id == 0 {
		c.JSON(http.StatusBadRequest, "err:"+err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Edit Success",
		"post":    post,
	})
}

// Update User
func PutUser(c *gin.Context) {
	user := pojo.User{}
	err := c.BindJSON(&user) //錯誤判斷,用指標指向user變量傳遞給BindJSON，以便在方法內部修改user變量
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error")
		return
	}
	user = pojo.UpdateUser(c.Param("id"), user) //取得URL內的參數id及將修改資料
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}

	c.JSON(http.StatusOK, user) //回傳user
}

// CreateUserList
func CreateUserList(c *gin.Context) {
	users := pojo.Users{}
	err := c.BindJSON(&users)
	if err != nil {
		c.String(400, "Error:%s", err.Error())
		return
	}
	c.JSON(http.StatusOK, users)
}

// Login User
func LoginUser(c *gin.Context) {
	type reqData struct {
		Email    *string `json:"Email"`
		Password *string `json:"Password"`
	}

	var req reqData
	c.ShouldBind(&req)

	// user := pojo.User{}
	// if err := c.ShouldBindJSON(&user); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// email := user.Email
	// password := user.Password
	if req.Email == nil || req.Password == nil {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	us := pojo.CheckUserPassword(*req.Email, *req.Password)

	token, err := middlewares.GenerateToken(us.Id, us.Name, us.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Login Successfully",
		"token":   token,
		"User":    us,
	})
}

func GetUserData(c *gin.Context) {
	errStr, claims := middlewares.ParseToken(c) //對應ParseToken(string, *jwt.MapClaims)
	if claims != nil {
		// Token Cliams Data
		var test jwt.MapClaims //宣告test為jwt.MapClaims型態
		test = *claims         //使test可以使用claims資料

		userId := int(test["userId"].(float64)) //UserId
		userName := test["userName"].(string)
		userEmail := test["userEmail"].(string)
		fmt.Println("User ID:", userId)
		fmt.Println("User Name:", userName)
		fmt.Println("User Email:", userEmail)

		c.JSON(http.StatusOK, gin.H{
			"userId":    userId,
			"userName":  userName,
			"userEmail": userEmail,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": errStr})
	}

}

// Logout User
func LogoutUser(c *gin.Context) {
	middlewares.ClearSession(c)
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout Successfully",
	})
}

// Check User Session
func CheckUserSession(c *gin.Context) {
	sessionId := middlewares.GetSession(c)
	if sessionId < 0 {
		c.JSON(http.StatusUnauthorized, "Error")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Check Session Successfully",
		"User":    middlewares.GetSession(c),
	})
}

// PostsContent
func UserPost(c *gin.Context) {
	post := pojo.Post{}                         //user的pojo.User結構變量體
	errStr, claims := middlewares.ParseToken(c) //對應ParseToken(string, *jwt.MapClaims)
	if claims != nil {
		// Token Cliams Data
		var test jwt.MapClaims //宣告test為jwt.MapClaims型態
		test = *claims         //使test可以使用claims資料

		userId := int(test["userId"].(float64)) //UserId
		userName := test["userName"].(string)
		userEmail := test["userEmail"].(string)
		fmt.Println("User NameId:", userId)
		fmt.Println("User Name:", userName)
		fmt.Println("User Email:", userEmail)

		post.User = userName
	} else {
		fmt.Println("error" + errStr)
	}
	err := c.BindJSON(&post) //錯誤判斷,用指標指向user變量傳遞給BindJSON，以便在方法內部修改user變量
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error :"+err.Error()) //判斷是否空值，則使用c.JSON方法返回一个HTTP狀態码
		return
	}

	//userList = append(userList, user)            //將uesr變量加入userList切片中，並傳回客戶端
	newPost := pojo.CreatePost(post)
	fmt.Println(newPost)
	c.JSON(http.StatusOK, "Success") //使用c.JSON方法返回一个HTTP狀態碼及成功字串
}

func UserPostImages(c *gin.Context) {
	image := pojo.Image{}
	// 解析表單數據，包括文件
	file, err := c.FormFile("File") //名稱需與前端form-data Key名稱相同
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() + "未上傳文件"})
		return
	}

	// 生成随機文件名
	// 用時間戳作為文件名
	timestamp := time.Now().UnixNano()
	filename := fmt.Sprintf("%d%s", timestamp, filepath.Ext(file.Filename))

	// 文件保存的相對路径
	uploadDir := "./Images"
	err = os.MkdirAll(uploadDir, os.ModePerm) //路徑中沒有資料夾則產生資料夾

	// Join完整的文件路径
	imgFullPath := filepath.Join(uploadDir, filename)

	// 將文件存到本地端
	err = c.SaveUploadedFile(file, imgFullPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error() + "保存文件失败"})
		return
	}

	// 將文件信息存到DB
	image.Img = filename
	newImage := pojo.CreatePostImg(image)
	c.JSON(http.StatusOK, newImage)
}

// update more file
func UserPostmoreImages(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() + "未上傳文件"})
		return
	}
	var images []pojo.Image
	files := form.File["File"]

	for _, file := range files {
		// 生成隨機文件名
		timestamp := time.Now().UnixNano()
		filename := fmt.Sprintf("%d%s", timestamp, filepath.Ext(file.Filename))

		// 指定文件保存路徑
		uploadDir := "./Images"
		err = os.MkdirAll(uploadDir, os.ModePerm)

		imgFullPath := filepath.Join(uploadDir, filename)

		// 將文件保存至本地端
		err := c.SaveUploadedFile(file, imgFullPath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "保存文件失败"})
			return
		}

		// 將文件保存至DB
		image := pojo.Image{
			Img: filename,
		}
		images = append(images, image)

	}
	c.JSON(http.StatusOK, images)
}
