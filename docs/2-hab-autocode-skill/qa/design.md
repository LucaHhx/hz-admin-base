# QA 技术方案 — hab-autocode-skill

## 测试策略

本需求的代码变更集中在认证层 (中间件 + 配置 + 路由)，核心业务逻辑不变。测试重点放在 API Key 认证中间件的正确性和安全边界。

## 测试范围

### 1. API Key 中间件单元测试

**测试文件:** `server/middleware/api_key_test.go`

| 用例 | 输入 | 期望结果 |
|------|------|----------|
| 有效 API Key | x-api-key 匹配配置 | 200，claims 中 AuthorityId=1, Username="api-key" |
| 无效 API Key | x-api-key 不匹配 | 401，返回 "invalid api key" |
| 未携带 API Key | 无 x-api-key 头 | fallback 到 JWT 认证流程 |
| 未配置 API Key | config 中 api-key 为空 | fallback 到 JWT 认证流程 |
| 空字符串 API Key | x-api-key: "" | fallback 到 JWT 认证流程 |

### 2. 路由隔离测试

| 用例 | 输入 | 期望结果 |
|------|------|----------|
| API Key 访问 AutoCode 路由 | x-api-key + POST /autoCode/getPackage | 200 |
| API Key 访问非 AutoCode 路由 | x-api-key + GET /user/getUserList | 401 (API Key 不生效) |
| JWT 访问 AutoCode 路由 | x-token + POST /autoCode/getPackage | 200 (原有流程不受影响) |

### 3. 端到端集成测试

按用户场景验证完整流程:

| 场景 | 操作 | 验证点 |
|------|------|--------|
| 场景 A | 使用 API Key 创建 Package | 返回成功，目录结构已创建 |
| 场景 B | 使用 API Key 创建模块代码 | 返回成功，文件已生成 |
| 场景 C | 使用 API Key 预览代码 | 返回预览内容，无文件写入 |
| 场景 D | 使用 API Key 回滚 | 返回成功，文件已移除 |
| 场景 E | 使用 API Key 查询表结构 | 返回数据库/表/列信息 |

### 4. 配置变更测试

| 用例 | 操作 | 期望结果 |
|------|------|----------|
| 未配置 api-key | config.yaml 无 api-key 字段 | 系统正常启动，API Key 认证不生效 |
| 配置后清空 | api-key 设为空字符串 | 等同于未配置 |

## 测试方法

- 中间件单元测试: 使用 `httptest` + `gin.CreateTestContext` 构造请求
- 集成测试: 手动 curl 验证 (因 AutoCode 涉及文件系统操作，自动化测试需要隔离环境)

## 不测试范围

- AutoCode 核心业务逻辑 (不在本需求修改范围内)
- 前端 (本需求不涉及前端修改)
- 现有 JWT 认证逻辑 (不修改)
