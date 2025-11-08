# Polymarket Go MCP

[![Go Version](https://img.shields.io/badge/go-1.24+-blue.svg)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![MCP Protocol](https://img.shields.io/badge/MCP-Protocol-blueviolet.svg)](https://spec.modelcontextprotocol.io)

A Go-based MCP (Model Context Protocol) server for accessing Polymarket prediction market data through Gamma API and Data API.

## Features

- 📊 **Market Data** - Access real-time market prices, volume, and liquidity
- 🔍 **Market Search** - Find markets by keywords and categories
- 📈 **Price History** - Get historical price data and trends
- 💰 **Portfolio Tracking** - Monitor positions and PnL
- 🚀 **High Performance** - Built with Go for fast and reliable performance
- 🔧 **Developer Friendly** - Multiple installation options for Go developers

## Quick Start

### For Go Developers (Recommended)

#### Option 1: Install from Source (Latest Version)
```bash
go install github.com/ivanzzeth/polymarket-go-mcp/cmd/polymarket-go-mcp@latest
```

#### Option 2: Build from Source
```bash
git clone https://github.com/ivanzzeth/polymarket-go-mcp
cd polymarket-go-mcp
go build -o polymarket-go-mcp ./cmd/polymarket-go-mcp
```

#### Option 3: Download Pre-built Binary
```bash
# Linux AMD64
curl -L https://github.com/ivanzzeth/polymarket-go-mcp/releases/latest/download/polymarket-go-mcp-linux-amd64 -o polymarket-go-mcp
chmod +x polymarket-go-mcp

# macOS ARM64
curl -L https://github.com/ivanzzeth/polymarket-go-mcp/releases/latest/download/polymarket-go-mcp-darwin-arm64 -o polymarket-go-mcp
chmod +x polymarket-go-mcp

# Windows AMD64
curl -L https://github.com/ivanzzeth/polymarket-go-mcp/releases/latest/download/polymarket-go-mcp-win-amd64.exe -o polymarket-go-mcp.exe
```

### For All Users

#### Option 4: Using npx (No Installation Required)
```bash
npx -y polymarket-go-mcp
```

#### Option 5: Using Docker
```bash
docker run --rm -it ivanzzeth/polymarket-go-mcp:latest
```

## Configuration

### Claude Desktop

Add to your `claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "polymarket": {
      "command": "polymarket-go-mcp"
    }
  }
}
```

### For Go Install Users

If you installed via `go install`, the binary will be in your `$GOPATH/bin` or `$GOBIN`. Make sure this directory is in your PATH:

```json
{
  "mcpServers": {
    "polymarket": {
      "command": "polymarket-go-mcp"
    }
  }
}
```

### For Source Build Users

If you built from source, use the full path:

```json
{
  "mcpServers": {
    "polymarket": {
      "command": "/path/to/polymarket-go-mcp/polymarket-go-mcp"
    }
  }
}
```

## Available Tools

### Market Data Tools
- `get_market_data` - Get detailed market information
- `search_markets` - Search markets by keyword
- `get_market_prices` - Get current market prices
- `get_market_history` - Get historical price data

### Portfolio Tools
- `get_positions` - Get user positions
- `get_pnl` - Calculate profit and loss
- `get_portfolio_value` - Get total portfolio value

### Analysis Tools
- `get_liquidity_data` - Get market liquidity information
- `get_volume_data` - Get trading volume statistics
- `get_market_trends` - Analyze market trends

## Development

### Prerequisites
- Go 1.24 or later
- Git

### Building from Source

```bash
# Clone the repository
git clone https://github.com/ivanzzeth/polymarket-go-mcp
cd polymarket-go-mcp

# Install dependencies
go mod download

# Build the binary
go build -o polymarket-go-mcp ./cmd/polymarket-go-mcp

# Test the build
./polymarket-go-mcp --version
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run integration tests
go test -tags=integration ./...
```

### Project Structure

```
polymarket-go-mcp/
├── cmd/
│   └── polymarket-go-mcp/
│       └── main.go          # CLI entry point
├── internal/
│   ├── mcp/                 # MCP protocol implementation
│   ├── gamma/               # Gamma API client
│   ├── data/                # Data API client
│   └── config/              # Configuration management
├── pkg/
│   └── utils/               # Utility functions
├── docs/                    # Documentation
├── scripts/                 # Build scripts
└── README.md
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `LOG_LEVEL` | Log level (debug, info, warn, error) | `info` |
| `GAMMA_API_BASE` | Gamma API base URL | `https://gamma-api.polymarket.com` |
| `DATA_API_BASE` | Data API base URL | `https://data-api.polymarket.com` |
| `CACHE_ENABLED` | Enable response caching | `true` |
| `CACHE_TTL` | Cache time-to-live | `10m` |

## Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/amazing-feature`
3. Commit your changes: `git commit -m 'Add amazing feature'`
4. Push to the branch: `git push origin feature/amazing-feature`
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

- 📖 [Documentation](docs/)
- 🐛 [Issue Tracker](https://github.com/ivanzzeth/polymarket-go-mcp/issues)
- 💬 [Discussions](https://github.com/ivanzzeth/polymarket-go-mcp/discussions)

## Acknowledgments

- [Model Context Protocol](https://spec.modelcontextprotocol.io) - The MCP specification
- [Polymarket](https://polymarket.com) - Prediction market platform
- [Gamma API](https://gamma-api.polymarket.com) - Market data API
- [Data API](https://data-api.polymarket.com) - Historical data API

---

**Made with ❤️ for the Go and MCP communities**