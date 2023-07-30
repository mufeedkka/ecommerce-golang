package main

import (
	// "back_end/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("ecommercedatabase.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect DB")
	}

	// db.AutoMigrate(&models.Products{})
}

func main() {
	r := gin.Default()

	// Define your routes here
	// Example route
	// r.GET("/users", getUsers)

	r.Run(":8081")
}
