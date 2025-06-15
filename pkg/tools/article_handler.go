package tools

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
	"net/url"

	"github.com/mark3labs/mcp-go/mcp"
)

// ArxivResponse represents the XML structure from arXiv API
type ArxivResponse struct {
	XMLName xml.Name `xml:"feed"`
	Entries []Entry  `xml:"entry"`
}

// Entry represents a single arXiv article entry
type Entry struct {
	Title     string   `xml:"title" json:"title"`
	ID        string   `xml:"id" json:"id"`
	Published string   `xml:"published" json:"published"`
	Updated   string   `xml:"updated" json:"updated"`
	Summary   string   `xml:"summary" json:"summary"`
	Authors   []Author `xml:"author" json:"authors"`
}

// Author represents an article author
type Author struct {
	Name string `xml:"name" json:"name"`
}

func GetNewArticlesToolHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	if req.Params.Arguments == nil {
		return mcp.NewToolResultError("No arguments provided. The user needs to provide a topic"), nil
	}

	args := req.Params.Arguments.(map[string]interface{})
	keyword, ok := args["keyword"].(string)
	if !ok {
		return mcp.NewToolResultError("No topic provided. The user needs to provide a topic"), nil
	}
	query := "all:" + keyword

	arxivURL := "http://export.arxiv.org/api/query"
	params := url.Values{}
	params.Add("search_query", query)
	params.Add("start", "0")
	params.Add("max_results", "10")
	params.Add("sortBy", "submittedDate")
	params.Add("sortOrder", "descending")

	response, err := http.Get(arxivURL + "?" + params.Encode())
	if err != nil {
		return mcp.NewToolResultError("Error fetching articles from arXiv: " + err.Error()), nil
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return mcp.NewToolResultError("Error reading response body: " + err.Error()), nil
	}

	// Parse XML response
	var arxivResp ArxivResponse
	if err := xml.Unmarshal(body, &arxivResp); err != nil {
		return mcp.NewToolResultError("Error parsing XML response: " + err.Error()), nil
	}

	// Convert to JSON
	jsonData, err := json.MarshalIndent(arxivResp, "", "  ")
	if err != nil {
		return mcp.NewToolResultError("Error converting to JSON: " + err.Error()), nil
	}

	return mcp.NewToolResultText(string(jsonData)), nil
}
