package main

import (
	"log"

	"github.com/harshitgarg/mcp-server/pkg/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create a new MCP server instance
	s := server.NewMCPServer(
		"BART Server",
		"1.0.0",
		server.WithToolCapabilities(true),
	)
	err := mcp.RegisterTools(s)
	if err != nil {
		log.Fatalf("Failed to register tools: %v", err)
	}

	sseServer := server.NewSSEServer(s,
		server.WithBaseURL("http://localhost:8080"),
		server.WithStaticBasePath("/mcp"),
	)

	// Start HTTP server on port 8080
	log.Printf("Starting server on :8080")
	if err := sseServer.Start(":8080"); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
