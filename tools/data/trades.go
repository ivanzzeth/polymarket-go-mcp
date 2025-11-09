package data

import (
	"context"
	"encoding/json"
	"fmt"

	polymarketdata "github.com/ivanzzeth/polymarket-go-data-client"
	"github.com/ivanzzeth/polymarket-go-mcp/client"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetTradesTool returns the MCP tool definition for getting trades
func GetTradesTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_trades",
		Description: "Get Polymarket trades data with filtering options",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"limit": map[string]any{
					"type":        "number",
					"description": "Number of trades to return",
				},
				"offset": map[string]any{
					"type":        "number",
					"description": "Number of trades to skip",
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
					"description": "Filter by user address",
				},
				"market": map[string]any{
					"type":        "string",
					"description": "Filter by market ID",
				},
				"side": map[string]any{
					"type":        "string",
					"description": "Filter by trade side (BUY, SELL)",
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
		},
	}
}

// GetTradesHandler handles the get_trades tool execution
func GetTradesHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]any) (*mcp.CallToolResult, any, error) {
	// Parse arguments directly into the library type
	var params polymarketdata.GetTradesParams
	argsBytes, _ := json.Marshal(args)
	if err := json.Unmarshal(argsBytes, &params); err != nil {
		return nil, nil, fmt.Errorf("failed to parse arguments: %w", err)
	}

	// Get client
	dataClient := client.GetDataClient()

	trades, err := dataClient.GetTrades(ctx, &params)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get trades: %w", err)
	}

	// Format response
	tradesJSON, err := json.MarshalIndent(trades, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal trades: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(tradesJSON)},
		},
	}, trades, nil
}

// GetLiveVolumeTool returns the MCP tool definition for getting live volume
func GetLiveVolumeTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_live_volume",
		Description: "Get Polymarket live volume data",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"limit": map[string]any{
					"type":        "number",
					"description": "Number of markets to return",
				},
				"offset": map[string]any{
					"type":        "number",
					"description": "Number of markets to skip",
				},
				"sort_by": map[string]any{
					"type":        "string",
					"description": "Field to sort by",
				},
				"sort_direction": map[string]any{
					"type":        "string",
					"description": "Sort direction (ASC, DESC)",
				},
			},
		},
	}
}

// GetLiveVolumeHandler handles the get_live_volume tool execution
func GetLiveVolumeHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]any) (*mcp.CallToolResult, any, error) {
	// Parse arguments directly into the library type
	var params polymarketdata.GetLiveVolumeParams
	argsBytes, _ := json.Marshal(args)
	if err := json.Unmarshal(argsBytes, &params); err != nil {
		return nil, nil, fmt.Errorf("failed to parse arguments: %w", err)
	}

	// Get client
	dataClient := client.GetDataClient()

	liveVolume, err := dataClient.GetLiveVolume(ctx, &params)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get live volume: %w", err)
	}

	// Format response
	liveVolumeJSON, err := json.MarshalIndent(liveVolume, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal live volume: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(liveVolumeJSON)},
		},
	}, liveVolume, nil
}
