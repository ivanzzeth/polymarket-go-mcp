package gamma

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ivanzzeth/polymarket-go-mcp/client"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// HealthCheckTool returns the MCP tool definition for health check
func HealthCheckTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "gamma_health_check",
		Description: "Check the health status of the Gamma API",
		InputSchema: map[string]any{
			"type":       "object",
			"properties": map[string]any{},
		},
	}
}

// HealthCheckHandler handles the gamma_health_check tool execution
func HealthCheckHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]any) (*mcp.CallToolResult, any, error) {
	// Get client
	gammaClient := client.GetGammaClient()

	health, err := gammaClient.HealthCheck(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to check gamma health: %w", err)
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
