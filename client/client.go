package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatClient struct {
	Model    Model     `json:"model"`
	Messages []Message `json:"messages"`
}

type Model struct {
	ApiKey string `json:"api_key"`
	ApiUrl string `json:"api_url"`
	Name   string `json:"name"`
}

type RequestPayload struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	MaxTokens   int       `json:"max_tokens"`
	Temperature float64   `json:"temperature"`
}

type ResponseChoice struct {
	Message struct {
		Content string `json:"content"`
	} `json:"message"`
}

func NewChatClient(apiKey string, apiUrl string, model string, message string) *ChatClient {
	return &ChatClient{
		Model: Model{
			ApiKey: apiKey,
			ApiUrl: apiUrl,
			Name:   model,
		},
		Messages: []Message{{Role: "system", Content: message}},
	}
}

func (c *ChatClient) SendMessage(userMessage string) (string, error) {
	// Append the user's message to the conversation history
	c.Messages = append(c.Messages, Message{Role: "user", Content: userMessage})

	// Create the request payload
	requestBody := RequestPayload{
		Model:       c.Model.Name,
		Messages:    c.Messages,
		MaxTokens:   2048,
		Temperature: 0.7,
	}
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("JSON serialization failed: %w", err)
	}
	req, err := http.NewRequest("POST", c.Model.ApiUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Set request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Model.ApiKey))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}
	var response struct {
		Choices []ResponseChoice `json:"choices"`
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", fmt.Errorf("failed to parse response JSON: %w", err)
	}
	if len(response.Choices) > 0 {
		c.Messages = append(c.Messages, Message{Role: "assistant", Content: response.Choices[0].Message.Content})
		return response.Choices[0].Message.Content, nil
	}
	return "", fmt.Errorf("no response received")
}
