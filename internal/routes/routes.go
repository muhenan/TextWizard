package routes

import (
	"TextWizard/internal/handlers"

	"github.com/gin-gonic/gin"
)

// InitializeRoutes sets up all the routes in one place
func InitializeRoutes(router *gin.Engine) {
	router.GET("/test", handlers.Test)
	router.POST("/summarize", handlers.SummarizeTextHandler)
}
