<template>
  <div id="userLayout" class="w-full h-full relative">
    <!-- 语言切换器 -->
    <div class="absolute top-4 right-4 z-[1000]">
      <el-select v-model="language" class="w-24" size="small" @change="handleLanguageChange">
        <el-option v-for="item in LanguageList" :key="item.value" :label="item.label" :value="item.value" />
      </el-select>
    </div>

    <div class="rounded-lg flex items-center justify-evenly w-full h-full md:w-screen md:h-screen md:bg-[#0f6bf8] bg-white">
      <div class="md:w-3/5 w-10/12 h-full flex items-center justify-evenly">
        <div class="h-full w-3/5 bg-white dark:bg-slate-900 absolute" />

        <!-- 登录表单 -->
        <div class="z-[999] pt-12 pb-10 md:w-96 w-full rounded-lg flex flex-col justify-between box-border">
          <div>
            <div class="flex items-center justify-center">
              <img class="w-24" src="@/assets/logo.png" alt="">
            </div>
            <div class="mb-9">
              <p class="text-center text-4xl font-bold">
                {{ $GIN_VUE_ADMIN.appName }}
              </p>
              <p class="text-center text-sm font-normal text-gray-500 mt-2.5">
                {{ $t('login.docsinfo') }}
              </p>
            </div>

            <el-form ref="loginForm" :model="formData" :rules="rules" @keyup.enter="handleMainAction">
              <!-- 用户名输入步骤 -->
              <div v-if="step === 'USERNAME'">
                <el-form-item prop="username" class="mb-6">
                  <div class="flex w-full gap-2">
                    <el-input
                      v-model="formData.username"
                      size="large"
                      :placeholder="$t('login.username')"
                      suffix-icon="user"
                      class="flex-1"
                    />

                    <el-tooltip :content="$t('login.next')" placement="top">
                      <el-button
                        size="large"
                        :loading="loading"
                        circle
                        class="next-btn"
                        aria-label="Next"
                        @click="handleNext"
                      >
                        <el-icon class="arrow-icon"><Right /></el-icon>
                      </el-button>
                    </el-tooltip>
                  </div>
                </el-form-item>
              </div>

              <!-- 认证步骤 -->
              <div v-if="step === 'AUTH'">
                <!-- 用户名显示 + 返回按钮 -->
                <el-form-item class="mb-6">
                  <div class="flex w-full gap-2 items-center">
                    <el-button
                      size="large"
                      type="default"
                      class="px-3"
                      @click="resetToUsername"
                    >
                      <el-icon><Back /></el-icon>
                    </el-button>
                    <div class="flex-1 text-center text-lg font-medium text-gray-700">
                      {{ securityState?.display_name || formData.username }}
                    </div>
                  </div>
                </el-form-item>

                <!-- TOTP验证码输入 -->
                <el-form-item v-if="showTotpInput" prop="code" class="mb-6">
                  <el-input
                    v-model="formData.code"
                    size="large"
                    :placeholder="$t('login.enterTotpCode')"
                    maxlength="6"
                    clearable
                    @input="handleCodeInput"
                  />
                </el-form-item>

                <!-- 密码输入（未绑定安全验证时） -->
                <el-form-item v-if="showPasswordInput" prop="password" class="mb-6">
                  <el-input
                    v-model="formData.password"
                    show-password
                    size="large"
                    type="password"
                    :placeholder="$t('login.pleaseEnterPassword')"
                  />
                </el-form-item>

                <!-- 图片验证码（密码登录时显示） -->
                <el-form-item v-if="showPasswordInput" prop="captcha" class="mb-6">
                  <div class="flex w-full gap-2">
                    <el-input
                      v-model="formData.captcha"
                      size="large"
                      :placeholder="$t('login.pleaseEnterCaptcha')"
                      class="flex-1"
                      maxlength="6"
                    />
                    <div class="w-32 h-10 border border-gray-300 rounded cursor-pointer flex items-center justify-center bg-white" @click="getCaptcha">
                      <img v-if="captchaUrl" :src="captchaUrl" alt="验证码" class="w-full h-full object-cover rounded">
                      <span v-else class="text-gray-400 text-sm">点击获取</span>
                    </div>
                  </div>
                </el-form-item>

                <!-- 登录按钮 -->
                <el-form-item class="mb-6">
                  <el-button
                    v-if="!securityState.has_passkey"
                    class="shadow shadow-active h-11 w-full"
                    type="primary"
                    size="large"
                    :loading="loading"
                    @click="handleLogin"
                  >
                    {{ getLoginButtonText }}
                  </el-button>
                </el-form-item>
              </div>

              <!-- Passkey登录按钮（总是显示） -->
              <el-form-item class="mb-6">
                <el-button

                  class="h-11 w-full"
                  type="default"
                  size="large"
                  :loading="passkeyLoading"
                  @click="handlePasskeyLogin"
                >
                  <el-icon class="mr-2"><Key /></el-icon>
                  {{ $t('login.passkeyLogin') }}
                </el-button>
              </el-form-item>
            </el-form>
          </div>
        </div>
      </div>

      <div class="hidden md:block w-1/2 h-full float-right bg-[#194bfb]">
        <img class="h-full w-full object-cover" src="@/assets/login.jpg" alt="banner">
      </div>
    </div>

    <!-- 绑定弹窗 -->
    <el-dialog v-model="showBindDialog" :title="$t('login.bindRequired')" width="500px" :close-on-click-modal="false">
      <div class="text-center mb-4">
        <p class="text-lg font-medium mb-2">
          {{ $t('login.bindRequired') }}
        </p>
        <p class="text-gray-500 text-sm">
          {{ $t('login.bindDescription') }}
        </p>
      </div>

      <el-tabs v-model="activeBindTab" class="bind-tabs">
        <!-- TOTP绑定 -->
        <el-tab-pane :label="$t('login.bindTotp')" name="totp">
          <div class="totp-bind-content">
            <div v-if="totpBindInfo" class="text-center">
              <!-- 二维码显示 -->
              <div class="qr-code-container mb-4">
                <canvas ref="qrCanvas" class="mx-auto border rounded-lg" />
              </div>

              <!-- 手动输入密钥 -->
              <div class="secret-info mb-4">
                <p class="text-sm text-gray-600 mb-2">
                  {{ $t('login.manualEntry') }}
                </p>
                <div class=" p-2 rounded text-center font-mono text-sm">
                  {{ totpBindInfo.secret }}
                </div>
              </div>

              <!-- 验证码输入 -->
              <el-form :model="bindForm" @submit.prevent="confirmBind">
                <el-form-item :label="$t('login.enterTotpCode')" required>
                  <el-input
                    v-model="bindForm.code"
                    :placeholder="$t('login.totpCode')"
                    maxlength="6"
                    clearable
                  />
                </el-form-item>
              </el-form>
            </div>
          </div>
        </el-tab-pane>

        <!-- Passkey绑定 -->
        <el-tab-pane :label="$t('login.bindPasskey')" name="passkey">
          <div class="passkey-bind-content">
            <div class="text-center">
              <div class="passkey-icon mb-4">
                <el-icon size="64" color="#409EFF">
                  <Key />
                </el-icon>
              </div>
              <div class="instructions text-sm text-gray-600 mb-4">
                <p class="mb-2">
                  {{ $t('login.passkeyDescription') }}
                </p>
                <p>{{ $t('login.passkeyInstructions') }}</p>
              </div>
              <el-button
                v-if="!passkeyCreated"
                type="primary"
                size="large"
                :loading="creatingPasskey"
                @click="createPasskey"
              >
                <el-icon class="mr-2"><Key /></el-icon>
                {{ $t('login.createPasskey') }}
              </el-button>
              <div v-else class="success-info">
                <el-icon class="text-green-500 mr-2" size="20">
                  <CircleCheck />
                </el-icon>
                <span class="text-green-600">{{ $t('login.passkeyCreated') }}</span>
              </div>
            </div>
          </div>
        </el-tab-pane>
      </el-tabs>

      <template #footer>
        <div class="dialog-footer">
          <el-button @click="cancelBind">{{ $t('login.cancel') }}</el-button>
          <el-button type="primary" :loading="binding" @click="confirmBind">
            {{ $t('login.confirm') }}
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { reactive, ref, computed, onMounted, nextTick } from 'vue'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'
import { useRouterStore } from '@/pinia/modules/router'
import { useI18n } from 'vue-i18n'
import { getDict } from '@/utils/dictionary'
import { Key, CircleCheck, Back, Right } from '@element-plus/icons-vue'
import router from '@/router/index'
import QRCode from 'qrcode'

