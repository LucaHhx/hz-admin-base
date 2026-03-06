# Frontend 技术方案 -- 项目清理 (project-cleanup)

> 需求: 1-project-cleanup | 角色: frontend

## 概述

清理前端项目中的 example 示例页面、插件页面（announcement/email）、相关 API 封装和路由配置，并统一项目标识为 HAB。

## 技术栈

无新增技术选型。本需求为纯清理操作，保持现有 Vue 技术栈不变。

## 清理方案详细设计

### 1. 移除 example 示例页面

#### 需要删除的目录/文件

| 类型 | 路径 | 说明 |
|------|------|------|
| 页面目录 | `web/src/view/example/` | 包含 index.vue, customer/customer.vue, upload/upload.vue, upload/scanUpload.vue, breakpoint/breakpoint.vue |
| API 封装 | `web/src/api/customer.js` | 客户管理 API |
| API 封装 | `web/src/api/fileUploadAndDownload.js` | 文件上传下载 API |
| API 封装 | `web/src/api/breakpoint.js` | 断点续传 API |
| API 封装 | `web/src/api/attachmentCategory.js` | 媒体库分类 API |

#### 需要修改的文件

| 文件 | 修改内容 |
|------|----------|
| `web/src/router/index.js` | 移除 example 相关路由（如 scanUpload 静态路由，第 33 行附近） |

### 2. 移除插件页面

#### 需要删除的目录

| 插件 | 路径 | 说明 |
|------|------|------|
| announcement | `web/src/plugin/announcement/` | 包含 api/info.js, form/info.vue, view/info.vue |
| email | `web/src/plugin/email/` | 包含 api/email.js, view/index.vue |

#### 需要修改的文件

| 文件 | 修改内容 |
|------|----------|
| `web/src/api/email.js` | 删除此文件（根目录 api 下也有一份 email API 封装） |

**注意**: 插件路由通常由后端动态下发菜单控制，前端无需清理静态路由配置。但需确认 `web/src/plugin/` 目录下删除 announcement 和 email 后是否为空目录，如果为空可保留作为插件扩展入口。

### 3. 统一项目标识

| 文件 | 修改内容 |
|------|----------|
| `web/package.json` | `name` 字段: `"hz-admin-base"` -> `"hab"` |

### 4. 前端编译验证

清理完成后执行前端构建验证:
- `npm run build` 或 `yarn build` 确认无编译错误
- 确认无 import 引用已删除的模块

## 清理顺序

```
步骤 1: 删除 example 页面目录 + 相关 API 文件
步骤 2: 清理 router/index.js 中 example 相关路由
步骤 3: 删除 plugin/announcement 和 plugin/email 目录
步骤 4: 删除根目录 api/email.js
步骤 5: 更新 package.json name 字段
步骤 6: 前端编译验证
```

## 关键决策

| 决策 | 选择 | 理由 |
|------|------|------|
| plugin 目录保留 | 删除 announcement/email 后保留空 plugin 目录 | 作为插件扩展入口，与后端 plugin-tool 保持对称 |
| API 文件清理 | 删除 example 相关的 4 个 API 文件 | 对应后端 API 已移除，保留会产生死代码 |

## 依赖与约束

- 前端路由主要由后端动态菜单驱动，静态路由中仅 scanUpload 需要清理
- 前端清理需与后端清理同步（后端 API 移除后，前端对应调用也需移除）
- package.json 的 name 字段修改不影响功能，仅影响项目标识
