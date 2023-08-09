package src

//對user集中管理
import (
	session "test/middlewares"
	"test/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("/users", session.SetSession())
	user.GET("/", service.FindALLUser)
	user.GET("/:id", service.FindByUserId)
	user.POST("/", service.PostUser)
	user.POST("/more", service.CreateUserList)
	//user.DELETE("/:id", service.DeleteUser)
	user.PUT("/:id", service.PutUser)
	//LoginUser
	user.POST("/login", service.LoginUser)

	// Check User session
	user.GET("/check", service.CheckUserSession)

	user.Use(session.AuthSession())
	{
		// delete user
		user.DELETE("/:id", service.DeleteUser)
		// LogoutUser
		user.GET("/logout", service.LogoutUser)
	}
}