// 新的API服务
import {
  securityState as securityStateApi,
  passwordVerify,
  totpLogin,
  passkeyAssertionOptions,
  passkeyAssertionVerify,
  totpBindInit,
  totpBindVerify,
  passkeyAttestationOptions,
  passkeyAttestationVerify
} from '@/api/auth'

import { captcha } from '@/api/user'

const { t, locale } = useI18n()

// 响应式数据
const language = ref('zh-CN')
const LanguageList = ref([])
const loginForm = ref(null)
const qrCanvas = ref(null)

// 状态管理
const step = ref('USERNAME') // 'USERNAME' | 'AUTH'
const loading = ref(false)
const passkeyLoading = ref(false)
const securityState = ref(null)

// 表单数据
const formData = reactive({
  username: localStorage.getItem('lastLoginUsername') || '',
  password: '',
  code: '',
  captcha: '',
  captchaId: ''
})

// 绑定相关
const showBindDialog = ref(false)
const activeBindTab = ref('passkey')
const binding = ref(false)
const bindSession = ref('')
const totpBindInfo = ref(null)
const bindForm = reactive({
  code: ''
})

// Passkey绑定
const passkeyCreated = ref(false)
const creatingPasskey = ref(false)
const passkeyData = ref(null)

// 验证码相关
const captchaUrl = ref('')

