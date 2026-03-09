# frontend 技术方案 — 登录模式配置

> 需求: login-mode-config | 角色: frontend

## 技术栈

- Vue 3 + Element Plus (现有框架)
- Pinia 状态管理
- Axios HTTP 请求

## 架构设计

### 整体方案

前端在登录页挂载时调用 `/auth/login-mode` 获取当前模式，根据模式值切换 UI 和交互流程:

```
onMounted → getLoginMode()
       │
       ▼
loginMode 值决定 UI:
  simple  → 用户名 + 密码，直接提交 passwordLogin，登录成功
  captcha → 用户名 + 密码 + 验证码，提交 passwordLogin，登录成功
  strict  → 当前分步流程不变 (USERNAME → AUTH → 绑定/TOTP/Passkey)
```

### API 新增

**文件**: `web/src/api/auth.js`

新增两个 API:

```javascript
// 获取登录模式（公开接口，无需 token）
export const getLoginMode = () => {
  return service({
    url: '/auth/login-mode',
    method: 'get'
  })
}

// simple/captcha 模式的密码登录
export const passwordLogin = (data) => {
  return service({
    url: '/auth/password/login',
    method: 'post',
    data
  })
}
```

注意: `passwordLogin` 复用现有的 `/auth/password/login` 端点（已存在于路由中），但后端行为会根据模式改变。

### 登录页改造

**文件**: `web/src/view/login/index.vue`

#### 新增状态

```javascript
const loginMode = ref('simple')  // 'simple' | 'captcha' | 'strict'
```

#### onMounted 初始化

```javascript
onMounted(async () => {
  // 获取登录模式
  const modeRes = await getLoginMode()
  if (modeRes.code === 0) {
    loginMode.value = modeRes.data.login_mode
  }
  // ... 现有的语言列表加载 ...
})
```

#### UI 切换逻辑

| 模式 | step 流程 | UI 展示 | 提交接口 | 登录成功处理 |
|------|-----------|---------|----------|-------------|
| simple | 无 step 切换，直接显示表单 | 用户名 + 密码 | passwordLogin | handleLoginSuccess(res.data) |
| captcha | 无 step 切换，直接显示表单 | 用户名 + 密码 + 验证码 | passwordLogin | handleLoginSuccess(res.data) |
| strict | USERNAME → AUTH | 当前完整分步流程 | securityState → passwordVerify/totpLogin/passkey | 现有逻辑不变 |

#### 模板结构

```html
<!-- simple/captcha 模式: 直接显示登录表单 -->
<div v-if="loginMode !== 'strict'">
  <el-form-item prop="username">
    <el-input v-model="formData.username" placeholder="用户名" />
  </el-form-item>
  <el-form-item prop="password">
    <el-input v-model="formData.password" type="password" placeholder="密码" />
  </el-form-item>
  <!-- 仅 captcha 模式显示验证码 -->
  <el-form-item v-if="loginMode === 'captcha'" prop="captcha">
    <el-input v-model="formData.captcha" placeholder="验证码" />
    <img :src="captchaUrl" @click="getCaptcha" />
  </el-form-item>
  <el-button @click="handleSimpleLogin">登录</el-button>
</div>

<!-- strict 模式: 保持现有分步流程 -->
<div v-else>
  <!-- 现有的 USERNAME / AUTH step 代码不变 -->
</div>
```

#### handleSimpleLogin 方法

```javascript
const handleSimpleLogin = async () => {
  loading.value = true
  try {
    const payload = {
      username: formData.username,
      password: formData.password
    }
    if (loginMode.value === 'captcha') {
      payload.captcha = formData.captcha
      payload.captchaId = formData.captchaId
    }
    const res = await passwordLogin(payload)
    if (res.code === 0) {
      await handleLoginSuccess(res.data)
    } else {
      ElMessage.error(res.msg)
      if (loginMode.value === 'captcha') {
        formData.captcha = ''
        await getCaptcha()
      }
    }
  } finally {
    loading.value = false
  }
}
```

#### Passkey 按钮处理

- **simple/captcha 模式**: 隐藏 Passkey 登录按钮（这两种模式不使用 Passkey）
- **strict 模式**: 保持 Passkey 按钮显示

#### captcha 模式验证码初始化

captcha 模式在页面挂载后自动获取验证码（复用现有的 `getCaptcha` 方法和 `/base/captcha` 接口）。

## 关键决策

1. **simple/captcha 使用同一个表单**: 两种模式的区别仅在于是否显示验证码输入框，共享同一个 `handleSimpleLogin` 处理函数，通过 `loginMode` 条件判断是否传递验证码参数。

2. **strict 模式代码零修改**: 将现有的 USERNAME/AUTH 分步流程整体包裹在 `v-else` 中，不修改任何内部逻辑，避免回归。

3. **loginMode 通过 API 获取而非硬编码**: 遵循 plan.md 中的设计决策，前端不硬编码模式值，始终从后端获取。

4. **复用 handleLoginSuccess**: simple/captcha 模式的 `passwordLogin` 返回的响应格式与 `TokenNext` 一致（包含 user/token/expiresAt），可直接复用现有的 `handleLoginSuccess` 方法。

## 涉及文件清单

| 文件 | 操作 | 说明 |
|------|------|------|
| `web/src/api/auth.js` | 修改 | 新增 getLoginMode / passwordLogin |
| `web/src/view/login/index.vue` | 修改 | 按 loginMode 切换 UI 和流程 |
