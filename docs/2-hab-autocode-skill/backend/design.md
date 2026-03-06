# Backend 技术方案 — hab-autocode-skill

## 技术栈

- Go + Gin (现有框架)
- 中间件机制: Gin middleware
- 配置管理: Viper (mapstructure yaml 绑定)
- Shell 脚本: bash (辅助 AI 调用)

---

## 一、架构设计

### 核心目标

在不修改现有 AutoCode 业务逻辑的前提下，新增 API Key 认证通道，使 AI Agent 可以通过 `x-api-key` 请求头直接访问 AutoCode 相关 API，绕过 JWT 登录流程。同时提供完整的 Skill 文档和辅助脚本供 AI 使用。

### 整体认证流程

```
请求 → CORS → [ApiKeyOrJWT 中间件] → [Casbin 中间件] → Handler
                  │
                  ├── 携带 x-api-key 且匹配 → 注入管理员 claims(AuthorityId=1) → Casbin 放行 → Handler
                  ├── 携带 x-api-key 但不匹配 → 401 拒绝
                  └── 未携带 x-api-key → fallback 到 JWTAuth() 标准流程
```

### 方案选择: 新建独立路由组 + ApiKeyOrJWT 组合中间件

不修改现有 PrivateGroup 的中间件链，而是为 AutoCode 路由创建一个新的路由组 `AutoCodeGroup`，该组使用 "API Key 或 JWT" 二选一的认证中间件。

理由:
1. 不影响现有路由的认证逻辑，其他模块仍走 JWTAuth + CasbinHandler
2. API Key 认证仅作用于 AutoCode 路由，不扩散到其他模块
3. Casbin 中间件对 `AuthorityId == 1` 直接放行 (`server/middleware/casbin_rbac.go:21`)，无需额外配置权限策略

---

## 二、API Key 认证中间件详细设计

### 2.1 配置结构体扩展

**文件:** `server/config/auto_code.go`

```go
type Autocode struct {
    Web    string `mapstructure:"web" json:"web" yaml:"web"`
    Root   string `mapstructure:"root" json:"root" yaml:"root"`
    Server string `mapstructure:"server" json:"server" yaml:"server"`
    Module string `mapstructure:"module" json:"module" yaml:"module"`
    AiPath string `mapstructure:"ai-path" json:"ai-path" yaml:"ai-path"`
    ApiKey string `mapstructure:"api-key" json:"api-key" yaml:"api-key"`
}
```

变更: 仅新增 `ApiKey` 字段，不修改已有字段和 `WebRoot()` 方法。

### 2.2 配置文件扩展

**文件:** `server/config.example.yaml`

```yaml
autocode:
  web: web/src
  server: server
  module: hab
  api-key: ""  # 留空表示禁用 API Key 认证；填入随机字符串启用
```

### 2.3 ApiKeyOrJWT 中间件实现

**新建文件:** `server/middleware/api_key.go`

```go
package middleware

import (
    "time"

    "hab/global"
    "hab/model/common/response"
    systemReq "hab/model/system/request"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "github.com/google/uuid"
)

// ApiKeyOrJWT 组合认证中间件
// 优先检查 x-api-key 请求头，命中则注入管理员 claims 并跳过 JWT 认证
// 未命中则 fallback 到标准 JWTAuth 流程
func ApiKeyOrJWT() gin.HandlerFunc {
    return func(c *gin.Context) {
        apiKey := c.GetHeader("x-api-key")
        configKey := global.HAB_CONFIG.AutoCode.ApiKey

        // 有有效 API Key 时，注入 claims 并跳过 JWT
        if configKey != "" && apiKey != "" {
            if apiKey != configKey {
                response.NoAuth("invalid api key", c)
                c.Abort()
                return
            }

            // 命中 API Key，注入超级管理员 claims
            claims := &systemReq.CustomClaims{
                BaseClaims: systemReq.BaseClaims{
                    UUID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
                    ID:          1,
                    Username:    "api-key",
                    NickName:    "API Key User",
                    AuthorityId: 1,
                },
                BufferTime: 0,
                RegisteredClaims: jwt.RegisteredClaims{
                    ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
                },
            }
            c.Set("claims", claims)
            c.Next()
            return
        }

        // 无 API Key 或未配置，走标准 JWT 流程
        JWTAuth()(c)
    }
}
```

