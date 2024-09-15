package handlers

import (
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	router.GET("/test", Test)
}

func Test(c *gin.Context) {
	c.String(200, "TextWizard is running")
}
