#!/bin/bash

echo "=== Testing get_markets Tool ==="

# Cross-platform timeout function
run_with_timeout() {
    local timeout_duration=$1
    shift
    
    if command -v timeout >/dev/null 2>&1; then
        # Linux: use timeout
        timeout "${timeout_duration}" "$@"
    elif command -v gtimeout >/dev/null 2>&1; then
        # macOS with GNU coreutils: use gtimeout
        gtimeout "${timeout_duration}" "$@"
    else
        # No timeout available, just run the command
        # Note: This is less safe but works on macOS without coreutils
        "$@"
    fi
}

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
} | run_with_timeout 15s ./polymarket-go-mcp 2>&1

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
} | run_with_timeout 15s ./polymarket-go-mcp 2>&1

RESULT=$?
if [ $RESULT -eq 0 ]; then
    echo "✅ get_markets tool test 2 passed"
else
    echo "❌ get_markets tool test 2 failed with exit code: $RESULT"
    exit 1
fi

echo ""
echo "=== get_markets Tool Tests Completed ==="