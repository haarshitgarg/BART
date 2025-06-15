package tools

import (
	"context"
	"io"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
)

func GetNewBooksToolHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	if req.Params.Arguments == nil {
		return mcp.NewToolResultError("No arguments provided. The user needs to provide a genre"), nil
	}
	args := req.Params.Arguments.(map[string]interface{})
	genre, ok := args["genre"].(string)
	if !ok {
		return mcp.NewToolResultError("No genre provided. The user needs to provide a genre"), nil
	}

	openLibraryURL := "https://openlibrary.org/subjects/" + genre + ".json?limit=10&sort=new"

	response, err := http.Get(openLibraryURL)
	if err != nil {
		return mcp.NewToolResultError("Error fetching books from Open Library: " + err.Error()), nil
	}
	// defer is a statement that delays the execution of a function until the surrounding function returns
	// In this case, response.Body.Close() will be called when GetNewBooksToolHandler() finishes executing
	// This ensures that we properly close the response body and free up system resources
	// even if an error occurs later in the function
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return mcp.NewToolResultError("Error reading response body: " + err.Error()), nil
	}

	return mcp.NewToolResultText(string(body)), nil
}
