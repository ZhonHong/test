package service

//定義尋找及新增的服務
import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
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

// 多筆上傳
func FindALLImage(c *gin.Context) {
	images := pojo.FindALLImages()
	c.JSON(http.StatusOK, images)
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

// UserPosts
func UserPost(c *gin.Context) {
	post := pojo.Post{}      //user的pojo.User結構變量體
	err := c.BindJSON(&post) //錯誤判斷,用指標指向user變量傳遞給BindJSON，以便在方法內部修改user變量
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error :"+err.Error()) //判斷是否空值，則使用c.JSON方法返回一个HTTP狀態码
		return
	}
	//userList = append(userList, user)            //將uesr變量加入userList切片中，並傳回客戶端
	newPost := pojo.CreatePost(post)

	// 将当前时间转换为字符串

	c.JSON(http.StatusOK, newPost) //使用c.JSON方法返回一个HTTP狀態碼及成功字串
}

func UserPostImages(c *gin.Context) {
	image := pojo.Image{}
	// 解析表單数據，包括文件
	file, err := c.FormFile("File")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() + "未上傳文件"})
		return
	}

	// 生成一个随机文件名或根据需要生成文件名
	// 这里使用了时间戳作为文件名，你可以根据需要生成唯一的文件名
	timestamp := time.Now().UnixNano()
	filename := fmt.Sprintf("%d%s", timestamp, filepath.Ext(file.Filename))

	// 文件保存的相對路径
	uploadDir := "./Images"
	err = os.MkdirAll(uploadDir, os.ModePerm) //路徑中沒有資料夾則產生資料夾

	// 拼接完整的文件保存路径
	imgFullPath := filepath.Join(uploadDir, filename)

	// 将文件保存到本地文件系统
	err = c.SaveUploadedFile(file, imgFullPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error() + "保存文件失败"})
		return
	}

	// 保存文件信息到数据库
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
		// 生成一个随机文件名或根据需要生成文件名
		timestamp := time.Now().UnixNano()
		filename := fmt.Sprintf("%d%s", timestamp, filepath.Ext(file.Filename))

		// 指定文件保存的路径
		uploadDir := "./Images"
		err = os.MkdirAll(uploadDir, os.ModePerm)

		// 拼接完整的文件保存路径
		imgFullPath := filepath.Join(uploadDir, filename)

		// 将文件保存到本地文件系统
		err := c.SaveUploadedFile(file, imgFullPath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "保存文件失败"})
			return
		}

		// 保存文件信息到数据库
		image := pojo.Image{
			Img: filename,
		}
		images = append(images, image)

	}
	c.JSON(http.StatusOK, images)
}
