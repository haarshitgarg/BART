package main

import (
	"context"
	"log"
	"flag"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	sseMode := flag.Bool("sse", false, "Run server in SSE mode")
	flag.Parse()
	// Create a new MCP server instance
	s := server.NewMCPServer(
		"BART Server",
		"1.0.0",
		server.WithToolCapabilities(true),
	)

	tool := mcp.NewTool(
		"example",
		mcp.WithDescription("An example tool"),
	)
	exampleToolHandler := func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return mcp.NewToolResultText("Hello, world!"), nil
	}
	s.AddTool(tool, exampleToolHandler)

	// Start HTTP server on port 8080
	if *sseMode {
		log.Printf("Starting server on :8080")
		sseServer := server.NewSSEServer(s,
			server.WithBaseURL("http://localhost:8080"),
			server.WithStaticBasePath("/mcp"),
		)
		if err := sseServer.Start(":8080"); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	} else {
		if err := server.ServeStdio(s); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}
}
