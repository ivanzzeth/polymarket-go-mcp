#!/bin/bash

# Polymarket Go MCP Runner
# Automatically downloads and runs the latest pre-built binary with caching

set -e

# Configuration
REPO="ivanzzeth/polymarket-go-mcp"
BINARY_NAME="polymarket-go-mcp"
CACHE_DIR="$HOME/.cache/polymarket-mcp"
CACHE_FILE="$CACHE_DIR/$BINARY_NAME"
CACHE_TIMESTAMP="$CACHE_DIR/.timestamp"
CACHE_TTL=3600  # 1 hour in seconds

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

# Function to check if cache is valid
is_cache_valid() {
    if [[ ! -f "$CACHE_FILE" ]] || [[ ! -f "$CACHE_TIMESTAMP" ]]; then
        return 1
    fi
    
    local current_time=$(date +%s)
    local cache_time=$(cat "$CACHE_TIMESTAMP" 2>/dev/null || echo 0)
    local time_diff=$((current_time - cache_time))
    
    if [[ $time_diff -lt $CACHE_TTL ]]; then
        return 0
    else
        return 1
    fi
}

# Function to download latest binary
download_binary() {
    local platform=$1
    local download_url=$(get_download_url "$platform")
    local temp_file=$(mktemp)
    
    print_info "Downloading latest binary for $platform..."
    print_info "URL: $download_url"
    
    # Download the binary
    if curl -L -f -s -o "$temp_file" "$download_url"; then
        # Make it executable
        chmod +x "$temp_file"
        
        # Create cache directory if it doesn't exist
        mkdir -p "$CACHE_DIR"
        
        # Move to cache
        mv "$temp_file" "$CACHE_FILE"
        
        # Update timestamp
        date +%s > "$CACHE_TIMESTAMP"
        
        print_success "Downloaded and cached latest binary"
        return 0
    else
        rm -f "$temp_file"
        print_error "Failed to download binary from $download_url"
        return 1
    fi
}

# Function to check for updates
check_for_updates() {
    local platform=$1
    
    if is_cache_valid; then
        print_info "Using cached binary (cache valid for $((CACHE_TTL/3600)) hours)"
        return 0
    fi
    
    print_info "Cache expired or missing, checking for updates..."
    download_binary "$platform"
}

# Function to ensure binary is available
ensure_binary() {
    local platform=$1
    
    # Check if we have a cached binary
    if [[ -f "$CACHE_FILE" ]]; then
        if is_cache_valid; then
            print_info "Using cached binary"
            return 0
        else
            print_warning "Cache expired, checking for updates..."
        fi
    fi
    
    # Download the binary
    if ! download_binary "$platform"; then
        print_error "Failed to download binary. Please check your internet connection."
        exit 1
    fi
}

# Function to run the MCP server
run_mcp() {
    local platform=$1
    shift
    local args="$@"
    
    # Ensure binary is available
    ensure_binary "$platform"
    
    # Run the binary
    print_info "Starting Polymarket MCP Server..."
    print_info "Binary: $CACHE_FILE"
    
    # Print environment variables for debugging
    if [[ -n "$DEBUG" ]]; then
        print_info "Environment variables:"
        env | grep -E "(POLYMARKET|DEBUG|LOG)" || print_info "  (No relevant environment variables found)"
    fi
    
    if [[ -n "$args" ]]; then
        print_info "Arguments: $args"
        # Use env to preserve all environment variables
        exec env "$CACHE_FILE" "$@"
    else
        # Use env to preserve all environment variables
        exec env "$CACHE_FILE"
    fi
}

# Function to show usage
usage() {
    echo "Polymarket Go MCP Runner"
    echo ""
    echo "Usage: $0 [options]"
    echo ""
    echo "Options:"
    echo "  --help, -h          Show this help message"
    echo "  --version, -v       Show version information"
    echo "  --update, -u        Force update the binary"
    echo "  --cache-info        Show cache information"
    echo "  --clear-cache       Clear the cache"
    echo "  --platform PLATFORM Specify platform (linux-amd64, darwin-arm64, etc.)"
    echo ""
    echo "Examples:"
    echo "  $0                    # Run with auto-detected platform"
    echo "  $0 --update           # Force update and run"
    echo "  $0 --platform linux-amd64  # Run with specific platform"
    echo "  $0 --help             # Show help"
    echo ""
    echo "This script automatically downloads and runs the latest pre-built"
    echo "Polymarket MCP server binary with intelligent caching."
}

