package gamma

import (
	"context"
	"encoding/json"
	"fmt"

	polymarketgamma "github.com/ivanzzeth/polymarket-go-gamma-client"
	"github.com/ivanzzeth/polymarket-go-mcp/client"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetMarketsTool returns the MCP tool definition for getting markets
func GetMarketsTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_markets",
		Description: "Get Polymarket gamma markets with comprehensive filtering and pagination options",
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
					"items":       map[string]any{"type": "number"},
					"description": "Filter by market IDs",
				},
				"slug": map[string]any{
					"type":        "array",
					"items":       map[string]any{"type": "string"},
					"description": "Filter by market slugs",
				},
				"clob_token_ids": map[string]any{
					"type":        "array",
					"items":       map[string]any{"type": "string"},
					"description": "Filter by CLOB token IDs",
				},
				"condition_ids": map[string]any{
					"type":        "array",
					"items":       map[string]any{"type": "string"},
					"description": "Filter by condition IDs",
				},
				"market_maker_address": map[string]any{
					"type":        "array",
					"items":       map[string]any{"type": "string"},
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
					"items":       map[string]any{"type": "string"},
					"description": "Filter by sports market types",
				},
				"rewards_min_size": map[string]any{
					"type":        "number",
					"description": "Minimum rewards size",
				},
				"question_ids": map[string]any{
					"type":        "array",
					"items":       map[string]any{"type": "string"},
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
		return nil, nil, fmt.Errorf("failed to parse arguments: %w", err)
	}

	// Get client
	gammaClient := client.GetGammaClient()

	markets, err := gammaClient.GetMarkets(ctx, &params)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get markets: %w", err)
	}

	// Format response
	marketsJSON, err := json.MarshalIndent(markets, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal markets: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(marketsJSON)},
		},
	}, markets, nil
}

// GetMarketByIDTool returns the MCP tool definition for getting a market by ID
func GetMarketByIDTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_market_by_id",
		Description: "Get a single Polymarket market by its ID",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"id": map[string]any{
					"type":        "string",
					"description": "The market ID to retrieve",
				},
			},
			"required": []string{"id"},
		},
	}
}

// GetMarketByIDHandler handles the get_market_by_id tool execution
func GetMarketByIDHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]any) (*mcp.CallToolResult, any, error) {
	// Parse arguments
	id, ok := args["id"].(string)
	if !ok {
		return nil, nil, fmt.Errorf("id parameter is required and must be a string")
	}

	// Parse query parameters if provided (excluding id)
	var params *polymarketgamma.GetMarketByIDQueryParams
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
			params = &polymarketgamma.GetMarketByIDQueryParams{}
		}
	} else {
		params = &polymarketgamma.GetMarketByIDQueryParams{}
	}

	// Get client
	gammaClient := client.GetGammaClient()

	// Get market by ID
	market, err := gammaClient.GetMarketByID(ctx, id, params)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get market by ID: %w", err)
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