**关键设计决策:**
- `AuthorityId: 1` — Casbin 中间件对此值直接放行，无需配置额外的权限策略
- `UUID` 使用固定值 `00000000-0000-0000-0000-000000000001`，避免与真实用户 UUID 冲突
- `ID: 1` 对应系统管理员用户 ID，确保 `utils.GetUserID(c)` 等辅助函数正常工作
- `Username: "api-key"` — 操作日志中可区分人工操作和 API Key 操作

### 2.4 路由初始化修改

**文件:** `server/initialize/router.go`

修改位置: 第 70-96 行区域

**修改前:**
```go
PrivateGroup := Router.Group(global.HAB_CONFIG.System.RouterPrefix)
PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
// ...
systemRouter.InitAutoCodeRouter(PrivateGroup, PublicGroup)      // 创建自动化代码
// ...
systemRouter.InitAutoCodeHistoryRouter(PrivateGroup)            // 自动化代码历史
```

**修改后:**
```go
PrivateGroup := Router.Group(global.HAB_CONFIG.System.RouterPrefix)
PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())

// AutoCode 专用路由组，支持 API Key 或 JWT 双重认证
AutoCodeGroup := Router.Group(global.HAB_CONFIG.System.RouterPrefix)
AutoCodeGroup.Use(middleware.ApiKeyOrJWT()).Use(middleware.CasbinHandler())
// ...
systemRouter.InitAutoCodeRouter(AutoCodeGroup, PublicGroup)      // 创建自动化代码 (改用 AutoCodeGroup)
// ...
systemRouter.InitAutoCodeHistoryRouter(AutoCodeGroup)            // 自动化代码历史 (改用 AutoCodeGroup)
```

**注意:** 仅将 `InitAutoCodeRouter` 和 `InitAutoCodeHistoryRouter` 的参数从 `PrivateGroup` 改为 `AutoCodeGroup`，其他路由保持不变。

---

## 三、Skill 目录结构设计

```
.claude/skills/hab-autocode/
├── SKILL.md              # 主文档 — AI 读取的完整操作指南
├── scripts/
│   ├── autocode.sh        # 封装 curl 调用的辅助脚本
│   └── config.sh          # 从 config.yaml 读取 API Key 和 server 地址
└── examples/
    ├── create-package.json     # 创建包的请求体示例
    ├── create-module.json      # 创建模块的请求体示例 (完整字段)
    └── create-module-simple.json # 创建模块的最简请求体示例
```

---

## 四、SKILL.md 内容设计

### 4.1 文档结构

```markdown
# HAB AutoCode Skill

> AI 自动化代码生成操作指南

## 前置条件
## 配置
## API 端点清单
## 操作流程
## 辅助脚本使用
## 错误处理与常见问题
## 安全注意事项
```

### 4.2 前置条件章节

```markdown
## 前置条件

1. HAB server 已启动并运行
2. `config.yaml` 中已配置 `autocode.api-key`
3. 数据库已初始化 (已执行 initdb)

### 配置 API Key

在 `server/config.yaml` 中:

\```yaml
autocode:
  web: web/src
  server: server
  module: hab
  api-key: "your-random-api-key-here"  # 生成: uuidgen 或 openssl rand -hex 32
\```

重启 server 使配置生效。
```

### 4.3 API 端点详细参数设计

#### Package 管理

**POST /autoCode/createPackage** — 创建新包

请求体 (`SysAutoCodePackageCreate`):
```json
{
  "packageName": "order",      // 必填，包名 (小写英文，不含 / \ ..)
  "label": "订单管理",           // 展示名
  "desc": "订单相关业务",         // 描述
  "template": "package"         // 必填，模板类型: "package" 或 "plugin"
}
```

成功响应:
```json
{
  "code": 0,
  "data": {},
  "msg": "Success"
}
```

