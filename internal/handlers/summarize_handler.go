package handlers

import (
	"TextWizard/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SummarizeTextHandler processes the incoming request for text summarization
func SummarizeTextHandler(c *gin.Context) {
	var requestBody struct {
		Text string `json:"text" binding:"required"`
		Note string `json:"note"`
	}

	// Parse the request body
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Call the summarization service
	summary, err := services.SummarizeText(requestBody.Text, requestBody.Note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the summary as a JSON response
	c.JSON(http.StatusOK, gin.H{"summary": summary})
}
