package gamma

import (
	"context"
	"encoding/json"
	"fmt"

	polymarketgamma "github.com/ivanzzeth/polymarket-go-gamma-client"
	"github.com/ivanzzeth/polymarket-go-mcp/client"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetRelatedTagsByIDTool returns the MCP tool definition for getting related tags by ID
func GetRelatedTagsByIDTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_related_tags_by_id",
		Description: "Get related tags by tag ID",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"tag_id": map[string]any{
					"type":        "string",
					"description": "The tag ID to get related tags for",
				},
				"limit": map[string]any{
					"type":        "number",
					"description": "Number of related tags to return",
				},
				"offset": map[string]any{
					"type":        "number",
					"description": "Number of related tags to skip",
				},
			},
			"required": []string{"tag_id"},
		},
	}
}

// GetRelatedTagsByIDHandler handles the get_related_tags_by_id tool execution
func GetRelatedTagsByIDHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]any) (*mcp.CallToolResult, any, error) {
	// Parse arguments
	tagID, ok := args["tag_id"].(string)
	if !ok {
		return nil, nil, fmt.Errorf("tag_id parameter is required and must be a string")
	}

	// Parse query parameters if provided (excluding tag_id)
	var params *polymarketgamma.GetRelatedTagsParams
	if len(args) > 1 {
		// Create a copy of args without tag_id
		queryArgs := make(map[string]any)
		for k, v := range args {
			if k != "tag_id" {
				queryArgs[k] = v
			}
		}
		if len(queryArgs) > 0 {
			argsBytes, _ := json.Marshal(queryArgs)
			if err := json.Unmarshal(argsBytes, &params); err != nil {
				return nil, nil, fmt.Errorf("failed to parse query parameters: %w", err)
			}
		} else {
			params = &polymarketgamma.GetRelatedTagsParams{}
		}
	} else {
		params = &polymarketgamma.GetRelatedTagsParams{}
	}

	// Get client
	gammaClient := client.GetGammaClient()

	// Get related tags by ID
	relatedTags, err := gammaClient.GetRelatedTagsByID(ctx, tagID, params)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get related tags by ID: %w", err)
	}

	// Format response
	relatedTagsJSON, err := json.MarshalIndent(relatedTags, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal related tags: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(relatedTagsJSON)},
		},
	}, relatedTags, nil
}

// GetRelatedTagsBySlugTool returns the MCP tool definition for getting related tags by slug
func GetRelatedTagsBySlugTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_related_tags_by_slug",
		Description: "Get related tags by tag slug",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"slug": map[string]any{
					"type":        "string",
					"description": "The tag slug to get related tags for",
				},
				"limit": map[string]any{
					"type":        "number",
					"description": "Number of related tags to return",
				},
				"offset": map[string]any{
					"type":        "number",
					"description": "Number of related tags to skip",
				},
			},
			"required": []string{"slug"},
		},
	}
}

// GetRelatedTagsBySlugHandler handles the get_related_tags_by_slug tool execution
func GetRelatedTagsBySlugHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]any) (*mcp.CallToolResult, any, error) {
	// Parse arguments
	slug, ok := args["slug"].(string)
	if !ok {
		return nil, nil, fmt.Errorf("slug parameter is required and must be a string")
	}

	// Parse query parameters if provided (excluding slug)
	var params *polymarketgamma.GetRelatedTagsParams
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
			params = &polymarketgamma.GetRelatedTagsParams{}
		}
	} else {
		params = &polymarketgamma.GetRelatedTagsParams{}
	}

	// Get client
	gammaClient := client.GetGammaClient()

	// Get related tags by slug
	relatedTags, err := gammaClient.GetRelatedTagsBySlug(ctx, slug, params)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get related tags by slug: %w", err)
	}

	// Format response
	relatedTagsJSON, err := json.MarshalIndent(relatedTags, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal related tags: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(relatedTagsJSON)},
		},
	}, relatedTags, nil
}