**POST /autoCode/getPackage** — 获取所有包

请求体: `{}`

成功响应:
```json
{
  "code": 0,
  "data": {
    "pkgs": [
      {
        "ID": 1,
        "packageName": "system",
        "template": "package",
        "label": "system包",
        "desc": "系统自动读取system包",
        "module": "hab"
      }
    ]
  },
  "msg": "Success"
}
```

**POST /autoCode/delPackage** — 删除包

请求体:
```json
{
  "id": 1
}
```

**GET /autoCode/getTemplates** — 获取可用模板类型

无请求参数。响应: `{"code": 0, "data": ["package", "plugin"], "msg": "Success"}`

#### 代码生成

**POST /autoCode/preview** — 预览代码 (不写入文件)

请求体 (`AutoCode` 结构体，核心字段):
```json
{
  "package": "order",
  "tableName": "orders",
  "structName": "Order",
  "packageName": "order",
  "abbreviation": "order",
  "humpPackageName": "order",
  "description": "订单管理",
  "gvaModel": true,
  "autoMigrate": true,
  "autoCreateApiToSql": true,
  "autoCreateMenuToSql": true,
  "autoCreateBtnAuth": true,
  "generateServer": true,
  "generateWeb": true,
  "fields": [
    {
      "fieldName": "OrderNo",
      "fieldDesc": "订单号",
      "fieldType": "string",
      "fieldJson": "orderNo",
      "dataType": "varchar",
      "dataTypeLong": "64",
      "comment": "订单编号",
      "columnName": "order_no",
      "fieldSearchType": "=",
      "form": true,
      "table": true,
      "desc": true,
      "require": true
    },
    {
      "fieldName": "Amount",
      "fieldDesc": "金额",
      "fieldType": "float64",
      "fieldJson": "amount",
      "dataType": "decimal",
      "dataTypeLong": "10,2",
      "comment": "订单金额",
      "columnName": "amount",
      "form": true,
      "table": true,
      "desc": true
    }
  ]
}
```

成功响应:
```json
{
  "code": 0,
  "data": {
    "autoCode": {
      "server/api/v1/order/order.go": "```go\n\n...\n\n```",
      "server/model/order/order.go": "```go\n\n...\n\n```",
      "server/router/order/order.go": "```go\n\n...\n\n```",
      "server/service/order/order.go": "```go\n\n...\n\n```",
      "web/src/api/order/order.js": "```js\n\n...\n\n```",
      "web/src/view/order/order/order.vue": "```vue\n\n...\n\n```"
    }
  },
  "msg": "预览成功"
}
```

**POST /autoCode/createTemp** — 实际生成代码

请求体: 与 preview 完全相同的 `AutoCode` 结构体。

成功响应:
```json
{
  "code": 0,
  "data": {},
  "msg": "Success"
}
```

**POST /autoCode/addFunc** — 为已有模块添加方法

请求体 (`AutoFunc`):
```json
{
  "package": "order",
  "structName": "Order",
  "packageName": "order",
  "humpPackageName": "order",
  "abbreviation": "order",
  "funcName": "ExportOrder",
  "router": "exportOrder",
  "method": "GET",
  "description": "导出订单",
  "funcDesc": "导出订单数据",
  "isAuth": true,
  "isPreview": false,
  "isAi": false
}
```

预览模式: 设置 `"isPreview": true`，返回 `{"api": "...", "server": "...", "js": "..."}`

#### 数据库查询

**GET /autoCode/getDB** — 获取数据库列表

查询参数: `?businessDB=` (可选)

响应:
```json
{
  "code": 0,
  "data": {
    "dbs": ["hab"],
    "dbList": [
      {"aliasName": "...", "dbName": "...", "disable": false, "dbtype": "mysql"}
    ]
  },
  "msg": "Success"
}
```

**GET /autoCode/getTables** — 获取表列表

查询参数: `?dbName=hab&businessDB=` (dbName 可选，默认当前数据库)

