# 部署策略

## 概述

本文档详细说明polymarket-mcp服务器的多种部署策略，确保用户能够选择最适合其环境的部署方式。

## 部署方案对比

| 部署方式 | 用户体验 | 技术要求 | 维护复杂度 | 推荐场景 |
|---------|----------|----------|------------|----------|
| **npm包** | ⭐⭐⭐⭐⭐ | Node.js | 低 | 个人用户、开发环境 |
| **Docker** | ⭐⭐⭐⭐ | Docker | 中 | 生产环境、团队使用 |
| **预编译二进制** | ⭐⭐⭐ | 无 | 低 | 系统管理员、CI/CD |
| **源码编译** | ⭐⭐ | Go工具链 | 高 | 开发者、定制需求 |

## 方案1：npm包部署（推荐）

### 优势
- 真正的npx体验
- 自动版本管理
- 跨平台支持
- 用户熟悉

### 配置示例

**Claude Desktop配置：**
```json
{
  "mcpServers": {
    "polymarket": {
      "command": "npx",
      "args": ["-y", "polymarket-mcp-server"]
    }
  }
}
```

**Windsurf配置：**
```json
{
  "mcpServers": {
    "polymarket": {
      "command": "npx",
      "args": ["-y", "polymarket-mcp-server"]
    }
  }
}
```

### 版本管理
```json
// 使用特定版本
{
  "command": "npx",
  "args": ["-y", "polymarket-mcp-server@1.2.3"]
}

// 使用最新版本
{
  "command": "npx", 
  "args": ["-y", "polymarket-mcp-server@latest"]
}
```

## 方案2：Docker部署

### 优势
- 环境隔离
- 一致性保证
- 易于扩展
- 生产就绪

### Docker镜像构建

```dockerfile
# Dockerfile
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

### 多架构Docker构建

```bash
#!/bin/bash
# build-docker.sh

set -e

VERSION=${1:-latest}
PLATFORMS="linux/amd64,linux/arm64"

echo "🚀 Building Docker image for $PLATFORMS..."

docker buildx create --use --name multiarch-builder

docker buildx build \
  --platform $PLATFORMS \
  --tag ivanzzeth/polymarket-mcp-server:$VERSION \
  --tag ivanzzeth/polymarket-mcp-server:latest \
  --push .

echo "✅ Docker image built and pushed successfully!"
```

### 配置示例

**Claude Desktop配置：**
```json
{
  "mcpServers": {
    "polymarket": {
      "command": "docker",
      "args": [
        "run", "--rm", "-i",
        "ivanzzeth/polymarket-mcp-server:latest"
      ]
    }
  }
}
```

**生产环境配置：**
```json
{
  "mcpServers": {
    "polymarket": {
      "command": "docker",
      "args": [
        "run", "--rm", "-i",
        "--memory=256m",
        "--cpus=0.5",
        "ivanzzeth/polymarket-mcp-server:latest"
      ]
    }
  }
}
```

## 方案3：预编译二进制部署

### 优势
- 无需额外依赖
- 性能最佳
- 系统集成友好

### 二进制下载脚本

```bash
#!/bin/bash
# install-binary.sh

set -e

VERSION=${1:-latest}
PLATFORM=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# 架构映射
case $ARCH in
    x86_64) ARCH="amd64" ;;
    aarch64) ARCH="arm64" ;;
    arm64) ARCH="arm64" ;;
    *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

# 平台映射
case $PLATFORM in
    linux) PLATFORM="linux" ;;
    darwin) PLATFORM="darwin" ;;
    *) echo "Unsupported platform: $PLATFORM"; exit 1 ;;
esac

BINARY_NAME="polymarket-mcp-server-${PLATFORM}-${ARCH}"
if [ "$PLATFORM" = "windows" ]; then
    BINARY_NAME="${BINARY_NAME}.exe"
fi

DOWNLOAD_URL="https://github.com/ivanzzeth/polymarket-go-mcp/releases/${VERSION}/download/${BINARY_NAME}"

echo "📥 Downloading polymarket-mcp-server ${VERSION} for ${PLATFORM}/${ARCH}..."
curl -L -o polymarket-mcp-server "$DOWNLOAD_URL"
chmod +x polymarket-mcp-server

echo "✅ Downloaded to: $(pwd)/polymarket-mcp-server"
```

### 配置示例

**系统级安装：**
```json
{
  "mcpServers": {
    "polymarket": {
      "command": "/usr/local/bin/polymarket-mcp-server"
    }
  }
}
```

**用户目录安装：**
```json
{
  "mcpServers": {
    "polymarket": {
      "command": "/home/user/.local/bin/polymarket-mcp-server"
    }
  }
}
```

## 方案4：源码编译部署

### 优势
- 完全控制
- 定制化构建
- 开发环境友好

### 构建步骤

```bash
# 克隆仓库
git clone https://github.com/ivanzzeth/polymarket-go-mcp
cd polymarket-go-mcp

