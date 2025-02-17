package main

import (
	"bufio"
	"fmt"
	"log"
	"mcpTermianl/client"
	"mcpTermianl/prompts"
	"mcpTermianl/server"
	"mcpTermianl/tool"
	"os"
)

func startConversationLoop(apiKey string, apiUrl string, model string) {
	// Initialize MCP servers
	mcpServers, err := server.NewTool()
	if err != nil {
		log.Fatalln(err)
	}
	prompt := prompts.GetPrompts(mcpServers.GetMCPServerStringList())
	chatClient := client.NewChatClient(apiKey, apiUrl, model, prompt)

	// Start the conversation loop
	scanner := bufio.NewScanner(os.Stdin)
	for {
		userInput := getUserInput(scanner)
		response, err := chatClient.SendMessage(userInput)
		if err != nil {
			log.Fatalf("Error sending message: %v", err)
		}
		// Process the response to check if it contains a ToolRequest
		if tool.HandleToolRequest(response, mcpServers, chatClient) {
			continue // If a ToolRequest is handled, continue to the next iteration
		}
		fmt.Println("Assistant:", response)
	}
}

// getUserInput function to retrieve user input
func getUserInput(scanner *bufio.Scanner) string {
	fmt.Print("You: ")
	scanner.Scan()
	return scanner.Text()
}

func main() {
	apiKey := ""
	apiUrl := "https://aiproxy.hzh.sealos.run/v1/chat/completions"
	model := "deepseek-chat"
	startConversationLoop(apiKey, apiUrl, model)
}