响应:
```json
{
  "code": 0,
  "data": {
    "tables": [
      {"tableName": "sys_users"}
    ]
  },
  "msg": "Success"
}
```

**GET /autoCode/getColumn** — 获取表列信息

查询参数: `?tableName=sys_users&dbName=hab&businessDB=`

响应:
```json
{
  "code": 0,
  "data": {
    "columns": [
      {
        "columnName": "id",
        "dataType": "bigint",
        "dataTypeLong": "20",
        "columnComment": "主键ID",
        "fieldName": "Id",
        "fieldType": "int64",
        "fieldJson": "id"
      }
    ]
  },
  "msg": "Success"
}
```

#### 历史和回滚

**POST /autoCode/getSysHistory** — 分页获取生成历史

请求体:
```json
{
  "page": 1,
  "pageSize": 10
}
```

响应:
```json
{
  "code": 0,
  "data": {
    "list": [
      {
        "ID": 1,
        "table": "orders",
        "package": "order",
        "structName": "Order",
        "description": "订单管理",
        "flag": 0
      }
    ],
    "total": 1,
    "page": 1,
    "pageSize": 10
  },
  "msg": "Success"
}
```

`flag` 字段: `0` = 未回滚, `1` = 已回滚

**POST /autoCode/getMeta** — 获取详细元信息

请求体:
```json
{
  "id": 1
}
```

响应: `{"code": 0, "data": {"meta": "<JSON字符串，原始AutoCode请求体>"}, "msg": "Success"}`

**POST /autoCode/rollback** — 回滚已生成代码

请求体 (`SysAutoHistoryRollBack`):
```json
{
  "id": 1,
  "deleteApi": true,
  "deleteMenu": true,
  "deleteTable": true
}
```

回滚操作包括:
- 移除生成的文件 (移动到 `rm_file/` 目录)
- 撤回 AST 注入的代码 (enter.go, router_biz.go 等)
- 删除自动创建的 API 记录 (当 `deleteApi: true`)
- 删除自动创建的菜单 (当 `deleteMenu: true`)
- 删除数据库表 (当 `deleteTable: true`)
- 将历史记录标记为 `flag = 1`

**POST /autoCode/delSysHistory** — 删除历史记录

请求体:
```json
{
  "id": 1
}
```

---

## 五、辅助脚本设计

### 5.1 config.sh — 配置读取脚本

**文件:** `.claude/skills/hab-autocode/scripts/config.sh`

```bash
#!/bin/bash
# 从 config.yaml 读取 API Key 和 server 配置
# 用法: source scripts/config.sh

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/../../../../" && pwd)"
CONFIG_FILE="$PROJECT_ROOT/server/config.yaml"

if [ ! -f "$CONFIG_FILE" ]; then
    echo "Error: config.yaml not found at $CONFIG_FILE" >&2
    exit 1
fi

# 读取 API Key (使用 grep + sed，不依赖 yq)
HAB_API_KEY=$(grep -A 10 '^autocode:' "$CONFIG_FILE" | grep 'api-key:' | sed 's/.*api-key:\s*"\?\([^"]*\)"\?.*/\1/' | tr -d ' ')

# 读取 server 端口
HAB_PORT=$(grep -A 5 '^system:' "$CONFIG_FILE" | grep 'addr:' | head -1 | sed 's/.*addr:\s*//' | tr -d ' ')

# 读取路由前缀
HAB_PREFIX=$(grep -A 10 '^system:' "$CONFIG_FILE" | grep 'router-prefix:' | sed 's/.*router-prefix:\s*"\?\([^"]*\)"\?.*/\1/' | tr -d ' ')

# 构建基础 URL
HAB_BASE_URL="http://localhost:${HAB_PORT:-9688}${HAB_PREFIX}"

export HAB_API_KEY HAB_PORT HAB_PREFIX HAB_BASE_URL
```

### 5.2 autocode.sh — 封装 curl 调用

**文件:** `.claude/skills/hab-autocode/scripts/autocode.sh`

