#!/bin/bash

# NPX-style runner for Polymarket Go MCP
# Usage: curl -sSL https://raw.githubusercontent.com/ivanzzeth/polymarket-go-mcp/main/scripts/npx-style.sh | bash -s -- [args]

set -e

# Configuration
REPO="ivanzzeth/polymarket-go-mcp"
BINARY_NAME="polymarket-go-mcp"
TEMP_DIR=$(mktemp -d)
BINARY_PATH="$TEMP_DIR/$BINARY_NAME"

# Cleanup function
cleanup() {
    rm -rf "$TEMP_DIR"
}
trap cleanup EXIT

# Function to detect platform
detect_platform() {
    local os
    local arch
    
    case "$(uname -s)" in
        Linux*)     os="linux" ;;
        Darwin*)    os="darwin" ;;
        CYGWIN*|MINGW*|MSYS*) os="windows" ;;
        *)          os="unknown" ;;
    esac
    
    case "$(uname -m)" in
        x86_64)     arch="amd64" ;;
        arm64|aarch64) arch="arm64" ;;
        *)          arch="unknown" ;;
    esac
    
    echo "${os}-${arch}"
}

# Function to get download URL for platform
get_download_url() {
    local platform=$1
    local extension=""
    
    if [[ $platform == *"windows"* ]]; then
        extension=".exe"
    fi
    
    echo "https://github.com/$REPO/releases/latest/download/$BINARY_NAME-$platform$extension"
}

# Function to download and run binary
download_and_run() {
    local platform=$1
    shift
    local download_url=$(get_download_url "$platform")
    
    echo "🚀 Downloading latest Polymarket MCP binary..."
    echo "📦 Platform: $platform"
    echo "🔗 URL: $download_url"
    
    # Download the binary
    if curl -L -f -s -o "$BINARY_PATH" "$download_url"; then
        # Make it executable
        chmod +x "$BINARY_PATH"
        
        echo "✅ Downloaded successfully"
        echo "🚀 Starting Polymarket MCP Server..."
        
        # Print environment variables for debugging
        if [[ -n "$DEBUG" ]]; then
            echo "🔧 Environment variables:"
            env | grep -E "(POLYMARKET|DEBUG|LOG)" || echo "  (No relevant environment variables found)"
        fi
        
        # Run the binary with all remaining arguments and preserve environment
        exec env "$BINARY_PATH" "$@"
    else
        echo "❌ Failed to download binary"
        exit 1
    fi
}

# Main script execution
main() {
    local platform=$(detect_platform)
    
    if [[ "$platform" == "unknown-unknown" ]]; then
        echo "❌ Unsupported platform"
        exit 1
    fi
    
    # Download and run
    download_and_run "$platform" "$@"
}

# Run main function with all arguments
main "$@"
