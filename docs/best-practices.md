# 最佳实践总结

## 概述

本文档总结polymarket-mcp服务器开发、部署和使用的最佳实践，为开发者提供完整的指导方案。

## 开发最佳实践

### Go MCP开发规范

#### 项目结构
```
polymarket-go-mcp/
├── cmd/
│   └── polymarket-mcp-server/
│       └── main.go          # 主程序入口
├── internal/
│   ├── mcp/
│   │   ├── server.go        # MCP服务器实现
│   │   └── tools.go         # 工具定义
│   ├── gamma/
│   │   ├── client.go        # Gamma API客户端
│   │   └── types.go         # Gamma数据类型
│   ├── data/
│   │   ├── client.go        # Data API客户端
│   │   └── types.go         # Data数据类型
│   └── config/
│       └── config.go        # 配置管理
├── pkg/
│   └── utils/
│       ├── cache.go         # 缓存工具
│       └── logger.go        # 日志工具
├── docs/                    # 文档目录
├── scripts/                 # 构建脚本
├── go.mod
├── go.sum
└── README.md
```

#### 代码组织原则

1. **清晰的包结构**
   - `internal/` 包含私有实现
   - `pkg/` 包含可复用的公共包
   - `cmd/` 包含可执行程序

2. **依赖注入**
```go
type Server struct {
    gammaClient *gamma.Client
    dataClient  *data.Client
    cache       *cache.Cache
    logger      *log.Logger
}

func NewServer(cfg *config.Config) *Server {
    return &Server{
        gammaClient: gamma.NewClient(cfg.Gamma),
        dataClient:  data.NewClient(cfg.Data),
        cache:       cache.New(cfg.Cache),
        logger:      log.New(cfg.LogLevel),
    }
}
```

3. **错误处理**
```go
func (s *Server) GetMarketData(ctx context.Context, marketID string) (*MarketData, error) {
    data, err := s.gammaClient.GetMarket(ctx, marketID)
    if err != nil {
        return nil, fmt.Errorf("failed to get market data: %w", err)
    }
    
    if data == nil {
        return nil, fmt.Errorf("market not found: %s", marketID)
    }
    
    return data, nil
}
```

### MCP协议最佳实践

#### 工具定义规范
```go
type ToolDefinition struct {
    Name        string                 `json:"name"`
    Description string                 `json:"description"`
    InputSchema map[string]interface{} `json:"inputSchema"`
}

// 示例工具定义
var GetMarketDataTool = ToolDefinition{
    Name:        "get_market_data",
    Description: "Get detailed market data including prices, volume, and liquidity",
    InputSchema: map[string]interface{}{
        "type": "object",
        "properties": map[string]interface{}{
            "market_id": map[string]interface{}{
                "type":        "string",
                "description": "The market ID to fetch data for",
            },
        },
        "required": []string{"market_id"},
    },
}
```

#### 资源定义规范
```go
type ResourceDefinition struct {
    URI         string `json:"uri"`
    Name        string `json:"name"`
    Description string `json:"description"`
    MimeType    string `json:"mimeType"`
}

// 示例资源定义
var MarketResource = ResourceDefinition{
    URI:         "polymarket://market/{market_id}",
    Name:        "Market Data",
    Description: "Access detailed market information",
    MimeType:    "application/json",
}
```

## 部署最佳实践

### 多平台构建策略

#### 构建矩阵配置
```yaml
# .github/workflows/build.yml
jobs:
  build:
    strategy:
      matrix:
        goos: [linux, darwin, windows]
        goarch: [amd64, arm64]
        include:
          - goos: windows
            ext: .exe
          - goos: linux
            ext: ""
          - goos: darwin  
            ext: ""
```

#### 版本管理
```bash
# 语义化版本管理
git tag v1.0.0
git push origin v1.0.0

# 预发布版本
git tag v1.1.0-rc.1
git push origin v1.1.0-rc.1
```

### 发布流程最佳实践

