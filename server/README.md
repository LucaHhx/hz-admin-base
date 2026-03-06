# HAB Server

HAB (Hz Admin Base) 后台管理系统服务端，基于 Gin + GORM 的通用后台管理基础模板。

## 快速开始

```bash
# 1. 复制配置文件
cp config.example.yaml config.yaml

# 2. 修改配置（默认 SQLite，开箱即用）

# 3. 启动服务
go run main.go
```

默认端口：
- 管理后台 API: `9688`
- 客户端 API: `9689`

## 项目结构

```
├── api/v1/          # API 层（handler）
│   ├── system/      # 系统管理 API
│   ├── business/    # 业务扩展 API（空壳）
│   └── api/         # 客户端 API（空壳）
├── config/          # 配置结构体（对应 config.yaml）
├── core/            # 核心组件初始化（zap, viper, server）
├── docs/            # Swagger 文档
├── global/          # 全局变量
├── initialize/      # 初始化（router, redis, gorm, validator）
├── middleware/      # Gin 中间件
├── model/           # 数据模型
│   ├── system/      # 系统模型
│   └── business/    # 业务模型（空壳）
├── plugin/          # 插件目录
│   └── plugin-tool/ # 通用插件工具
├── router/          # 路由层
├── service/         # 业务逻辑层
├── resource/        # 静态资源（模板、Excel）
└── utils/           # 工具函数
```

## 业务扩展

在以下空壳目录中添加业务代码：

- `api/v1/business/` - 业务 API handler
- `router/business/` - 业务路由注册
- `service/business/` - 业务逻辑
- `model/business/` - 业务数据模型

## 配置说明

默认使用 SQLite 数据库，无需额外安装数据库服务。如需切换为 MySQL：

```yaml
system:
  db-type: mysql    # 改为 mysql
  use-redis: true   # 按需启用 Redis
```
