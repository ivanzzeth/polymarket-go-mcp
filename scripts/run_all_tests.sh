#!/bin/bash

echo "=== Running All MCP Server Tests ==="
echo ""

# Make all scripts executable
chmod +x scripts/*.sh

# Test 1: MCP Protocol
echo "📋 Test 1: MCP Protocol"
./scripts/test_mcp_protocol.sh
if [ $? -ne 0 ]; then
    echo "❌ MCP Protocol test failed"
    exit 1
fi
echo ""

# Test 2: health_check tool
echo "📋 Test 2: health_check Tool"
./scripts/test_health_check.sh
if [ $? -ne 0 ]; then
    echo "❌ health_check tool test failed"
    exit 1
fi
echo ""

# Test 3: get_markets tool
echo "📋 Test 3: get_markets Tool"
./scripts/test_get_markets.sh
if [ $? -ne 0 ]; then
    echo "❌ get_markets tool test failed"
    exit 1
fi
echo ""

echo "🎉 All tests passed successfully!"
echo ""
echo "=== Test Summary ==="
echo "✅ MCP Protocol: Server initialization and tool listing"
echo "✅ health_check: Server health status verification"
echo "✅ get_markets: Market data retrieval with filters"
echo ""
echo "The Polymarket MCP server is working correctly!"