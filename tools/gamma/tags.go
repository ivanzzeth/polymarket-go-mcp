package gamma

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ivanzzeth/polymarket-go-mcp/client"
	polymarketgamma "github.com/ivanzzeth/polymarket-go-gamma-client"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetTagsTool returns the MCP tool definition for getting tags
func GetTagsTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_tags",
		Description: "Get Polymarket tags with filtering options",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"limit": map[string]any{
					"type":        "number",
					"description": "Number of tags to return",
				},
				"offset": map[string]any{
					"type":        "number",
					"description": "Number of tags to skip",
				},
				"id": map[string]any{
					"type":        "array",
					"items":       map[string]any{"type": "number"},
					"description": "Filter by tag IDs",
				},
				"slug": map[string]any{
					"type":        "array",
					"items":       map[string]any{"type": "string"},
					"description": "Filter by tag slugs",
				},
			},
		},
	}
}

// GetTagsHandler handles the get_tags tool execution
func GetTagsHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]any) (*mcp.CallToolResult, any, error) {
	// Parse arguments directly into the library type
	var params polymarketgamma.GetTagsParams
	argsBytes, _ := json.Marshal(args)
	if err := json.Unmarshal(argsBytes, &params); err != nil {
		return nil, nil, fmt.Errorf("failed to parse arguments: %w", err)
	}

	// Get client
	gammaClient := client.GetGammaClient()

	// Get tags
	tags, err := gammaClient.GetTags(ctx, &params)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get tags: %w", err)
	}

	// Format response
	tagsJSON, err := json.MarshalIndent(tags, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal tags: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(tagsJSON)},
		},
	}, tags, nil
}

