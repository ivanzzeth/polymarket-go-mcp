package gamma

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ivanzzeth/polymarket-go-gamma-client"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetMarketsTool returns the MCP tool definition for getting markets
func GetMarketsTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_markets",
		Description: "Get Polymarket markets with comprehensive filtering and pagination options",
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
				"order": map[string]any{
					"type":        "string",
					"description": "Comma-separated list of fields to order by",
				},
				"ascending": map[string]any{
					"type":        "boolean",
					"description": "Sort order (true for ascending)",
				},
				"id": map[string]any{
					"type":        "array",
					"items": map[string]any{"type": "number"},
					"description": "Filter by market IDs",
				},
				"slug": map[string]any{
					"type":        "array",
					"items": map[string]any{"type": "string"},
					"description": "Filter by market slugs",
				},
				"clob_token_ids": map[string]any{
					"type":        "array",
					"items": map[string]any{"type": "string"},
					"description": "Filter by CLOB token IDs",
				},
				"condition_ids": map[string]any{
					"type":        "array",
					"items": map[string]any{"type": "string"},
					"description": "Filter by condition IDs",
				},
				"market_maker_address": map[string]any{
					"type":        "array",
					"items": map[string]any{"type": "string"},
					"description": "Filter by market maker addresses",
				},
				"liquidity_num_min": map[string]any{
					"type":        "number",
					"description": "Minimum liquidity amount",
				},
				"liquidity_num_max": map[string]any{
					"type":        "number",
					"description": "Maximum liquidity amount",
				},
				"volume_num_min": map[string]any{
					"type":        "number",
					"description": "Minimum volume amount",
				},
				"volume_num_max": map[string]any{
					"type":        "number",
					"description": "Maximum volume amount",
				},
				"start_date_min": map[string]any{
					"type":        "string",
					"description": "Minimum start date (ISO 8601)",
				},
				"start_date_max": map[string]any{
					"type":        "string",
					"description": "Maximum start date (ISO 8601)",
				},
				"end_date_min": map[string]any{
					"type":        "string",
					"description": "Minimum end date (ISO 8601)",
				},
				"end_date_max": map[string]any{
					"type":        "string",
					"description": "Maximum end date (ISO 8601)",
				},
				"tag_id": map[string]any{
					"type":        "number",
					"description": "Filter by tag ID",
				},
				"related_tags": map[string]any{
					"type":        "boolean",
					"description": "Include related tags",
				},
				"cyom": map[string]any{
					"type":        "boolean",
					"description": "Create Your Own Market filter",
				},
				"uma_resolution_status": map[string]any{
					"type":        "string",
					"description": "UMA resolution status filter",
				},
				"game_id": map[string]any{
					"type":        "string",
					"description": "Filter by game ID",
				},
				"sports_market_types": map[string]any{
					"type":        "array",
					"items": map[string]any{"type": "string"},
					"description": "Filter by sports market types",
				},
				"rewards_min_size": map[string]any{
					"type":        "number",
					"description": "Minimum rewards size",
				},
				"question_ids": map[string]any{
					"type":        "array",
					"items": map[string]any{"type": "string"},
					"description": "Filter by question IDs",
				},
				"include_tag": map[string]any{
					"type":        "boolean",
					"description": "Include tag information",
				},
				"closed": map[string]any{
					"type":        "boolean",
					"description": "Include closed markets",
				},
			},
		},
	}
}

// GetMarketsHandler handles the get_markets tool execution
func GetMarketsHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]any) (*mcp.CallToolResult, any, error) {
	// Parse arguments directly into the library type
	var params polymarketgamma.GetMarketsParams
	argsBytes, _ := json.Marshal(args)
	if err := json.Unmarshal(argsBytes, &params); err != nil {
		return nil, nil, err
	}

	// Create Gamma client with proper HTTP client
	client := polymarketgamma.NewClient(&http.Client{})
	
	markets, err := client.GetMarkets(ctx, &params)
	if err != nil {
		return nil, nil, err
	}

	// Format response
	marketsJSON, _ := json.MarshalIndent(markets, "", "  ")
	
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(marketsJSON)},
		},
	}, markets, nil
}