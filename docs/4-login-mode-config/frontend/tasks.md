# 任务清单

> 计划: login-mode-config/frontend | 创建: 2026-03-09

| # | 任务 | 状态 | 开始日期 | 完成日期 | 备注 |
|---|------|------|----------|----------|------|
| 1 | auth.js 新增 getLoginMode API (GET /auth/login-mode) | 已完成 | 2026-03-09 | 2026-03-09 | 公开接口, 无需 token |
| 2 | auth.js 新增 passwordLogin API (POST /auth/password/login) | 已完成 | 2026-03-09 | 2026-03-09 | 复用已有端点 |
| 3 | index.vue 新增 loginMode 状态, onMounted 时调用 getLoginMode | 已完成 | 2026-03-09 | 2026-03-09 | 默认值 simple |
| 4 | index.vue 实现 simple 模式 UI: 用户名+密码直接登录 | 已完成 | 2026-03-09 | 2026-03-09 | 隐藏验证码和 Passkey 按钮 |
| 5 | index.vue 实现 captcha 模式 UI: 用户名+密码+验证码登录 | 已完成 | 2026-03-09 | 2026-03-09 | 自动获取验证码 |
| 6 | index.vue 实现 handleSimpleLogin 方法, 调用 passwordLogin | 已完成 | 2026-03-09 | 2026-03-09 | 成功后复用 handleLoginSuccess |
| 7 | index.vue strict 模式保持现有分步流程不变 | 已完成 | 2026-03-09 | 2026-03-09 | 用 v-if/v-else 隔离 |
