package data

import (
	"context"
	"encoding/json"
	"fmt"

	polymarketdata "github.com/ivanzzeth/polymarket-go-data-client"
	"github.com/ivanzzeth/polymarket-go-mcp/client"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetOpenInterestTool returns the MCP tool definition for getting open interest
func GetOpenInterestTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_open_interest",
		Description: "Get Polymarket open interest data",
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
				"market": map[string]any{
					"type":        "string",
					"description": "Filter by market ID",
				},
			},
		},
	}
}

// GetOpenInterestHandler handles the get_open_interest tool execution
func GetOpenInterestHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]any) (*mcp.CallToolResult, any, error) {
	// Parse arguments directly into the library type
	var params polymarketdata.GetOpenInterestParams
	argsBytes, _ := json.Marshal(args)
	if err := json.Unmarshal(argsBytes, &params); err != nil {
		return nil, nil, fmt.Errorf("failed to parse arguments: %w", err)
	}

	// Get client
	dataClient := client.GetDataClient()

	openInterest, err := dataClient.GetOpenInterest(ctx, &params)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get open interest: %w", err)
	}

	// Format response
	openInterestJSON, err := json.MarshalIndent(openInterest, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal open interest: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(openInterestJSON)},
		},
	}, openInterest, nil
}

// GetHoldersTool returns the MCP tool definition for getting holders
func GetHoldersTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_holders",
		Description: "Get Polymarket holders data with filtering options",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"limit": map[string]any{
					"type":        "number",
					"description": "Number of holders to return",
				},
				"offset": map[string]any{
					"type":        "number",
					"description": "Number of holders to skip",
				},
				"sort_by": map[string]any{
					"type":        "string",
					"description": "Field to sort by",
				},
				"sort_direction": map[string]any{
					"type":        "string",
					"description": "Sort direction (ASC, DESC)",
				},
				"market": map[string]any{
					"type":        "string",
					"description": "Filter by market ID",
				},
				"user": map[string]any{
					"type":        "string",
					"description": "Filter by user address",
				},
			},
		},
	}
}

// GetHoldersHandler handles the get_holders tool execution
func GetHoldersHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]any) (*mcp.CallToolResult, any, error) {
	// Parse arguments directly into the library type
	var params polymarketdata.GetHoldersParams
	argsBytes, _ := json.Marshal(args)
	if err := json.Unmarshal(argsBytes, &params); err != nil {
		return nil, nil, fmt.Errorf("failed to parse arguments: %w", err)
	}

	// Get client
	dataClient := client.GetDataClient()

	holders, err := dataClient.GetHolders(ctx, &params)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get holders: %w", err)
	}

	// Format response
	holdersJSON, err := json.MarshalIndent(holders, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal holders: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(holdersJSON)},
		},
	}, holders, nil
}

// GetTradedMarketsCountTool returns the MCP tool definition for getting traded markets count
func GetTradedMarketsCountTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_traded_markets_count",
		Description: "Get Polymarket traded markets count data",
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
			},
		},
	}
}

// GetTradedMarketsCountHandler handles the get_traded_markets_count tool execution
func GetTradedMarketsCountHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]any) (*mcp.CallToolResult, any, error) {
	// Parse arguments directly into the library type
	var params polymarketdata.GetTradedMarketsCountParams
	argsBytes, _ := json.Marshal(args)
	if err := json.Unmarshal(argsBytes, &params); err != nil {
		return nil, nil, fmt.Errorf("failed to parse arguments: %w", err)
	}

	// Get client
	dataClient := client.GetDataClient()

	tradedMarketsCount, err := dataClient.GetTradedMarketsCount(ctx, &params)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get traded markets count: %w", err)
	}

	// Format response
	tradedMarketsCountJSON, err := json.MarshalIndent(tradedMarketsCount, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal traded markets count: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(tradedMarketsCountJSON)},
		},
	}, tradedMarketsCount, nil
}

// DataHealthCheckTool returns the MCP tool definition for data health check
func DataHealthCheckTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "data_health_check",
		Description: "Check the health status of the Data API",
		InputSchema: map[string]any{
			"type":       "object",
			"properties": map[string]any{},
		},
	}
}

// DataHealthCheckHandler handles the data_health_check tool execution
func DataHealthCheckHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]any) (*mcp.CallToolResult, any, error) {
	// Get client
	dataClient := client.GetDataClient()

	health, err := dataClient.HealthCheck(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to check data health: %w", err)
	}

	// Format response
	healthJSON, err := json.MarshalIndent(health, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal health response: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(healthJSON)},
		},
	}, health, nil
}
