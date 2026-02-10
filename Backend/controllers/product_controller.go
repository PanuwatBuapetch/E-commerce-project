package controllers

import (
	"ecommerce-backend/models"
	"ecommerce-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /api/products
func GetProducts(c *gin.Context) {
	products, err := services.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if products == nil {
		products = []models.Product{}
	}
	c.JSON(http.StatusOK, products)
}

// POST /api/products
func AddProduct(c *gin.Context) {
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newProduct, err := services.CreateProduct(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newProduct)
}

// PUT /api/products/:id
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedProduct, err := services.UpdateProduct(id, input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, updatedProduct)
}

// DELETE /api/products/:id
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteProduct(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
