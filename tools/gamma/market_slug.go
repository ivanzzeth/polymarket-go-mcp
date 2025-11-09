package gamma

import (
	"context"
	"encoding/json"
	"fmt"

	polymarketgamma "github.com/ivanzzeth/polymarket-go-gamma-client"
	"github.com/ivanzzeth/polymarket-go-mcp/client"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetMarketBySlugTool returns the MCP tool definition for getting a market by slug
func GetMarketBySlugTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_market_by_slug",
		Description: "Get a single Polymarket market by its slug",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"slug": map[string]any{
					"type":        "string",
					"description": "The market slug to retrieve",
				},
			},
			"required": []string{"slug"},
		},
	}
}

// GetMarketBySlugHandler handles the get_market_by_slug tool execution
func GetMarketBySlugHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]any) (*mcp.CallToolResult, any, error) {
	// Parse arguments
	slug, ok := args["slug"].(string)
	if !ok {
		return nil, nil, fmt.Errorf("slug parameter is required and must be a string")
	}

	// Parse query parameters if provided (excluding slug)
	var params *polymarketgamma.GetMarketByIDQueryParams
	if len(args) > 1 {
		// Create a copy of args without slug
		queryArgs := make(map[string]any)
		for k, v := range args {
			if k != "slug" {
				queryArgs[k] = v
			}
		}
		if len(queryArgs) > 0 {
			argsBytes, _ := json.Marshal(queryArgs)
			if err := json.Unmarshal(argsBytes, &params); err != nil {
				return nil, nil, fmt.Errorf("failed to parse query parameters: %w", err)
			}
		} else {
			params = &polymarketgamma.GetMarketByIDQueryParams{}
		}
	} else {
		params = &polymarketgamma.GetMarketByIDQueryParams{}
	}

	// Get client
	gammaClient := client.GetGammaClient()

	// Get market by slug
	market, err := gammaClient.GetMarketBySlug(ctx, slug, params)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get market by slug: %w", err)
	}

	// Format response
	marketJSON, err := json.MarshalIndent(market, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal market: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(marketJSON)},
		},
	}, market, nil
}
