package server

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Args struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

type MCPServer struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Type        string `json:"type"`
	Host        string `json:"host"`
	Description string `json:"description"`
	Args        []Args `json:"args"`
	Response    string `json:"response"`
}

type Tool struct {
	Server []MCPServer `json:"server"`
}

func NewTool() (*Tool, error) {
	// Open the local server.json file
	file, err := os.Open("server.json")
	if err != nil {
		return nil, fmt.Errorf("failed to open server.json file: %v", err)
	}
	defer file.Close()

	// Create an instance of the Tool type
	var tool Tool

	// Decode the JSON file into the server struct
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tool)
	if err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %v", err)
	}

	// Return the decoded Tool object
	return &tool, nil
}

// GetMCPServerStringList takes a pointer to a *[]MCPServer object list and returns a concatenated string
func (t *Tool) GetMCPServerStringList() string {
	servers := &t.Server
	var result []string
	for _, server := range *servers {
		// Build the string representation of each MCPServer
		serverInfo := fmt.Sprintf("Name: %s, Version: %s, Type: %s, Host: %s, Description: %s, Response: %s",
			server.Name, server.Version, server.Type, server.Host, server.Description, server.Response)

		// Add Args information if available
		if len(server.Args) > 0 {
			var argsInfo []string
			for _, arg := range server.Args {
				argsInfo = append(argsInfo, fmt.Sprintf("Key: %s, Value: %s, Description: %s", arg.Key, arg.Value, arg.Description))
			}
			serverInfo = fmt.Sprintf("%s, Args: [%s]", serverInfo, strings.Join(argsInfo, "; "))
		}

		// Add the MCPServer's information to the result list
		result = append(result, serverInfo)
	}
	// Combine all MCPServer information into a single string
	return strings.Join(result, "\n")
}
