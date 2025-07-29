package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/client/transport"
	"github.com/mark3labs/mcp-go/mcp"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Create stdio transport
	stdioTransport := transport.NewStdio("go", nil, "run", "/home/author/go/go-mcp/main.go")
	// stdioTransport, err := transport.NewStreamableHTTP("http://localhost:8082/mcp")
	// if err != nil {
	// 	log.Fatalf("Failed to create stdio transport: %v", err)
	// }

	// Create client with the transport
	c := client.NewClient(stdioTransport)

	// Start the client
	if err := c.Start(ctx); err != nil {
		log.Fatalf("Failed to start client: %v", err)
	}
	defer c.Close()

	// Initialize the client
	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "Hello World Client",
		Version: "1.0.0",
	}
	initRequest.Params.Capabilities = mcp.ClientCapabilities{}

	serverInfo, err := c.Initialize(ctx, initRequest)
	if err != nil {
		log.Fatalf("Failed to initialize: %v", err)
	}

	fmt.Printf("Connected to server: %s (version %s)\n",
		serverInfo.ServerInfo.Name,
		serverInfo.ServerInfo.Version)

	// List available tools
	if serverInfo.Capabilities.Tools != nil {
		toolsRequest := mcp.ListToolsRequest{}
		toolsResult, err := c.ListTools(ctx, toolsRequest)
		if err != nil {
			log.Fatalf("Failed to list tools: %v", err)
		}

		fmt.Printf("Available tools: %d\n", len(toolsResult.Tools))
		for _, tool := range toolsResult.Tools {
			fmt.Printf("- %s: %s\n", tool.Name, tool.Description)
		}

		// Call a tool
		callRequest := mcp.CallToolRequest{}
		callRequest.Params.Name = "hello_world"
		callRequest.Params.Arguments = map[string]interface{}{
			"name": "John Doe",
		}

		result, err := c.CallTool(ctx, callRequest)
		if err != nil {
			log.Fatalf("Failed to call tool: %v", err)
		}

		// Print the result
		for _, content := range result.Content {
			if textContent, ok := content.(mcp.TextContent); ok {
				fmt.Printf("Result: %s\n", textContent.Text)
			}
		}
	}
}
