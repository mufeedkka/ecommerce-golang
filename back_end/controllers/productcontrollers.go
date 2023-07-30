import controllers

func AdminGetAllProducts(c *gin.Context) {
	var products []models.Product
	err := db.Find(&products).Error
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch products"})
		return
	}

	c.JSON(200, products)
}