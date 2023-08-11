package controllers

import (
	"fmt"
	"net/http"

	"git.com/ecommerce/initializer"
	"git.com/ecommerce/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func UserSignUp(c *gin.Context) {
	fmt.Println("kkkaaaaaaaa")
	var user models.Users

	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	// Validation of req.body.
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errArray []string
		for _, e := range validationErrors {
			field := e.Field()
			tag := e.Tag()
			if tag == "required" {
				errArray = append(errArray, field+" "+tag)
			} else {
				errArray = append(errArray, field+" should be "+tag)
			}
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errArray[0],
		})
		c.Abort()
		return
	}

	// Hash Passowrd.
	if err := user.HashPassword(user.Password); err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	// Create User
	result := initializer.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to create user. Try again!!",
		})
		c.Abort()
		return
	}

	// go utils.SendEmailWithoutHTML([]string{user.Email}, "Successfull Registered!!", "Thankyou for being our partner. Let's make some collaborations")

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "User created successfully",
	})
}