// 计算属性
const showTotpInput = computed(() => {
  return securityState.value?.has_totp && !securityState.value?.has_passkey
})

const showPasswordInput = computed(() => {
  return !securityState.value?.has_totp && !securityState.value?.has_passkey
})

const getLoginButtonText = computed(() => {
  if (showTotpInput.value) {
    return t('login.verifyCode')
  }
  return t('login.login')
})

// 表单验证规则
const rules = computed(() => ({
  username: [
    { required: true, message: t('login.usernameRequired'), trigger: 'blur' }
  ],
  password: showPasswordInput.value ? [
    { required: true, message: t('login.passwordRequired'), trigger: 'blur' }
  ] : [],
  code: showTotpInput.value ? [
    { required: true, message: t('login.codeRequired'), trigger: 'blur' },
    { len: 6, message: t('login.codeLength'), trigger: 'blur' }
  ] : [],
  captcha: showPasswordInput.value ? [
    { required: true, message: t('login.captchaRequired'), trigger: 'blur' },
    { min: 4, message: t('login.captchaLength'), trigger: 'blur' }
  ] : []
}))

// Store
const userStore = useUserStore()
const routerStore = useRouterStore()

// 生命周期
onMounted(async() => {
  try {
    const res = await getDict('Language')
    LanguageList.value = res
  } catch (error) {
    console.error('Failed to load language list:', error)
  }
})

// 方法
const handleLanguageChange = (val) => {
  locale.value = val
}

const handleMainAction = () => {
  if (step.value === 'USERNAME') {
    handleNext()
  } else if (step.value === 'AUTH') {
    handleLogin()
  }
}

const handleNext = async() => {
  if (!formData.username) {
    ElMessage.error(t('login.enterUsername'))
    return
  }

  try {
    loading.value = true
    const res = await securityStateApi({ username: formData.username })

    if (res.code === 0) {
      securityState.value = res.data
      step.value = 'AUTH'
      localStorage.setItem('lastLoginUsername', formData.username)

      // 如果需要密码登录，自动获取验证码
      if (!res.data.has_totp && !res.data.has_passkey) {
        await getCaptcha()
      }
    } else {
      ElMessage.error(res.msg || t('login.checkUserFailed'))
    }
  } catch (error) {
    console.error('Security state check failed:', error)
    ElMessage.error(t('login.checkUserFailed'))
  } finally {
    loading.value = false
  }
}

