#!/bin/bash

echo "=== Listing All Available MCP Tools ==="

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
echo ""

# List all tools
echo "=== Available Tools ==="
OUTPUT=$({
    # Initialize
    echo '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{"roots":{"listChanged":true},"tools":{"listChanged":true}},"clientInfo":{"name":"tool-lister","version":"1.0.0"}}}'
    sleep 0.5
    
    # List tools
    echo '{"jsonrpc":"2.0","id":2,"method":"tools/list"}'
    sleep 1
} | run_with_timeout 10s ./polymarket-go-mcp 2>&1)

# Extract and format the tools list
if command -v jq >/dev/null 2>&1; then
    # Use jq to format JSON nicely
    echo "$OUTPUT" | grep '"result"' | jq -r '.result.tools[] | "  • \(.name): \(.description)"' 2>/dev/null || echo "$OUTPUT" | grep -A 1000 '"result"'
else
    # Fallback: just show the raw JSON result
    echo "$OUTPUT" | grep -A 1000 '"result"'
fi

echo ""
echo "=== Tool Listing Completed ==="

