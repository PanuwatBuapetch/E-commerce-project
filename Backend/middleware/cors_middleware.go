package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware อนุญาตให้ Frontend ยิง API ข้าม Port ได้
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// อนุญาตทุก Origin (*) หรือจะระบุเจาะจงเป็น "http://localhost:3000" ก็ได้เพื่อความปลอดภัย
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		// ถ้าเป็น Request แบบ OPTIONS (Pre-flight) ให้ตอบกลับทันทีว่า "อนุญาตนะ"
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
