# QA 技术方案 — hab-autocode-test

## 技术栈

- 测试工具: bash + curl + jq (命令行 API 测试)
- 服务环境: HAB server (Go + Gin，本地运行)
- 数据库: SQLite (HAB 默认配置)
- 辅助工具: diff (文件对比), go build (编译验证)

## 测试环境准备

### 前置条件

1. HAB server 源码位于 `/Users/luca/work/gva-base/server/`
2. Skill 文件通过符号链接: `.claude/skills` -> `/Users/luca/work/hz-agents/.claude/skills`
3. 需要可用的 `config.yaml`（基于 `config.example.yaml` 创建，填入 `api-key`）
4. SQLite 数据库已初始化

### 环境启动流程

```bash
# 1. 确认符号链接有效
ls -la .claude/skills/hab-autocode/SKILL.md

# 2. 配置 API Key
cd server
cp config.example.yaml config.yaml  # 如果不存在
# 编辑 config.yaml，在 autocode 段设置 api-key

# 3. 启动 HAB 服务
go run main.go &

# 4. 等待服务就绪
sleep 3
curl -s http://localhost:9688/health || echo "服务未启动"
```

### 环境清理

测试完成后需清理测试产生的数据：
- 回滚所有测试创建的代码模块
- 删除测试创建的 Package
- 停止 HAB 服务

## 测试策略

### 分层测试

| 层级 | 测试内容 | 方法 | 依赖 |
|------|----------|------|------|
| L1 静态验证 | 脚本语法、JSON 格式、文件存在性 | bash -n / jq / ls | 无需启动服务 |
| L2 单接口测试 | 14 个 API 端点逐一调用 | curl + jq 解析响应 | 需要 HAB 服务运行 |
| L3 认证测试 | API Key 三种场景 | curl 不同 header 组合 | 需要 HAB 服务运行 |
| L4 端到端测试 | 完整业务流程 | 按顺序调用多个 API | 需要 HAB 服务运行 |

### 测试执行顺序

L1 -> L2 -> L3 -> L4，前一层通过后再执行下一层。

## 详细测试设计

### 1. 静态验证 (L1)

#### 1.1 脚本语法检查

```bash
# config.sh 语法检查
bash -n .claude/skills/hab-autocode/scripts/config.sh

# autocode.sh 语法检查
bash -n .claude/skills/hab-autocode/scripts/autocode.sh

# 执行权限检查
test -x .claude/skills/hab-autocode/scripts/config.sh
test -x .claude/skills/hab-autocode/scripts/autocode.sh
```

#### 1.2 示例 JSON 验证

```bash
# 格式正确性
jq . .claude/skills/hab-autocode/examples/create-package.json
jq . .claude/skills/hab-autocode/examples/create-module.json
jq . .claude/skills/hab-autocode/examples/create-module-simple.json

# 必填字段检查 -- create-package.json
jq -e '.packageName and .template' .claude/skills/hab-autocode/examples/create-package.json

# 必填字段检查 -- create-module.json
jq -e '.package and .tableName and .structName and .fields' .claude/skills/hab-autocode/examples/create-module.json
```

#### 1.3 SKILL.md 文档存在性

```bash
test -f .claude/skills/hab-autocode/SKILL.md
wc -l .claude/skills/hab-autocode/SKILL.md  # 应有足够内容
```

### 2. 单接口测试 (L2)

每个测试用例格式：发送请求 -> 检查 HTTP 状态码 -> 检查响应 JSON 中 `code` 字段为 0 -> 检查返回数据结构。

#### 2.1 Package 管理

| 用例 ID | API | 方法 | 期望 | 验证点 |
|---------|-----|------|------|--------|
| PKG-01 | /autoCode/getPackage | POST {} | code=0 | data.pkgs 为数组 |
| PKG-02 | /autoCode/getTemplates | GET | code=0 | data 包含 "package" |
| PKG-03 | /autoCode/createPackage | POST | code=0 | 创建 test-pkg 包成功 |
| PKG-04 | /autoCode/getPackage | POST {} | code=0 | pkgs 中包含 test-pkg |
| PKG-05 | /autoCode/delPackage | POST | code=0 | 删除 test-pkg 成功 |
| PKG-06 | /autoCode/createPackage (重复) | POST | code!=0 | 创建已存在的包返回错误 |

#### 2.2 数据库查询

| 用例 ID | API | 方法 | 期望 | 验证点 |
|---------|-----|------|------|--------|
| DB-01 | /autoCode/getDB | GET | code=0 | data.dbs 为数组 |
| DB-02 | /autoCode/getTables | GET | code=0 | data.tables 为数组 |
| DB-03 | /autoCode/getColumn?tableName=sys_users | GET | code=0 | data.columns 包含 id 列 |

#### 2.3 代码生成

| 用例 ID | API | 方法 | 期望 | 验证点 |
|---------|-----|------|------|--------|
| GEN-01 | /autoCode/preview | POST (module JSON) | code=0 | data.autoCode 包含文件内容 |
| GEN-02 | /autoCode/createTemp | POST (module JSON) | code=0 | 文件实际生成 |
| GEN-03 | /autoCode/addFunc | POST (func JSON, isPreview=true) | code=0 | 返回预览代码 |

