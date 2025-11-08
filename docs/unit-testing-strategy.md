# Polymarket MCP Server Unit Testing Strategy

## Overview
This document outlines the unit testing strategy for the Polymarket MCP server using Go's built-in testing framework.

## Test Structure

### 1. MCP Protocol Tests
```go
// TestMCPProtocol - Tests MCP server initialization and protocol compliance
func TestMCPProtocol(t *testing.T) {
    // Test server initialization
    // Test tool listing
    // Test protocol version handling
}
```

### 2. Tool Handler Tests
```go
// TestHealthCheckHandler - Tests health check tool functionality
func TestHealthCheckHandler(t *testing.T) {
    // Test with valid request
    // Test response format
    // Test error handling
}

// TestGetMarketsHandler - Tests market data retrieval
func TestGetMarketsHandler(t *testing.T) {
    // Test with different parameters
    // Test error scenarios
    // Test response parsing
}
```

### 3. Integration Tests
```go
// TestGammaClientIntegration - Tests external API integration
func TestGammaClientIntegration(t *testing.T) {
    // Test API client creation
    // Test request/response handling
    // Test error scenarios
}
```

## Test Categories

### Unit Tests
- **Handler Functions**: Test individual tool handlers in isolation
- **Client Logic**: Test API client methods with mocked responses
- **Data Parsing**: Test JSON marshaling/unmarshaling

### Integration Tests  
- **External APIs**: Test actual API calls (with proper setup)
- **MCP Protocol**: Test full MCP message flow

### Mock Strategy
```go
// Mock Gamma API client for testing
type MockGammaClient struct {
    GetMarketsFunc func(ctx context.Context, params *GetMarketsParams) ([]Market, error)
}

func (m *MockGammaClient) GetMarkets(ctx context.Context, params *GetMarketsParams) ([]Market, error) {
    return m.GetMarketsFunc(ctx, params)
}
```

## Test Files Structure
```
tests/
├── mcp_protocol_test.go     # MCP protocol tests
├── tools_test.go            # Tool handler tests
├── gamma_client_test.go     # Gamma API client tests
├── integration_test.go      # Integration tests
└── test_utils.go            # Test utilities and mocks
```

## Key Test Scenarios

### Health Check Tool
- ✅ Valid request returns healthy status
- ✅ Response contains expected content
- ✅ Error handling for malformed requests

### Get Markets Tool
- ✅ Default parameters return markets
- ✅ Filter parameters work correctly  
- ✅ Error handling for API failures
- ✅ Response format validation
- ✅ Pagination parameters

### Error Scenarios
- Network timeouts
- Invalid parameters
- API rate limiting
- JSON parsing errors

## Test Data
- Mock market data for consistent testing
- Error scenarios for robustness testing
- Edge cases for parameter validation

## Running Tests
```bash
# Run all tests
go test ./...

# Run specific test
go test -v -run TestHealthCheckHandler

# Run with coverage
go test -cover ./...
```

## Benefits of This Approach
1. **Maintainability**: Tests are version-controlled with code
2. **CI/CD Integration**: Easy to integrate with GitHub Actions
3. **Fast Execution**: Unit tests run quickly
4. **Isolation**: Tests don't depend on external services
5. **Documentation**: Tests serve as living documentation