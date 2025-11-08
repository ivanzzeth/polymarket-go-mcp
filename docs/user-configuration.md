# 用户配置指南

## 概述

本文档提供polymarket-mcp服务器在各种MCP客户端中的详细配置指南，确保用户能够快速上手使用。

## 快速开始

### 最简单的配置方式

**使用npx方式（推荐）：**
```json
{
  "polymarket-mcp": {
    "command": "npx",
    "args": ["-y", "polymarket-mcp-server"]
  }
}
```

这个配置适用于所有支持MCP协议的客户端，无需任何手动安装步骤。

## 客户端特定配置

### Claude Desktop

#### 配置文件位置
- **macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
- **Windows**: `%APPDATA%\Claude\claude_desktop_config.json`
- **Linux**: `~/.config/Claude/claude_desktop_config.json`

#### 完整配置示例
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

#### 带环境变量的配置
```json
{
  "mcpServers": {
    "polymarket": {
      "command": "npx",
      "args": ["-y", "polymarket-mcp-server"],
      "env": {
        "LOG_LEVEL": "info",
        "API_TIMEOUT": "30s"
      }
    }
  }
}
```

### Windsurf

#### 配置文件位置
- 项目根目录下的 `.windsurfrules` 文件

#### 配置示例
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

### Cursor

#### 配置文件位置
- 项目根目录下的 `.cursor/mcp.json` 文件

#### 配置示例
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

### 其他MCP客户端

对于其他支持MCP协议的客户端，配置方式类似：

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

## 高级配置选项

### 版本管理

#### 使用特定版本
```json
{
  "polymarket-mcp": {
    "command": "npx",
    "args": ["-y", "polymarket-mcp-server@1.2.3"]
  }
}
```

#### 使用最新版本
```json
{
  "polymarket-mcp": {
    "command": "npx",
    "args": ["-y", "polymarket-mcp-server@latest"]
  }
}
```

#### 使用预发布版本
```json
{
  "polymarket-mcp": {
    "command": "npx",
    "args": ["-y", "polymarket-mcp-server@next"]
  }
}
```

### 环境变量配置

#### 调试模式
```json
{
  "polymarket-mcp": {
    "command": "npx",
    "args": ["-y", "polymarket-mcp-server"],
    "env": {
      "DEBUG": "true",
      "LOG_LEVEL": "debug"
    }
  }
}
```

#### 生产环境配置
```json
{
  "polymarket-mcp": {
    "command": "npx",
    "args": ["-y", "polymarket-mcp-server"],
    "env": {
      "LOG_LEVEL": "warn",
      "API_TIMEOUT": "60s"
    }
  }
}
```

#### 自定义API端点
```json
{
  "polymarket-mcp": {
    "command": "npx",
    "args": ["-y", "polymarket-mcp-server"],
    "env": {
      "GAMMA_API_BASE": "https://custom-gamma-api.example.com",
      "DATA_API_BASE": "https://custom-data-api.example.com"
    }
  }
}
```

### 替代部署方式配置

#### Docker方式
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

#### 预编译二进制方式
```json
{
  "polymarket-mcp": {
    "command": "/path/to/polymarket-mcp-server"
  }
}
```

#### 源码编译方式
```json
{
  "polymarket-mcp": {
    "command": "go",
    "args": [
      "run",
      "github.com/ivanzzeth/polymarket-go-mcp"
    ]
  }
}
```

## 配置验证

### 测试配置

配置完成后，可以通过以下方式验证：

1. **重启MCP客户端** - 重新加载配置
2. **检查工具列表** - 确认polymarket工具可用
3. **测试简单工具** - 使用 `health_check` 工具测试连接

### 故障排除

#### 常见问题

**问题1：npx命令未找到**
```bash
# 解决方案：安装Node.js
# 访问 https://nodejs.org 下载安装
```

**问题2：权限错误**
```bash
# 解决方案：检查执行权限
chmod +x /path/to/polymarket-mcp-server
```

**问题3：网络连接问题**
```json
{
  "polymarket-mcp": {
    "command": "npx",
    "args": ["-y", "polymarket-mcp-server"],
    "env": {
      "HTTP_PROXY": "http://proxy.example.com:8080",
      "HTTPS_PROXY": "http://proxy.example.com:8080"
    }
  }
}
```

#### 调试模式

启用详细日志来诊断问题：

```json
{
  "polymarket-mcp": {
    "command": "npx",
    "args": ["-y", "polymarket-mcp-server"],
    "env": {
      "DEBUG": "true",
      "LOG_LEVEL": "trace"
    }
  }
}
```

## 多服务器配置

### 同时配置多个MCP服务器

```json
{
  "mcpServers": {
    "polymarket": {
      "command": "npx",
      "args": ["-y", "polymarket-mcp-server"]
    },
    "weather": {
      "command": "npx",
      "args": ["-y", "weather-mcp-server"]
    },
    "filesystem": {
      "command": "npx", 
      "args": ["-y", "filesystem-mcp-server"]
    }
  }
}
```

### 项目特定配置

对于特定项目，可以在项目目录中创建本地配置：

**项目根目录下的 `.claude/mcp_servers.json`:**
```json
{
  "polymarket": {
    "command": "npx",
    "args": ["-y", "polymarket-mcp-server"]
  }
}
```

## 性能优化配置

### 资源限制

#### 内存限制（Docker方式）
```json
{
  "polymarket-mcp": {
    "command": "docker",
    "args": [
      "run", "--rm", "-i",
      "--memory=256m",
      "--cpus=0.5",
      "ivanzzeth/polymarket-mcp-server:latest"
    ]
  }
}
```

#### 超时设置
```json
{
  "polymarket-mcp": {
    "command": "npx",
    "args": ["-y", "polymarket-mcp-server"],
    "env": {
      "REQUEST_TIMEOUT": "30s",
      "CACHE_TTL": "5m"
    }
  }
}
```

### 缓存配置

```json
{
  "polymarket-mcp": {
    "command": "npx",
    "args": ["-y", "polymarket-mcp-server"],
    "env": {
      "CACHE_ENABLED": "true",
      "CACHE_SIZE": "100MB",
      "CACHE_TTL": "10m"
    }
  }
}
```

## 安全配置

### API密钥配置

如果需要访问受保护的API端点：

```json
{
  "polymarket-mcp": {
    "command": "npx",
    "args": ["-y", "polymarket-mcp-server"],
    "env": {
      "API_KEY": "your-api-key-here",
      "API_SECRET": "your-api-secret-here"
    }
  }
}
```

### 网络限制

```json
{
  "polymarket-mcp": {
    "command": "npx",
    "args": ["-y", "polymarket-mcp-server"],
    "env": {
      "ALLOWED_DOMAINS": "gamma-api.polymarket.com,data-api.polymarket.com"
    }
  }
}
```

## 配置模板

### 开发环境模板
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

### 生产环境模板
```json
{
  "polymarket-mcp": {
    "command": "npx",
    "args": ["-y", "polymarket-mcp-server@latest"],
    "env": {
      "LOG_LEVEL": "info",
      "CACHE_ENABLED": "true",
      "CACHE_TTL": "10m"
    }
  }
}
```

### 企业环境模板
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

## 更新和维护

### 检查更新
```bash
npx polymarket-mcp-server@latest --version
```

### 清理缓存
```bash
# 清理npx缓存
npx clear-npx-cache

# 或者手动删除
rm -rf ~/.npm/_npx
```

通过遵循本指南，用户可以轻松配置polymarket-mcp服务器，享受无缝的预测市场数据访问体验。