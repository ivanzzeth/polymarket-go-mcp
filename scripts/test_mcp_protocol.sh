#!/bin/bash

echo "=== MCP Protocol Test Suite ==="
echo "Testing Polymarket MCP Server..."

# Check if server binary exists
if [ ! -f "./polymarket-go-mcp" ]; then
    echo "❌ Server binary not found. Building..."
    go build -o polymarket-go-mcp .
    if [ $? -ne 0 ]; then
        echo "❌ Failed to build server"
        exit 1
    fi
fi

echo "✅ Server binary ready"

# Test 1: Initialize and list tools
echo ""
echo "=== Test 1: Initialize and List Tools ==="
{
    echo '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{"roots":{"listChanged":true},"tools":{"listChanged":true}},"clientInfo":{"name":"test-client","version":"1.0.0"}}}'
    sleep 0.5
    echo '{"jsonrpc":"2.0","id":2,"method":"tools/list"}'
    sleep 0.5
} | timeout 10s ./polymarket-go-mcp 2>&1

if [ $? -eq 0 ]; then
    echo "✅ Test 1 passed: Server initialized and tools listed"
else
    echo "❌ Test 1 failed"
    exit 1
fi

echo ""
echo "=== All MCP Protocol Tests Completed ==="