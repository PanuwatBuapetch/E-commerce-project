package main

import (
	"ecommerce-backend/configs"
	"ecommerce-backend/controllers"
	"ecommerce-backend/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.ConnectDB()
	r := gin.Default()

	// ✅ 2. เรียกใช้ CORS Middleware ตรงนี้ (ก่อนกำหนด Routes)
	r.Use(middleware.CORSMiddleware())

	// Routes ต่างๆ
	r.GET("/api/products", controllers.GetProducts)
	r.POST("/api/orders", controllers.CreateOrder)

	r.Run(":8080")
}
