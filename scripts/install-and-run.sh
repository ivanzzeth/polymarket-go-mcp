#!/bin/bash

# One-liner installer and runner for Polymarket Go MCP
# Usage: curl -sSL https://raw.githubusercontent.com/ivanzzeth/polymarket-go-mcp/main/scripts/install-and-run.sh | bash

set -e

# Configuration
REPO="ivanzzeth/polymarket-go-mcp"
BINARY_NAME="polymarket-go-mcp"
INSTALL_DIR="$HOME/.local/bin"
CACHE_DIR="$HOME/.cache/polymarket-mcp"
CACHE_FILE="$CACHE_DIR/$BINARY_NAME"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

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

# Function to download and install binary
download_and_install() {
    local platform=$1
    local download_url=$(get_download_url "$platform")
    local temp_file=$(mktemp)
    
    print_info "Downloading latest Polymarket MCP binary for $platform..."
    print_info "URL: $download_url"
    
    # Download the binary
    if curl -L -f -s -o "$temp_file" "$download_url"; then
        # Make it executable
        chmod +x "$temp_file"
        
        # Create install directory if it doesn't exist
        mkdir -p "$INSTALL_DIR"
        
        # Install to local bin
        mv "$temp_file" "$INSTALL_DIR/$BINARY_NAME"
        
        print_success "Installed Polymarket MCP to $INSTALL_DIR/$BINARY_NAME"
        
        # Add to PATH if not already there
        if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
            print_warning "Adding $INSTALL_DIR to PATH in ~/.bashrc"
            echo "export PATH=\"\$PATH:$INSTALL_DIR\"" >> "$HOME/.bashrc"
            print_info "Please run: source ~/.bashrc"
        fi
        
        return 0
    else
        rm -f "$temp_file"
        print_error "Failed to download binary from $download_url"
        return 1
    fi
}

# Function to run the MCP server
run_mcp() {
    local platform=$1
    shift
    
    # Check if binary is already installed
    if command -v "$BINARY_NAME" >/dev/null 2>&1; then
        print_info "Found existing installation, running Polymarket MCP Server..."
        # Use env to preserve all environment variables
        exec env "$BINARY_NAME" "$@"
    else
        # Download and run directly from cache
        download_and_install "$platform"
        print_info "Starting Polymarket MCP Server..."
        # Use env to preserve all environment variables
        exec env "$INSTALL_DIR/$BINARY_NAME" "$@"
    fi
}

# Main script execution
main() {
    local platform=$(detect_platform)
    
    if [[ "$platform" == "unknown-unknown" ]]; then
        print_error "Unable to detect platform. Unsupported system."
        exit 1
    fi
    
    print_info "Platform: $platform"
    
    # Run the MCP server
    run_mcp "$platform" "$@"
}

# Run main function with all arguments
main "$@"
