#!/bin/bash

echo "=== Testing get_markets Tool ==="

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

# Test get_markets tool with different parameters
echo ""
echo "=== Test 1: get_markets with default parameters ==="
{
    # Initialize
    echo '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{"roots":{"listChanged":true},"tools":{"listChanged":true}},"clientInfo":{"name":"test-client","version":"1.0.0"}}}'
    sleep 0.5
    
    # Call get_markets tool with default parameters
    echo '{"jsonrpc":"2.0","id":2,"method":"tools/call","params":{"name":"get_markets","arguments":{"limit":5}}}'
    sleep 2
} | timeout 15s ./polymarket-go-mcp 2>&1

RESULT=$?
if [ $RESULT -eq 0 ]; then
    echo "✅ get_markets tool test 1 passed"
else
    echo "❌ get_markets tool test 1 failed with exit code: $RESULT"
    exit 1
fi

echo ""
echo "=== Test 2: get_markets with closed markets filter ==="
{
    # Initialize
    echo '{"jsonrpc":"2.0","id":3,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{"roots":{"listChanged":true},"tools":{"listChanged":true}},"clientInfo":{"name":"test-client","version":"1.0.0"}}}'
    sleep 0.5
    
    # Call get_markets tool with closed markets
    echo '{"jsonrpc":"2.0","id":4,"method":"tools/call","params":{"name":"get_markets","arguments":{"limit":3,"closed":true}}}'
    sleep 2
} | timeout 15s ./polymarket-go-mcp 2>&1

RESULT=$?
if [ $RESULT -eq 0 ]; then
    echo "✅ get_markets tool test 2 passed"
else
    echo "❌ get_markets tool test 2 failed with exit code: $RESULT"
    exit 1
fi

echo ""
echo "=== get_markets Tool Tests Completed ==="