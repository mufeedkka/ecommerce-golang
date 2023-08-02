package routes

import (
	"git.com/ecommerce/controllers"
	"github.com/gin-gonic/gin"
)

func AdminRoute(router *gin.Engine) {

	admin := router.Group("")
	{
		admin.POST("/createproduct", controllers.AdminCreateProduct)
		// admin.POST("/updateproduct/:productid", controllers.AdminUpdateProduct)
		admin.GET("/", controllers.Ping)
	}
}
