package initializer

import (
	"log"
	"os"

	"git.com/ecommerce/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnection() {
	var err error
	DB, err = gorm.Open(sqlite.Open("ecommerce.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database \n", err.Error())
		os.Exit(2)
	}
	err = DB.AutoMigrate(&models.Products{})
	log.Println("Connected to database successfully")

}
