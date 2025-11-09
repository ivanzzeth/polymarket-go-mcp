#!/bin/bash

echo "=== Testing health_check Tool ==="

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
} | run_with_timeout 10s ./polymarket-go-mcp 2>&1

RESULT=$?
if [ $RESULT -eq 0 ]; then
    echo "✅ health_check tool test passed"
else
    echo "❌ health_check tool test failed with exit code: $RESULT"
    exit 1
fi

echo ""
echo "=== health_check Tool Test Completed ==="