#### 2.4 历史和回滚

| 用例 ID | API | 方法 | 期望 | 验证点 |
|---------|-----|------|------|--------|
| HIS-01 | /autoCode/getSysHistory | POST {page:1,pageSize:10} | code=0 | data.list 为数组 |
| HIS-02 | /autoCode/getMeta | POST {id: N} | code=0 | data.meta 包含原始请求 |
| HIS-03 | /autoCode/rollback | POST {id: N, deleteTable: true, ...} | code=0 | 文件已移除 |
| HIS-04 | /autoCode/delSysHistory | POST {id: N} | code=0 | 历史记录已删除 |

### 3. 认证测试 (L3)

| 用例 ID | 场景 | Header | 期望 |
|---------|------|--------|------|
| AUTH-01 | 有效 API Key | x-api-key: <正确值> | 200, code=0 |
| AUTH-02 | 无效 API Key | x-api-key: wrong-key | 401, "invalid api key" |
| AUTH-03 | 未携带 Key 且无 JWT | 无 x-api-key, 无 x-token | 401, 未登录提示 |
| AUTH-04 | API Key 访问非 AutoCode 路由 | x-api-key + GET /user/getUserList | 401, API Key 不生效 |

### 4. 端到端流程测试 (L4)

完整流程按以下顺序执行，每步验证成功后才进入下一步：

```
Step 1: 查询现有包列表 (getPackage)
  -> 记录当前包数量

Step 2: 创建测试包 "e2e-test" (createPackage)
  -> 验证返回 code=0
  -> 验证 server/api/v1/e2e_test/enter.go 等文件已创建

Step 3: 查询数据库表结构 (getDB -> getTables)
  -> 记录可用表信息

Step 4: 预览代码 (preview)
  -> 使用 create-module-simple.json 为模板，package 改为 "e2e-test"
  -> 验证返回预览内容

Step 5: 生成代码 (createTemp)
  -> 验证返回 code=0
  -> 验证生成的文件存在 (server/ 和 web/ 下)
  -> go build ./... 编译验证

Step 6: 查询生成历史 (getSysHistory)
  -> 确认新记录存在，flag=0

Step 7: 回滚 (rollback)
  -> deleteApi=true, deleteMenu=true, deleteTable=true
  -> 验证文件已移至 rm_file/
  -> 验证历史记录 flag=1

Step 8: 清理 -- 删除测试包 (delPackage)
  -> 验证 e2e-test 包已删除

Step 9: 最终验证 -- 确认环境恢复到测试前状态
  -> getPackage 确认包数量与 Step 1 一致
  -> go build ./... 编译通过
```

### 5. autocode.sh 脚本命令测试

验证脚本封装的每个命令都能正常工作：

| 命令 | 测试方法 | 期望 |
|------|----------|------|
| ./autocode.sh packages | 直接执行 | 返回 JSON，code=0 |
| ./autocode.sh templates | 直接执行 | 返回模板列表 |
| ./autocode.sh create-package '...' | 传入 JSON | 创建包成功 |
| ./autocode.sh delete-package N | 传入 ID | 删除包成功 |
| ./autocode.sh preview file.json | 传入文件 | 返回预览内容 |
| ./autocode.sh get-db | 直接执行 | 返回数据库列表 |
| ./autocode.sh get-tables | 直接执行 | 返回表列表 |
| ./autocode.sh get-columns sys_users | 传入表名 | 返回列信息 |
| ./autocode.sh history | 直接执行 | 返回历史列表 |
| ./autocode.sh help | 直接执行 | 输出帮助信息 |

## 测试报告模板

```markdown
# hab-autocode-test 测试报告

## 测试环境
- HAB 版本: x.x.x
- Go 版本: x.x.x
- 数据库: SQLite
- 测试日期: YYYY-MM-DD

## 测试结果总览
| 类别 | 总数 | 通过 | 失败 | 跳过 |
|------|------|------|------|------|
| 静态验证 | N | | | |
| 单接口测试 | N | | | |
| 认证测试 | N | | | |
| 端到端测试 | N | | | |
| **合计** | **N** | | | |

## 详细结果
[逐一列出每个用例的执行结果]

## 验收清单对照
[与 2-hab-autocode-skill 验收清单 17 项逐一对照]

## 发现的问题
| 编号 | 描述 | 严重程度 | 状态 |
|------|------|----------|------|

## 结论
[通过 / 有条件通过 / 未通过]
```

## 依赖与约束

- 需要 HAB 服务在本地运行（`go run main.go` 或编译后的二进制）
- 需要配置有效的 `api-key`
- 测试过程中会创建和删除 Package / 模块，需确保不与正在使用的数据冲突
- 端到端测试会修改文件系统（生成代码文件），回滚后需验证清理干净
- Skill 文件通过符号链接访问，测试前需确认链接有效
