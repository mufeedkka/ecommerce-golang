package controllers

import (
	"fmt"
	"net/http"

	"git.com/ecommerce/initializer"
	"git.com/ecommerce/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProductStruct struct {
	ShortName          string  `json:"shortName"`
	LongName           string  `json:"longName"`
	Cost               int     `json:"cost"`
	Price              int     `json:"price"`
	DiscountType       *string `json:"discountType"`
	DiscountPrice      *int    `json:"discountPrice"`
	Description        string  `json:"description"`
	Stock              int     `json:"stock"`
	DealerName         string  `json:"dealerName"`
	DealerPlace        string  `json:"dealerPlace"`
	ProductDestination string  `json:"productDestination"`
}

func AdminCreateProduct(c *gin.Context) {

	var product models.Products

	if err := c.ShouldBind(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	validate := validator.New()

	if err := validate.Struct(product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	result := initializer.DB.Create(&product)

	if result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to add product",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"product": product,
	})

}

func Ping(c *gin.Context) {
	c.JSON(200, "pong")
}
