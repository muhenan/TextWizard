package handlers

import "github.com/gin-gonic/gin"

// Test is a simple health check endpoint
func Test(c *gin.Context) {
	c.String(200, "TextWizard is running")
}
