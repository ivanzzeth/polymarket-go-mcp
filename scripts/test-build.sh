#!/bin/bash

# Test build script for GitHub Actions workflow
# This script simulates the build process used in the GitHub Actions workflow

set -e

echo "Testing build process for polymarket-go-mcp"

# Clean previous builds
rm -rf dist
mkdir -p dist

# Test building for different platforms
platforms=(
    "linux amd64 linux-amd64"
    "linux arm64 linux-arm64" 
    "darwin amd64 darwin-amd64"
    "darwin arm64 darwin-arm64"
    "windows amd64 win-amd64 .exe"
)

for platform in "${platforms[@]}"; do
    IFS=' ' read -r -a parts <<< "$platform"
    os="${parts[0]}"
    arch="${parts[1]}"
    target="${parts[2]}"
    ext="${parts[3]:-}"
    
    echo "Building for $os/$arch -> $target$ext"
    
    GOOS="$os" GOARCH="$arch" go build -o "polymarket-go-mcp$ext" .
    
    if [ -f "polymarket-go-mcp$ext" ]; then
        mv "polymarket-go-mcp$ext" "dist/polymarket-go-mcp-$target$ext"
        echo "✓ Successfully built polymarket-go-mcp-$target$ext"
        
        # Test the binary if not Windows
        if [ "$os" != "windows" ]; then
            chmod +x "dist/polymarket-go-mcp-$target$ext"
            echo "✓ Binary is executable"
        fi
    else
        echo "✗ Failed to build for $target"
        exit 1
    fi
done

echo ""
echo "All builds completed successfully!"
echo "Built binaries:"
ls -la dist/
