package gamma

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ivanzzeth/polymarket-go-mcp/client"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetSportsMetadataTool returns the MCP tool definition for getting sports metadata
func GetSportsMetadataTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_sports_metadata",
		Description: "Get sports metadata including images and resolution sources",
		InputSchema: map[string]any{
			"type":       "object",
			"properties": map[string]any{},
		},
	}
}

// GetSportsMetadataHandler handles the get_sports_metadata tool execution
func GetSportsMetadataHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]any) (*mcp.CallToolResult, any, error) {
	// Get client
	gammaClient := client.GetGammaClient()

	sportsMetadata, err := gammaClient.GetSportsMetadata(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get sports metadata: %w", err)
	}

	// Format response
	sportsJSON, err := json.MarshalIndent(sportsMetadata, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal sports metadata: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(sportsJSON)},
		},
	}, sportsMetadata, nil
}
