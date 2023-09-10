package src

//對user集中管理
import (
	session "test/middlewares"
	"test/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {

	api := r.Group("/api", session.SetSession())
	{
		api.GET("/", service.FindALLUser)
		api.GET("/:id", service.FindByUserId)
		api.POST("/Register", service.PostUser)
		api.POST("/more", service.CreateUserList)
		//user.DELETE("/:id", service.DeleteUser)
		api.PUT("/:id", service.PutUser)
		//LoginUser
		api.POST("/Login", service.LoginUser)
		api.GET("/getUserData", service.GetUserData)
	}
	// Check User session
	api.GET("/check", service.CheckUserSession)

	api.Use(session.AuthSession())
	{
		// delete user
		api.DELETE("/:id", service.DeleteUser)
		// LogoutUser
		api.GET("/logout", service.LogoutUser)
	}
}
