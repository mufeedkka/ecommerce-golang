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
	fmt.Println(&product, "test")
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

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var body ProductStruct

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	var product models.Products

	result := initializer.DB.Find(&product, "id = ?", id)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "No data found",
		})
		c.Abort()
		return
	}
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": result.Error,
		})
		c.Abort()
		return
	}
	product.ShortName = body.ShortName
	product.LongName = body.LongName

	initializer.DB.Save(&product)

	c.JSON(http.StatusOK, gin.H{
		"update":  true,
		"product": product,
	})
}

func GetAllProducts(c *gin.Context) {
	var product []models.Products

	result := initializer.DB.Find(&product)

	if result.Error != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": result.Error,
		})
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "No data found",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"products": product,
	})
}

func DeletProduct(c *gin.Context) {
	id := c.Param("id")
	var product []models.Products

	initializer.DB.Delete(&product, id)

	c.JSON(http.StatusOK, gin.H{
		"delete": true,
		// "Product": product,
	})
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")

	var product models.Products

	result := initializer.DB.Find(&product, "id = ?", id)

	if result.Error != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": result.Error,
		})
		c.Abort()
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "No data found",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"Product": product,
	})
}

func Ping(c *gin.Context) {
	c.JSON(200, "pong")
}
