package routes

import (
	"git.com/ecommerce/controllers"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	user := router.Group("/user")
	{
		user.POST("/signup", controllers.UserSignUp)
		// user.POST("/signin", controllers.UserSignIn)
		// user.GET("/validate", middlewares.UserAuth, controllers.ValidateUser)
		// user.POST("/addprofile", middlewares.UserAuth, controllers.AddUserProfile)
	}
}
