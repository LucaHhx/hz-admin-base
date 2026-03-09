# QA 技术任务

> 计划: hab-autocode-test | 角色: qa | 创建: 2026-03-09

| # | 任务 | 状态 | 开始日期 | 完成日期 | 备注 |
|---|------|------|----------|----------|------|
| 1 | 准备测试环境 -- 确认符号链接有效、创建 config.yaml 并配置 api-key、启动 HAB 服务、验证服务就绪 | 已完成 | 2026-03-09 | 2026-03-09 | 对应 L2 任务 #1 前置; HAB服务已启动(PID 6312,端口9688),创建config.yaml符号链接指向config.local.yaml |
| 2 | 执行脚本语法检查 -- bash -n 检查 config.sh 和 autocode.sh 语法、验证执行权限 | 已完成 | 2026-03-09 | 2026-03-09 | 对应 L2 任务 #1; config.sh/autocode.sh语法OK+可执行权限OK |
| 3 | 执行 config.sh 功能测试 -- source config.sh 后验证 HAB_API_KEY、HAB_PORT、HAB_BASE_URL 环境变量值正确 | 已完成 | 2026-03-09 | 2026-03-09 | 对应 L2 任务 #1; 发现2个缺陷: (1)config.sh硬编码config.yaml但实际是config.local.yaml (2)sed正则在macOS不兼容,API Key解析值错误 |
| 4 | 执行示例 JSON 验证 -- jq 格式检查 3 个 JSON 文件、验证必填字段存在、与 SKILL.md 字段说明交叉比对 | 已完成 | 2026-03-09 | 2026-03-09 | 对应 L2 任务 #2; 3个JSON格式合法,必填字段存在 |
| 5 | 测试 Package 管理 4 个端点 -- 按用例 PKG-01~PKG-06 逐一执行 curl 调用并验证响应 | 已完成 | 2026-03-09 | 2026-03-09 | 对应 L2 任务 #3; PKG-01~06全部通过,包括创建/查询/重复创建/删除 |
| 6 | 测试数据库查询 3 个端点 -- 按用例 DB-01~DB-03 逐一执行 curl 调用并验证响应 | 已完成 | 2026-03-09 | 2026-03-09 | 对应 L2 任务 #4; DB-01~03全部通过,getDB返回6个数据库,getTables返回26个表,getColumn返回sys_users的26列 |
| 7 | 测试代码生成 3 个端点 -- 按用例 GEN-01~GEN-03 执行预览、生成、添加方法测试 | 已完成 | 2026-03-09 | 2026-03-09 | 对应 L2 任务 #5，需先创建测试包; GEN-01~03全部通过: preview返回17个文件,createTemp生成文件已验证存在,addFunc预览返回api/js/server代码 |
| 8 | 测试历史和回滚 4 个端点 -- 按用例 HIS-01~HIS-04 执行历史查询、元数据、回滚、删除历史测试 | 已完成 | 2026-03-09 | 2026-03-09 | 对应 L2 任务 #6，依赖任务 #7 生成的记录; HIS-01~04通过: 查询历史/获取Meta/回滚(文件级)/删除历史OK; 注意rollback时deleteTable和deleteMenu失败(记录不存在),需在完整流程中验证 |
| 9 | 测试 API Key 认证机制 -- 按用例 AUTH-01~AUTH-04 测试有效 Key、无效 Key、无 Key、非 AutoCode 路由 | 已完成 | 2026-03-09 | 2026-03-09 | 对应 L2 任务 #7; AUTH-01~04全部通过: 有效Key可调用(200),无效Key被拒(code=7),无Key走JWT(code=7,未登录),非AutoCode私有路由被拒(401); 注意getUserList是公开路由不受认证保护(已有设计非API Key问题) |
| 10 | 执行端到端流程测试 -- 按 L4 设计的 Step 1~9 完整执行创建包到回滚清理的全流程，每步验证通过后继续 | 已完成 | 2026-03-09 | 2026-03-09 | 对应 L2 任务 #8，依赖任务 #5-#9 通过; E2E Step1~9完成: 创建包/预览/生成代码/编译通过/回滚/清理全流程OK; 注意回滚时deleteTable和deleteMenu失败(表和菜单未实际创建,可能因服务未重启导致autoMigrate未生效),文件级回滚正常 |
| 11 | 测试 autocode.sh 脚本各命令 -- 逐一执行 packages/templates/create-package/preview/get-db/get-tables/get-columns/history/help 命令 | 已完成 | 2026-03-09 | 2026-03-09 | 对应 L2 任务 #1; autocode.sh无法正常工作(help命令也退出码1): 原因是config.sh有3个缺陷: (1)sed正则macOS不兼容 (2)grep router-prefix缺失时pipefail导致退出 (3)硬编码config.yaml但实际是config.local.yaml |
| 12 | 对照验收清单并输出报告 -- 与 2-hab-autocode-skill 验收清单 17 项逐一对照，填写测试报告模板 | 已完成 | 2026-03-09 | 2026-03-09 | 对应 L2 任务 #9; 验收报告已写入log.md: 39个用例,36通过/3失败; 17项验收清单15通过/1失败/1部分通过; 发现8个问题(4个脚本缺陷+2个回滚问题+2个文档问题) |
| 13 | 清理测试环境 -- 回滚所有测试数据、删除测试包、停止 HAB 服务、确认环境恢复原状 | 已完成 | 2026-03-09 | 2026-03-09 | 最后执行; HAB服务已停止,config.yaml符号链接已删除,残留测试目录已清理,编译验证通过 |
