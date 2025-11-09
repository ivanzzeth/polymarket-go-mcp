# Version Management Guide

This document describes the version management best practices and automated release process for polymarket-go-mcp.

## Overview

The project uses a comprehensive version management system that includes:

- **Semantic Versioning** (SemVer) for consistent version numbering
- **Automated release scripts** for version updates and git tagging
- **GitHub Actions integration** for automated binary builds and releases
- **Changelog management** for tracking changes

## Release Script

The main release automation script is located at `scripts/release.sh`.

### Usage

```bash
# Make the script executable
chmod +x scripts/release.sh

# Run a release
./scripts/release.sh <version> "<release_message>"
```

### Examples

```bash
# Major release with breaking changes
./scripts/release.sh v2.0.0 "Complete rewrite with new API structure"

# Minor release with new features
./scripts/release.sh v1.1.0 "Add portfolio tracking and PnL calculation"

# Patch release with bug fixes
./scripts/release.sh v1.0.1 "Fix bug in market search functionality"
```

## Release Process

### 1. Pre-release Checklist

Before running a release, ensure:

- [ ] All tests pass: `go test ./...`
- [ ] Code builds successfully: `go build .`
- [ ] Working directory is clean: `git status`
- [ ] Changelog is updated with recent changes
- [ ] Documentation is up to date

### 2. Automated Release Steps

The release script performs the following actions:

1. **Version Validation**
   - Validates semantic version format (vMAJOR.MINOR.PATCH)
   - Checks if version is actually changing

2. **Working Directory Check**
   - Ensures no uncommitted changes exist
   - Prevents accidental releases from dirty state

3. **Build and Test**
   - Builds the project to verify compilation
   - Runs all tests to ensure quality

4. **Version Update**
   - Updates `constants/constants.go` with new version
   - Verifies the update was successful

5. **Changelog Management**
   - Creates or updates `CHANGELOG.md`
   - Adds new version section with release message
   - Follows Keep a Changelog format

6. **Git Operations**
   - Commits version and changelog changes
   - Creates annotated git tag with release message

### 3. Post-release Steps

After the script completes:

1. **Review Changes**
   ```bash
   git log --oneline -n 3
   git show <tag_name>
   ```

2. **Push Changes**
   ```bash
   git push origin main
   git push origin <tag_name>
   ```

3. **Monitor GitHub Actions**
   - Check the Actions tab in GitHub repository
   - Verify build and release process completes successfully
   - Confirm binaries are available in GitHub Releases

## Semantic Versioning

The project follows [Semantic Versioning 2.0.0](https://semver.org/):

- **MAJOR** version (X.0.0): Incompatible API changes
- **MINOR** version (0.X.0): New functionality in backward compatible manner
- **PATCH** version (0.0.X): Backward compatible bug fixes

### Version Examples

- `v1.0.0`: Initial stable release
- `v1.1.0`: Added new features (backward compatible)
- `v1.1.1`: Bug fixes (backward compatible)
- `v2.0.0`: Breaking changes (incompatible API)

## Changelog Format

The changelog follows the [Keep a Changelog](https://keepachangelog.com/) format:

```markdown
# Changelog

## [Unreleased]

## [1.1.0] - 2024-01-15

### Added
- Portfolio tracking functionality
- PnL calculation tools

### Changed
- Improved error handling in market data API

### Fixed
- Memory leak in long-running sessions
```

## GitHub Actions Integration

When a git tag is pushed, GitHub Actions automatically:

1. **Builds binaries** for all supported platforms:
   - Linux AMD64
   - Linux ARM64
   - macOS AMD64
   - macOS ARM64
   - Windows AMD64

2. **Creates GitHub Release** with:
   - Automatic release notes generation
   - All pre-built binaries attached
   - Tag-based versioning

3. **Makes binaries available** at:
   ```
   https://github.com/ivanzzeth/polymarket-go-mcp/releases/latest/download/polymarket-go-mcp-<platform>
   ```

## Best Practices

### Version Management

1. **Use Semantic Versioning** consistently
2. **Update version constants** in code before release
3. **Maintain comprehensive changelog** for all releases
4. **Test thoroughly** before releasing
5. **Use descriptive release messages**

### Release Frequency

- **Patch releases**: As needed for critical bug fixes
- **Minor releases**: Monthly for new features
- **Major releases**: When breaking changes are necessary

### Quality Assurance

- Always run tests before release
- Verify builds on multiple platforms when possible
- Review changelog entries for accuracy
- Test the actual binary in target environments

## Troubleshooting

### Common Issues

1. **Version validation fails**
   - Ensure version follows `vX.Y.Z` format
   - Check for typos in version string

2. **Working directory not clean**
   - Commit or stash changes before release
   - Use `git status` to check current state

3. **Build or test failures**
   - Fix compilation errors before release
   - Ensure all tests pass

4. **GitHub Actions failure**
   - Check Actions tab for detailed error logs
   - Verify repository permissions for releases

### Recovery Steps

If a release fails:

1. **Delete the local tag** (if created):
   ```bash
   git tag -d <tag_name>
   ```

2. **Reset the commit** (if needed):
   ```bash
   git reset --hard HEAD~1
   ```

3. **Fix the issues** and retry the release process

## Related Files

- `scripts/release.sh` - Main release automation script
- `constants/constants.go` - Version constants definition
- `CHANGELOG.md` - Project changelog (auto-generated)
- `.github/workflows/release.yml` - GitHub Actions release workflow
- `scripts/test-build.sh` - Build verification script

## Additional Tools

### GitHub CLI

You can also use GitHub CLI for release management:

```bash
# Create release with GitHub CLI
gh release create v1.0.0 --title "v1.0.0" --notes "Initial release"

# View releases
gh release list

# Download assets
gh release download v1.0.0
```

This comprehensive version management system ensures consistent, automated, and reliable releases for the polymarket-go-mcp project.
