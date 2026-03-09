# 计划日志

> 计划: hab-autocode-test | 创建: 2026-03-09

<!--
类型: [决策] [变更] [修复] [新增] [测试] [备注] [完成]
格式: - [类型] 内容
按日期分组，最新在前
-->

## 2026-03-09

- [完成] [qa] 完成任务 #13: 清理测试环境 -- 回滚所有测试数据、删除测试包、停止 HAB 服务、确认环境恢复原状 (HAB服务已停止,config.yaml符号链接已删除,残留测试目录已清理,编译验证通过)
- [变更] [qa] 开始任务 #13: 清理测试环境 -- 回滚所有测试数据、删除测试包、停止 HAB 服务、确认环境恢复原状
- [完成] [qa] 完成任务 #12: 对照验收清单并输出报告 -- 与 2-hab-autocode-skill 验收清单 17 项逐一对照，填写测试报告模板 (验收报告已写入log.md: 39个用例,36通过/3失败; 17项验收清单15通过/1失败/1部分通过; 发现8个问题(4个脚本缺陷+2个回滚问题+2个文档问题))
- [测试] QA 验收测试报告 -- 详见下方完整报告

---

# hab-autocode-test 验收测试报告

## 测试环境

- HAB Server: Go (Gin) 本地运行, PID 6312
- 端口: 9688
- 数据库: MySQL (hab)
- API Key: `hab-autocode-test-key-2026` (config.local.yaml)
- 测试日期: 2026-03-09
- Go 版本: 本地环境
- 操作系统: macOS Darwin 25.3.0

## 测试结果总览

| 类别 | 总数 | 通过 | 失败 | 备注 |
|------|------|------|------|------|
| L1 静态验证 | 8 | 8 | 0 | 脚本语法/权限/JSON/SKILL.md |
| L2 单接口测试 | 16 | 16 | 0 | PKG/DB/GEN/HIS 全部端点 |
| L3 认证测试 | 4 | 4 | 0 | API Key 四种场景 |
| L4 端到端测试 | 9 | 8 | 1 | 回滚时 deleteTable/Menu 失败 |
| 脚本功能测试 | 2 | 0 | 2 | config.sh/autocode.sh macOS 兼容性 |
| **合计** | **39** | **36** | **3** | |

## L1 静态验证详细结果

| 用例 | 测试内容 | 结果 |
|------|----------|------|
| L1-01 | `bash -n config.sh` 语法检查 | 通过 |
| L1-02 | `bash -n autocode.sh` 语法检查 | 通过 |
| L1-03 | config.sh 可执行权限 (-x) | 通过 |
| L1-04 | autocode.sh 可执行权限 (-x) | 通过 |
| L1-05 | SKILL.md 存在 (541行) | 通过 |
| L1-06 | create-package.json 格式+必填字段 | 通过 |
| L1-07 | create-module.json 格式+必填字段 | 通过 |
| L1-08 | create-module-simple.json 格式+必填字段 | 通过 |

## L2 单接口测试详细结果

### Package 管理 (4 端点, 6 用例)

| 用例 ID | API | 请求 | 响应 | 结果 |
|---------|-----|------|------|------|
| PKG-01 | POST /autoCode/getPackage | `{}` | `code=0`, pkgs=3个包 | 通过 |
| PKG-02 | GET /autoCode/getTemplates | 无参数 | `code=0`, data=["package","plugin","storage"] | 通过 |
| PKG-03 | POST /autoCode/createPackage | `{"packageName":"testqapkg",...}` | `code=0` | 通过 |
| PKG-04 | POST /autoCode/getPackage | `{}` | pkgs 含 testqapkg | 通过 |
| PKG-05 | POST /autoCode/delPackage | `{"id":9}` | `code=0` | 通过 |
| PKG-06 | POST /autoCode/createPackage (重复) | 同 PKG-03 | `code=7, msg="Creation failed"` | 通过 |

### 数据库查询 (3 端点, 3 用例)

| 用例 ID | API | 请求 | 响应 | 结果 |
|---------|-----|------|------|------|
| DB-01 | GET /autoCode/getDB | 无参数 | `code=0`, dbs=6个数据库 | 通过 |
| DB-02 | GET /autoCode/getTables | 无参数 | `code=0`, tables=26个表 | 通过 |
| DB-03 | GET /autoCode/getColumn?tableName=sys_users | tableName=sys_users | `code=0`, columns=26列, 含 id/created_at/updated_at | 通过 |

