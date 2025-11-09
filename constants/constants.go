package constants

// Server constants
const (
	ServerName = "polymarket-go-mcp"
	Version    = "0.0.4"
)

// Tool names
const (
	ToolHealthCheck   = "health_check"
	ToolGetMarkets    = "get_markets"
	ToolGetMarketByID = "get_market_by_id"
	ToolGetTags       = "get_tags"
)

// Tool descriptions
const (
	HealthCheckDescription   = "Check if the Polymarket MCP server is working"
	GetMarketsDescription    = "Get Polymarket markets with comprehensive filtering and pagination options"
	GetMarketByIDDescription = "Get a single Polymarket market by its ID"
	GetTagsDescription       = "Get Polymarket tags with filtering options"
)

// Error messages
const (
	ErrNoContentInResult   = "no content in result"
	ErrExpectedTextContent = "expected text content"
	ErrFailedToParseJSON   = "failed to parse JSON"
)

// Default values
const (
	DefaultLimit  = 10
	DefaultOffset = 0
)
