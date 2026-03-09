# 变更日志

格式参考 [Keep a Changelog](https://keepachangelog.com/)。

## [Unreleased]

### 新增

- 添加 API Key 认证中间件，支持 AutoCode API 无需登录直接调用
- 创建 hab-autocode skill（SKILL.md / scripts / examples），AI 可通过 Skill 调用 AutoCode API 完成自动化代码生成
- 创建 cmd-autocode 命令，提供交互式代码生成向导（探索阶段 + 执行阶段）
- 项目清理完成：移除 example 模块、announcement/email 插件，统一项目标识为 HAB

### 变更

- 项目标识从 hz-admin-base 统一重命名为 hab
- config.example.yaml 默认使用 SQLite，Redis 改为可选组件

### 修复

### 移除

- 移除 example 模块（客户管理、文件上传下载示例、断点续传示例）前后端代码
- 移除 announcement 公告插件和 email 邮件插件
- 移除 .agents/skills/ 下废弃的旧 skill 文件

## [2.7.9] - 2026-03-06

### 新增

- 初始化项目文档，补充 L1 项目级文档（project.md / tasks.md / CHANGELOG.md）
- 从业务项目剥离为基础底座项目 (hz-admin-base)

### 变更

- 移除业务相关逻辑和数据，保留通用后台管理功能
