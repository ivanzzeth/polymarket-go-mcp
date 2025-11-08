# npx风格实现方案

## 概述

本文档详细说明如何实现真正的npx风格体验，让用户能够通过简单的MCP配置直接使用polymarket-mcp服务器，无需任何手动安装步骤。

## 目标

用户只需要在MCP配置中添加：
```json
{
  "polymarket-mcp": {
    "command": "npx",
    "args": ["-y", "polymarket-mcp-server"]
  }
}
```

然后就能立即使用，无需任何手动安装。

## 技术方案

### 方案1：npm包发布（推荐）

#### 包结构
```
polymarket-mcp-server/
├── package.json
├── bin/
│   ├── polymarket-mcp-server-linux
│   ├── polymarket-mcp-server-macos
│   ├── polymarket-mcp-server-win.exe
│   └── index.js              # 平台检测脚本
└── README.md
```

#### package.json配置
```json
{
  "name": "polymarket-mcp-server",
  "version": "1.0.0",
  "description": "Polymarket MCP Server - Access prediction market data through MCP protocol",
  "bin": {
    "polymarket-mcp-server": "./bin/index.js"
  },
  "files": [
    "bin/"
  ],
  "keywords": [
    "mcp",
    "polymarket",
    "prediction-markets",
    "gamma-api",
    "data-api"
  ],
  "author": "ivanzzeth",
  "license": "MIT",
  "repository": {
    "type": "git",
    "url": "https://github.com/ivanzzeth/polymarket-go-mcp"
  },
  "engines": {
    "node": ">=14"
  }
}
```

#### 平台检测脚本 (bin/index.js)
```javascript
#!/usr/bin/env node
const { platform, arch } = process;
const path = require('path');
const { spawn } = require('child_process');
const { existsSync } = require('fs');

// 根据平台选择正确的二进制
function getBinaryName() {
    if (platform === 'win32') {
        return 'polymarket-mcp-server-win.exe';
    } else if (platform === 'darwin') {
        return 'polymarket-mcp-server-macos';
    } else {
        return 'polymarket-mcp-server-linux';
    }
}

function getBinaryPath() {
    const binaryName = getBinaryName();
    return path.join(__dirname, binaryName);
}

function main() {
    const binaryPath = getBinaryPath();
    
    if (!existsSync(binaryPath)) {
        console.error(`Error: Binary not found for platform ${platform}`);
        console.error(`Expected path: ${binaryPath}`);
        process.exit(1);
    }

    // 设置执行权限（非Windows平台）
    if (platform !== 'win32') {
        const { chmodSync } = require('fs');
        try {
            chmodSync(binaryPath, '755');
        } catch (error) {
            // 权限设置失败不影响执行
        }
    }

    // 启动二进制
    const child = spawn(binaryPath, process.argv.slice(2), {
        stdio: 'inherit',
        env: process.env
    });

    child.on('close', (code) => {
        process.exit(code);
    });

    child.on('error', (error) => {
        console.error('Failed to start polymarket-mcp-server:', error);
        process.exit(1);
    });
}

main();
```

### 方案2：Docker方式

#### Docker配置
```dockerfile
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o polymarket-mcp-server

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/polymarket-mcp-server .
ENTRYPOINT ["./polymarket-mcp-server"]
```

#### 用户配置
```json
{
  "polymarket-mcp": {
    "command": "docker",
    "args": [
      "run", "--rm", "-i",
      "ivanzzeth/polymarket-mcp-server:latest"
    ]
  }
}
```

### 方案3：远程脚本执行

#### 运行脚本 (run.sh)
```bash
#!/bin/bash
set -e

# 检测平台
detect_platform() {
    case "$(uname -s)" in
        Linux*)     echo "linux";;
        Darwin*)    echo "darwin";;
        CYGWIN*|MINGW*|MSYS*) echo "windows";;
        *)          echo "unknown"
    esac
}

# 检测架构
detect_arch() {
    case "$(uname -m)" in
        x86_64*)    echo "amd64";;
        aarch64*)   echo "arm64";;
        arm64*)     echo "arm64";;
        *)          echo "unknown"
    esac
}

# 下载二进制
download_binary() {
    local platform=$1
    local arch=$2
    local version=${3:-"latest"}
    
    local download_url="https://github.com/ivanzzeth/polymarket-go-mcp/releases/${version}/download/polymarket-mcp-server-${platform}-${arch}"
    
    if [ "$platform" = "windows" ]; then
        download_url="${download_url}.exe"
    fi
    
    echo "Downloading polymarket-mcp-server for ${platform}-${arch}..."
    curl -L -o polymarket-mcp-server "$download_url"
    chmod +x polymarket-mcp-server
}

main() {
    local platform=$(detect_platform)
    local arch=$(detect_arch)
    
    if [ "$platform" = "unknown" ] || [ "$arch" = "unknown" ]; then
        echo "Unsupported platform: $(uname -s) $(uname -m)"
        exit 1
    fi
    
    download_binary "$platform" "$arch"
    
    # 执行服务器
    exec ./polymarket-mcp-server "$@"
}

main "$@"
```

#### 用户配置
```json
{
  "polymarket-mcp": {
    "command": "bash",
    "args": [
      "-c",
      "curl -fsSL https://raw.githubusercontent.com/ivanzzeth/polymarket-go-mcp/main/run.sh | bash -s -- --stdio"
    ]
  }
}
```

## 构建和发布流程

### 多平台构建
```bash
#!/bin/bash
# build-release.sh

set -e

VERSION=${1:-$(git describe --tags --abbrev=0)}
PLATFORMS=("linux" "darwin" "windows")
ARCHS=("amd64" "arm64")

mkdir -p release

for platform in "${PLATFORMS[@]}"; do
    for arch in "${ARCHS[@]}"; do
        echo "Building for $platform/$arch..."
        
        output_name="polymarket-mcp-server-$platform-$arch"
        if [ "$platform" = "windows" ]; then
            output_name="$output_name.exe"
        fi
        
        GOOS=$platform GOARCH=$arch go build -o "release/$output_name"
        
        # 创建checksum
        shasum -a 256 "release/$output_name" > "release/$output_name.sha256"
    done
done

echo "Build complete. Files in release/ directory"
```

### GitHub Actions自动发布
```yaml
name: Release

on:
  push:
    tags:
      - 'v*'

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.24'
          
      - name: Build
        run: ./build-release.sh
        
      - name: Create Release
        uses: softprops/action-gh-release@v1
        with:
          files: release/*
```

## 推荐方案

**首选方案：npm包发布**
- 真正的npx体验
- 自动版本管理
- 跨平台支持
- 用户熟悉

**备选方案：Docker**
- 环境隔离
- 一致性保证
- 但需要用户安装Docker

## 下一步行动

1. 实现多平台构建脚本
2. 创建npm包结构
3. 设置自动发布流程
4. 测试各平台兼容性
5. 发布到npm registry

通过此方案，用户将获得真正的npx风格体验，只需简单配置即可立即使用polymarket-mcp服务器。