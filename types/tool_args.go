package types

// Tool parameter types - only define what's needed for MCP tool inputs
// We use the original library types directly when they already provide all needed parameters
// Only define new types here when the library doesn't provide suitable types for MCP tools

type GetMarketByIDArgs struct {
	ID string `json:"id" jsonschema:"Market ID"`
}

type SearchMarketsArgs struct {
	Query string `json:"query" jsonschema:"Search query string"`
	Limit *int   `json:"limit,omitempty" jsonschema:"Number of results"`
}