package main

import (
	"os"

	"git.com/ecommerce/initializer"
	"git.com/ecommerce/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	initializer.DatabaseConnection()

	r := gin.Default()
	routes.AdminRoute(r)
	r.Run(":8080" + os.Getenv("PORT"))
}