# Function to show version
show_version() {
    if [[ -f "$CACHE_FILE" ]]; then
        print_info "Cached binary version:"
        "$CACHE_FILE" --version 2>/dev/null || echo "  Version information not available"
    else
        print_info "No cached binary found"
    fi
    
    print_info "Latest release:"
    local latest_url="https://github.com/$REPO/releases/latest"
    local latest_tag=$(curl -s -I "$latest_url" | grep -i "location:" | sed 's/.*\/tag\///' | tr -d '\r')
    echo "  $latest_tag"
}

# Function to show cache info
show_cache_info() {
    if [[ -f "$CACHE_FILE" ]]; then
        local size=$(du -h "$CACHE_FILE" | cut -f1)
        local timestamp=$(cat "$CACHE_TIMESTAMP" 2>/dev/null || echo "unknown")
        local date_str=$(date -d "@$timestamp" 2>/dev/null || echo "unknown")
        
        print_info "Cache Information:"
        echo "  Location: $CACHE_FILE"
        echo "  Size: $size"
        echo "  Cached: $date_str"
        echo "  Platform: $(detect_platform)"
        
        if is_cache_valid; then
            local current_time=$(date +%s)
            local cache_time=$(cat "$CACHE_TIMESTAMP")
            local time_diff=$((current_time - cache_time))
            local time_remaining=$((CACHE_TTL - time_diff))
            local hours=$((time_remaining / 3600))
            local minutes=$(( (time_remaining % 3600) / 60 ))
            
            print_success "Cache is valid for ${hours}h ${minutes}m"
        else
            print_warning "Cache is expired or invalid"
        fi
    else
        print_info "No cached binary found"
    fi
}

# Function to clear cache
clear_cache() {
    if [[ -f "$CACHE_FILE" ]]; then
        rm -f "$CACHE_FILE"
        print_success "Removed cached binary: $CACHE_FILE"
    fi
    
    if [[ -f "$CACHE_TIMESTAMP" ]]; then
        rm -f "$CACHE_TIMESTAMP"
        print_success "Removed cache timestamp"
    fi
    
    if [[ -d "$CACHE_DIR" ]] && [[ -z "$(ls -A "$CACHE_DIR")" ]]; then
        rmdir "$CACHE_DIR"
        print_success "Removed empty cache directory"
    fi
}

# Main script execution
main() {
    local platform
    local force_update=false
    
    # Parse command line arguments
    while [[ $# -gt 0 ]]; do
        case $1 in
            --help|-h)
                usage
                exit 0
                ;;
            --version|-v)
                show_version
                exit 0
                ;;
            --update|-u)
                force_update=true
                shift
                ;;
            --cache-info)
                show_cache_info
                exit 0
                ;;
            --clear-cache)
                clear_cache
                exit 0
                ;;
            --platform)
                if [[ -n "$2" ]]; then
                    platform="$2"
                    shift 2
                else
                    print_error "Platform argument requires a value"
                    exit 1
                fi
                ;;
            *)
                # Unknown argument, pass to binary
                break
                ;;
        esac
    done
    
    # Detect platform if not specified
    if [[ -z "$platform" ]]; then
        platform=$(detect_platform)
        if [[ "$platform" == "unknown-unknown" ]]; then
            print_error "Unable to detect platform. Please specify with --platform"
            exit 1
        fi
        print_info "Detected platform: $platform"
    fi
    
    # Force update if requested
    if [[ "$force_update" == "true" ]]; then
        print_info "Forcing update..."
        clear_cache
    fi
    
    # Run the MCP server
    run_mcp "$platform" "$@"
}

# Run main function with all arguments
main "$@"
