# hab-autocode-test

> 对 hab-autocode skill 和 cmd-autocode 命令进行完整功能测试，验证所有交付物的正确性和可用性

## 目标

1. 验证 hab-autocode skill 的所有 API 端点可正常调用，返回结果符合预期
2. 验证辅助脚本（autocode.sh / config.sh）可正确执行，参数传递无误
3. 验证示例 JSON 文件的完整性和格式正确性
4. 验证 cmd-autocode 命令的端到端流程（探索阶段 + 执行阶段）可正常运行
5. 与 2-hab-autocode-skill 需求验收清单逐项对照，确认所有验收标准已满足

## 范围

**包含:**

- SKILL.md 文档中所有 13+ 个 API 端点的实际调用测试（需启动 HAB 服务）
- scripts/autocode.sh 的各子命令执行测试（get-packages, get-db, get-tables, get-columns, preview, create, rollback 等）
- scripts/config.sh 的配置读取测试（API Key、服务地址、端口）
- examples/ 目录下 3 个示例 JSON 的格式和字段完整性验证
- cmd-autocode.md 命令流程的可执行性验证
- 2-hab-autocode-skill 验收清单中"分析阶段"、"API Key 认证"、"Skill 交付物"、"功能验证"四组共 17 项的逐项对照
- 错误场景测试：无效 API Key、缺少参数、包名冲突、回滚不存在的记录等

**不包含:**

- 不修改 hab-autocode skill 的代码或文档（测试中发现的问题记录但不修复）
- 不修改 AutoCode 核心业务逻辑
- 不涉及前端功能测试
- 不涉及性能测试或压力测试

## 核心用户场景

- 场景 A: 测试人员启动 HAB 服务后，使用 autocode.sh 脚本逐一调用所有 API 端点，验证返回状态码和数据结构正确
- 场景 B: 测试人员检查 config.sh 是否能正确从 config.yaml / config.local.yaml 读取 API Key 和服务地址
- 场景 C: 测试人员使用示例 JSON 文件作为输入，调用 createPackage 和 preview 接口，验证请求格式与 API 兼容
- 场景 D: 测试人员模拟 cmd-autocode 的完整流程——从查询包列表、创建包、配置字段、预览代码、生成代码到回滚撤回
- 场景 E: 测试人员使用无效 API Key 调用接口，验证认证失败的错误提示和状态码
- 场景 F: 测试人员对照 2-hab-autocode-skill 验收清单，逐项标记通过/未通过，输出最终验收报告

## 时间线

| 里程碑 | 目标日期 | 状态 |
|--------|----------|------|
| 测试计划与用例设计 | 2026-03-09 | 进行中 |
| API 端点逐一调用测试 | - | 待办 |
| 脚本与示例文件测试 | - | 待办 |
| 端到端流程测试 | - | 待办 |
| 验收清单对照与报告输出 | - | 待办 |

## 背景

2-hab-autocode-skill 需求已完成 QA 静态验收（代码审查、编译验证、文档完整性检查），但尚未进行实际运行环境下的动态功能测试。Skill 文件通过符号链接存在于 hz-agents 仓库（`/Users/luca/work/hz-agents/`），需要在 gva-base 项目中启动 HAB 服务后，对所有 API 端点进行真实调用测试，确保文档描述的操作流程在实际环境中可行。

## 验收清单

### 脚本与文件验证
- [ ] autocode.sh 所有子命令可正常执行（无语法错误、参数解析正确）
- [ ] config.sh 可从 config.yaml 正确读取 api-key 和 addr/port 配置
- [ ] examples/ 下 3 个 JSON 文件格式合法，字段与 SKILL.md 中的 API 参数说明一致

### API 端点调用测试
- [ ] Package 管理 4 个端点全部返回正确响应（createPackage / delPackage / getPackage / getTemplates）
- [ ] 代码生成 3 个端点全部返回正确响应（preview / createTemp / addFunc）
- [ ] 数据库查询 3 个端点全部返回正确响应（getDB / getTables / getColumn）
- [ ] 历史和回滚 4 个端点全部返回正确响应（getSysHistory / getMeta / rollback / delSysHistory）
- [ ] API Key 认证正常工作：有效 Key 可调用、无效 Key 被拒绝、未携带 Key 回退到 JWT

### 端到端流程测试
- [ ] 完整流程可执行：创建包 → 配置字段 → 预览代码 → 生成代码 → 编译通过 → 回滚撤回
- [ ] 回滚后生成的文件和数据库记录已正确清理

### 验收清单对照
- [ ] 2-hab-autocode-skill 验收清单"分析阶段" 4 项全部确认通过
- [ ] 2-hab-autocode-skill 验收清单"API Key 认证" 2 项全部确认通过
- [ ] 2-hab-autocode-skill 验收清单"Skill 交付物" 5 项全部确认通过
- [ ] 2-hab-autocode-skill 验收清单"功能验证" 7 项全部确认通过
- [ ] 输出最终测试报告，包含各项测试结果、发现的问题和建议
