# 任务清单

> 计划: project-cleanup/backend | 创建: 2026-03-06

| # | 任务 | 状态 | 开始日期 | 完成日期 | 备注 |
|---|------|------|----------|----------|------|
| 1 | Go module 重命名: go.mod 中 `hz-admin-base` -> `hab` | 待办 |  |  | L2#5 |
| 2 | 全局 import 替换: server/ 下所有 .go 文件中 `"hz-admin-base/` -> `"hab/` (255 文件 638 处) | 待办 |  |  | L2#5 |
| 3 | 编译验证: 重命名后 `go build` 确认通过 | 待办 |  |  | L2#5 |
| 4 | 删除 example 目录: api/v1/example, model/example, service/example, router/example, source/example | 待办 |  |  | L2#1 |
| 5 | 清理 example 引用: api/v1/enter.go 移除 ExampleApiGroup 及 import | 待办 |  |  | L2#1 |
| 6 | 清理 example 引用: router/enter.go 移除 Example 字段及 import | 待办 |  |  | L2#1 |
| 7 | 清理 example 引用: service/enter.go 移除 ExampleServiceGroup 及 import | 待办 |  |  | L2#1 |
| 8 | 清理 example 引用: initialize/router.go 移除 exampleRouter 变量及 3 行路由注册 | 待办 |  |  | L2#1 |
| 9 | 清理 example 引用: initialize/gorm.go 移除 example import 及 5 个 model 的 AutoMigrate | 待办 |  |  | L2#1 |
| 10 | 清理 example 引用: initialize/ensure_tables.go 移除 example import 及 4 个 model | 待办 |  |  | L2#1 |
| 11 | 编译验证: example 模块移除后 `go build` 确认通过 | 待办 |  |  | L2#1 |
| 12 | 删除插件目录: plugin/announcement 和 plugin/email | 待办 |  |  | L2#2 |
| 13 | 清理插件引用: plugin_biz_v1.go 移除 email import，清空 bizPluginV1 函数体 | 待办 |  |  | L2#2 |
| 14 | 清理插件引用: plugin_biz_v2.go 移除 announcement import，清空 bizPluginV2 函数体 | 待办 |  |  | L2#2 |
| 15 | 编译验证: 插件移除后 `go build` 确认通过 | 待办 |  |  | L2#2 |
| 16 | 清理 business 残留: 删除 model/business/request 空目录，创建 model/business/enter.go 空壳 | 待办 |  |  | L2#4 |
| 17 | 编译验证: business 目录清理后 `go build` 确认通过 | 待办 |  |  | L2#4 |
| 18 | 清理种子数据: source/system/menu.go 移除 example 相关菜单项 (4 行) | 待办 |  |  | L2#1 |
| 19 | 清理种子数据: source/system/api.go 移除 example 相关 API 种子数据 (分片上传/文件上传下载/客户/断点续传插件版/媒体库分类) | 待办 |  |  | L2#1 |
| 20 | 清理种子数据: source/system/casbin.go 移除 example 相关 casbin 规则 (fileUploadAndDownload/customer/simpleUploader/attachmentCategory) | 待办 |  |  | L2#1 |
| 21 | 清理种子数据: source/system/authorities_menus.go 调整菜单索引切片（因 example 菜单移除导致索引变化） | 待办 |  |  | L2#1 |
| 22 | 编译验证: 种子数据清理后 `go build` 确认通过 | 待办 |  |  | L2#1 |
| 23 | 更新 config.example.yaml: db-name/module/prefix 中 `hz-admin-base` -> `hab` (3 处) | 待办 |  |  | L2#5 |
| 24 | 更新 .gitignore: `/server/hz-admin-base` -> `/server/hab` (2 处) | 待办 |  |  | L2#5 |
| 25 | 更新 Dockerfile: WORKDIR/COPY/CMD 中 `hz-admin-base` -> `hab` (3 处) | 待办 |  |  | L2#5 |
| 26 | 更新 .drone.yml: 全局替换 `hz-admin-base` -> `hab` (8 处) | 待办 |  |  | L2#5 |
| 27 | 更新 config.example.yaml: db-type 改为 sqlite, use-redis 改为 false (模板默认配置) | 待办 |  |  | L2#6 |
| 28 | 确认废弃 skill 已清理: .claude/skills/ 下 code-generator/mysql-operator/skill-creator 已删除 | 待办 |  |  | L2#7 |
| 29 | 更新 README.md: 改为通用模板说明（项目定位/快速开始/目录结构/扩展方式） | 待办 |  |  | L2#8 |
| 30 | 最终编译验证: `go build` + 启动测试确认项目可正常运行 | 待办 |  |  | L2#9 |
