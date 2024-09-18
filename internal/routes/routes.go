package routes

import (
	"TextWizard/internal/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置CORS头
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// 处理 OPTIONS 请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}

// InitializeRoutes sets up all the routes in one place
func InitializeRoutes(router *gin.Engine) {
	router.Use(CORSMiddleware())

	router.GET("/test", handlers.Test)
	router.POST("/summarize", handlers.SummarizeTextHandler)
}
