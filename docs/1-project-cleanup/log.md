# 计划日志

> 计划: project-cleanup | 创建: 2026-03-06

<!--
类型: [决策] [变更] [修复] [新增] [测试] [备注] [完成]
格式: - [类型] 内容
按日期分组，最新在前
-->

## 2026-03-06

- [新增] 创建 frontend/qa 角色目录，补充前端清理技术方案和 QA 验证方案
- [决策] config.example.yaml 默认 db-type 改为 sqlite、use-redis 改为 false，降低模板项目部署门槛，Redis 为可选依赖
- [变更] PM 评审: 范围补充前端清理（前端页面、路由、API 调用、组件）; 新增 Redis 可选约束（默认 SQLite）; 任务描述调整为功能级别; 任务顺序重排（前端清理提前到任务3）
- [变更] PM 评审: 补充任务 #8 前端残留代码清理、#9 移除 Redis 强制依赖，与 plan.md 范围对齐
- [决策] 保留 plugin/plugin-tool 通用工具目录和 initialize/plugin.go 插件入口框架；保留 business 和 api 两套空壳入口目录作为业务扩展点
- [决策] 清理顺序: 先重命名模块 -> 再删 example -> 再删插件 -> 清理 business -> 清理种子数据 -> 更新配置。每步编译验证，避免中间状态失败
- [决策] Go 模块名选定为 hab (简短、与项目名 HAB 一致)，全局替换 hz-admin-base -> hab，影响 255 个文件 638 处引用
- [决策] 用户确认清理范围: (1) example 模块全部移除 (2) business/api 入口目录保留空壳结构 (3) announcement/email 插件全部移除 (4) 项目标识统一重命名为 HAB
- [新增] 创建计划