const resetToUsername = () => {
  step.value = 'USERNAME'
  securityState.value = null
  formData.password = ''
  formData.code = ''
  formData.captcha = ''
  formData.captchaId = ''
  captchaUrl.value = ''
}

// 获取验证码
const getCaptcha = async() => {
  try {
    const res = await captcha()
    if (res.code === 0) {
      captchaUrl.value = res.data.picPath
      formData.captchaId = res.data.captchaId
    } else {
      ElMessage.error(t('login.getCaptchaFailed'))
    }
  } catch (error) {
    console.error('Get captcha failed:', error)
    ElMessage.error(t('login.getCaptchaFailed'))
  }
}

const handleLogin = async() => {
  try {
    const valid = await loginForm.value.validate()
    if (!valid) return

    loading.value = true

    let res
    if (showPasswordInput.value) {
      // 密码验证（未绑定用户）
      res = await passwordVerify({
        username: formData.username,
        password: formData.password,
        captcha: formData.captcha,
        captchaId: formData.captchaId
      })

      if (res.code === 0) {
        // 验证成功，存储绑定会话token
        bindSession.value = res.data.bind_session
        // 需要绑定
        await initBindDialog()
      } else {
        ElMessage.error(res.msg || t('login.loginFailed'))
        // 如果是验证码相关错误，重新获取验证码
        if (res.msg && (res.msg.includes('验证码') || res.msg.includes('captcha'))) {
          formData.captcha = ''
          await getCaptcha()
        }
      }
    } else if (showTotpInput.value) {
      // TOTP登录
      res = await totpLogin({
        username: formData.username,
        code: formData.code
      })

      if (res.code === 0) {
        await handleLoginSuccess(res.data)
      } else {
        ElMessage.error(res.msg || t('login.loginFailed'))
      }
    }
  } catch (error) {
    console.error('Login failed:', error)
    ElMessage.error(t('login.loginFailed'))
  } finally {
    loading.value = false
  }
}

const handlePasskeyLogin = async() => {
  try {
    passkeyLoading.value = true

    // 获取断言选项
    const optionsRes = await passkeyAssertionOptions({
      // username: formData.username || undefined
    })

    if (optionsRes.code !== 0) {
      ElMessage.error(optionsRes.msg || t('login.passkeyFailed'))
      return
    }

    const { publicKey } = optionsRes.data

    // 转换challenge
    publicKey.challenge = base64urlToArrayBuffer(publicKey.challenge)
    if (publicKey.allowCredentials && Array.isArray(publicKey.allowCredentials) && publicKey.allowCredentials.length > 0) {
      publicKey.allowCredentials.forEach(cred => {
        if (cred.id) {
          cred.id = base64urlToArrayBuffer(cred.id)
        }
      })
    } else {
      // 如果allowCredentials为空或无效，设为undefined让WebAuthn自动发现
      delete publicKey.allowCredentials
    }

    // 调用WebAuthn API
    const credential = await navigator.credentials.get({ publicKey })
    if (!credential) {
      throw new Error('Passkey authentication cancelled')
    }

    // 提交验证
    const verifyRes = await passkeyAssertionVerify({
      id: credential.id,
      rawId: arrayBufferToBase64url(credential.rawId),
      type: credential.type,
      response: {
        clientDataJSON: arrayBufferToBase64url(credential.response.clientDataJSON),
        authenticatorData: arrayBufferToBase64url(credential.response.authenticatorData),
        signature: arrayBufferToBase64url(credential.response.signature),
        userHandle: credential.response.userHandle ? arrayBufferToBase64url(credential.response.userHandle) : ''
      }
    })

    if (verifyRes.code === 0) {
      await handleLoginSuccess(verifyRes.data)
    } else {
      ElMessage.error(verifyRes.msg || t('login.passkeyFailed'))
    }
  } catch (error) {
    console.error('Passkey login failed:', error)
    if (error.name === 'NotAllowedError') {
      ElMessage.error(t('login.passkeyNotAllowed'))
    } else {
      ElMessage.error(t('login.passkeyFailed'))
    }
  } finally {
    passkeyLoading.value = false
  }
}

