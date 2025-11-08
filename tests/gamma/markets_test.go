package gamma

import (
	"context"
	"testing"

	"github.com/ivanzzeth/polymarket-go-mcp/tools/gamma"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestGetMarkets(t *testing.T) {
	ctx := context.Background()
	req := &mcp.CallToolRequest{}
	args := map[string]any{
		"limit": float64(5), // JSON numbers are float64 in Go
	}

	handler := gamma.GetMarketsHandler
	result, data, err := handler(ctx, req, args)

	if err != nil {
		t.Fatalf("failed to get markets: %v", err)
	}

	if result == nil {
		t.Fatal("result is nil")
	}

	if len(result.Content) == 0 {
		t.Fatal("result content is empty")
	}

	if data == nil {
		t.Fatal("data is nil")
	}

	t.Logf("result: %+v", result)
	t.Logf("data: %+v", data)
}
