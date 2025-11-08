package gamma

import (
	"context"
	"encoding/json"
	"fmt"

	polymarketgamma "github.com/ivanzzeth/polymarket-go-gamma-client"
	"github.com/ivanzzeth/polymarket-go-mcp/client"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetTeamsTool returns the MCP tool definition for getting teams
func GetTeamsTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_teams",
		Description: "Get Polymarket teams with filtering options",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"limit": map[string]any{
					"type":        "number",
					"description": "Number of teams to return",
				},
				"offset": map[string]any{
					"type":        "number",
					"description": "Number of teams to skip",
				},
				"id": map[string]any{
					"type":        "array",
					"items":       map[string]any{"type": "number"},
					"description": "Filter by team IDs",
				},
				"slug": map[string]any{
					"type":        "array",
					"items":       map[string]any{"type": "string"},
					"description": "Filter by team slugs",
				},
			},
		},
	}
}

// GetTeamsHandler handles the get_teams tool execution
func GetTeamsHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]any) (*mcp.CallToolResult, any, error) {
	// Parse arguments directly into the library type
	var params polymarketgamma.GetTeamsParams
	argsBytes, _ := json.Marshal(args)
	if err := json.Unmarshal(argsBytes, &params); err != nil {
		return nil, nil, fmt.Errorf("failed to parse arguments: %w", err)
	}

	// Get client
	gammaClient := client.GetGammaClient()

	teams, err := gammaClient.GetTeams(ctx, &params)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get teams: %w", err)
	}

	// Format response
	teamsJSON, err := json.MarshalIndent(teams, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal teams: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(teamsJSON)},
		},
	}, teams, nil
}
