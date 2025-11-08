# Go开发者快速指南

## 概述

本文档专门为Go开发者提供polymarket-go-mcp服务器的快速安装和使用指南，重点介绍Go开发者喜欢的原生安装方式。

## 为什么选择Go原生方式？

作为Go开发者，您可能更喜欢：
- 使用熟悉的Go工具链
- 直接从源码编译，确保最新功能
- 避免额外的依赖（如Node.js、Docker）
- 更好的性能和资源控制

## 快速安装

### 方式1：go install（推荐）

这是最简单的方式，直接从GitHub安装最新版本：

```bash
go install github.com/ivanzzeth/polymarket-go-mcp/cmd/polymarket-go-mcp@latest
```

安装完成后，二进制文件将位于：
- `$GOPATH/bin/polymarket-go-mcp`（如果设置了GOPATH）
- `$GOBIN/polymarket-go-mcp`（如果设置了GOBIN）
- `~/go/bin/polymarket-go-mcp`（默认位置）

### 方式2：源码编译

如果您需要自定义构建或开发：

```bash
# 克隆仓库
git clone https://github.com/ivanzzeth/polymarket-go-mcp
cd polymarket-go-mcp

# 安装依赖
go mod download

# 构建
go build -o polymarket-go-mcp .

# 验证构建
./polymarket-go-mcp --version
```

### 方式3：多平台交叉编译

如果您需要为其他平台构建：

```bash
# Linux AMD64
GOOS=linux GOARCH=amd64 go build -o polymarket-go-mcp-linux-amd64 .

# macOS ARM64
GOOS=darwin GOARCH=arm64 go build -o polymarket-go-mcp-darwin-arm64 .

# Windows AMD64
GOOS=windows GOARCH=amd64 go build -o polymarket-go-mcp-win-amd64.exe .
```

## 配置MCP客户端

### Claude Desktop配置

```json
{
  "mcpServers": {
    "polymarket": {
      "command": "polymarket-go-mcp"
    }
  }
}
```

### 确保二进制在PATH中

如果使用`go install`，确保Go的bin目录在PATH中：

```bash
# 检查是否在PATH中
which polymarket-go-mcp

# 如果不在PATH中，添加到PATH
export PATH=$PATH:$(go env GOPATH)/bin

# 永久添加到shell配置文件中
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc  # 或 ~/.zshrc
```

### 使用绝对路径

如果不想修改PATH，可以使用绝对路径：

```json
{
  "mcpServers": {
    "polymarket": {
      "command": "/home/user/go/bin/polymarket-go-mcp"
    }
  }
}
```

## 开发环境设置

### 项目结构

```
polymarket-go-mcp/
├── main.go                  # 主程序入口
├── go.mod                   # Go模块定义
├── go.sum                   # 依赖锁定文件
├── README.md               # 项目说明
├── test_server.sh          # 测试脚本
├── clients/                # API客户端
│   └── gamma_client.go     # Gamma客户端实现
├── tools/                  # MCP工具实现
│   ├── gamma/              # Gamma相关工具
│   │   └── markets.go      # 市场数据工具
│   └── data/               # Data相关工具
│       └── placeholder.go  # 数据工具占位符
├── resources/              # MCP资源实现
│   └── market_resource.go  # 市场资源
├── types/                  # 类型定义
│   └── tool_args.go        # 工具参数类型
└── docs/                   # 完整文档
```

### 开发工作流

```bash
# 1. 克隆项目
git clone https://github.com/ivanzzeth/polymarket-go-mcp
cd polymarket-go-mcp

# 2. 安装依赖
go mod download

# 3. 运行测试
go test ./...

# 4. 构建并测试
go build -o polymarket-go-mcp .
./polymarket-go-mcp --help

# 5. 开发时实时构建和测试
go run . --version
```

### 调试配置

创建开发环境配置：

```json
{
  "mcpServers": {
    "polymarket": {
      "command": "go",
      "args": ["run", "."],
      "env": {
        "DEBUG": "true",
        "LOG_LEVEL": "debug"
      }
    }
  }
}
```

## 高级用法

### 环境变量配置

```bash
# 设置环境变量
export LOG_LEVEL=debug
export CACHE_ENABLED=false

# 然后运行
polymarket-go-mcp
```

### 命令行参数

```bash
# 查看帮助
polymarket-go-mcp --help

# 查看版本
polymarket-go-mcp --version

# 健康检查
polymarket-go-mcp --health

# 指定日志级别
polymarket-go-mcp --log-level debug
```

### 自定义构建标签

如果您需要条件编译：

```bash
# 构建包含调试信息的版本
go build -tags=debug -o polymarket-go-mcp-debug .

# 构建性能优化版本
go build -tags=optimize -o polymarket-go-mcp-optimized .
```

## 性能优化

### 编译优化

```bash
# 启用优化和去除调试信息
go build -ldflags="-s -w" -o polymarket-go-mcp .

# 进一步优化（Go 1.20+）
go build -gcflags="-l=4" -o polymarket-go-mcp .
```

### 内存分析

```bash
# 构建包含内存分析支持的版本
go build -tags=memprofile -o polymarket-go-mcp-profile .
```

## 故障排除

### 常见问题

**问题1：命令未找到**
```bash
# 解决方案：确保Go bin目录在PATH中
export PATH=$PATH:$(go env GOPATH)/bin
```

**问题2：权限错误**
```bash
# 解决方案：添加执行权限
chmod +x polymarket-go-mcp
```

**问题3：依赖问题**
```bash
# 解决方案：清理并重新安装依赖
go clean -modcache
go mod download
```

**问题4：版本冲突**
```bash
# 解决方案：更新到最新版本
go install github.com/ivanzzeth/polymarket-go-mcp/cmd/polymarket-go-mcp@latest
```

### 调试技巧

```bash
# 启用详细日志
polymarket-go-mcp --log-level trace

# 检查二进制信息
file polymarket-go-mcp
ldd polymarket-go-mcp  # Linux
otool -L polymarket-go-mcp  # macOS

# 检查Go版本兼容性
go version
```

## 贡献指南

### 开发流程

1. **Fork仓库**
   ```bash
   git clone https://github.com/your-username/polymarket-go-mcp
   cd polymarket-go-mcp
   ```

2. **创建功能分支**
   ```bash
   git checkout -b feature/your-feature
   ```

3. **运行测试**
   ```bash
   go test ./...
   go vet ./...
   golangci-lint run
   ```

4. **提交更改**
   ```bash
   git add .
   git commit -m "feat: add your feature"
   git push origin feature/your-feature
   ```

### 代码规范

- 遵循Go标准代码格式：`gofmt -s -w .`
- 运行静态分析：`golangci-lint run`
- 编写单元测试覆盖核心功能
- 使用有意义的提交信息

## 总结

作为Go开发者，您有多种选择来使用polymarket-go-mcp服务器：

- **快速使用**：`go install` 直接安装
- **开发调试**：源码编译和`go run`
- **生产部署**：优化构建和交叉编译

选择最适合您需求的方式，享受Go语言带来的高性能和开发效率！