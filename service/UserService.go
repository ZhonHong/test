package service

//定義尋找及新增的服務
import (
	"log"
	"net/http"
	"test/middlewares"

	//"os/user"

	"test/pojo"

	"github.com/gin-gonic/gin"
)

//var userList = []pojo.User{}

// Get User，當使用FindALLUser函數時，將userList以JSON格式返回客戶端
func FindALLUser(c *gin.Context) {
	//c.JSON(http.StatusOK, userList) //接受gin.Context類型的參數c，並使用c.JSON方法將userList以JSON格式返回客戶端
	users := pojo.FindALLUsers()
	c.JSON(http.StatusOK, users)
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
	name := c.PostForm("name")
	password := c.PostForm("password")
	user := pojo.CheckUserPassword(name, password)
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	middlewares.SaveSession(c, user.Id) //將user.Id傳入
	c.JSON(http.StatusOK, gin.H{
		"message": "Login Successfully",
		"User":    user,
		"Session": middlewares.GetSession(c),
	})
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
	if sessionId == 0 {
		c.JSON(http.StatusUnauthorized, "Error")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Check Session Successfully",
		"User":    middlewares.GetSession(c),
	})
}
