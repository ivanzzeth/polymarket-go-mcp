#!/bin/bash

echo "=== Testing health_check Tool ==="

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

# Test health_check tool
echo ""
echo "=== Test: health_check Tool ==="
{
    # Initialize
    echo '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{"roots":{"listChanged":true},"tools":{"listChanged":true}},"clientInfo":{"name":"test-client","version":"1.0.0"}}}'
    sleep 0.5
    
    # Call health_check tool
    echo '{"jsonrpc":"2.0","id":2,"method":"tools/call","params":{"name":"health_check","arguments":{}}}'
    sleep 1
} | timeout 10s ./polymarket-go-mcp 2>&1

RESULT=$?
if [ $RESULT -eq 0 ]; then
    echo "✅ health_check tool test passed"
else
    echo "❌ health_check tool test failed with exit code: $RESULT"
    exit 1
fi

echo ""
echo "=== health_check Tool Test Completed ==="