package main

import (
	"log"
	"net/http"
)

func main() {
	// Initialize the MCP server
	log.Println("Starting MCP server...")

	// Basic HTTP server setup
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("MCP Server is running!"))
	})

	// Start the server
	port := ":8080"
	log.Printf("Server listening on port %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
} 