// import { nextTick } from 'vue'

const handleLoginSuccess = async(loginData) => {
  await userStore.setUserInfo(loginData.user)
  await userStore.setToken(loginData.token)

  ElMessage.success(t('login.loginSuccess'))

  // 1) 拉取并生成动态路由
  await routerStore.SetAsyncRouter()

  // 2) 注册路由（含 layout 与可能的 notLayout 顶层）
  const routes = routerStore.asyncRouters
  for (const r of routes) {
    router.addRoute(r)
    // 如需更稳，且 r.name === 'layout'，可把 children 再逐一 add 到 'layout'
  }

  // 3) 等一拍，确保路由树刷新
  await nextTick()

  // 4) 用“name”跳，如果没有就尝试当作 path 跳
  const target = userStore.userInfo?.authority?.defaultRouter
  const names = router.getRoutes().map(r => r.name)
  if (target && router.hasRoute(target)) {
    await router.replace({ name: target })
  } else if (target && target.startsWith('/')) {
    await router.replace({ path: target })
  } else {
    console.warn('[login] defaultRouter not found:', target, names)
    ElMessage.error(t('login.noPermission'))
  }
}

// 绑定相关方法
const initBindDialog = async() => {
  showBindDialog.value = true

  // 初始化TOTP绑定信息
  try {
    const res = await totpBindInit()
    if (res.code === 0) {
      totpBindInfo.value = res.data
      await nextTick()
      if (qrCanvas.value) {
        await QRCode.toCanvas(qrCanvas.value, res.data.otpauth_url)
      }
    }
  } catch (error) {
    console.error('TOTP bind init failed:', error)
  }
}

const createPasskey = async() => {
  try {
    creatingPasskey.value = true

    const res = await passkeyAttestationOptions({
      displayName: formData.username
    })

    if (res.code !== 0) {
      ElMessage.error(res.msg || t('login.passkeyCreateFailed'))
      return
    }

    const { publicKey } = res.data

    // 转换数据
    publicKey.challenge = base64urlToArrayBuffer(publicKey.challenge)
    publicKey.user.id = base64urlToArrayBuffer(publicKey.user.id)

    const credential = await navigator.credentials.create({ publicKey })
    if (!credential) {
      throw new Error('Passkey creation cancelled')
    }

    passkeyData.value = {
      id: credential.id,
      rawId: arrayBufferToBase64url(credential.rawId),
      type: credential.type,
      response: {
        clientDataJSON: arrayBufferToBase64url(credential.response.clientDataJSON),
        attestationObject: arrayBufferToBase64url(credential.response.attestationObject)
      }
    }

    passkeyCreated.value = true
    ElMessage.success(t('login.passkeyCreated'))
  } catch (error) {
    console.error('Passkey creation failed:', error)
    ElMessage.error(t('login.passkeyCreateFailed'))
  } finally {
    creatingPasskey.value = false
  }
}

const confirmBind = async() => {
  try {
    binding.value = true

    if (activeBindTab.value === 'totp') {
      if (!bindForm.code || bindForm.code.length !== 6) {
        ElMessage.error(t('login.codeRequired'))
        return
      }

      const res = await totpBindVerify({
        code: bindForm.code,
        secret: totpBindInfo.value?.secret,
        bind_session: bindSession.value
      })

      if (res.code === 0) {
        ElMessage.success(t('login.bindSuccess'))
        showBindDialog.value = false
        resetBindDialog()
        // 绑定成功后自动登录
        await handleNext()
      } else {
        ElMessage.error(res.msg || t('login.bindFailed'))
      }
    } else if (activeBindTab.value === 'passkey') {
      if (!passkeyCreated.value || !passkeyData.value) {
        ElMessage.error(t('login.createPasskeyFirst'))
        return
      }

      const res = await passkeyAttestationVerify({
        ...passkeyData.value,
        bind_session: bindSession.value
      })

      if (res.code === 0) {
        ElMessage.success(t('login.bindSuccess'))
        showBindDialog.value = false
        resetBindDialog()
        // 绑定成功后自动登录
        await handleNext()
      } else {
        ElMessage.error(res.msg || t('login.bindFailed'))
      }
    }
  } catch (error) {
    console.error('Bind failed:', error)
    ElMessage.error(t('login.bindFailed'))
  } finally {
    binding.value = false
  }
}

