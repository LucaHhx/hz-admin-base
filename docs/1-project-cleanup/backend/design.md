# Backend 技术方案 -- 项目清理 (project-cleanup)

> 需求: project-cleanup | 角色: backend

## 概述

将 hz-admin-base 项目清理为干净的通用后台管理基础模板（HAB）。本方案覆盖 example 模块移除、插件移除、business 目录规范化、项目标识重命名、废弃 skill 清理等全部后端清理工作。

## 技术栈

无新增技术选型。本需求为纯清理操作，保持现有 Go (Gin + GORM) 技术栈不变。

## 项目标识重命名方案

### 新标识命名

| 项目 | 旧值 | 新值 |
|------|------|------|
| Go module name | `hz-admin-base` | `hab` |
| 前端 package name | `hz-admin-base` | `hab` |
| 数据库名 (config.example.yaml) | `hz-admin-base` | `hab` |
| 日志前缀 (config.example.yaml) | `[hz-admin-base]` | `[hab]` |
| autocode module (config.example.yaml) | `hz-admin-base` | `hab` |
| Dockerfile WORKDIR | `/srv/hz-admin-base` | `/srv/hab` |
| Dockerfile CMD | `./hz-admin-base` | `./hab` |
| .drone.yml 路径/构建产物 | `hz-admin-base` | `hab` |
| .gitignore 二进制忽略 | `/server/hz-admin-base` | `/server/hab` |

### import 路径替换

Go 模块名从 `hz-admin-base` 改为 `hab` 后，所有 Go 源文件中的 import 路径需要全局替换：
- `"hz-admin-base/xxx"` -> `"hab/xxx"`

**影响范围**: 255 个文件，638 处引用。使用 `sed` 全局替换完成。

## 清理方案详细设计

### 1. 移除 example 模块

example 模块包含客户管理、文件上传下载、断点续传、媒体库分类四个示例功能。

#### 需要删除的目录/文件

| 层级 | 路径 | 文件数 |
|------|------|--------|
| API 层 | `server/api/v1/example/` | 5 个文件 (enter.go + 4 个 handler) |
| Model 层 | `server/model/example/` | 8 个文件 (模型 + request + response) |
| Service 层 | `server/service/example/` | 5 个文件 (enter.go + 4 个 service) |
| Router 层 | `server/router/example/` | 4 个文件 (enter.go + 3 个 router) |
| Source 种子数据 | `server/source/example/` | 1 个文件 (file_upload_download.go) |

#### 需要修改的引用文件

| 文件 | 修改内容 |
|------|----------|
| `server/api/v1/enter.go` | 移除 `example` import 和 `ExampleApiGroup` 字段 |
| `server/router/enter.go` | 移除 `example` import 和 `Example` 字段 |
| `server/service/enter.go` | 移除 `example` import 和 `ExampleServiceGroup` 字段 |
| `server/initialize/router.go` | 移除 `exampleRouter` 变量及其三行路由注册 (CustomerRouter, FileUploadAndDownloadRouter, AttachmentCategoryRouter) |
| `server/initialize/gorm.go` | 移除 `example` import 和 RegisterTables() 中 5 个 example model (ExaFile, ExaCustomer, ExaFileChunk, ExaFileUploadAndDownload, ExaAttachmentCategory) |
| `server/initialize/ensure_tables.go` | 移除 `example` import 和 MigrateTable/TableCreated 中 4 个 example model |
| `server/source/system/menu.go` | 移除示例文件菜单项 (第 65-68 行: example 父菜单 + upload/breakpoint/customer 子菜单)，注意需同步调整后续菜单项的 ParentId 编号 |
| `server/source/system/api.go` | 移除分片上传、文件上传与下载、客户、断点续传(插件版)、媒体库分类的 API 种子数据 |
| `server/source/system/casbin.go` | 移除 fileUploadAndDownload、customer、simpleUploader、attachmentCategory 相关的 casbin 规则 |
| `server/source/system/authorities_menus.go` | 移除 example 菜单后需调整菜单索引切片逻辑 |

### 2. 移除插件模块

#### 需要删除的目录

| 插件 | 路径 |
|------|------|
| announcement | `server/plugin/announcement/` (含 api/config/gen/initialize/model/plugin/router/service 子目录) |
| email | `server/plugin/email/` (含 api/config/global/model/router/service/utils 子目录) |

#### 需要修改的引用文件

| 文件 | 修改内容 |
|------|----------|
| `server/initialize/plugin_biz_v1.go` | 移除 `email` import，清空 `bizPluginV1` 函数体（保留空函数签名，因为 `plugin.go` 调用它） |
| `server/initialize/plugin_biz_v2.go` | 移除 `announcement` import，清空 `bizPluginV2` 函数体 |
| `server/initialize/plugin.go` | `InstallPlugin` 函数保留（作为插件扩展入口），内部 `bizPluginV1`/`bizPluginV2` 调用保留但已无实际内容 |

**注意**: `server/plugin/plugin-tool/` 是通用插件工具，不属于业务插件，**保留不删除**。