```bash
#!/bin/bash
# HAB AutoCode API 封装脚本
# 用法: ./autocode.sh <command> [args...]
#
# 命令:
#   packages                    获取所有包
#   templates                   获取模板列表
#   create-package <json>       创建包
#   delete-package <id>         删除包
#   preview <json_file>         预览代码
#   create <json_file>          创建代码
#   add-func <json>             添加方法
#   get-db                      获取数据库列表
#   get-tables [dbName]         获取表列表
#   get-columns <tableName>     获取表字段
#   history [page] [pageSize]   获取生成历史
#   meta <id>                   获取 Meta 信息
#   rollback <id> [flags]       回滚代码
#   delete-history <id>         删除历史记录

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "$SCRIPT_DIR/config.sh"

if [ -z "$HAB_API_KEY" ]; then
    echo "Error: API Key not configured in config.yaml" >&2
    exit 1
fi

# 通用 curl 函数
api_post() {
    local path="$1"
    local data="${2:-{}}"
    curl -s -X POST "${HAB_BASE_URL}${path}" \
        -H "Content-Type: application/json" \
        -H "x-api-key: ${HAB_API_KEY}" \
        -d "$data"
}

api_get() {
    local path="$1"
    curl -s -X GET "${HAB_BASE_URL}${path}" \
        -H "x-api-key: ${HAB_API_KEY}"
}

api_post_file() {
    local path="$1"
    local file="$2"
    curl -s -X POST "${HAB_BASE_URL}${path}" \
        -H "Content-Type: application/json" \
        -H "x-api-key: ${HAB_API_KEY}" \
        -d @"$file"
}

CMD="${1:-help}"
shift || true

case "$CMD" in
    packages)
        api_post "/autoCode/getPackage"
        ;;
    templates)
        api_get "/autoCode/getTemplates"
        ;;
    create-package)
        api_post "/autoCode/createPackage" "$1"
        ;;
    delete-package)
        api_post "/autoCode/delPackage" "{\"id\": $1}"
        ;;
    preview)
        api_post_file "/autoCode/preview" "$1"
        ;;
    create)
        api_post_file "/autoCode/createTemp" "$1"
        ;;
    add-func)
        api_post "/autoCode/addFunc" "$1"
        ;;
    get-db)
        api_get "/autoCode/getDB"
        ;;
    get-tables)
        local_db="${1:-}"
        api_get "/autoCode/getTables?dbName=${local_db}"
        ;;
    get-columns)
        api_get "/autoCode/getColumn?tableName=$1"
        ;;
    history)
        page="${1:-1}"
        pageSize="${2:-10}"
        api_post "/autoCode/getSysHistory" "{\"page\": $page, \"pageSize\": $pageSize}"
        ;;
    meta)
        api_post "/autoCode/getMeta" "{\"id\": $1}"
        ;;
    rollback)
        id="$1"
        flags="${2:-}"
        delete_api="true"
        delete_menu="true"
        delete_table="true"
        if [ "$flags" = "keep-table" ]; then
            delete_table="false"
        fi
        api_post "/autoCode/rollback" "{\"id\": $id, \"deleteApi\": $delete_api, \"deleteMenu\": $delete_menu, \"deleteTable\": $delete_table}"
        ;;
    delete-history)
        api_post "/autoCode/delSysHistory" "{\"id\": $1}"
        ;;
    help|*)
        echo "Usage: $0 <command> [args...]"
        echo ""
        echo "Commands:"
        echo "  packages                    Get all packages"
        echo "  templates                   Get template types"
        echo "  create-package <json>       Create a package"
        echo "  delete-package <id>         Delete a package"
        echo "  preview <json_file>         Preview generated code"
        echo "  create <json_file>          Generate code"
        echo "  add-func <json>             Add method to module"
        echo "  get-db                      List databases"
        echo "  get-tables [dbName]         List tables"
        echo "  get-columns <tableName>     List columns"
        echo "  history [page] [pageSize]   List generation history"
        echo "  meta <id>                   Get meta info"
        echo "  rollback <id> [flags]       Rollback generated code"
        echo "  delete-history <id>         Delete history record"
        ;;
esac
```

---