#### npm包发布
```bash
#!/bin/bash
# scripts/publish-npm.sh

set -e

# 验证构建
npm run prepublishOnly

# 更新版本
npm version patch

# 构建所有平台
./scripts/build-binaries.sh

# 发布到npm
npm publish

# 验证发布
npx polymarket-mcp-server@latest --version
```

#### Docker镜像发布
```bash
#!/bin/bash
# scripts/publish-docker.sh

set -e

VERSION=${1:-latest}

# 构建多架构镜像
docker buildx create --use
docker buildx build \
  --platform linux/amd64,linux/arm64 \
  --tag ivanzzeth/polymarket-mcp-server:$VERSION \
  --tag ivanzzeth/polymarket-mcp-server:latest \
  --push .

# 验证镜像
docker run --rm ivanzzeth/polymarket-mcp-server:latest --version
```

## 配置最佳实践

### 环境特定配置

#### 开发环境
```json
{
  "polymarket-mcp": {
    "command": "npx",
    "args": ["-y", "polymarket-mcp-server"],
    "env": {
      "DEBUG": "true",
      "LOG_LEVEL": "debug",
      "CACHE_ENABLED": "false"
    }
  }
}
```

#### 生产环境
```json
{
  "polymarket-mcp": {
    "command": "npx", 
    "args": ["-y", "polymarket-mcp-server@latest"],
    "env": {
      "LOG_LEVEL": "info",
      "CACHE_ENABLED": "true",
      "CACHE_TTL": "10m",
      "REQUEST_TIMEOUT": "30s"
    }
  }
}
```

#### 企业环境
```json
{
  "polymarket-mcp": {
    "command": "docker",
    "args": [
      "run", "--rm", "-i",
      "--memory=512m",
      "--cpus=1.0",
      "--env=LOG_LEVEL=warn",
      "ivanzzeth/polymarket-mcp-server:latest"
    ]
  }
}
```

### 安全配置

#### API密钥管理
```json
{
  "polymarket-mcp": {
    "command": "npx",
    "args": ["-y", "polymarket-mcp-server"],
    "env": {
      "API_KEY": "${POLYMARKET_API_KEY}",
      "API_SECRET": "${POLYMARKET_API_SECRET}"
    }
  }
}
```

#### 网络限制
```json
{
  "polymarket-mcp": {
    "command": "npx",
    "args": ["-y", "polymarket-mcp-server"],
    "env": {
      "ALLOWED_DOMAINS": "gamma-api.polymarket.com,data-api.polymarket.com",
      "HTTP_PROXY": "http://corporate-proxy:8080"
    }
  }
}
```

## 性能优化最佳实践

### 缓存策略

#### 内存缓存
```go
type CacheConfig struct {
    Enabled    bool          `json:"enabled"`
    Size       string        `json:"size"`       // e.g., "100MB"
    TTL        time.Duration `json:"ttl"`        // e.g., "10m"
    Cleanup    time.Duration `json:"cleanup"`    // e.g., "1m"
}

func NewCache(cfg *CacheConfig) *Cache {
    if !cfg.Enabled {
        return &Cache{enabled: false}
    }
    
    return &Cache{
        enabled: true,
        data:    make(map[string]*cacheItem),
        maxSize: parseSize(cfg.Size),
        ttl:     cfg.TTL,
    }
}
```

#### 请求合并
```go
type RequestBatcher struct {
    requests chan *BatchRequest
    timeout  time.Duration
    maxSize  int
}

func (b *RequestBatcher) Do(ctx context.Context, key string, fn func() (interface{}, error)) (interface{}, error) {
    // 合并相同请求，避免重复API调用
}
```

### 资源管理

#### 连接池
```go
type ClientPool struct {
    clients chan *http.Client
    factory func() *http.Client
}

func (p *ClientPool) Get() *http.Client {
    select {
    case client := <-p.clients:
        return client
    default:
        return p.factory()
    }
}

func (p *ClientPool) Put(client *http.Client) {
    select {
    case p.clients <- client:
    default:
        // 池已满，丢弃连接
    }
}
```

