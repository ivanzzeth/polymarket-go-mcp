package data

import (
	"context"
	"encoding/json"
	"fmt"

	polymarketdata "github.com/ivanzzeth/polymarket-go-data-client"
	"github.com/ivanzzeth/polymarket-go-mcp/client"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetPositionsTool returns the MCP tool definition for getting positions
func GetPositionsTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_positions",
		Description: "Get Polymarket positions data with filtering options",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"limit": map[string]any{
					"type":        "number",
					"description": "Number of positions to return",
				},
				"offset": map[string]any{
					"type":        "number",
					"description": "Number of positions to skip",
				},
				"sort_by": map[string]any{
					"type":        "string",
					"description": "Field to sort by",
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
				"filter_type": map[string]any{
					"type":        "string",
					"description": "Filter by type (CASH, etc.)",
				},
			},
			"required": []string{"user"},
		},
	}
}

// GetPositionsHandler handles the get_positions tool execution
func GetPositionsHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]any) (*mcp.CallToolResult, any, error) {
	// Parse arguments directly into the library type
	var params polymarketdata.GetPositionsParams
	argsBytes, _ := json.Marshal(args)
	if err := json.Unmarshal(argsBytes, &params); err != nil {
		return nil, nil, fmt.Errorf("failed to parse arguments: %w", err)
	}

	// Get client
	dataClient := client.GetDataClient()

	positions, err := dataClient.GetPositions(ctx, &params)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get positions: %w", err)
	}

	// Format response
	positionsJSON, err := json.MarshalIndent(positions, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal positions: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(positionsJSON)},
		},
	}, positions, nil
}

// GetClosedPositionsTool returns the MCP tool definition for getting closed positions
func GetClosedPositionsTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_closed_positions",
		Description: "Get Polymarket closed positions data with filtering options",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"limit": map[string]any{
					"type":        "number",
					"description": "Number of closed positions to return",
				},
				"offset": map[string]any{
					"type":        "number",
					"description": "Number of closed positions to skip",
				},
				"sort_by": map[string]any{
					"type":        "string",
					"description": "Field to sort by (REALIZEDPNL, etc.)",
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
			},
			"required": []string{"user"},
		},
	}
}

// GetClosedPositionsHandler handles the get_closed_positions tool execution
func GetClosedPositionsHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]any) (*mcp.CallToolResult, any, error) {
	// Parse arguments directly into the library type
	var params polymarketdata.GetClosedPositionsParams
	argsBytes, _ := json.Marshal(args)
	if err := json.Unmarshal(argsBytes, &params); err != nil {
		return nil, nil, fmt.Errorf("failed to parse arguments: %w", err)
	}

	// Get client
	dataClient := client.GetDataClient()

	closedPositions, err := dataClient.GetClosedPositions(ctx, &params)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get closed positions: %w", err)
	}

	// Format response
	closedPositionsJSON, err := json.MarshalIndent(closedPositions, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal closed positions: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(closedPositionsJSON)},
		},
	}, closedPositions, nil
}

// GetPositionsValueTool returns the MCP tool definition for getting positions value
func GetPositionsValueTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_positions_value",
		Description: "Get Polymarket positions value data",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"limit": map[string]any{
					"type":        "number",
					"description": "Number of users to return",
				},
				"offset": map[string]any{
					"type":        "number",
					"description": "Number of users to skip",
				},
				"sort_by": map[string]any{
					"type":        "string",
					"description": "Field to sort by (CURRENT, etc.)",
				},
				"sort_direction": map[string]any{
					"type":        "string",
					"description": "Sort direction (ASC, DESC)",
				},
				"user": map[string]any{
					"type":        "string",
					"description": "Filter by user address (Required: User Profile Address 0x-prefixed, 40 hex chars)",
				},
			},
			"required": []string{"user"},
		},
	}
}

// GetPositionsValueHandler handles the get_positions_value tool execution
func GetPositionsValueHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]any) (*mcp.CallToolResult, any, error) {
	// Parse arguments directly into the library type
	var params polymarketdata.GetValueParams
	argsBytes, _ := json.Marshal(args)
	if err := json.Unmarshal(argsBytes, &params); err != nil {
		return nil, nil, fmt.Errorf("failed to parse arguments: %w", err)
	}

	// Get client
	dataClient := client.GetDataClient()

	positionsValue, err := dataClient.GetPositionsValue(ctx, &params)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get positions value: %w", err)
	}

	// Format response
	positionsValueJSON, err := json.MarshalIndent(positionsValue, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal positions value: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(positionsValueJSON)},
		},
	}, positionsValue, nil
}
