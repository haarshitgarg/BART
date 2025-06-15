package mcp

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterTools registers all MCP tools with the server
func RegisterTools(s *server.MCPServer) error {
	// Example tool registration
	exampleTool := mcp.NewTool(
		"example",
		mcp.WithDescription("An example tool"),
	)

	// Create the tool handler
	handler := func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return mcp.NewToolResultText("Example tool response"), nil
	}

	// Add the tool to the server
	s.AddTool(exampleTool, handler)
	return nil
} 