const cancelBind = () => {
  showBindDialog.value = false
  resetBindDialog()
  ElMessage.info(t('login.bindCancelled'))
}

const resetBindDialog = () => {
  activeBindTab.value = 'totp'
  bindSession.value = ''
  totpBindInfo.value = null
  bindForm.code = ''
  passkeyCreated.value = false
  passkeyData.value = null
}

const handleCodeInput = (value) => {
  // 自动提交6位验证码
  if (value.length === 6 && /^\d{6}$/.test(value)) {
    handleLogin()
  }
}

// WebAuthn工具函数
const base64urlToArrayBuffer = (base64url) => {
  const base64 = base64url.replace(/-/g, '+').replace(/_/g, '/')
  const padded = base64 + '='.repeat((4 - (base64.length % 4)) % 4)
  const binary = atob(padded)
  const bytes = new Uint8Array(binary.length)
  for (let i = 0; i < binary.length; i++) {
    bytes[i] = binary.charCodeAt(i)
  }
  return bytes.buffer
}

const arrayBufferToBase64url = (buffer) => {
  const bytes = new Uint8Array(buffer)
  let binary = ''
  for (let i = 0; i < bytes.byteLength; i++) {
    binary += String.fromCharCode(bytes[i])
  }
  return btoa(binary).replace(/\+/g, '-').replace(/\//g, '_').replace(/=+$/g, '')
}
</script>

<style scoped>
.bind-tabs :deep(.el-tabs__content) {
  min-height: 300px;
}

.qr-code-container {
  display: flex;
  justify-content: center;
  align-items: center;
}

.passkey-icon {
  color: #409EFF;
}

.success-info {
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 16px;
  font-weight: 500;
}

.instructions p {
  margin-bottom: 8px;
  line-height: 1.5;
}

.secret-info .bg-gray-100 {
  word-break: break-all;
  user-select: all;
}

.totp-bind-content,
.passkey-bind-content {
  padding: 20px 10px;
}
/* Premium Next Button Styles */
.next-btn {
  width: 44px;
  height: 44px;
  border: none;
  background: linear-gradient(135deg, #5b8cff, #2a62f2);
  color: #fff;
  box-shadow: 0 6px 16px rgba(25, 75, 251, 0.3), 0 2px 6px rgba(0, 0, 0, 0.08);
  transition: transform 0.15s ease, box-shadow 0.2s ease, filter 0.2s ease;
}

.next-btn:hover:not(.is-loading) {
  transform: translateY(-1px);
  box-shadow: 0 10px 24px rgba(25, 75, 251, 0.35), 0 4px 10px rgba(0, 0, 0, 0.1);
  filter: brightness(1.03);
}

.next-btn:active:not(.is-loading) {
  transform: translateY(0);
  box-shadow: 0 6px 16px rgba(25, 75, 251, 0.25), 0 2px 6px rgba(0, 0, 0, 0.08);
}

/* Subtle animated arrow nudge on hover */
.arrow-icon {
  transition: transform 0.2s ease;
  font-size: 18px;
}
.next-btn:hover .arrow-icon {
  transform: translateX(2px);
}

/* Focus ring for accessibility */
.next-btn:focus-visible {
  outline: 3px solid rgba(25, 75, 251, 0.35);
  outline-offset: 2px;
}

/* Loading state tweaks */
.next-btn.is-loading {
  opacity: 0.9;
  cursor: default;
}
</style>
