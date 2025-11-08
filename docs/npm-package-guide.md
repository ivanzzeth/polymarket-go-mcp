# npm包发布指南

## 概述

本文档详细说明如何将polymarket-mcp服务器打包为npm包并发布到npm registry，实现真正的npx风格体验。

## npm包结构

### 目录结构
```
polymarket-mcp-server-npm/
├── package.json              # npm包配置
├── bin/
│   ├── index.js              # 平台检测脚本
│   ├── polymarket-mcp-server-linux-amd64
│   ├── polymarket-mcp-server-linux-arm64
│   ├── polymarket-mcp-server-darwin-amd64
│   ├── polymarket-mcp-server-darwin-arm64
│   ├── polymarket-mcp-server-win-amd64.exe
│   └── polymarket-mcp-server-win-arm64.exe
├── README.md                 # npm包说明文档
└── .npmignore               # npm发布忽略文件
```

### package.json配置

```json
{
  "name": "polymarket-mcp-server",
  "version": "1.0.0",
  "description": "Polymarket MCP Server - Access prediction market data through MCP protocol",
  "main": "bin/index.js",
  "bin": {
    "polymarket-mcp-server": "./bin/index.js"
  },
  "files": [
    "bin/"
  ],
  "scripts": {
    "prepublishOnly": "node scripts/verify-binaries.js",
    "test": "echo \"No tests specified\" && exit 0"
  },
  "keywords": [
    "mcp",
    "polymarket",
    "prediction-markets",
    "gamma-api",
    "data-api",
    "model-context-protocol"
  ],
  "author": "ivanzzeth <your-email@example.com>",
  "license": "MIT",
  "repository": {
    "type": "git",
    "url": "https://github.com/ivanzzeth/polymarket-go-mcp"
  },
  "homepage": "https://github.com/ivanzzeth/polymarket-go-mcp",
  "bugs": {
    "url": "https://github.com/ivanzzeth/polymarket-go-mcp/issues"
  },
  "engines": {
    "node": ">=14"
  },
  "os": [
    "darwin",
    "linux",
    "win32"
  ],
  "cpu": [
    "x64",
    "arm64"
  ]
}
```

### 平台检测脚本 (bin/index.js)

```javascript
#!/usr/bin/env node
const { platform, arch } = process;
const path = require('path');
const { spawn } = require('child_process');
const { existsSync } = require('fs');

// 平台映射
const PLATFORM_MAP = {
  'win32': 'win',
  'darwin': 'darwin',
  'linux': 'linux'
};

// 架构映射
const ARCH_MAP = {
  'x64': 'amd64',
  'arm64': 'arm64'
};

function getBinaryName() {
  const platformName = PLATFORM_MAP[platform];
  const archName = ARCH_MAP[arch];
  
  if (!platformName || !archName) {
    throw new Error(`Unsupported platform: ${platform} ${arch}`);
  }
  
  let binaryName = `polymarket-mcp-server-${platformName}-${archName}`;
  
  if (platform === 'win32') {
    binaryName += '.exe';
  }
  
  return binaryName;
}

function getBinaryPath() {
  const binaryName = getBinaryName();
  return path.join(__dirname, binaryName);
}

function verifyBinary() {
  const binaryPath = getBinaryPath();
  
  if (!existsSync(binaryPath)) {
    console.error(`❌ Error: Binary not found for platform ${platform} ${arch}`);
    console.error(`Expected path: ${binaryPath}`);
    console.error('Available binaries:');
    
    const fs = require('fs');
    const files = fs.readdirSync(__dirname);
    files.forEach(file => {
      if (file.startsWith('polymarket-mcp-server-')) {
        console.error(`  - ${file}`);
      }
    });
    
    process.exit(1);
  }
  
  return binaryPath;
}

function setBinaryPermissions(binaryPath) {
  if (platform !== 'win32') {
    const { chmodSync } = require('fs');
    try {
      chmodSync(binaryPath, '755');
      console.log(`✅ Set executable permissions for ${binaryPath}`);
    } catch (error) {
      console.warn(`⚠️  Could not set permissions: ${error.message}`);
    }
  }
}

function main() {
  try {
    const binaryPath = verifyBinary();
    setBinaryPermissions(binaryPath);
    
    console.log(`🚀 Starting polymarket-mcp-server (${platform}/${arch})...`);
    
    const child = spawn(binaryPath, process.argv.slice(2), {
      stdio: 'inherit',
      env: process.env
    });

    child.on('close', (code) => {
      process.exit(code);
    });

    child.on('error', (error) => {
      console.error('❌ Failed to start polymarket-mcp-server:', error);
      process.exit(1);
    });
    
    // 处理退出信号
    process.on('SIGINT', () => {
      child.kill('SIGINT');
    });
    
    process.on('SIGTERM', () => {
      child.kill('SIGTERM');
    });
    
  } catch (error) {
    console.error('❌ Fatal error:', error.message);
    process.exit(1);
  }
}

// 如果直接运行此脚本
if (require.main === module) {
  main();
}

module.exports = { getBinaryName, getBinaryPath };
```

## 构建流程

### 多平台构建脚本

