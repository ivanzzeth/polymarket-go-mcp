package gamma

import (
	"context"
	"encoding/json"
	"fmt"

	polymarketgamma "github.com/ivanzzeth/polymarket-go-gamma-client"
	"github.com/ivanzzeth/polymarket-go-mcp/client"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetSeriesTool returns the MCP tool definition for getting series
func GetSeriesTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_series",
		Description: "Get Polymarket series with comprehensive filtering and pagination options",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"limit": map[string]any{
					"type":        "number",
					"description": "Number of series to return",
				},
				"offset": map[string]any{
					"type":        "number",
					"description": "Number of series to skip",
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
					"description": "Filter by series IDs",
				},
				"slug": map[string]any{
					"type":        "array",
					"items":       map[string]any{"type": "string"},
					"description": "Filter by series slugs",
				},
				"tag_id": map[string]any{
					"type":        "number",
					"description": "Filter by tag ID",
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
				"closed": map[string]any{
					"type":        "boolean",
					"description": "Include closed series",
				},
			},
		},
	}
}

// GetSeriesHandler handles the get_series tool execution
func GetSeriesHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]any) (*mcp.CallToolResult, any, error) {
	// Parse arguments directly into the library type
	var params polymarketgamma.GetSeriesParams
	argsBytes, _ := json.Marshal(args)
	if err := json.Unmarshal(argsBytes, &params); err != nil {
		return nil, nil, fmt.Errorf("failed to parse arguments: %w", err)
	}

	// Get client
	gammaClient := client.GetGammaClient()

	series, err := gammaClient.GetSeries(ctx, &params)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get series: %w", err)
	}

	// Format response
	seriesJSON, err := json.MarshalIndent(series, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal series: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(seriesJSON)},
		},
	}, series, nil
}

// GetSeriesByIDTool returns the MCP tool definition for getting a series by ID
func GetSeriesByIDTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_series_by_id",
		Description: "Get a single Polymarket series by its ID",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"id": map[string]any{
					"type":        "string",
					"description": "The series ID to retrieve",
				},
			},
			"required": []string{"id"},
		},
	}
}

// GetSeriesByIDHandler handles the get_series_by_id tool execution
func GetSeriesByIDHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]any) (*mcp.CallToolResult, any, error) {
	// Parse arguments
	id, ok := args["id"].(string)
	if !ok {
		return nil, nil, fmt.Errorf("id parameter is required and must be a string")
	}

	// Parse query parameters if provided (excluding id)
	var params *polymarketgamma.GetSeriesByIDQueryParams
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
			params = &polymarketgamma.GetSeriesByIDQueryParams{}
		}
	} else {
		params = &polymarketgamma.GetSeriesByIDQueryParams{}
	}

	// Get client
	gammaClient := client.GetGammaClient()

	// Get series by ID
	series, err := gammaClient.GetSeriesByID(ctx, id, params)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get series by ID: %w", err)
	}

	// Format response
	seriesJSON, err := json.MarshalIndent(series, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal series: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(seriesJSON)},
		},
	}, series, nil
}
