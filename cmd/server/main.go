package main

import (
	"flag"
	"log"

	"github.com/mark3labs/mcp-go/server"

	"github.com/haarshitgarg/BART/pkg/tools"
)

func main() {
	sseMode := flag.Bool("sse", false, "Run server in SSE mode")
	httpMode:= flag.Bool("http", false, "Run server in HTTP mode")

	flag.Parse()

	if *sseMode && *httpMode {
		log.Printf("Cannot set both sse and http mode together. Choose one of http or sse")
		return
	}
	// Create a new MCP server instance
	s := server.NewMCPServer(
		"BART Server",
		"1.0.0",
		server.WithToolCapabilities(true),
	)

	if err := tools.RegisterTools(s); err != nil {
		log.Fatalf("Failed to register tools: %v", err)
	}

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
	} else if *httpMode {
		log.Printf("Starting HTTP server on :8080")
		httpServer := server.NewStreamableHTTPServer(s,
			server.WithEndpointPath("/mcp"),
		)
		if err := httpServer.Start(":8080"); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
		log.Printf("HTTP server started on http://localhost:8080/mcp")
	} else {
		log.Printf("Starting stdio server")
		if err := server.ServeStdio(s); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}
}
