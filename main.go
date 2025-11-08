package main

import (
	"context"

	"github.com/ivanzzeth/polymarket-go-mcp/constants"
	"github.com/ivanzzeth/polymarket-go-mcp/tools"
	"github.com/ivanzzeth/polymarket-go-mcp/tools/gamma"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// healthCheckHandler handles health check requests
func healthCheckHandler(ctx context.Context, req *mcp.CallToolRequest, input map[string]any) (*mcp.CallToolResult, map[string]any, error) {
	return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: `{"status": "healthy", "message": "Polymarket MCP Server is running"}`},
			},
		}, map[string]any{
			"status":  "healthy",
			"message": "Polymarket MCP Server is running",
		}, nil
}

func main() {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    constants.ServerName,
		Version: constants.ServerVersion,
	}, nil)

	// Add health check tool
	mcp.AddTool(server, &mcp.Tool{
		Name:        constants.ToolHealthCheck,
		Description: constants.HealthCheckDescription,
	}, healthCheckHandler)

	// Add gamma tools
	mcp.AddTool(server, gamma.GetMarketsTool(), tools.RecoverPanicWrapper(gamma.GetMarketsHandler))
	mcp.AddTool(server, gamma.GetMarketByIDTool(), tools.RecoverPanicWrapper(gamma.GetMarketByIDHandler))
	// TODO: Uncomment when GetTags is confirmed to exist
	// mcp.AddTool(server, gamma.GetTagsTool(), tools.RecoverPanicWrapper(gamma.GetTagsHandler))

	// Run the server
	server.Run(context.Background(), &mcp.StdioTransport{})
}
