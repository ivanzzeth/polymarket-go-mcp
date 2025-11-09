# Polymarket Top 10 用户行为分析报告

## 分析目标
通过分析Polymarket平台上持仓价值最高的前10名用户，了解他们的交易行为模式，识别潜在的学习机会和交易策略。

## 数据来源
- **Gamma API**: 市场、事件、标签等元数据
- **Data API**: 用户活动、交易、持仓等实时数据
- **Subgraph**: 用户地址和链上数据

## 分析步骤

### 1. 识别Top 10用户
通过持仓价值排序获取前10名用户

**工具调用**: `get_positions_value`
```json
{
  "limit": 10,
  "sort_by": "CURRENT",
  "sort_direction": "DESC"
}
```

**预期结果**: 获取持仓价值最高的10个用户地址和持仓信息

### 2. 用户交易行为分析
