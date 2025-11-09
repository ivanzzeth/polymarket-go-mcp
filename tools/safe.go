package tools

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func RecoverPanicWrapper(fn mcp.ToolHandlerFor[map[string]any, any]) mcp.ToolHandlerFor[map[string]any, any] {
	return func(ctx context.Context, req *mcp.CallToolRequest, args map[string]any) (toolResult *mcp.CallToolResult, result any, err error) {
		defer func() {
			if r := recover(); r != nil {
				toolResult, result, err = nil, nil, fmt.Errorf("panic: %v", r)
			}
		}()
		return fn(ctx, req, args)
	}
}
