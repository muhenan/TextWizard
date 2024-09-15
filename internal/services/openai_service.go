package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// ChatCompletionRequest is the structure of the request sent to the OpenAI API.
type ChatCompletionRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

// Message defines the structure of a message sent in the API request.
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatCompletionResponse represents the response from the OpenAI API.
type ChatCompletionResponse struct {
	Choices []Choice `json:"choices"`
}

// Choice is a part of the response from the OpenAI API.
type Choice struct {
	Message Message `json:"message"`
}

// SummarizeText sends a request to the OpenAI API to summarize a given text.
func SummarizeText(text string, note string) (string, error) {
	// Load the API key from the environment variable
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("API key is missing")
	}

	// Define the request body
	requestBody := ChatCompletionRequest{
		Model: "gpt-4o-mini", // Use the appropriate model
		Messages: []Message{
			{Role: "system", Content: "You are a helpful assistant specialized in summarizing text."},
			{Role: "user", Content: fmt.Sprintf("Please summarize the following text: \n%s \n note: %s", text, note)},
		},
	}

	// Marshal the request body into JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Parse the response
	var completionResponse ChatCompletionResponse
	if err := json.Unmarshal(body, &completionResponse); err != nil {
		return "", err
	}

	// Extract the summary from the response
	if len(completionResponse.Choices) > 0 {
		return completionResponse.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no summary returned")
}
