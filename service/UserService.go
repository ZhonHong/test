package service

//定義尋找及新增的服務
import (
	"log"
	"net/http"
	"strconv"
	"test/pojo"

	"github.com/gin-gonic/gin"
)

var userList = []pojo.User{}

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
	userList = append(userList, user)            //將uesr變量加入userList切片中，並傳回客戶端
	c.JSON(http.StatusOK, "Successfullu Posted") //使用c.JSON方法返回一个HTTP狀態码及成功字串

}

func DeleteUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id")) //將回傳值轉為string並取得URL內的參數id
	for _, user := range userList {
		log.Println(user) //golang中的陣列無法讀取資料後確認資料相同後刪除，只能用index刪除

		userList = append(userList[:userId], userList[userId+1:]...)
		c.JSON(http.StatusOK, "Successfully deleted")
		return

	}
	c.JSON(http.StatusNotFound, "Error")
}

func PutUser(c *gin.Context) {
	beforeUser := pojo.User{}
	err := c.BindJSON(&beforeUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error")
	}
	userId, _ := strconv.Atoi(c.Param("id")) //將回傳值轉為string並取得URL內的參數id
	for key, user := range userList {
		if userId == user.Id {
			userList[key] = beforeUser
			log.Println(userList[key])
			c.JSON(http.StatusOK, "Successfully")
			return
		}
	}
	c.JSON(http.StatusNotFound, "Error")
}
