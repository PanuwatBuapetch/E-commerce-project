package controllers

import (
	"ecommerce-backend/models"
	"ecommerce-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST /api/orders
func CreateOrder(c *gin.Context) {
	var input models.Order

	// 1. รับค่า JSON จากหน้าบ้าน (total_amount, user_id)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. ส่งให้ Service บันทึกลง Database
	orderID, err := services.CreateOrder(input.TotalAmount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order: " + err.Error()})
		return
	}

	// 3. ส่งผลลัพธ์กลับไปให้หน้าบ้าน (Next.js)
	c.JSON(http.StatusCreated, gin.H{
		"message":  "Order created successfully",
		"order_id": orderID,
	})
}