// GetRelatedTagsDetailByIDTool returns the MCP tool definition for getting detailed related tags by ID
func GetRelatedTagsDetailByIDTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_related_tags_detail_by_id",
		Description: "Get detailed tag information for related tags by ID",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"tag_id": map[string]any{
					"type":        "string",
					"description": "The tag ID to get detailed related tags for",
				},
				"limit": map[string]any{
					"type":        "number",
					"description": "Number of detailed related tags to return",
				},
				"offset": map[string]any{
					"type":        "number",
					"description": "Number of detailed related tags to skip",
				},
			},
			"required": []string{"tag_id"},
		},
	}
}

// GetRelatedTagsDetailByIDHandler handles the get_related_tags_detail_by_id tool execution
func GetRelatedTagsDetailByIDHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]any) (*mcp.CallToolResult, any, error) {
	// Parse arguments
	tagID, ok := args["tag_id"].(string)
	if !ok {
		return nil, nil, fmt.Errorf("tag_id parameter is required and must be a string")
	}

	// Parse query parameters if provided (excluding tag_id)
	var params *polymarketgamma.GetRelatedTagsParams
	if len(args) > 1 {
		// Create a copy of args without tag_id
		queryArgs := make(map[string]any)
		for k, v := range args {
			if k != "tag_id" {
				queryArgs[k] = v
			}
		}
		if len(queryArgs) > 0 {
			argsBytes, _ := json.Marshal(queryArgs)
			if err := json.Unmarshal(argsBytes, &params); err != nil {
				return nil, nil, fmt.Errorf("failed to parse query parameters: %w", err)
			}
		} else {
			params = &polymarketgamma.GetRelatedTagsParams{}
		}
	} else {
		params = &polymarketgamma.GetRelatedTagsParams{}
	}

	// Get client
	gammaClient := client.GetGammaClient()

	// Get detailed related tags by ID
	relatedTags, err := gammaClient.GetRelatedTagsDetailByID(ctx, tagID, params)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get detailed related tags by ID: %w", err)
	}

	// Format response
	relatedTagsJSON, err := json.MarshalIndent(relatedTags, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal detailed related tags: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(relatedTagsJSON)},
		},
	}, relatedTags, nil
}

// GetRelatedTagsDetailBySlugTool returns the MCP tool definition for getting detailed related tags by slug
func GetRelatedTagsDetailBySlugTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_related_tags_detail_by_slug",
		Description: "Get detailed tag information for related tags by slug",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"slug": map[string]any{
					"type":        "string",
					"description": "The tag slug to get detailed related tags for",
				},
				"limit": map[string]any{
					"type":        "number",
					"description": "Number of detailed related tags to return",
				},
				"offset": map[string]any{
					"type":        "number",
					"description": "Number of detailed related tags to skip",
				},
			},
			"required": []string{"slug"},
		},
	}
}

// GetRelatedTagsDetailBySlugHandler handles the get_related_tags_detail_by_slug tool execution
func GetRelatedTagsDetailBySlugHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]any) (*mcp.CallToolResult, any, error) {
	// Parse arguments
	slug, ok := args["slug"].(string)
	if !ok {
		return nil, nil, fmt.Errorf("slug parameter is required and must be a string")
	}

	// Parse query parameters if provided (excluding slug)
	var params *polymarketgamma.GetRelatedTagsParams
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
			params = &polymarketgamma.GetRelatedTagsParams{}
		}
	} else {
		params = &polymarketgamma.GetRelatedTagsParams{}
	}

	// Get client
	gammaClient := client.GetGammaClient()

	// Get detailed related tags by slug
	relatedTags, err := gammaClient.GetRelatedTagsDetailBySlug(ctx, slug, params)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get detailed related tags by slug: %w", err)
	}

	// Format response
	relatedTagsJSON, err := json.MarshalIndent(relatedTags, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal detailed related tags: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(relatedTagsJSON)},
		},
	}, relatedTags, nil
}