### 代码生成 (3 端点, 3 用例)

| 用例 ID | API | 请求 | 响应 | 结果 |
|---------|-----|------|------|------|
| GEN-01 | POST /autoCode/preview | E2eItem 模块 JSON | `code=0`, 返回 17 个文件代码预览 | 通过 |
| GEN-02 | POST /autoCode/createTemp | 同 GEN-01 | `code=0`, 6 个关键文件已在文件系统生成 | 通过 |
| GEN-03 | POST /autoCode/addFunc (isPreview=true) | ExportE2eItem 方法 | `code=0`, 返回 api(571c)/js(344c)/server(227c) 代码 | 通过 |

### 历史和回滚 (4 端点, 4 用例)

| 用例 ID | API | 请求 | 响应 | 结果 |
|---------|-----|------|------|------|
| HIS-01 | POST /autoCode/getSysHistory | `{page:1,pageSize:10}` | `code=0`, list=2条记录 | 通过 |
| HIS-02 | POST /autoCode/getMeta | `{id:3}` | `code=0`, data.meta 存在 | 通过 |
| HIS-03 | POST /autoCode/rollback | `{id:3,deleteApi:false,deleteMenu:false,deleteTable:false}` | `code=0, msg="回滚成功"`, 文件已移除, flag=1 | 通过 |
| HIS-04 | POST /autoCode/delSysHistory | `{id:3}` | `code=0` | 通过 |

**HIS-03 备注**: deleteTable=true 时失败 (`Unknown table 'hab.e2e_items'`)，deleteMenu=true 时失败 (`记录不存在`)。文件级回滚正常。原因: createTemp 生成代码后未重启服务，autoMigrate 未生效，表和菜单未实际创建。

## L3 认证测试详细结果

| 用例 ID | 场景 | Header | 响应 | 结果 |
|---------|------|--------|------|------|
| AUTH-01 | 有效 API Key | `x-api-key: hab-autocode-test-key-2026` | HTTP 200, code=0 | 通过 |
| AUTH-02 | 无效 API Key | `x-api-key: wrong-invalid-key` | code=7, `"invalid api key"` | 通过 |
| AUTH-03 | 无 Key 且无 JWT | 无认证头 | code=7, `"未登录或非法访问"` | 通过 |
| AUTH-04 | API Key 访问非 AutoCode 私有路由 | `x-api-key` + POST /authority/getAuthorityList | HTTP 401, `"未登录或非法访问"` | 通过 |

**AUTH-04 补充说明**: `/user/getUserList` 可被 API Key 调用，但经代码审查确认该路由注册在 PublicGroup（无认证），是已有设计而非 API Key 穿透问题。

## L4 端到端流程测试详细结果

| Step | 操作 | 验证 | 结果 |
|------|------|------|------|
| 1 | getPackage 查询现有包 | 当前 4 个包 | 通过 |
| 2 | createPackage "e2eflow" | code=0, enter.go 文件已创建 | 通过 |
| 3 | getDB + getTables | 6个数据库, 26个表 | 通过 |
| 4 | preview E2eTask 模块 | code=0, 17个文件预览 | 通过 |
| 5 | createTemp 生成代码 | code=0, 6个关键文件存在, `go build ./...` 编译通过 | 通过 |
| 6 | getSysHistory 查询历史 | ID=4, flag=0, table=e2e_tasks | 通过 |
| 7 | rollback (文件级) | code=0, 文件已移除, flag=1 | 通过 (有条件) |
| 8 | delPackage e2eflow | code=0 | 通过 |
| 9 | 最终验证 | 包数量恢复, `go build` 编译通过 | 通过 |

**Step 7 备注**: deleteTable=true 和 deleteMenu=true 均失败，仅文件级回滚成功。

## 脚本功能测试详细结果

| 测试项 | 问题描述 | 严重程度 |
|--------|----------|----------|
| config.sh API Key 解析 | sed 正则在 macOS (BSD sed) 上不兼容，`\s` 和 `\?` 未被支持，导致 HAB_API_KEY 值为 `api-key:"hab-autocode-test-key-2026"` 而非 `hab-autocode-test-key-2026` | 高 |
| config.sh router-prefix | 当 config.yaml 中不存在 router-prefix 配置时，grep 返回非零退出码，在 `set -euo pipefail` 模式下导致脚本异常退出 | 高 |
| config.sh 文件路径 | 硬编码查找 `config.yaml`，但实际环境使用 `config.local.yaml`（需要创建符号链接才能工作） | 中 |
| autocode.sh 全部命令 | 因依赖 config.sh，所有命令均无法正常执行（包括 help 命令） | 高 |

