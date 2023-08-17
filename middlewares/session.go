package middlewares

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const userkey = "session_id"

// Use cookie to store session id
func SetSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(userkey))
	store.Options(sessions.Options{
		MaxAge:   3600, //設置過期時間為1小時(以秒作為單位)
		HttpOnly: true, //只能在 HTTP 請求中訪問 Cookie
	})
	return sessions.Sessions("mysession", store)
}

// User Auth Session Middle
func AuthSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c) //創建session預設值
		sessionID := session.Get(userkey)
		if sessionID == nil {
			// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			// 	"message:": "此頁面需要登入",
			// })
			c.JSON(http.StatusBadRequest, gin.H{
				"message:": "此頁面需要登入",
			})
			return
		}
		c.Next()
	}
}

// Save Session for User
func SaveSession(c *gin.Context, userID int) {
	session := sessions.Default(c)
	session.Set(userkey, userID)
	session.Save()
}

// Clear Session for User
func ClearSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

// Get Session for User查看Session狀態
func GetSession(c *gin.Context) int {
	session := sessions.Default(c)
	sessionID := session.Get(userkey)
	if sessionID == nil {
		return 0
	}
	return sessionID.(int)
}

// Check Session for User
func CheckSession(c *gin.Context) bool {
	session := sessions.Default(c)
	sessionID := session.Get(userkey)
	// if sessionID == nil{
	// 	return false
	// }
	// return true
	return sessionID != nil
}