```bash
#!/bin/bash
# scripts/build-binaries.sh

set -e

echo "🚀 Building polymarket-mcp-server for multiple platforms..."

# 创建临时构建目录
BUILD_DIR="tmp-build"
mkdir -p $BUILD_DIR

# 平台和架构配置
PLATFORMS=("linux" "darwin" "windows")
ARCHS=("amd64" "arm64")

# 清理之前的构建
rm -rf bin/*
mkdir -p bin

# 构建每个平台和架构的组合
for platform in "${PLATFORMS[@]}"; do
  for arch in "${ARCHS[@]}"; do
    echo "📦 Building for $platform/$arch..."
    
    output_name="polymarket-mcp-server-${platform}-${arch}"
    if [ "$platform" = "windows" ]; then
      output_name="${output_name}.exe"
    fi
    
    # 构建Go二进制
    GOOS=$platform GOARCH=$arch go build -o "$BUILD_DIR/$output_name"
    
    # 复制到bin目录
    cp "$BUILD_DIR/$output_name" "bin/"
    
    echo "✅ Built: bin/$output_name"
  done
done

# 复制平台检测脚本
cp scripts/platform-wrapper.js bin/index.js
chmod +x bin/index.js

# 清理临时文件
rm -rf $BUILD_DIR

echo "🎉 All binaries built successfully!"
echo "📁 Output directory: bin/"
```

### 二进制验证脚本

```javascript
// scripts/verify-binaries.js
const { existsSync } = require('fs');
const { getBinaryName, getBinaryPath } = require('../bin/index.js');

console.log('🔍 Verifying npm package binaries...');

const platforms = ['win32', 'darwin', 'linux'];
const architectures = ['x64', 'arm64'];

let allBinariesExist = true;

for (const platform of platforms) {
  for (const arch of architectures) {
    // 临时设置process.platform和process.arch来测试
    const originalPlatform = process.platform;
    const originalArch = process.arch;
    
    Object.defineProperty(process, 'platform', { value: platform });
    Object.defineProperty(process, 'arch', { value: arch });
    
    try {
      const binaryName = getBinaryName();
      const binaryPath = require('path').join(__dirname, '..', 'bin', binaryName);
      
      if (existsSync(binaryPath)) {
        console.log(`✅ ${platform}/${arch}: ${binaryName}`);
      } else {
        console.log(`❌ ${platform}/${arch}: Missing ${binaryName}`);
        allBinariesExist = false;
      }
    } catch (error) {
      console.log(`❌ ${platform}/${arch}: ${error.message}`);
      allBinariesExist = false;
    }
    
    // 恢复原始值
    Object.defineProperty(process, 'platform', { value: originalPlatform });
    Object.defineProperty(process, 'arch', { value: originalArch });
  }
}

if (!allBinariesExist) {
  console.error('❌ Some binaries are missing. Package verification failed.');
  process.exit(1);
}

console.log('✅ All binaries verified successfully!');
```

## 发布流程

### 1. 准备发布

```bash
# 切换到npm包目录
cd polymarket-mcp-server-npm

# 安装依赖（如果需要）
npm install

# 运行测试和验证
npm run prepublishOnly

# 更新版本号
npm version patch  # 或 minor, major

# 构建所有二进制文件
./scripts/build-binaries.sh
```

### 2. 发布到npm

```bash
# 登录npm（第一次发布需要）
npm login

# 发布包
npm publish

# 或者发布为beta版本
npm publish --tag beta
```

### 3. 验证发布

```bash
# 在临时目录测试安装
mkdir test-install && cd test-install
npx polymarket-mcp-server --version

# 测试MCP配置
echo '{
  "polymarket-mcp": {
    "command": "npx",
    "args": ["-y", "polymarket-mcp-server"]
  }
}' > test-config.json
```

## 自动发布流程

### GitHub Actions工作流

```yaml
name: Publish to npm

on:
  release:
    types: [published]

jobs:
  build-and-publish:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Setup Node.js
        uses: actions/setup-node@v4
        with:
          node-version: '18'
          registry-url: 'https://registry.npmjs.org'
          
      - name: Setup Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.24'
          
      - name: Build binaries
        run: |
          cd polymarket-mcp-server-npm
          ./scripts/build-binaries.sh
          
      - name: Verify package
        run: |
          cd polymarket-mcp-server-npm
          npm run prepublishOnly
          
      - name: Publish to npm
        run: |
          cd polymarket-mcp-server-npm
          npm publish
        env:
          NODE_AUTH_TOKEN: ${{ secrets.NPM_TOKEN }}
```

## 版本管理

### 语义化版本
- **主版本号**：不兼容的API修改
- **次版本号**：向后兼容的功能性新增
- **修订号**：向后兼容的问题修正

### 发布标签
```bash
# 发布稳定版本
npm publish

# 发布预发布版本
npm publish --tag next

# 发布测试版本
npm publish --tag beta
```

## 用户使用指南

### 基本使用
```json
{
  "polymarket-mcp": {
    "command": "npx",
    "args": ["-y", "polymarket-mcp-server"]
  }
}
```

### 指定版本
```json
{
  "polymarket-mcp": {
    "command": "npx",
    "args": ["-y", "polymarket-mcp-server@1.0.0"]
  }
}
```

### 使用最新版本
```json
{
  "polymarket-mcp": {
    "command": "npx",
    "args": ["-y", "polymarket-mcp-server@latest"]
  }
}
```

## 故障排除

### 常见问题

1. **二进制文件缺失**
   - 确保所有平台的二进制文件都包含在bin目录
   - 运行验证脚本检查

2. **权限问题**
   - 确保二进制文件有执行权限
   - 在非Windows平台设置chmod +x

3. **平台不支持**
   - 检查package.json中的os和cpu字段
   - 确保支持用户的操作系统和架构

通过遵循此指南，您可以成功将polymarket-mcp服务器发布为npm包，为用户提供真正的npx风格体验。