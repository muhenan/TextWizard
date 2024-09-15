package main

import (
	"TextWizard/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Load configurations
	// config.LoadConfig()

	// Initialize routes
	handlers.InitializeRoutes(r)

	// Start the server on port 8080
	r.Run(":8080")
}
