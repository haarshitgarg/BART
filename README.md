# BART Server

A MCP server for books and research papers

## Setup

1. Make sure you have Go 1.21 or later installed
2. Clone this repository
3. Run the server:
   ```bash
   go run cmd/server/main.go -sse

   OR

   go run cmd/server/main.go -http

   OR 

   go run cmd/server/main.go
   ```

The server will start on port 8080.

## Development

The following tools have been implemented:

1. **Get New Books Tool**
   - Fetches the latest books published in the last 30 days
   - Uses Open Library API
   - Requires a genre parameter to filter books
   - Returns up to 10 most recent books in the specified genre

2. **Get New Articles Tool**
   - Fetches the latest research articles from arXiv
   - Uses arXiv API
   - Requires a keyword parameter to search articles
   - Returns up to 10 most recent articles matching the keyword
   - Includes article title, authors, publication date, and summary

## Future Plans

1. **Enhanced Search Capabilities**
   - Add support for multiple keywords and genres
   - Implement advanced filtering options
   - Add date range selection for articles and books

2. **Additional Data Sources**
   - Integrate with more academic databases (e.g., Google Scholar, PubMed)
   - Add support for more book platforms (e.g., Google Books, Amazon)

3. **User Features**
   - Add user authentication
   - Implement bookmarking functionality
   - Add personal reading lists
   - Enable article/book recommendations

4. **API Improvements**
   - Add rate limiting
   - Implement caching for better performance
   - Add pagination support for large result sets

5. **Documentation**
   - Add API documentation
   - Create usage examples
   - Add contribution guidelines
