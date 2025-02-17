package tool

import (
	"encoding/json"
	"fmt"
	"log"
	"mcpTermianl/client"
	"mcpTermianl/req"
	"mcpTermianl/server"
)

// ToolRequest struct to represent the "server" and "arguments" in the JSON response
type ToolRequest struct {
	Name      string                 `json:"name"`
	Arguments map[string]interface{} `json:"arguments"`
}

// HandleToolRequest function to process and check if the response is a ToolRequest type
func HandleToolRequest(response string, mcpServers *server.Tool, chatClient *client.ChatClient) bool {
	var request ToolRequest
	// Try to unmarshal the response into a ToolRequest struct
	err := json.Unmarshal([]byte(response), &request)
	if err != nil {
		return false // If it's not a ToolRequest, return false to continue normal processing
	}

	// Search for the corresponding mcpServer and send a request to it
	for _, mcpServer := range mcpServers.Server {
		if request.Name == mcpServer.Name {
			request, _ := req.BuildRequest(mcpServer.Host, request.Arguments)
			response, err := chatClient.SendMessage(request)
			if err != nil {
				log.Fatalf("Error sending message: %v", err)
			}
			fmt.Println("Assistant:", response)
			return true
		}
	}
	return false
}
