package clients

import (
	"context"
	"net/http"

	"github.com/ivanzzeth/polymarket-go-gamma-client"
)

// GammaClient wraps the polymarket gamma client with additional functionality
type GammaClient struct {
	client *polymarketgamma.Client
}

// NewGammaClient creates a new Gamma client wrapper
func NewGammaClient() *GammaClient {
	return &GammaClient{
		client: polymarketgamma.NewClient(http.DefaultClient),
	}
}

// GetMarkets wraps the GetMarkets method with additional error handling
func (c *GammaClient) GetMarkets(ctx context.Context, params *polymarketgamma.GetMarketsParams) ([]*polymarketgamma.Market, error) {
	return c.client.GetMarkets(ctx, params)
}

// GetMarketByID gets a specific market by ID
func (c *GammaClient) GetMarketByID(ctx context.Context, marketID string) (*polymarketgamma.Market, error) {
	params := &polymarketgamma.GetMarketByIDQueryParams{}
	market, err := c.client.GetMarketByID(ctx, marketID, params)
	if err != nil {
		return nil, err
	}
	return market, nil
}

// SearchMarkets searches markets by query
func (c *GammaClient) SearchMarkets(ctx context.Context, query string, limit int) (*polymarketgamma.SearchResponse, error) {
	params := &polymarketgamma.SearchParams{
		Query: query,
		Limit: limit,
	}
	return c.client.Search(ctx, params)
}

// GetEvents gets events with optional filtering
func (c *GammaClient) GetEvents(ctx context.Context, limit int, series *string) ([]*polymarketgamma.Event, error) {
	params := &polymarketgamma.GetEventsParams{
		Limit: limit,
	}
	if series != nil {
		params.Series = series
	}
	return c.client.GetEvents(ctx, params)
}