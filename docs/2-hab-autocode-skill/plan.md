# hab-autocode-skill

> 创建 AI 自动化代码生成 Skill，让 AI 能通过 HAB 的 AutoCode API 自动创建包、生成模块代码，并支持回滚

## 目标

1. 分析并封装 HAB 项目现有的 AutoCode API（Package 创建、模板代码生成、回滚等），形成完整的 AI 可调用的 Skill
2. 创建 `.claude/skills/hab-autocode/SKILL.md`，提供给 AI 使用的完整操作指南
3. 支持通过配置文件中的 API Key 进行身份认证，绕过严格的登录流程
4. 覆盖创建包、创建模块、预览代码、回滚撤回等全部自动化操作

## 范围

**包含:**

- 分析现有 AutoCode 相关 API（Package CRUD、Template Preview/Create、History/Rollback、DB/Tables/Columns 查询）
- 分析代码生成器的模板系统（package/plugin 模板、AST 注入机制）
- 分析自动化代码生产的完整逻辑（文件生成、数据库迁移、API/菜单自动创建）
- 实现 API Key 认证机制，允许配置文件中的 key 直接访问受保护的 AutoCode API
- 创建完整的 Skill 文档（`.claude/skills/hab-autocode/SKILL.md`），包含：
  - API 端点清单及参数说明
  - 操作流程指南（创建包 → 创建模块 → 预览 → 确认）
  - 回滚操作指南
  - API Key 配置和使用方式
  - 错误处理和常见问题
- 分析并封装撤回/回滚机制的 API（包创建回滚、模块创建回滚）

**不包含:**

- 不修改现有 AutoCode 核心业务逻辑
- 不修改前端代码
- 不涉及新的代码生成模板开发

## 核心用户场景

- 场景 A: AI 读取配置文件中的 API Key，通过 curl 调用 HAB API 创建一个新的 Package（如 "order"），自动生成 api/router/service/model 的目录结构和 enter.go
- 场景 B: AI 在已有 Package 中创建一个新的 CRUD 模块（如 "Order" 结构体），自动生成前后端代码、数据库迁移、API 注册、菜单配置
- 场景 C: AI 先使用 Preview API 预览将要生成的代码，确认无误后再调用 Create API 实际生成
- 场景 D: AI 发现生成的代码有问题，使用 Rollback API 撤回已生成的代码、API 注册、菜单配置，恢复到生成前的状态
- 场景 E: AI 查询数据库中的表结构信息（GetDB/GetTables/GetColumn），根据已有表自动填充字段配置后生成代码
- 场景 F: AI 为已有模块添加新的自定义方法（AddFunc），扩展 CRUD 之外的业务接口
- 场景 G (端到端): AI 从零创建一个完整的业务模块（如订单管理），完整流程为：查询现有 Package 列表 → 创建新 Package（如 "order"） → 查询数据库表结构获取字段信息 → 配置模块字段 → 预览生成代码 → 确认并生成代码（含数据库迁移、API 注册、菜单配置） → 验证生成结果 → 如有问题则回滚

## Skill 交付物清单

| 交付物 | 路径 | 说明 |
|--------|------|------|
| Skill 主文档 | `.claude/skills/hab-autocode/SKILL.md` | AI 操作指南，包含所有 API 调用说明、操作流程、错误处理 |
| API 调用脚本 | `.claude/skills/hab-autocode/scripts/autocode.sh` | 封装 curl 调用的辅助脚本，简化 API 请求 |
| 配置读取脚本 | `.claude/skills/hab-autocode/scripts/config.sh` | 从 config.yaml 读取 API Key 和服务地址的辅助脚本 |
| 示例请求文件 | `.claude/skills/hab-autocode/examples/` | 各 API 的示例请求 JSON，供 AI 参考和复用 |

## 涉及的 API 端点