# 安装依赖
go mod download

# 构建
go build -o polymarket-mcp-server

# 测试
./polymarket-mcp-server --version
```

### 配置示例

```json
{
  "mcpServers": {
    "polymarket": {
      "command": "/path/to/polymarket-go-mcp/polymarket-mcp-server"
    }
  }
}
```

## 环境特定配置

### 开发环境
```json
{
  "mcpServers": {
    "polymarket": {
      "command": "npx",
      "args": ["-y", "polymarket-mcp-server"],
      "env": {
        "DEBUG": "true",
        "LOG_LEVEL": "debug"
      }
    }
  }
}
```

### 生产环境
```json
{
  "mcpServers": {
    "polymarket": {
      "command": "docker",
      "args": [
        "run", "--rm", "-i",
        "--memory=512m",
        "--cpus=1.0",
        "--env=LOG_LEVEL=info",
        "ivanzzeth/polymarket-mcp-server:latest"
      ]
    }
  }
}
```

### CI/CD环境
```json
{
  "mcpServers": {
    "polymarket": {
      "command": "/opt/polymarket-mcp-server",
      "env": {
        "LOG_LEVEL": "warn"
      }
    }
  }
}
```

## 自动部署流程

### GitHub Actions自动构建

```yaml
name: Multi-Platform Deployment

on:
  push:
    branches: [main]
    tags: ['v*']
  release:
    types: [published]

jobs:
  build-binaries:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        goos: [linux, darwin, windows]
        goarch: [amd64, arm64]
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v4
        with:
          go-version: '1.24'
      
      - name: Build
        run: |
          GOOS=${{ matrix.goos }} GOARCH=${{ matrix.goarch }} go build -o polymarket-mcp-server-${{ matrix.goos }}-${{ matrix.goarch }}
          if [ "${{ matrix.goos }}" = "windows" ]; then
            mv polymarket-mcp-server-${{ matrix.goos }}-${{ matrix.goarch }} polymarket-mcp-server-${{ matrix.goos }}-${{ matrix.goarch }}.exe
          fi
          
      - name: Upload artifacts
        uses: actions/upload-artifact@v4
        with:
          name: polymarket-mcp-server-${{ matrix.goos }}-${{ matrix.goarch }}
          path: polymarket-mcp-server-${{ matrix.goos }}-${{ matrix.goarch }}*

  build-docker:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v3
        
      - name: Login to Docker Hub
        uses: docker/login-action@v3
        with:
          username: ${{ secrets.DOCKER_USERNAME }}
          password: ${{ secrets.DOCKER_PASSWORD }}
          
      - name: Build and push
        uses: docker/build-push-action@v5
        with:
          context: .
          platforms: linux/amd64,linux/arm64
          push: true
          tags: |
            ivanzzeth/polymarket-mcp-server:latest
            ivanzzeth/polymarket-mcp-server:${{ github.ref_name }}

  publish-npm:
    runs-on: ubuntu-latest
    needs: [build-binaries]
    steps:
      - uses: actions/checkout@v4
      
      - name: Setup Node.js
        uses: actions/setup-node@v4
        with:
          node-version: '18'
          registry-url: 'https://registry.npmjs.org'
          
      - name: Download binaries
        uses: actions/download-artifact@v4
        with:
          path: artifacts
          
      - name: Prepare npm package
        run: |
          mkdir -p polymarket-mcp-server-npm/bin
          cp -r artifacts/* polymarket-mcp-server-npm/bin/
          cp scripts/* polymarket-mcp-server-npm/
          cp package.json polymarket-mcp-server-npm/
          
      - name: Publish to npm
        run: |
          cd polymarket-mcp-server-npm
          npm publish
        env:
          NODE_AUTH_TOKEN: ${{ secrets.NPM_TOKEN }}
```

## 监控和日志

### 健康检查
```bash
# 检查服务器状态
polymarket-mcp-server --health

# 检查版本
polymarket-mcp-server --version

# 查看帮助
polymarket-mcp-server --help
```

### 日志配置
```json
{
  "mcpServers": {
    "polymarket": {
      "command": "npx",
      "args": ["-y", "polymarket-mcp-server"],
      "env": {
        "LOG_LEVEL": "info",
        "LOG_FORMAT": "json"
      }
    }
  }
}
```

## 推荐部署策略

### 个人用户
- **首选**: npm包方式
- **备选**: 预编译二进制

### 团队/企业
- **首选**: Docker方式
- **备选**: 预编译二进制

### 开发者
- **首选**: 源码编译
- **备选**: npm包方式

通过提供多种部署方案，polymarket-mcp服务器能够适应各种使用场景和环境需求，确保最佳的用户体验和系统兼容性。