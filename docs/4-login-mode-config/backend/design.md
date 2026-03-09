# backend 技术方案 — 登录模式配置

> 需求: login-mode-config | 角色: backend

## 技术栈

- Go + Gin (现有框架)
- Viper 配置管理 (已集成 `config.local.yaml`)
- golang-jwt/jwt v5 (通过 `TokenNext` 签发标准 JWT)

## 架构设计

### 整体方案

通过在 `System` 配置结构体中新增 `LoginMode` 字段，让后端根据配置值决定 `PasswordLogin` 的验证逻辑。前端通过新增的公开接口 `/auth/login-mode` 获取当前模式，自行适配 UI。

```
config.local.yaml (login-mode)
       │
       ▼
System.LoginMode ──► GetLoginMode API (公开,无需token)
       │
       ▼
PasswordLogin ──► 按模式切换:
  simple  → 跳过验证码, 直接 TokenNext 签发 JWT
  captcha → 要求验证码, 然后 TokenNext 签发 JWT
  strict  → 保持当前行为 (验证码 + pwd.SetPwdToken)
```

### 配置字段

**文件**: `server/config/system.go`

在 `System` 结构体末尾新增:

```go
LoginMode string `mapstructure:"login-mode" json:"login-mode" yaml:"login-mode"` // 登录模式: simple|captcha|strict
```

**默认值处理**: 若 `LoginMode` 为空字符串，视为 `simple`。

### 新增接口: GetLoginMode

**文件**: `server/api/v1/system/sys_user.go`

```go
func (b *BaseApi) GetLoginMode(c *gin.Context) {
    mode := global.HAB_CONFIG.System.LoginMode
    if mode == "" {
        mode = "simple"
    }
    response.OkWithDetailed(gin.H{"login_mode": mode}, "Success", c)
}
```

**路由注册**: `server/router/system/sys_base.go`

在 `authRouter` 组中新增:
```go
authRouter.GET("login-mode", baseApi.GetLoginMode)
```

注意: 使用 GET 方法，因为这是一个无副作用的查询接口，且无需认证。

### 修改 PasswordLogin 逻辑

**文件**: `server/api/v1/system/sys_user.go` — `PasswordLogin` 方法

当前实现分析:
- 第 189-198 行: 硬编码要求验证码
- 第 226 行: 使用 `pwd.SetPwdToken(user.ID)` 生成 token (基于 Redis 的临时 token，非标准 JWT)
- 第 228-232 行: 手动构造 `LoginResponse`

**改造逻辑**:

```go
func (b *BaseApi) PasswordLogin(c *gin.Context) {
    // ... 参数绑定和用户名密码非空校验 (保持不变) ...

    // 获取登录模式
    loginMode := global.HAB_CONFIG.System.LoginMode
    if loginMode == "" {
        loginMode = "simple"
    }

    // 验证码校验 — 仅 captcha 和 strict 模式需要
    if loginMode != "simple" {
        if req.Captcha == "" || req.CaptchaId == "" {
            response.FailWithMessage("请输入验证码", c)
            return
        }
        if !store.Verify(req.CaptchaId, req.Captcha, true) {
            response.FailWithMessage("验证码错误", c)
            return
        }
    }

    // 用户名密码验证 (保持不变)
    u := &system.SysUser{Username: req.Username, Password: req.Password}
    user, err := userService.Login(u)
    // ... 错误处理、Enable 检查、权限检查 (保持不变) ...

    // Token 签发 — strict 模式保持原行为，其他模式用 TokenNext
    if loginMode == "strict" {
        token := pwd.SetPwdToken(user.ID)
        response.OkWithDetailed(systemRes.LoginResponse{
            User:      *user,
            Token:     token,
            ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
        }, "登录成功", c)
    } else {
        b.TokenNext(c, *user)
    }
}
```

## 关键决策

1. **simple/captcha 使用 `TokenNext` 而非 `SetPwdToken`**: `SetPwdToken` 生成的是基于 Redis 的临时 token，仅用于密码验证后的绑定流程中间态。simple/captcha 模式是"一步到位"的登录，应签发标准 JWT（通过 `TokenNext`），与 TOTP/Passkey 登录成功后的行为一致。

2. **GetLoginMode 使用 GET 方法**: 这是一个幂等的状态查询，不需要请求体，GET 更符合 REST 语义。

3. **strict 模式保持原行为不变**: 不修改 strict 分支的任何逻辑，避免引入回归。strict 模式的完整流程是: PasswordLogin(验证码+密码) → 返回 pwdToken → 前端进入绑定/TOTP/Passkey 流程。

4. **配置默认值 `simple`**: 与 plan.md 一致。在代码中双重兜底 — yaml 默认值 + 代码判空兜底。

## 涉及文件清单

| 文件 | 操作 | 说明 |
|------|------|------|
| `server/config/system.go` | 修改 | 新增 LoginMode 字段 |
| `server/config.local.yaml` | 已完成 | 已有 `login-mode: simple` |
| `server/router/system/sys_base.go` | 修改 | authRouter 新增 GET login-mode 路由 |
| `server/api/v1/system/sys_user.go` | 修改 | 新增 GetLoginMode 方法; 改造 PasswordLogin |
