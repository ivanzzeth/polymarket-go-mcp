#!/bin/bash

# Release script for polymarket-go-mcp
# Automates version updates, git tagging, and release process

set -e

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

# Function to validate semantic version
validate_version() {
    local version=$1
    if [[ ! $version =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
        print_error "Invalid version format: $version"
        print_error "Version must follow semantic versioning: vMAJOR.MINOR.PATCH (e.g., v1.0.0)"
        exit 1
    fi
}

# Function to get current version from constants
get_current_version() {
    local version_line=$(grep -E 'Version\s*=\s*"[0-9]+\.[0-9]+\.[0-9]+"' constants/constants.go)
    if [[ $version_line =~ \"([0-9]+\.[0-9]+\.[0-9]+)\" ]]; then
        echo "v${BASH_REMATCH[1]}"
    else
        print_error "Could not find current version in constants/constants.go"
        exit 1
    fi
}

# Function to update version in constants file
update_version() {
    local old_version=$1
    local new_version=$2
    
    print_info "Updating version from $old_version to $new_version in constants/constants.go"
    
    # Remove 'v' prefix for the constants file
    local version_without_v=${new_version#v}
    
    # Debug: show what we're looking for
    print_info "Looking for version pattern in constants/constants.go"
    grep -n "Version" constants/constants.go
    
    # Update the version in constants file
    sed -i "s/Version = \"[0-9]\+\.[0-9]\+\.[0-9]\+\"/Version = \"$version_without_v\"/" constants/constants.go
    
    # Debug: show the updated file
    print_info "Updated constants/constants.go:"
    grep -n "Version" constants/constants.go
    
    # Verify the update
    local updated_version=$(get_current_version)
    if [[ "$updated_version" == "$new_version" ]]; then
        print_success "Version updated successfully in constants/constants.go"
    else
        print_error "Failed to update version. Expected: $new_version, Got: $updated_version"
        print_error "Current file content:"
        grep "Version" constants/constants.go
        exit 1
    fi
}

# Function to create changelog entry
create_changelog_entry() {
    local version=$1
    local message=$2
    
    if [[ ! -f CHANGELOG.md ]]; then
        print_warning "CHANGELOG.md not found, creating new one"
        cat > CHANGELOG.md << EOF
# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## $version - $(date +%Y-%m-%d)

### Added
- $message

EOF
        print_success "Created CHANGELOG.md"
    else
        # Insert new version section after [Unreleased]
        sed -i "/## \[Unreleased\]/a \\
\\
## $version - $(date +%Y-%m-%d)\\
\\
### Added\\
- $message\\
" CHANGELOG.md
        print_success "Updated CHANGELOG.md"
    fi
}

# Function to check if working directory is clean
check_clean_working_dir() {
    if [[ -n $(git status --porcelain) ]]; then
        print_error "Working directory is not clean. Please commit or stash changes before releasing."
        git status --porcelain
        exit 1
    fi
    print_success "Working directory is clean"
}

# Function to build and test
build_and_test() {
    print_info "Building and testing the project..."
    
    # Build the project
    if go build -o polymarket-go-mcp .; then
        print_success "Build successful"
    else
        print_error "Build failed"
        exit 1
    fi
    
    # Run tests
    if go test ./...; then
        print_success "Tests passed"
    else
        print_error "Tests failed"
        exit 1
    fi
    
    # Clean up build artifact
    rm -f polymarket-go-mcp
}

# Main release function
release() {
    local new_version=$1
    local release_message=$2
    
    print_info "Starting release process for version: $new_version"
    
    # Validate version format
    validate_version "$new_version"
    
    # Get current version
    local current_version=$(get_current_version)
    print_info "Current version: $current_version"
    
    # Check if version is actually changing
    if [[ "$current_version" == "$new_version" ]]; then
        print_warning "Version is already $new_version. No changes needed."
        exit 0
    fi
    
    # Check working directory
    check_clean_working_dir
    
    # Build and test
    build_and_test
    
    # Update version
    update_version "$current_version" "$new_version"
    
    # Create changelog entry
    create_changelog_entry "$new_version" "$release_message"
    
    # Commit changes
    print_info "Committing version changes..."
    git add constants/constants.go CHANGELOG.md
    git commit -m "chore: release $new_version
    
$release_message"
    
    # Create git tag
    print_info "Creating git tag: $new_version"
    git tag -a "$new_version" -m "Release $new_version
    
$release_message"
    
    print_success "Release preparation completed!"
    print_info ""
    print_info "Next steps:"
    print_info "1. Review the changes: git log --oneline -n 3"
    print_info "2. Push the changes: git push origin main"
    print_info "3. Push the tag: git push origin $new_version"
    print_info ""
    print_info "After pushing the tag, GitHub Actions will automatically:"
    print_info "- Build binaries for all platforms"
    print_info "- Create a GitHub Release"
    print_info "- Upload pre-built binaries"
}

# Function to show usage
usage() {
    echo "Usage: $0 <version> <release_message>"
    echo ""
    echo "Examples:"
    echo "  $0 v1.0.0 \"Initial release with basic market data tools\""
    echo "  $0 v1.1.0 \"Add portfolio tracking and PnL calculation\""
    echo "  $0 v1.0.1 \"Fix bug in market search functionality\""
    echo ""
    echo "Version must follow semantic versioning: vMAJOR.MINOR.PATCH"
    echo "  - MAJOR: Breaking changes"
    echo "  - MINOR: New features (backward compatible)"
    echo "  - PATCH: Bug fixes (backward compatible)"
    exit 1
}

# Main script execution
if [[ $# -lt 2 ]]; then
    usage
fi

new_version=$1
shift
release_message="$*"

release "$new_version" "$release_message"
