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

	// Configuring CORS middleware
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:3000"} // Replace with your Next.js frontend URL
	// config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	// r.Use(cors.New(config))

	routes.AdminRoute(r)
	routes.UserRouter(r)
	r.Run(":8080" + os.Getenv("PORT"))
}
