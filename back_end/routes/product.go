package routes

import (
	"git.com/ecommerce/controllers"
	"github.com/gin-gonic/gin"
)

func AdminRoute(router *gin.Engine) {

	admin := router.Group("")
	{
		admin.POST("/createproduct", controllers.AdminCreateProduct)
		admin.GET("/allproducts", controllers.GetAllProducts)
		admin.PATCH("/updateproduct/:id", controllers.UpdateProduct)
		admin.DELETE("/deleteproduct/:id", controllers.DeletProduct)
		admin.GET("/product/:id", controllers.GetProductByID)
		admin.GET("/", controllers.Ping)
	}
}