## 六、示例请求 JSON 设计

### 6.1 create-package.json

```json
{
  "packageName": "order",
  "label": "订单管理",
  "desc": "订单相关业务模块",
  "template": "package"
}
```

### 6.2 create-module.json (完整字段)

```json
{
  "package": "order",
  "tableName": "orders",
  "structName": "Order",
  "packageName": "order",
  "abbreviation": "order",
  "humpPackageName": "order",
  "description": "订单管理",
  "businessDB": "",
  "gvaModel": true,
  "autoMigrate": true,
  "autoCreateApiToSql": true,
  "autoCreateMenuToSql": true,
  "autoCreateBtnAuth": true,
  "onlyTemplate": false,
  "generateServer": true,
  "generateWeb": true,
  "fields": [
    {
      "fieldName": "OrderNo",
      "fieldDesc": "订单号",
      "fieldType": "string",
      "fieldJson": "orderNo",
      "dataType": "varchar",
      "dataTypeLong": "64",
      "comment": "订单编号",
      "columnName": "order_no",
      "fieldSearchType": "=",
      "form": true,
      "table": true,
      "desc": true,
      "require": true
    },
    {
      "fieldName": "Amount",
      "fieldDesc": "金额",
      "fieldType": "float64",
      "fieldJson": "amount",
      "dataType": "decimal",
      "dataTypeLong": "10,2",
      "comment": "订单金额",
      "columnName": "amount",
      "form": true,
      "table": true,
      "desc": true
    },
    {
      "fieldName": "Status",
      "fieldDesc": "状态",
      "fieldType": "int",
      "fieldJson": "status",
      "dataType": "int",
      "dataTypeLong": "4",
      "comment": "订单状态 1:待支付 2:已支付 3:已完成",
      "columnName": "status",
      "fieldSearchType": "=",
      "form": true,
      "table": true,
      "desc": true,
      "defaultValue": "1"
    }
  ]
}
```

### 6.3 create-module-simple.json (最简示例)

```json
{
  "package": "order",
  "tableName": "orders",
  "structName": "Order",
  "packageName": "order",
  "abbreviation": "order",
  "humpPackageName": "order",
  "description": "订单",
  "gvaModel": true,
  "autoMigrate": true,
  "autoCreateApiToSql": true,
  "autoCreateMenuToSql": true,
  "autoCreateBtnAuth": true,
  "generateServer": true,
  "generateWeb": true,
  "fields": [
    {
      "fieldName": "Name",
      "fieldDesc": "名称",
      "fieldType": "string",
      "fieldJson": "name",
      "dataType": "varchar",
      "dataTypeLong": "255",
      "comment": "名称",
      "columnName": "name",
      "form": true,
      "table": true,
      "desc": true
    }
  ]
}
```

---

## 七、SKILL.md 完整操作流程设计

### 7.1 创建包 → 创建模块 → 预览 → 确认 完整流程

```
步骤 1: 获取已有包列表
  POST /autoCode/getPackage → 确认目标包是否已存在

步骤 2: 如果包不存在，创建包
  POST /autoCode/createPackage
  → 自动创建 api/v1/<pkg>/enter.go, router/<pkg>/enter.go, service/<pkg>/enter.go 等目录结构

步骤 3: 查询数据库表结构 (可选，用于基于已有表生成)
  GET /autoCode/getDB → 获取数据库
  GET /autoCode/getTables?dbName=<db> → 获取表
  GET /autoCode/getColumn?tableName=<table>&dbName=<db> → 获取字段

步骤 4: 预览代码
  POST /autoCode/preview → 检查生成的代码是否符合预期

步骤 5: 确认生成
  POST /autoCode/createTemp → 实际生成代码文件、注册 API、创建菜单、迁移数据库

步骤 6: 验证
  检查生成的文件是否存在，确认功能可用
```

### 7.2 回滚流程