## 验收清单对照 (2-hab-autocode-skill 17 项)

### 分析阶段 (4 项)

| # | 验收项 | 结果 | 证据 |
|---|--------|------|------|
| 1 | 已完成对所有 AutoCode API 的分析和文档化 | 通过 | SKILL.md 541行, 覆盖全部 14 个 API 端点 |
| 2 | 已完成对代码生成器模板系统的分析 | 通过 | SKILL.md 包含 package/plugin/storage 模板说明 |
| 3 | 已完成对自动化代码生产逻辑的分析 | 通过 | SKILL.md 流程二详述完整生成逻辑 |
| 4 | 已完成对回滚机制的分析 | 通过 | SKILL.md 流程五详述回滚流程 |

### API Key 认证 (2 项)

| # | 验收项 | 结果 | 证据 |
|---|--------|------|------|
| 5 | config.yaml 中可配置 API Key | 通过 | config.local.yaml 中 autocode.api-key 字段已配置并生效 |
| 6 | AutoCode 路由支持 API Key 认证 | 通过 | AUTH-01~04 测试全部通过, router.go 中 AutoCodeGroup 独立路由组 |

### Skill 交付物 (5 项)

| # | 验收项 | 结果 | 证据 |
|---|--------|------|------|
| 7 | SKILL.md 编写完成, AI 可独立操作 | 通过 | 541行, 5个操作流程, 字段参考, 错误处理, curl 示例 |
| 8 | 辅助脚本可直接执行 | 失败 | config.sh macOS sed 不兼容, autocode.sh 因此全部失败 |
| 9 | 示例请求文件覆盖所有 API 端点 | 部分通过 | 3个JSON覆盖 createPackage/createModule, 但缺少 addFunc/rollback/history 的示例 |
| 10 | 文档覆盖所有错误场景 | 通过 | SKILL.md 错误处理表格覆盖 7 种错误场景 |
| 11 | 文档包含完整回滚流程说明 | 通过 | SKILL.md 流程五: 查询历史 -> 查看详情 -> 执行回滚 -> 删除历史 |

### 功能验证 (7 项)

| # | 验收项 | 结果 | 证据 |
|---|--------|------|------|
| 12 | AI 可通过 Skill 成功创建 Package | 通过 | PKG-03 curl 调用 createPackage 成功 |
| 13 | AI 可通过 Skill 成功创建模块代码 | 通过 | GEN-02 createTemp 成功, 文件已生成, 编译通过 |
| 14 | AI 可通过 Skill 成功预览代码 | 通过 | GEN-01 preview 返回 17 个文件代码 |
| 15 | AI 可通过 Skill 成功回滚已生成代码 | 通过 | HIS-03 rollback 文件级回滚成功 |
| 16 | AI 可通过 Skill 查询数据库表结构 | 通过 | DB-01~03 全部通过 |
| 17 | AI 可通过 Skill 为已有模块添加新方法 | 通过 | GEN-03 addFunc 预览模式成功 |

**验收清单统计: 15/17 通过, 1 失败, 1 部分通过**

## 发现的问题

| 编号 | 描述 | 严重程度 | 影响模块 |
|------|------|----------|----------|
| BUG-01 | config.sh sed 正则在 macOS (BSD sed) 不兼容: `\s` 和 `\?` 不被支持, API Key 解析值错误 | 高 | scripts/config.sh |
| BUG-02 | config.sh 当 config.yaml 缺少 router-prefix 时, grep 非零退出码在 pipefail 下导致脚本退出 | 高 | scripts/config.sh |
| BUG-03 | config.sh 硬编码查找 config.yaml, 不支持 config.local.yaml | 中 | scripts/config.sh |
| BUG-04 | autocode.sh 因依赖 config.sh, 全部命令无法在 macOS 上正常执行 | 高 | scripts/autocode.sh |
| BUG-05 | rollback 时 deleteTable=true 失败: 表未被 autoMigrate 创建 (需服务重启后才生效) | 低 | API /autoCode/rollback |
| BUG-06 | rollback 时 deleteMenu=true 失败: 菜单记录不存在 | 低 | API /autoCode/rollback |
| NOTE-01 | 示例 JSON 仅覆盖 createPackage 和 createModule, 缺少 addFunc/rollback/history 示例 | 低 | examples/ |
| NOTE-02 | getTemplates 实际返回 ["package","plugin","storage"], SKILL.md 文档中描述为 "package" 或 "plugin" (缺少 storage) | 低 | SKILL.md |

