package data

import (
	"context"
	"encoding/json"
	"fmt"

	polymarketdata "github.com/ivanzzeth/polymarket-go-data-client"
	"github.com/ivanzzeth/polymarket-go-mcp/client"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetActivityTool returns the MCP tool definition for getting activity
func GetActivityTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_activity",
		Description: "Get Polymarket activity data with filtering options",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"limit": map[string]any{
					"type":        "number",
					"description": "Number of activities to return",
				},
				"offset": map[string]any{
					"type":        "number",
					"description": "Number of activities to skip",
				},
				"sort_by": map[string]any{
					"type":        "string",
					"description": "Field to sort by (TIMESTAMP)",
				},
				"sort_direction": map[string]any{
					"type":        "string",
					"description": "Sort direction (ASC, DESC)",
				},
				"user": map[string]any{
					"type":        "string",
					"description": "Filter by user address (Required: User Profile Address 0x-prefixed, 40 hex chars)",
				},
				"market": map[string]any{
					"type":        "string",
					"description": "Filter by market ID",
				},
				"type": map[string]any{
					"type":        "string",
					"description": "Filter by activity type (TRADE, etc.)",
				},
				"start_date": map[string]any{
					"type":        "string",
					"description": "Start date filter (ISO 8601)",
				},
				"end_date": map[string]any{
					"type":        "string",
					"description": "End date filter (ISO 8601)",
				},
			},
			"required": []string{"user"},
		},
	}
}

// GetActivityHandler handles the get_activity tool execution
func GetActivityHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]any) (*mcp.CallToolResult, any, error) {
	// Create a copy of args without the "type" field to avoid unmarshaling issues
	filteredArgs := make(map[string]any)
	for k, v := range args {
		if k != "type" {
			filteredArgs[k] = v
		}
	}

	// Parse arguments directly into the library type
	var params polymarketdata.GetActivityParams
	argsBytes, _ := json.Marshal(filteredArgs)
	if err := json.Unmarshal(argsBytes, &params); err != nil {
		return nil, nil, fmt.Errorf("failed to parse arguments: %w", err)
	}

	// Get client
	dataClient := client.GetDataClient()

	activity, err := dataClient.GetActivity(ctx, &params)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get activity: %w", err)
	}

	// Format response
	activityJSON, err := json.MarshalIndent(activity, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal activity: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(activityJSON)},
		},
	}, activity, nil
}
