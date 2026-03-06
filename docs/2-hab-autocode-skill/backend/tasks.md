# Backend 技术任务

> 计划: hab-autocode-skill | 角色: backend | 创建: 2026-03-06

| # | 任务 | 状态 | 开始日期 | 完成日期 | 备注 |
|---|------|------|----------|----------|------|
| 1 | 扩展 Autocode 配置结构体 -- 在 `server/config/auto_code.go` 的 Autocode 结构体中新增 `ApiKey string` 字段 (mapstructure: "api-key") | 待办 | | | 对应 L2 任务 #5 |
| 2 | 更新配置示例文件 -- 在 `server/config.example.yaml` 的 autocode 段新增 `api-key: ""` 配置项并添加注释说明 | 待办 | | | 对应 L2 任务 #5 |
| 3 | 实现 ApiKeyOrJWT 中间件 -- 新建 `server/middleware/api_key.go`，实现: (1) 检查 x-api-key 请求头 (2) 与 config 中的 key 比对 (3) 命中时注入管理员 claims (AuthorityId=1, ID=1, Username="api-key", UUID=固定值) (4) 未命中时 fallback 到 JWTAuth() | 待办 | | | 对应 L2 任务 #6 |
| 4 | 修改路由初始化 -- 在 `server/initialize/router.go` 中新增 AutoCodeGroup 路由组 (ApiKeyOrJWT + CasbinHandler)，将 InitAutoCodeRouter 和 InitAutoCodeHistoryRouter 从 PrivateGroup 迁移到 AutoCodeGroup | 待办 | | | 对应 L2 任务 #7 |
| 5 | 创建 config.sh 脚本 -- 新建 `.claude/skills/hab-autocode/scripts/config.sh`，从 config.yaml 读取 api-key、server port、router-prefix，导出 HAB_API_KEY / HAB_BASE_URL 环境变量 | 待办 | | | 对应 L2 任务 #8 |
| 6 | 创建 autocode.sh 脚本 -- 新建 `.claude/skills/hab-autocode/scripts/autocode.sh`，封装所有 AutoCode API 的 curl 调用 (packages/create-package/preview/create/rollback/get-db/get-tables/get-columns 等命令) | 待办 | | | 对应 L2 任务 #8 |
| 7 | 创建示例 JSON 文件 -- 新建 `.claude/skills/hab-autocode/examples/` 目录，包含 create-package.json、create-module.json、create-module-simple.json 三个示例请求体 | 待办 | | | 对应 L2 任务 #8 |
| 8 | 编写 SKILL.md 主文档 -- 创建 `.claude/skills/hab-autocode/SKILL.md`，包含: 前置条件、API Key 配置、所有 API 端点详细参数和响应格式、完整操作流程 (创建包→创建模块→预览→确认→回滚)、脚本使用说明、AutoCode/AutoCodeField/AutoFunc 字段参考表、错误处理和常见问题、安全注意事项 | 待办 | | | 对应 L2 任务 #8 |
| 9 | 手动集成测试 -- 配置 api-key 后验证: (1) x-api-key 可访问 AutoCode API (2) 不带 key 走 JWT 认证 (3) 错误 key 返回 401 (4) API Key 不能访问非 AutoCode 路由 (5) 使用 autocode.sh 脚本创建包和预览代码 | 待办 | | | 对应 L2 任务 #9 |
