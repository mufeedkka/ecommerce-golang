package routes

import (
	"github.com/gin-gonic/gin"
)

func AdminRoute(router *gin.Engine) {
	admin := router.Group("/api/admin")
	{
		admin.GET("/getproducts", controllers.AdminGetAllProducts)
	}
}
