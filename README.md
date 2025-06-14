# MCP Server

A Model Context Protocol (MCP) server implementation in Go.

## Project Structure

```
.
├── main.go           # Main server entry point
├── go.mod           # Go module definition
├── internal/        # Private application code
│   └── mcp/        # MCP protocol implementation
│       └── server.go
└── README.md        # This file
```

## Setup

1. Make sure you have Go 1.21 or later installed
2. Clone this repository
3. Run the server:
   ```bash
   go run main.go
   ```

The server will start on port 8080.

## Development

This is a work in progress. The current implementation provides a basic HTTP server structure that will be extended to support the MCP protocol.
