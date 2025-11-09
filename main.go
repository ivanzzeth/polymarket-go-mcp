package main

import (
	"context"

	"github.com/ivanzzeth/polymarket-go-mcp/constants"
	"github.com/ivanzzeth/polymarket-go-mcp/tools"
	"github.com/ivanzzeth/polymarket-go-mcp/tools/data"
	"github.com/ivanzzeth/polymarket-go-mcp/tools/gamma"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// healthCheckHandler handles health check requests
func healthCheckHandler(ctx context.Context, req *mcp.CallToolRequest, input map[string]any) (*mcp.CallToolResult, map[string]any, error) {
	return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: `{"status": "healthy", "message": "Polymarket Go MCP Server is running"}`},
			},
		}, map[string]any{
			"status":  "healthy",
			"message": "Polymarket Go MCP Server is running",
		}, nil
}

func main() {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    constants.ServerName,
		Version: constants.Version,
	}, nil)

	// Add health check tool
	mcp.AddTool(server, &mcp.Tool{
		Name:        constants.ToolHealthCheck,
		Description: constants.HealthCheckDescription,
	}, healthCheckHandler)

	// Add gamma tools
	mcp.AddTool(server, gamma.GetMarketsTool(), tools.RecoverPanicWrapper(gamma.GetMarketsHandler))
	mcp.AddTool(server, gamma.GetMarketByIDTool(), tools.RecoverPanicWrapper(gamma.GetMarketByIDHandler))
	mcp.AddTool(server, gamma.GetMarketBySlugTool(), tools.RecoverPanicWrapper(gamma.GetMarketBySlugHandler))
	mcp.AddTool(server, gamma.GetTagsTool(), tools.RecoverPanicWrapper(gamma.GetTagsHandler))
	mcp.AddTool(server, gamma.GetTagByIDTool(), tools.RecoverPanicWrapper(gamma.GetTagByIDHandler))
	mcp.AddTool(server, gamma.GetTagBySlugTool(), tools.RecoverPanicWrapper(gamma.GetTagBySlugHandler))
	mcp.AddTool(server, gamma.GetRelatedTagsByIDTool(), tools.RecoverPanicWrapper(gamma.GetRelatedTagsByIDHandler))
	mcp.AddTool(server, gamma.GetRelatedTagsBySlugTool(), tools.RecoverPanicWrapper(gamma.GetRelatedTagsBySlugHandler))
	mcp.AddTool(server, gamma.GetRelatedTagsDetailByIDTool(), tools.RecoverPanicWrapper(gamma.GetRelatedTagsDetailByIDHandler))
	mcp.AddTool(server, gamma.GetRelatedTagsDetailBySlugTool(), tools.RecoverPanicWrapper(gamma.GetRelatedTagsDetailBySlugHandler))
	mcp.AddTool(server, gamma.GetEventsTool(), tools.RecoverPanicWrapper(gamma.GetEventsHandler))
	mcp.AddTool(server, gamma.GetEventByIDTool(), tools.RecoverPanicWrapper(gamma.GetEventByIDHandler))
	mcp.AddTool(server, gamma.GetEventBySlugTool(), tools.RecoverPanicWrapper(gamma.GetEventBySlugHandler))
	mcp.AddTool(server, gamma.GetSeriesTool(), tools.RecoverPanicWrapper(gamma.GetSeriesHandler))
	mcp.AddTool(server, gamma.GetSeriesByIDTool(), tools.RecoverPanicWrapper(gamma.GetSeriesByIDHandler))
	mcp.AddTool(server, gamma.GetTeamsTool(), tools.RecoverPanicWrapper(gamma.GetTeamsHandler))
	mcp.AddTool(server, gamma.GetSportsMetadataTool(), tools.RecoverPanicWrapper(gamma.GetSportsMetadataHandler))
	mcp.AddTool(server, gamma.SearchTool(), tools.RecoverPanicWrapper(gamma.SearchHandler))
	mcp.AddTool(server, gamma.HealthCheckTool(), tools.RecoverPanicWrapper(gamma.HealthCheckHandler))

	// Add data tools
	mcp.AddTool(server, data.GetActivityTool(), tools.RecoverPanicWrapper(data.GetActivityHandler))
	mcp.AddTool(server, data.GetTradesTool(), tools.RecoverPanicWrapper(data.GetTradesHandler))
	mcp.AddTool(server, data.GetLiveVolumeTool(), tools.RecoverPanicWrapper(data.GetLiveVolumeHandler))
	mcp.AddTool(server, data.GetPositionsTool(), tools.RecoverPanicWrapper(data.GetPositionsHandler))
	mcp.AddTool(server, data.GetClosedPositionsTool(), tools.RecoverPanicWrapper(data.GetClosedPositionsHandler))
	mcp.AddTool(server, data.GetPositionsValueTool(), tools.RecoverPanicWrapper(data.GetPositionsValueHandler))
	mcp.AddTool(server, data.GetOpenInterestTool(), tools.RecoverPanicWrapper(data.GetOpenInterestHandler))
	mcp.AddTool(server, data.GetHoldersTool(), tools.RecoverPanicWrapper(data.GetHoldersHandler))
	mcp.AddTool(server, data.GetTradedMarketsCountTool(), tools.RecoverPanicWrapper(data.GetTradedMarketsCountHandler))
	mcp.AddTool(server, data.DataHealthCheckTool(), tools.RecoverPanicWrapper(data.DataHealthCheckHandler))

	// Run the server
	server.Run(context.Background(), &mcp.StdioTransport{})
}