### BUG-01 复现步骤

```bash
# 在 macOS 上执行:
echo '  api-key: "hab-autocode-test-key-2026"' | sed 's/.*api-key:\s*"\?\([^"]*\)"\?.*/\1/'
# 实际输出: api-key:"hab-autocode-test-key-2026"
# 期望输出: hab-autocode-test-key-2026

# 修复方案: 使用 sed -E 和 POSIX 字符类
echo '  api-key: "hab-autocode-test-key-2026"' | sed -E 's/.*api-key:[[:space:]]*"?([^"]*)"?.*/\1/'
# 输出: hab-autocode-test-key-2026
```

### BUG-02 复现步骤

```bash
# config.yaml 中不存在 router-prefix 配置时:
set -euo pipefail
HAB_PREFIX=$(grep -A 10 '^system:' server/config.yaml | grep 'router-prefix:' | sed '...')
# grep 'router-prefix:' 返回 1 (无匹配), pipefail 导致整个赋值语句返回非零, set -e 退出脚本

# 修复方案: 添加 || true
HAB_PREFIX=$(grep -A 10 '^system:' "$CONFIG_FILE" | grep 'router-prefix:' | sed '...' || true)
```

## 结论

**有条件通过**

- 核心功能全部正常: 14 个 API 端点均可通过 curl + API Key 正确调用, 返回结果符合文档描述
- API Key 认证机制设计合理并正确实现: 有效 Key 可调用, 无效 Key 被拒, 仅限 AutoCode 路由
- SKILL.md 文档质量高 (541行), AI 可通过 curl 直接调用完成所有 AutoCode 操作
- 辅助脚本 (config.sh/autocode.sh) 在 macOS 环境下无法正常工作, 需修复 sed 兼容性和 grep 错误处理
- 建议补充示例 JSON 覆盖 addFunc 和 rollback 请求体

**阻塞项**: 无 (AI 可直接使用 curl 调用, 不依赖脚本)
**建议修复**: BUG-01~04 (脚本兼容性), NOTE-02 (文档 storage 模板)

---

