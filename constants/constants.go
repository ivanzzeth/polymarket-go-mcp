package constants

// Server constants
const (
	ServerName    = "polymarket-go-mcp"
	ServerVersion = "0.1.0"
)

// Tool names
const (
	ToolHealthCheck = "health_check"
	ToolGetMarkets  = "get_markets"
)

// Tool descriptions
const (
	HealthCheckDescription = "Check if the Polymarket MCP server is working"
	GetMarketsDescription  = "Get Polymarket markets with comprehensive filtering and pagination options"
)

// Error messages
const (
	ErrNoContentInResult = "no content in result"
	ErrExpectedTextContent = "expected text content"
	ErrFailedToParseJSON = "failed to parse JSON"
)

// Default values
const (
	DefaultLimit = 10
	DefaultOffset = 0
)