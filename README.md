# HAB (Hz Admin Base)

通用后台管理基础项目，提供开箱即用的用户管理、权限控制、代码生成等企业级后台管理能力，可作为业务系统的基础底座快速扩展。

## 技术栈

| 层级 | 技术 |
|------|------|
| **后端** | Go 1.24+、Gin、GORM、Casbin、JWT、Zap、Viper |
| **前端** | Vue 3、Element Plus、Pinia、Vue Router、Vite |
| **数据库** | SQLite（默认）、MySQL、PostgreSQL、MSSQL、Oracle |
| **可选组件** | Redis、MongoDB、云存储（阿里云 OSS / MinIO / AWS S3 / 腾讯 COS） |

## 核心功能

- **用户与权限** — RBAC 角色管理、菜单权限、按钮级权限、列级权限、数据过滤
- **代码生成器** — 可视化配置自动生成 CRUD 模块，支持历史记录和回滚
- **多运行模式** — 后台管理 / API 服务 / 全部，灵活适配部署场景
- **安全体系** — JWT 鉴权、IP 限流、API Key 认证、三种登录模式（simple / captcha / strict）
- **国际化** — 中文 / 英文双语支持
- **文件管理** — 多存储后端、断点续传、附件分类
- **系统运维** — 操作记录、数据字典、参数配置、定时任务

## 快速开始

### 环境要求

- Go 1.24+
- Node.js 18+

### 启动后端

```bash
cd server
cp config.example.yaml config.yaml
go run main.go
```

默认端口：管理后台 API `9688`，客户端 API `9689`。

默认使用 SQLite，无需安装数据库，开箱即用。

### 启动前端

```bash
cd web
npm install
npm run serve
```

默认管理员账号：`admin` / `123456`

## 项目结构

```
├── server/              # Go 后端
│   ├── api/v1/          # API 接口层
│   │   ├── system/      #   系统管理 API
│   │   ├── business/    #   业务扩展 API（空壳）
│   │   └── api/         #   客户端 API（空壳）
│   ├── model/           # 数据模型
│   ├── service/         # 业务逻辑层
│   ├── router/          # 路由定义
│   ├── middleware/      # 中间件（JWT / CORS / Casbin / 限流）
│   ├── initialize/      # 初始化（数据库 / 路由 / 定时任务）
│   ├── config/          # 配置结构体
│   ├── cmd/             # CLI 工具（代码生成 / 密码哈希）
│   └── utils/           # 工具函数
│
├── web/                 # Vue 3 前端
│   └── src/
│       ├── api/         # API 请求模块
│       ├── view/        # 页面视图
│       ├── components/  # 通用组件
│       ├── pinia/       # 状态管理
│       ├── router/      # 路由配置
│       ├── i18n/        # 国际化
│       └── utils/       # 工具函数
│
├── docs/                # 项目文档
└── .claude/             # AI 开发辅助（hz-agents）
```

## 业务扩展

在以下空壳目录中添加业务代码即可扩展：

```
server/api/v1/business/    → 业务 API handler
server/router/business/    → 业务路由注册
server/service/business/   → 业务逻辑
server/model/business/     → 业务数据模型
```

也可使用内置的**代码生成器**通过可视化配置自动生成 CRUD 模块。

## 数据库切换

默认 SQLite，切换 MySQL 只需修改 `config.yaml`：

```yaml
system:
  db-type: mysql
  use-redis: true   # 按需启用

mysql:
  addr: 127.0.0.1:3306
  db-name: hab
  username: root
  password: your-password
```

支持的数据库：SQLite、MySQL、PostgreSQL、MSSQL、Oracle。

## AI 辅助开发（hz-agents）

本项目集成了 [hz-agents](https://github.com/LucaHhx/hz-agents) 多智能体开发框架，通过 `.claude/` 目录下的符号链接接入。

提供 6 个专业 AI 角色（PM / Tech Lead / UI / Frontend / Backend / QA）和 14 个斜杠命令，覆盖从需求规划到验收测试的完整开发流程：

```bash
# 初始化新项目
/hz-init

# 需求文档评审
/unify-doc-review 用户系统

# 启动全团队开发
/unify-dev 用户系统

# 单独调用角色
/review-pm 用户系统       # PM 评审需求
/review-tech 用户系统     # Tech Lead 技术设计
/dev-frontend 用户系统    # 前端开发
/dev-backend 用户系统     # 后端开发
/review-qa 用户系统       # QA 验收测试
```

## 配置说明

完整配置参考 [`server/config.example.yaml`](server/config.example.yaml)，最小化配置参考 [`server/config.minimal.yaml`](server/config.minimal.yaml)。

关键配置项：

| 配置项 | 默认值 | 说明 |
|--------|--------|------|
| `system.db-type` | sqlite | 数据库类型 |
| `system.use-redis` | false | 是否启用 Redis |
| `system.addr` | 9688 | 管理后台 API 端口 |
| `system.client-addr` | 9689 | 客户端 API 端口 |
| `jwt.expires-time` | 7d | Token 过期时间 |
| `system.ip-limit-count` | 15000 | 每小时 IP 限流次数 |

## 许可证

MIT License