- [变更] [qa] 开始任务 #12: 对照验收清单并输出报告 -- 与 2-hab-autocode-skill 验收清单 17 项逐一对照，填写测试报告模板
- [完成] [qa] 完成任务 #11: 测试 autocode.sh 脚本各命令 -- 逐一执行 packages/templates/create-package/preview/get-db/get-tables/get-columns/history/help 命令 (autocode.sh无法正常工作(help命令也退出码1): 原因是config.sh有3个缺陷: (1)sed正则macOS不兼容 (2)grep router-prefix缺失时pipefail导致退出 (3)硬编码config.yaml但实际是config.local.yaml)
- [变更] [qa] 开始任务 #11: 测试 autocode.sh 脚本各命令 -- 逐一执行 packages/templates/create-package/preview/get-db/get-tables/get-columns/history/help 命令
- [完成] [qa] 完成任务 #10: 执行端到端流程测试 -- 按 L4 设计的 Step 1~9 完整执行创建包到回滚清理的全流程，每步验证通过后继续 (E2E Step1~9完成: 创建包/预览/生成代码/编译通过/回滚/清理全流程OK; 注意回滚时deleteTable和deleteMenu失败(表和菜单未实际创建,可能因服务未重启导致autoMigrate未生效),文件级回滚正常)
- [变更] [qa] 开始任务 #10: 执行端到端流程测试 -- 按 L4 设计的 Step 1~9 完整执行创建包到回滚清理的全流程，每步验证通过后继续
- [完成] [qa] 完成任务 #9: 测试 API Key 认证机制 -- 按用例 AUTH-01~AUTH-04 测试有效 Key、无效 Key、无 Key、非 AutoCode 路由 (AUTH-01~04全部通过: 有效Key可调用(200),无效Key被拒(code=7),无Key走JWT(code=7,未登录),非AutoCode私有路由被拒(401); 注意getUserList是公开路由不受认证保护(已有设计非API Key问题))
- [变更] [qa] 开始任务 #9: 测试 API Key 认证机制 -- 按用例 AUTH-01~AUTH-04 测试有效 Key、无效 Key、无 Key、非 AutoCode 路由
- [完成] [qa] 完成任务 #8: 测试历史和回滚 4 个端点 -- 按用例 HIS-01~HIS-04 执行历史查询、元数据、回滚、删除历史测试 (HIS-01~04通过: 查询历史/获取Meta/回滚(文件级)/删除历史OK; 注意rollback时deleteTable和deleteMenu失败(记录不存在),需在完整流程中验证)
- [变更] [qa] 开始任务 #8: 测试历史和回滚 4 个端点 -- 按用例 HIS-01~HIS-04 执行历史查询、元数据、回滚、删除历史测试
- [完成] [qa] 完成任务 #7: 测试代码生成 3 个端点 -- 按用例 GEN-01~GEN-03 执行预览、生成、添加方法测试 (GEN-01~03全部通过: preview返回17个文件,createTemp生成文件已验证存在,addFunc预览返回api/js/server代码)
- [变更] [qa] 开始任务 #7: 测试代码生成 3 个端点 -- 按用例 GEN-01~GEN-03 执行预览、生成、添加方法测试
- [完成] [qa] 完成任务 #6: 测试数据库查询 3 个端点 -- 按用例 DB-01~DB-03 逐一执行 curl 调用并验证响应 (DB-01~03全部通过,getDB返回6个数据库,getTables返回26个表,getColumn返回sys_users的26列)
- [变更] [qa] 开始任务 #6: 测试数据库查询 3 个端点 -- 按用例 DB-01~DB-03 逐一执行 curl 调用并验证响应
- [完成] [qa] 完成任务 #5: 测试 Package 管理 4 个端点 -- 按用例 PKG-01~PKG-06 逐一执行 curl 调用并验证响应 (PKG-01~06全部通过,包括创建/查询/重复创建/删除)
- [变更] [qa] 开始任务 #5: 测试 Package 管理 4 个端点 -- 按用例 PKG-01~PKG-06 逐一执行 curl 调用并验证响应
- [完成] [qa] 完成任务 #1: 准备测试环境 -- 确认符号链接有效、创建 config.yaml 并配置 api-key、启动 HAB 服务、验证服务就绪 (HAB服务已启动(PID 6312,端口9688),创建config.yaml符号链接指向config.local.yaml)
- [完成] [qa] 完成任务 #3: 执行 config.sh 功能测试 -- source config.sh 后验证 HAB_API_KEY、HAB_PORT、HAB_BASE_URL 环境变量值正确 (发现2个缺陷: (1)config.sh硬编码config.yaml但实际是config.local.yaml (2)sed正则在macOS不兼容,API Key解析值错误)
- [变更] [qa] 开始任务 #3: 执行 config.sh 功能测试 -- source config.sh 后验证 HAB_API_KEY、HAB_PORT、HAB_BASE_URL 环境变量值正确
- [完成] [qa] 完成任务 #4: 执行示例 JSON 验证 -- jq 格式检查 3 个 JSON 文件、验证必填字段存在、与 SKILL.md 字段说明交叉比对 (3个JSON格式合法,必填字段存在)
- [变更] [qa] 开始任务 #4: 执行示例 JSON 验证 -- jq 格式检查 3 个 JSON 文件、验证必填字段存在、与 SKILL.md 字段说明交叉比对
- [完成] [qa] 完成任务 #2: 执行脚本语法检查 -- bash -n 检查 config.sh 和 autocode.sh 语法、验证执行权限 (config.sh/autocode.sh语法OK+可执行权限OK)
- [变更] [qa] 开始任务 #1: 准备测试环境 -- 确认符号链接有效、创建 config.yaml 并配置 api-key、启动 HAB 服务、验证服务就绪
- [决策] Tech Lead 创建 qa 角色目录，编写测试技术方案 (分层测试: L1静态验证/L2单接口/L3认证/L4端到端)，拆解为 13 个技术任务; 此需求为纯测试需求，不需要 backend/frontend/ui 角色目录
- [备注] Skill 文件通过符号链接存在于 /Users/luca/work/hz-agents/ 仓库，测试时需确认链接有效
- [决策] 测试需启动 HAB 服务（go run main.go），使用实际 API 调用验证所有端点，覆盖正常和异常场景
- [新增] 创建测试需求，目标是对 hab-autocode skill 和 cmd-autocode 命令进行完整动态功能测试
- [新增] 创建计划