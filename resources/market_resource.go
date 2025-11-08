package resources

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivanzzeth/polymarket-go-gamma-client"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// MarketResource provides access to market data as MCP resources
type MarketResource struct {
	client *polymarketgamma.Client
}

// NewMarketResource creates a new market resource handler
func NewMarketResource() *MarketResource {
	return &MarketResource{
		client: polymarketgamma.NewClient(http.DefaultClient),
	}
}

// GetMarketResource gets market data by ID
func (r *MarketResource) GetMarketResource(ctx context.Context, marketID string) (*mcp.Resource, error) {
	params := &polymarketgamma.GetMarketByIDQueryParams{}
	market, err := r.client.GetMarketByID(ctx, marketID, params)
	if err != nil {
		return nil, fmt.Errorf("failed to get market: %w", err)
	}

	// Format market data as JSON
	marketJSON, err := json.MarshalIndent(market, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal market data: %w", err)
	}

	return &mcp.Resource{
		URI: fmt.Sprintf("market://%s", marketID),
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: string(marketJSON),
			},
		},
	}, nil
}

// ListMarketResources lists available market resources (placeholder)
func (r *MarketResource) ListMarketResources(ctx context.Context) ([]*mcp.Resource, error) {
	// This would typically list available markets
	// For now, return an empty list as this is a complex operation
	return []*mcp.Resource{}, nil
}