```
步骤 1: 查询历史
  POST /autoCode/getSysHistory → 找到要回滚的记录 ID

步骤 2: 查看详情 (可选)
  POST /autoCode/getMeta → 确认是要回滚的目标

步骤 3: 执行回滚
  POST /autoCode/rollback
  → deleteApi: true   删除注册的 API 记录
  → deleteMenu: true  删除创建的菜单
  → deleteTable: true 删除创建的数据库表
  → 文件会被移动到 rm_file/ 目录 (非永久删除)
  → AST 注入的代码会被撤回

步骤 4: (可选) 删除历史记录
  POST /autoCode/delSysHistory
```

### 7.3 错误处理

| 错误场景 | 错误信息 | 解决方案 |
|----------|----------|----------|
| 包名重复 | "存在相同PackageName" | 先 getPackage 查询，不重复创建 |
| 包名是 Go 关键字 | "<name>为go的关键字!" | 改用其他名称 |
| 结构体重复 | "已经创建过此数据结构,请勿重复创建!" | 先查历史，必要时先 rollback 再重新创建 |
| 包结构异常 | "package结构异常,缺少..." | 检查对应 enter.go 文件是否存在 |
| 包名含非法字符 | "包名不合法" | 不要在包名中使用 / \ .. |
| API Key 无效 | 401 "invalid api key" | 检查 config.yaml 中的 api-key 配置 |
| 未认证 | 401 "未登录或非法访问" | 确认请求头中携带了 x-api-key |

---

## 八、AutoCode 字段说明

### AutoCode 结构体字段参考

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| package | string | 是 | 包名，需先通过 createPackage 创建 |
| tableName | string | 是 | 数据库表名 |
| structName | string | 是 | Go 结构体名称 (首字母大写) |
| packageName | string | 是 | 文件名 (小写) |
| abbreviation | string | 是 | 简称 (小写，用于路由前缀) |
| humpPackageName | string | 是 | 驼峰文件名 (通常等于 packageName) |
| description | string | 是 | 中文描述 |
| businessDB | string | 否 | 业务数据库名 (多库时使用) |
| gvaModel | bool | 否 | 是否使用默认 Model (ID, CreatedAt 等) |
| autoMigrate | bool | 否 | 是否自动迁移建表 |
| autoCreateApiToSql | bool | 否 | 是否自动注册 API |
| autoCreateMenuToSql | bool | 否 | 是否自动创建菜单 |
| autoCreateBtnAuth | bool | 否 | 是否自动创建按钮权限 |
| onlyTemplate | bool | 否 | 是否只生成模板文件 (不注入 gorm/router) |
| generateServer | bool | 否 | 是否生成后端代码 |
| generateWeb | bool | 否 | 是否生成前端代码 |

### AutoCodeField 字段参考

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| fieldName | string | 是 | Go 字段名 (首字母大写) |
| fieldDesc | string | 是 | 中文描述 |
| fieldType | string | 是 | Go 类型: string, int, float64, bool, time.Time, json, array, picture, file, richtext, video, enum |
| fieldJson | string | 是 | JSON 序列化名 (小驼峰) |
| dataType | string | 是 | 数据库类型: varchar, int, bigint, decimal, text, datetime, tinyint |
| dataTypeLong | string | 是 | 数据库类型长度: "255", "10,2", "20" 等 |
| comment | string | 否 | 数据库字段注释 |
| columnName | string | 是 | 数据库列名 (下划线) |
| fieldSearchType | string | 否 | 搜索条件: "=", "!=", ">", "<", ">=", "<=", "LIKE", "BETWEEN" |
| form | bool | 否 | 前端新建/编辑表单显示 |
| table | bool | 否 | 前端表格列显示 |
| desc | bool | 否 | 前端详情显示 |
| require | bool | 否 | 是否必填 |
| defaultValue | string | 否 | 默认值 |
| sort | bool | 否 | 是否支持排序 |
| primaryKey | bool | 否 | 是否主键 (gvaModel=false 时需要) |
| fieldIndexType | string | 否 | 索引类型 |

