# 计划日志

> 计划: login-mode-config | 创建: 2026-03-09

<!--
类型: [决策] [变更] [修复] [新增] [测试] [备注] [完成]
格式: - [类型] 内容
按日期分组，最新在前
-->

## 2026-03-09

- [修复] UI 视觉审查修复：strict 模式验证码图片高度 h-10→h-11、圆角 rounded→rounded-lg；右侧装饰区域响应式宽度 md:w-1/2 lg:w-3/5；captcha 模式验证码边框 border-gray-300→border-gray-200；绑定弹窗默认 Tab passkey→totp
- [修复] PasswordLogin 中硬编码中文消息替换为错误码(ErrCaptchaRequired/ErrCaptchaInvalid/ErrUsernamePasswordEmpty/ErrPermissionDenied)，支持前端 i18n 翻译
- [修复] config.example.yaml 添加 login-mode 配置示例字段
- [修复] 移除登录页重复错误提示：request.js 中间件已统一处理 API 错误消息，前端 handleSimpleLogin/handleNext/handleLogin/handlePasskeyLogin/getCaptcha/confirmBind 等方法中的重复 ElMessage.error 调用已清除
- [测试] API 测试发现严重 Bug: PasswordLogin 第 236 行检查 AuthorityName=='账户状态'，但数据库中不存在该角色，导致所有用户 simple/captcha 模式登录返回 code:7 '用户权限不正确'
- [变更] [qa] 开始任务 #1: 测试 GET /auth/login-mode 接口 (三种配置值 + 默认值 + 无认证)
- [完成] [frontend] 完成任务 #7: index.vue strict 模式保持现有分步流程不变
- [变更] [frontend] 开始任务 #7: index.vue strict 模式保持现有分步流程不变
- [完成] [frontend] 完成任务 #6: index.vue 实现 handleSimpleLogin 方法, 调用 passwordLogin
- [变更] [frontend] 开始任务 #6: index.vue 实现 handleSimpleLogin 方法, 调用 passwordLogin
- [完成] [frontend] 完成任务 #5: index.vue 实现 captcha 模式 UI: 用户名+密码+验证码登录
- [变更] [frontend] 开始任务 #5: index.vue 实现 captcha 模式 UI: 用户名+密码+验证码登录
- [完成] [frontend] 完成任务 #4: index.vue 实现 simple 模式 UI: 用户名+密码直接登录
- [变更] [frontend] 开始任务 #4: index.vue 实现 simple 模式 UI: 用户名+密码直接登录
- [完成] [frontend] 完成任务 #3: index.vue 新增 loginMode 状态, onMounted 时调用 getLoginMode
- [完成] 后端实现完成: System.LoginMode 配置字段、GET /auth/login-mode 接口、PasswordLogin 按模式切换验证码和 Token 签发逻辑
- [完成] [backend] 完成任务 #6: 改造 PasswordLogin: strict 模式保持原有 SetPwdToken 逻辑
- [完成] [backend] 完成任务 #5: 改造 PasswordLogin: simple/captcha 模式调用 TokenNext 签发 JWT
- [完成] [backend] 完成任务 #4: 改造 PasswordLogin: 按 loginMode 条件跳过验证码
- [变更] [backend] 开始任务 #6: 改造 PasswordLogin: strict 模式保持原有 SetPwdToken 逻辑
- [变更] [backend] 开始任务 #5: 改造 PasswordLogin: simple/captcha 模式调用 TokenNext 签发 JWT
- [变更] [backend] 开始任务 #4: 改造 PasswordLogin: 按 loginMode 条件跳过验证码
- [完成] [backend] 完成任务 #3: 实现 GetLoginMode 接口 (`server/api/v1/system/sys_user.go`)
- [变更] [backend] 开始任务 #3: 实现 GetLoginMode 接口 (`server/api/v1/system/sys_user.go`)
- [完成] [backend] 完成任务 #2: authRouter 新增 GET /auth/login-mode 路由 (`server/router/system/sys_base.go`)
- [变更] [frontend] 开始任务 #3: index.vue 新增 loginMode 状态, onMounted 时调用 getLoginMode
- [变更] [backend] 开始任务 #2: authRouter 新增 GET /auth/login-mode 路由 (`server/router/system/sys_base.go`)
- [完成] [frontend] 完成任务 #2: auth.js 新增 passwordLogin API (POST /auth/password/login)
- [完成] [frontend] 完成任务 #1: auth.js 新增 getLoginMode API (GET /auth/login-mode)
- [完成] [backend] 完成任务 #1: System 结构体新增 LoginMode 字段 (`server/config/system.go`)
- [变更] [backend] 开始任务 #1: System 结构体新增 LoginMode 字段 (`server/config/system.go`)
- [变更] [frontend] 开始任务 #1: auth.js 新增 getLoginMode API (GET /auth/login-mode)
- [完成] [ui] 完成任务 #1: 登录模式配置 UI 设计稿（merge.html + design.md + Introduction.md + Resources） (完成三种登录模式 UI 设计稿、设计系统文档、前端实现指南及全部资源交付)
- [决策] Tech Lead 完成 L3 技术文档: simple/captcha 模式使用 TokenNext 签发标准 JWT; strict 模式保持 SetPwdToken 不变; GetLoginMode 使用 GET 方法; 前端通过 v-if/v-else 隔离 strict 与 simple/captcha 模式
- [变更] [ui] 开始任务 #1: 登录模式配置 UI 设计稿（merge.html + design.md + Introduction.md + Resources）
- [变更] PM评审: 1)L1 tasks.md 补充登录模式配置需求条目; 2)plan.md 范围描述从技术实现改为业务功能级; 3)plan.md 关键设计决策改为关键业务决策,移除技术细节; 4)L2 tasks.md 从技术任务改为功能级任务拆解
- [新增] 创建计划