### 3. 清理 business 目录残留

当前状态分析:
- `server/model/business/request/` -- 目录存在但已空（b_batch_pay.go 已在 git 中标记为 deleted）
- `server/api/v1/business/enter.go` -- 已是干净的空壳
- `server/router/business/enter.go` -- 已是干净的空壳
- `server/service/business/enter.go` -- 已是干净的空壳

#### 操作

1. 删除 `server/model/business/request/` 空目录
2. 创建规范的 model/business 空壳结构

#### 规范空壳结构设计

business 和 api 两套入口目录应保持统一的空壳模式:

**business 目录**（后台管理业务扩展）:
```
server/api/v1/business/enter.go      # type ApiGroup struct{}
server/router/business/enter.go      # type RouterGroup struct{}
server/service/business/enter.go     # type ServiceGroup struct{}
server/model/business/enter.go       # package business (空, 留给业务模型)
```

**api 目录**（客户端 API 扩展）:
```
server/api/v1/api/enter.go           # type ApiGroup struct{} (已存在)
server/router/api/enter.go           # type RouterGroup struct{} (已存在)
```

### 4. 清理废弃的 .claude/skills 文件

git status 显示以下文件已标记为 deleted:
- `.claude/skills/code-generator`
- `.claude/skills/mysql-operator`
- `.claude/skills/skill-creator`

这些文件已在 git 暂存区中标记为删除。确认清理状态即可，无需额外操作。

### 5. 更新配置文件

#### config.example.yaml

| 位置 | 旧值 | 新值 |
|------|------|------|
| mysql.db-name | `hz-admin-base` | `hab` |
| autocode.module | `hz-admin-base` | `hab` |
| zap.prefix | `'[hz-admin-base]'` | `'[hab]'` |

#### .gitignore

| 行号 | 旧值 | 新值 |
|------|------|------|
| 32 | `/server/hz-admin-base` | `/server/hab` |
| 33 | `/server/hz-admin-base/config.yaml` | `/server/hab/config.yaml` |

#### Dockerfile

| 行号 | 旧值 | 新值 |
|------|------|------|
| 14 | `WORKDIR /srv/hz-admin-base` | `WORKDIR /srv/hab` |
| 15 | `COPY . /srv/hz-admin-base` | `COPY . /srv/hab` |
| 19 | `CMD ["./hz-admin-base", "-type=all"]` | `CMD ["./hab", "-type=all"]` |

#### .drone.yml

全局替换 `hz-admin-base` -> `hab`（8 处引用，涉及构建路径、产物名称、部署路径）。

### 6. 前端 package.json

仅修改 `name` 字段: `"hz-admin-base"` -> `"hab"`。

## 清理顺序建议

为避免中间状态编译失败，建议按以下顺序执行:

```
步骤 1: Go module 重命名 (go.mod + 全局 import 替换)
         -> 确保编译通过
步骤 2: 移除 example 模块 (删除目录 + 清理引用)
         -> 确保编译通过
步骤 3: 移除插件模块 (删除目录 + 清理引用)
         -> 确保编译通过
步骤 4: 清理 business 目录残留 + 规范空壳
         -> 确保编译通过
步骤 5: 清理种子数据 (source/system 中的 menu/api/casbin)
         -> 确保编译通过
步骤 6: 更新配置文件 (.gitignore, config.example.yaml, Dockerfile, .drone.yml)
步骤 7: 更新前端 package.json
步骤 8: 清理废弃 skill 文件 (确认 git 状态)
步骤 9: 更新 README
步骤 10: 最终编译验证 (go build + 启动测试)
```

**理由**: 步骤 1 (重命名) 必须先做，因为后续删除文件操作中的 import 路径已经是新的。如果先删文件再改 import，中间状态会因为旧 import 路径找不到模块而失败。每步之后编译验证，确保增量可控。

## 关键决策

| 决策 | 选择 | 理由 |
|------|------|------|
| 新模块名 | `hab` | 简短、与项目名 HAB 一致，import 路径简洁 |
| plugin.go 保留 | 保留 InstallPlugin 空壳 | 作为插件扩展入口，router.go 中虽注释了调用但保留架构 |
| plugin-tool 保留 | 不删除 | 通用插件工具，非业务插件，其他新插件可能依赖 |
| 文件上传下载移除 | 随 example 一起移除 | 属于示例代码，如需此功能可作为新模块重新实现 |
| 种子数据清理 | 移除 example 相关条目 | 保持种子数据与代码一致，避免初始化时引用不存在的路由/模型 |

## 依赖与约束

- 本需求不涉及数据库迁移或 Schema 变更
- 种子数据的修改仅影响数据库初始化（首次 init 时），不影响已有运行环境
- menu.go 中移除 example 菜单后，菜单 ID 会发生变化（因为 ID 由插入顺序决定），authorities_menus.go 中的硬编码索引需要同步调整
- 文件上传下载功能在 example 模块中实现，移除后系统将不再提供文件管理功能
