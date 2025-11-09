# GitHub Actions 自动发布指南

本文档说明如何使用 GitHub Actions 工作流自动构建和发布预构建的二进制文件。

## 概述

GitHub Actions 工作流会在以下情况下自动触发：
- 推送以 `v` 开头的标签时（例如 `v1.0.0`）
- 手动触发（通过 GitHub 界面）

工作流会为以下平台构建二进制文件：
- Linux AMD64
- Linux ARM64  
- macOS AMD64
- macOS ARM64
- Windows AMD64

## 使用方法

### 1. 创建新版本

要创建新版本，请按照以下步骤操作：

```bash
# 1. 确保所有更改已提交
git add .
git commit -m "准备发布 v1.0.0"

# 2. 创建标签
git tag v1.0.0

# 3. 推送标签到 GitHub
git push origin v1.0.0
```

推送标签后，GitHub Actions 会自动：
- 为所有支持的平台构建二进制文件
- 创建 GitHub Release
- 上传所有预构建的二进制文件

### 2. 手动触发

你也可以在 GitHub 仓库的 Actions 标签页中手动触发工作流：
1. 进入仓库的 "Actions" 标签页
2. 选择 "Release" 工作流
3. 点击 "Run workflow" 按钮

### 3. 下载预构建的二进制文件

发布后，用户可以通过以下方式下载二进制文件：

```bash
# Linux AMD64
curl -L https://github.com/ivanzzeth/polymarket-go-mcp/releases/latest/download/polymarket-go-mcp-linux-amd64 -o polymarket-go-mcp
chmod +x polymarket-go-mcp

# macOS ARM64
curl -L https://github.com/ivanzzeth/polymarket-go-mcp/releases/latest/download/polymarket-go-mcp-darwin-arm64 -o polymarket-go-mcp
chmod +x polymarket-go-mcp

# Windows AMD64
curl -L https://github.com/ivanzzeth/polymarket-go-mcp/releases/latest/download/polymarket-go-mcp-win-amd64.exe -o polymarket-go-mcp.exe
```

## 工作流配置

工作流配置文件位于 `.github/workflows/release.yml`，包含以下主要部分：

### 触发条件
```yaml
on:
  push:
    tags:
      - 'v*'
  workflow_dispatch:  # 允许手动触发
```

### 构建矩阵
```yaml
strategy:
  matrix:
    platform:
      - os: linux
        arch: amd64
        target: linux-amd64
      - os: linux
        arch: arm64
        target: linux-arm64
      # ... 其他平台
```

### 构建步骤
1. **检出代码** - 获取最新代码
2. **设置 Go 环境** - 配置 Go 1.24.x
3. **构建二进制文件** - 为每个平台交叉编译
4. **测试二进制文件** - 验证构建结果
5. **上传制品** - 保存构建结果
6. **创建发布** - 生成 GitHub Release 并上传所有二进制文件

## 本地测试

你可以使用提供的测试脚本来验证构建过程：

```bash
chmod +x scripts/test-build.sh
./scripts/test-build.sh
```

这个脚本会模拟 GitHub Actions 的构建过程，为所有平台构建二进制文件。

## 故障排除

### 构建失败
- 确保 `go.mod` 文件正确配置
- 检查所有依赖项可用
- 验证 Go 版本兼容性

### 发布失败
- 确保有足够的权限创建 Release
- 检查网络连接
- 验证标签格式正确

### 二进制文件无法运行
- 确保为 Unix 系统设置了可执行权限
- 验证目标平台兼容性
- 检查动态链接库依赖

## 最佳实践

1. **版本管理**
   - 使用语义化版本控制（SemVer）
   - 在发布前充分测试
   - 更新 CHANGELOG.md

2. **构建优化**
   - 保持构建过程快速
   - 最小化二进制文件大小
   - 使用适当的编译标志

3. **发布流程**
   - 在发布前创建预发布版本
   - 提供清晰的发布说明
   - 验证所有平台的二进制文件

## 相关文件

- `.github/workflows/release.yml` - GitHub Actions 工作流配置
- `scripts/test-build.sh` - 本地构建测试脚本
- `README.md` - 用户文档，包含下载说明
