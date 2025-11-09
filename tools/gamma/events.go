package gamma

import (
	"context"
	"encoding/json"
	"fmt"

	polymarketgamma "github.com/ivanzzeth/polymarket-go-gamma-client"
	"github.com/ivanzzeth/polymarket-go-mcp/client"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetEventsTool returns the MCP tool definition for getting events
func GetEventsTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_events",
		Description: "Get Polymarket events with comprehensive filtering and pagination options",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"limit": map[string]any{
					"type":        "number",
					"description": "Number of events to return",
				},
				"offset": map[string]any{
					"type":        "number",
					"description": "Number of events to skip",
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
					"description": "Filter by event IDs",
				},
				"slug": map[string]any{
					"type":        "array",
					"items":       map[string]any{"type": "string"},
					"description": "Filter by event slugs",
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
				"game_id": map[string]any{
					"type":        "string",
					"description": "Filter by game ID",
				},
				"sports_market_types": map[string]any{
					"type":        "array",
					"items":       map[string]any{"type": "string"},
					"description": "Filter by sports market types",
				},
				"closed": map[string]any{
					"type":        "boolean",
					"description": "Include closed events",
				},
			},
		},
	}
}

// GetEventsHandler handles the get_events tool execution
func GetEventsHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]any) (*mcp.CallToolResult, any, error) {
	// Parse arguments directly into the library type
	var params polymarketgamma.GetEventsParams
	argsBytes, _ := json.Marshal(args)
	if err := json.Unmarshal(argsBytes, &params); err != nil {
		return nil, nil, fmt.Errorf("failed to parse arguments: %w", err)
	}

	// Get client
	gammaClient := client.GetGammaClient()

	events, err := gammaClient.GetEvents(ctx, &params)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get events: %w", err)
	}

	// Format response
	eventsJSON, err := json.MarshalIndent(events, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal events: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(eventsJSON)},
		},
	}, events, nil
}

// GetEventByIDTool returns the MCP tool definition for getting an event by ID
func GetEventByIDTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_event_by_id",
		Description: "Get a single Polymarket event by its ID",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"id": map[string]any{
					"type":        "string",
					"description": "The event ID to retrieve",
				},
			},
			"required": []string{"id"},
		},
	}
}

// GetEventByIDHandler handles the get_event_by_id tool execution
func GetEventByIDHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]any) (*mcp.CallToolResult, any, error) {
	// Parse arguments
	id, ok := args["id"].(string)
	if !ok {
		return nil, nil, fmt.Errorf("id parameter is required and must be a string")
	}

	// Parse query parameters if provided (excluding id)
	var params *polymarketgamma.GetEventByIDQueryParams
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
			params = &polymarketgamma.GetEventByIDQueryParams{}
		}
	} else {
		params = &polymarketgamma.GetEventByIDQueryParams{}
	}

	// Get client
	gammaClient := client.GetGammaClient()

	// Get event by ID
	event, err := gammaClient.GetEventByID(ctx, id, params)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get event by ID: %w", err)
	}

	// Format response
	eventJSON, err := json.MarshalIndent(event, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal event: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(eventJSON)},
		},
	}, event, nil
}

// GetEventBySlugTool returns the MCP tool definition for getting an event by slug
func GetEventBySlugTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_event_by_slug",
		Description: "Get a single Polymarket event by its slug",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"slug": map[string]any{
					"type":        "string",
					"description": "The event slug to retrieve",
				},
			},
			"required": []string{"slug"},
		},
	}
}

// GetEventBySlugHandler handles the get_event_by_slug tool execution
func GetEventBySlugHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]any) (*mcp.CallToolResult, any, error) {
	// Parse arguments
	slug, ok := args["slug"].(string)
	if !ok {
		return nil, nil, fmt.Errorf("slug parameter is required and must be a string")
	}

	// Parse query parameters if provided (excluding slug)
	var params *polymarketgamma.GetEventBySlugQueryParams
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
			params = &polymarketgamma.GetEventBySlugQueryParams{}
		}
	} else {
		params = &polymarketgamma.GetEventBySlugQueryParams{}
	}

	// Get client
	gammaClient := client.GetGammaClient()

	// Get event by slug
	event, err := gammaClient.GetEventBySlug(ctx, slug, params)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get event by slug: %w", err)
	}

	// Format response
	eventJSON, err := json.MarshalIndent(event, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal event: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(eventJSON)},
		},
	}, event, nil
}
