package routes

import (
	"ecommerce-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/products", controllers.GetProducts)
		api.POST("/products", controllers.AddProduct)
		api.PUT("/products/:id", controllers.UpdateProduct)
		api.DELETE("/products/:id", controllers.DeleteProduct)
	}
}
