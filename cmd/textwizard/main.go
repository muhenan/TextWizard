package main

import (
	"TextWizard/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Load configurations
	// config.LoadConfig()

	// Initialize routes
	routes.InitializeRoutes(r)

	// Start the server on port 8080
	r.Run(":8080")
}
