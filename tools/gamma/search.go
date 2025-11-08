package gamma

import (
	"context"
	"encoding/json"
	"fmt"

	polymarketgamma "github.com/ivanzzeth/polymarket-go-gamma-client"
	"github.com/ivanzzeth/polymarket-go-mcp/client"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// SearchTool returns the MCP tool definition for searching
func SearchTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "search",
		Description: "Search Polymarket content with comprehensive search options",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"query": map[string]any{
					"type":        "string",
					"description": "Search query string",
				},
				"limit": map[string]any{
					"type":        "number",
					"description": "Number of results to return",
				},
				"offset": map[string]any{
					"type":        "number",
					"description": "Number of results to skip",
				},
				"type": map[string]any{
					"type":        "string",
					"description": "Type of content to search (market, event, series, tag)",
				},
				"tag_id": map[string]any{
					"type":        "number",
					"description": "Filter by tag ID",
				},
				"closed": map[string]any{
					"type":        "boolean",
					"description": "Include closed content",
				},
			},
			"required": []string{"query"},
		},
	}
}

// SearchHandler handles the search tool execution
func SearchHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]any) (*mcp.CallToolResult, any, error) {
	// Parse arguments directly into the library type
	var params polymarketgamma.SearchParams
	argsBytes, _ := json.Marshal(args)
	if err := json.Unmarshal(argsBytes, &params); err != nil {
		return nil, nil, fmt.Errorf("failed to parse arguments: %w", err)
	}

	// Get client
	gammaClient := client.GetGammaClient()

	searchResult, err := gammaClient.Search(ctx, &params)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to search: %w", err)
	}

	// Format response
	searchJSON, err := json.MarshalIndent(searchResult, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal search results: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(searchJSON)},
		},
	}, searchResult, nil
}
