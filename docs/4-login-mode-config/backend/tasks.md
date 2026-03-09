# 任务清单

> 计划: login-mode-config/backend | 创建: 2026-03-09

| # | 任务 | 状态 | 开始日期 | 完成日期 | 备注 |
|---|------|------|----------|----------|------|
| 1 | System 结构体新增 LoginMode 字段 (`server/config/system.go`) | 已完成 | 2026-03-09 | 2026-03-09 | mapstructure:"login-mode", 默认 simple |
| 2 | authRouter 新增 GET /auth/login-mode 路由 (`server/router/system/sys_base.go`) | 已完成 | 2026-03-09 | 2026-03-09 | 公开接口, 无需认证 |
| 3 | 实现 GetLoginMode 接口 (`server/api/v1/system/sys_user.go`) | 已完成 | 2026-03-09 | 2026-03-09 | 返回 {login_mode: string} |
| 4 | 改造 PasswordLogin: 按 loginMode 条件跳过验证码 | 已完成 | 2026-03-09 | 2026-03-09 | simple 跳过, captcha/strict 保留 |
| 5 | 改造 PasswordLogin: simple/captcha 模式调用 TokenNext 签发 JWT | 已完成 | 2026-03-09 | 2026-03-09 | 替换 pwd.SetPwdToken |
| 6 | 改造 PasswordLogin: strict 模式保持原有 SetPwdToken 逻辑 | 已完成 | 2026-03-09 | 2026-03-09 | 确保不引入回归 |