### AutoFunc 字段参考

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| package | string | 是 | 包名 |
| structName | string | 是 | 结构体名 |
| packageName | string | 是 | 文件名 |
| humpPackageName | string | 是 | 驼峰文件名 |
| abbreviation | string | 是 | 简称 |
| funcName | string | 是 | 方法名 (首字母大写) |
| router | string | 是 | 路由路径 (小驼峰) |
| method | string | 是 | HTTP 方法: GET, POST, PUT, DELETE |
| description | string | 否 | 描述 |
| funcDesc | string | 否 | 方法介绍 |
| isAuth | bool | 否 | 是否需要鉴权路由 |
| isPreview | bool | 否 | true 时只预览不注入 |
| isAi | bool | 否 | AI 模式 (可自定义代码) |
| apiFunc | string | 否 | 自定义 API handler 代码 (isAi=true 时) |
| serverFunc | string | 否 | 自定义 service 代码 (isAi=true 时) |
| jsFunc | string | 否 | 自定义前端 JS 代码 (isAi=true 时) |

---

## 九、安全性设计

### API Key 存储
- 存储在 `config.yaml` 中，与 JWT signing-key 同级别保护
- 不存入数据库，避免被前端管理界面暴露
- 建议使用 `uuidgen` 或 `openssl rand -hex 32` 生成
- `config.yaml` 已在 `.gitignore` 中，不会提交到版本控制

### API Key 传输
- 通过 `x-api-key` 请求头传输，不出现在 URL 参数中
- 生产环境必须使用 HTTPS

### 访问范围限制
- API Key 仅对 AutoCode 相关路由生效 (通过独立路由组 AutoCodeGroup 实现物理隔离)
- 不能用于访问用户管理、权限管理等其他敏感接口

### 审计与追溯
- API Key 请求注入的 `Username: "api-key"`，操作日志中可区分人工操作和 AI 操作
- AutoCode 所有创建操作都有历史记录 (`sys_auto_code_histories` 表)，支持回滚

### 风险与缓解

| 风险 | 缓解措施 |
|------|----------|
| API Key 泄露 | 仅在 config.yaml 中配置，不提交到版本控制；可随时更换 |
| API Key 被滥用 | 仅限 AutoCode 路由；操作有历史记录可追溯和回滚 |
| config.yaml 权限不当 | 文档中提醒设置文件权限 600 |
| 代码注入风险 | AutoCode 使用模板生成，包名已做路径穿越校验 (api handler 中检查 / \ ..) |

---

## 十、影响范围汇总

### 新建文件

| 文件 | 说明 |
|------|------|
| `server/middleware/api_key.go` | ApiKeyOrJWT 组合中间件 |
| `.claude/skills/hab-autocode/SKILL.md` | AI 操作指南主文档 |
| `.claude/skills/hab-autocode/scripts/config.sh` | 配置读取脚本 |
| `.claude/skills/hab-autocode/scripts/autocode.sh` | API 调用封装脚本 |
| `.claude/skills/hab-autocode/examples/create-package.json` | 示例请求体 |
| `.claude/skills/hab-autocode/examples/create-module.json` | 示例请求体 |
| `.claude/skills/hab-autocode/examples/create-module-simple.json` | 示例请求体 |

### 修改文件

| 文件 | 修改内容 |
|------|----------|
| `server/config/auto_code.go` | Autocode 结构体新增 `ApiKey` 字段 |
| `server/config.example.yaml` | autocode 段新增 `api-key` 配置项 |
| `server/initialize/router.go` | 新增 AutoCodeGroup 路由组，迁移 AutoCode 路由注册 |

### 不修改的文件

- 所有 AutoCode API handler (auto_code.go, auto_code_package.go, auto_code_template.go, auto_code_history.go)
- 所有 AutoCode service
- 所有 AutoCode router 定义
- Casbin 权限策略
- 前端代码

---

## 十一、依赖与约束

- 不依赖新的外部包，全部使用现有依赖 (gin, jwt, uuid)
- Shell 脚本仅依赖 bash, curl, grep, sed (标准 Unix 工具)
- 兼容现有 JWT 认证流程，两种认证方式互不干扰
- 当 `api-key` 配置为空时，系统行为与修改前完全一致
