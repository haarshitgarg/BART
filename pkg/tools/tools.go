package tools

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterTools registers all MCP tools with the server
func RegisterTools(s *server.MCPServer) error {

	// GET NEW BOOKS TOOL
	tool := mcp.NewTool(
		"get_new_books",
		mcp.WithDescription("Get latest books published in the last 30 days"),
		mcp.WithString(
			"genre",
			mcp.Description("The genre of books to the user is interested in"),
			mcp.Required(),
		),
	)
	s.AddTool(tool, GetNewBooksToolHandler)

	// GET NEW ARTICLES TOOL
	tool = mcp.NewTool(
		"get_new_articles_by_keyword",
		mcp.WithDescription("Get latest articles published by keyword"),
		mcp.WithString(
			"keyword",
			mcp.Description("The keyword of the articles to the user is interested in"),
			mcp.Required(),
		),
	)
	s.AddTool(tool, GetNewArticlesToolHandler)

	return nil
} 