## 监控和日志最佳实践

### 结构化日志
```go
type Logger struct {
    level  LogLevel
    writer io.Writer
}

func (l *Logger) Info(msg string, fields ...Field) {
    if l.level >= InfoLevel {
        l.log(InfoLevel, msg, fields...)
    }
}

func (l *Logger) Error(msg string, err error, fields ...Field) {
    if l.level >= ErrorLevel {
        fields = append(fields, Field{Key: "error", Value: err.Error()})
        l.log(ErrorLevel, msg, fields...)
    }
}
```

### 健康检查
```go
type HealthChecker struct {
    components []HealthCheckable
}

func (h *HealthChecker) Check() HealthStatus {
    status := HealthStatus{Healthy: true}
    
    for _, component := range h.components {
        componentStatus := component.HealthCheck()
        status.Components = append(status.Components, componentStatus)
        
        if !componentStatus.Healthy {
            status.Healthy = false
        }
    }
    
    return status
}
```

## 测试最佳实践

### 单元测试
```go
func TestGetMarketData(t *testing.T) {
    // 模拟依赖
    mockGamma := &MockGammaClient{}
    mockData := &MockDataClient{}
    
    server := NewServer(Config{
        Gamma: GammaConfig{BaseURL: "http://test-gamma"},
        Data:  DataConfig{BaseURL: "http://test-data"},
    })
    server.gammaClient = mockGamma
    server.dataClient = mockData
    
    // 设置期望
    mockGamma.On("GetMarket", "test-market").Return(&Market{ID: "test-market"}, nil)
    
    // 执行测试
    result, err := server.GetMarketData(context.Background(), "test-market")
    
    // 验证结果
    assert.NoError(t, err)
    assert.Equal(t, "test-market", result.ID)
    mockGamma.AssertExpectations(t)
}
```

### 集成测试
```go
func TestServerIntegration(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping integration test in short mode")
    }
    
    // 启动测试服务器
    server := startTestServer(t)
    defer server.Stop()
    
    // 执行集成测试
    client := NewTestClient(server.URL)
    result, err := client.GetMarketData("test-market")
    
    assert.NoError(t, err)
    assert.NotNil(t, result)
}
```

## 文档最佳实践

### README结构
```markdown
# Polymarket MCP Server

## 快速开始
```json
{
  "polymarket-mcp": {
    "command": "npx",
    "args": ["-y", "polymarket-mcp-server"]
  }
}
```

## 功能特性
- 市场数据查询
- 价格信息获取
- 流动性分析

## 配置指南
[详细配置说明](docs/user-configuration.md)

## 开发指南  
[开发文档](docs/development.md)
```

### API文档
```go
// GetMarketData retrieves detailed market information including
// current prices, volume, and liquidity data.
//
// Parameters:
//   - marketID: The unique identifier of the market
//
// Returns:
//   - *MarketData: Detailed market information
//   - error: Any error that occurred during the request
//
// Example:
//   data, err := client.GetMarketData("0x123...")
//   if err != nil {
//       return err
//   }
func (c *Client) GetMarketData(marketID string) (*MarketData, error) {
    // implementation
}
```

## 持续集成最佳实践

### GitHub Actions配置
```yaml
name: CI

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v4
        with:
          go-version: '1.24'
      
      - name: Run tests
        run: go test -v ./...
        
      - name: Run lint
        run: golangci-lint run
        
  build:
    runs-on: ubuntu-latest
    needs: test
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v4
        with:
          go-version: '1.24'
          
      - name: Build
        run: go build -o polymarket-mcp-server
```

## 总结

通过遵循这些最佳实践，polymarket-mcp服务器能够：

1. **提供优秀的用户体验** - 通过npx一键安装
2. **确保代码质量** - 通过清晰的架构和测试
3. **支持多种部署方式** - 适应不同环境需求
4. **提供完整文档** - 便于用户和开发者使用
5. **实现持续交付** - 通过自动化流程保证质量

这些实践为构建高质量、易用的MCP服务器提供了完整的指导框架。