### Package 管理
| API | 方法 | 路径 | 说明 |
|-----|------|------|------|
| 创建 Package | POST | `/autoCode/createPackage` | 创建新的 Package/Plugin 包 |
| 删除 Package | POST | `/autoCode/delPackage` | 删除指定 Package |
| 获取所有 Package | POST | `/autoCode/getPackage` | 获取已有 Package 列表 |
| 获取模板列表 | GET | `/autoCode/getTemplates` | 获取可用模板类型 |

### 代码生成
| API | 方法 | 路径 | 说明 |
|-----|------|------|------|
| 预览代码 | POST | `/autoCode/preview` | 预览生成代码（不写入文件） |
| 创建代码 | POST | `/autoCode/createTemp` | 实际生成并保存代码 |
| 添加方法 | POST | `/autoCode/addFunc` | 为已有模块添加新方法 |

### 数据库查询
| API | 方法 | 路径 | 说明 |
|-----|------|------|------|
| 获取数据库列表 | GET | `/autoCode/getDB` | 获取配置的数据库 |
| 获取表列表 | GET | `/autoCode/getTables` | 获取指定库的所有表 |
| 获取表列 | GET | `/autoCode/getColumn` | 获取表的列信息 |

### 历史和回滚
| API | 方法 | 路径 | 说明 |
|-----|------|------|------|
| 获取历史列表 | POST | `/autoCode/getSysHistory` | 分页获取生成历史 |
| 获取 Meta 信息 | POST | `/autoCode/getMeta` | 获取代码生成的详细元信息 |
| 回滚代码 | POST | `/autoCode/rollback` | 撤回已生成的代码 |
| 删除历史 | POST | `/autoCode/delSysHistory` | 删除历史记录 |

## API Key 认证方案（业务需求）

- AI 需要一种无需人工登录即可调用 AutoCode API 的认证方式
- 通过配置文件中的 API Key 进行身份认证，避免复杂的登录流程
- API Key 仅用于 AutoCode 相关操作，不影响其他系统功能
- API Key 持有者拥有等同于管理员的 AutoCode 操作权限

> [备注] 具体的技术实现方案（中间件设计、请求头格式、认证流程等）由开发团队决定

## 时间线

| 里程碑 | 目标日期 | 状态 |
|--------|----------|------|
| 需求确认与文档完善 | 2026-03-06 | 已完成 |
| API Key 中间件开发 | - | 待办 |
| Skill 文档编写 | - | 待办 |
| 集成测试验证 | - | 待办 |

## 验收清单

### 分析阶段
- [ ] 已完成对所有 AutoCode API 的分析和文档化
- [ ] 已完成对代码生成器模板系统的分析
- [ ] 已完成对自动化代码生产逻辑的分析
- [ ] 已完成对回滚机制的分析

### API Key 认证
- [ ] `config.yaml` 中可配置 API Key
- [ ] AutoCode 路由支持 API Key 认证，无需登录即可调用

### Skill 交付物
- [ ] `SKILL.md` 编写完成，AI 读取后能独立完成所有 AutoCode 操作，无需额外指导
- [ ] 辅助脚本（autocode.sh、config.sh）可直接执行，无需额外安装或配置
- [ ] 示例请求文件覆盖所有 API 端点
- [ ] 文档覆盖所有错误场景（认证失败、参数错误、包已存在、回滚失败等）及对应处理方式
- [ ] 文档包含完整的回滚流程说明（何时回滚、如何回滚、回滚后的状态验证）

### 功能验证
- [ ] AI 可通过 Skill 成功创建 Package
- [ ] AI 可通过 Skill 成功创建模块代码（含自动迁移、API 注册、菜单创建）
- [ ] AI 可通过 Skill 成功预览代码
- [ ] AI 可通过 Skill 成功回滚已生成的代码
- [ ] AI 可通过 Skill 查询数据库表结构
- [ ] AI 可通过 Skill 为已有模块添加新方法
- [ ] AI 可通过 Skill 完成端到端场景（从零创建完整业务模块）
