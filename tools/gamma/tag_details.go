package gamma

import (
	"context"
	"encoding/json"
	"fmt"

	polymarketgamma "github.com/ivanzzeth/polymarket-go-gamma-client"
	"github.com/ivanzzeth/polymarket-go-mcp/client"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetTagByIDTool returns the MCP tool definition for getting a tag by ID
func GetTagByIDTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_tag_by_id",
		Description: "Get a single Polymarket tag by its ID",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"id": map[string]any{
					"type":        "string",
					"description": "The tag ID to retrieve",
				},
			},
			"required": []string{"id"},
		},
	}
}

// GetTagByIDHandler handles the get_tag_by_id tool execution
func GetTagByIDHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]any) (*mcp.CallToolResult, any, error) {
	// Parse arguments
	id, ok := args["id"].(string)
	if !ok {
		return nil, nil, fmt.Errorf("id parameter is required and must be a string")
	}

	// Parse query parameters if provided (excluding id)
	var params *polymarketgamma.GetTagByIDQueryParams
	if len(args) > 1 {
		// Create a copy of args without id
		queryArgs := make(map[string]any)
		for k, v := range args {
			if k != "id" {
				queryArgs[k] = v
			}
		}
		if len(queryArgs) > 0 {
			argsBytes, _ := json.Marshal(queryArgs)
			if err := json.Unmarshal(argsBytes, &params); err != nil {
				return nil, nil, fmt.Errorf("failed to parse query parameters: %w", err)
			}
		} else {
			params = &polymarketgamma.GetTagByIDQueryParams{}
		}
	} else {
		params = &polymarketgamma.GetTagByIDQueryParams{}
	}

	// Get client
	gammaClient := client.GetGammaClient()

	// Get tag by ID
	tag, err := gammaClient.GetTagByID(ctx, id, params)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get tag by ID: %w", err)
	}

	// Format response
	tagJSON, err := json.MarshalIndent(tag, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal tag: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(tagJSON)},
		},
	}, tag, nil
}

// GetTagBySlugTool returns the MCP tool definition for getting a tag by slug
func GetTagBySlugTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_tag_by_slug",
		Description: "Get a single Polymarket tag by its slug",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"slug": map[string]any{
					"type":        "string",
					"description": "The tag slug to retrieve",
				},
			},
			"required": []string{"slug"},
		},
	}
}

// GetTagBySlugHandler handles the get_tag_by_slug tool execution
func GetTagBySlugHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]any) (*mcp.CallToolResult, any, error) {
	// Parse arguments
	slug, ok := args["slug"].(string)
	if !ok {
		return nil, nil, fmt.Errorf("slug parameter is required and must be a string")
	}

	// Parse query parameters if provided (excluding slug)
	var params *polymarketgamma.GetTagBySlugQueryParams
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
			params = &polymarketgamma.GetTagBySlugQueryParams{}
		}
	} else {
		params = &polymarketgamma.GetTagBySlugQueryParams{}
	}

	// Get client
	gammaClient := client.GetGammaClient()

	// Get tag by slug
	tag, err := gammaClient.GetTagBySlug(ctx, slug, params)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get tag by slug: %w", err)
	}

	// Format response
	tagJSON, err := json.MarshalIndent(tag, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal tag: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(tagJSON)},
		},
	}, tag, nil
}
