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
                {{ $t('login.description') }}
              </p>
            </div>

            <el-form ref="loginForm" :model="loginFormData" :rules="rules" @keyup.enter="handleLogin">
              <el-form-item prop="username" class="mb-6">
                <el-input v-model="loginFormData.username" size="large" :placeholder="$t('login.username')" suffix-icon="user" />
              </el-form-item>
              <el-form-item prop="password" class="mb-6">
                <el-input
                  v-model="loginFormData.password"
                  show-password
                  size="large"
                  type="password"
                  :placeholder="needGoogleAuth ? $t('login.passwordOptional') : $t('login.password')"
                />
              </el-form-item>
              <el-form-item v-if="loginFormData.openCaptcha" prop="captcha" class="mb-6">
                <div class="flex w-full justify-between">
                  <el-input v-model="loginFormData.captcha" :placeholder="$t('login.captcha')" size="large" class="flex-1 mr-5" />
                  <div class="w-1/3 h-11 bg-[#c3d4f2] rounded">
                    <img v-if="picPath" class="w-full h-full" :src="picPath" :alt="$t('login.captcha')" @click="getCaptcha">
                  </div>
                </div>
              </el-form-item>
              <el-form-item v-if="needGoogleAuth" prop="googleAuthCode" class="mb-6">
                <el-input v-model="loginFormData.googleAuthCode" :placeholder="$t('login.enterGoogleAuthCode')" size="large" />
              </el-form-item>
              <el-form-item class="mb-6">
                <el-button class="shadow shadow-active h-11 w-full" type="primary" size="large" @click="handleLogin">
                  {{ $t('login.login') }}
                </el-button>
              </el-form-item>
              <el-form-item class="mb-6">
                <el-button class="h-11 w-full" type="default" size="large" @click="handlePasskeyLogin">
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

    <!-- 安全绑定弹窗 -->
    <el-dialog v-model="showSecurityBindDialog" :title="$t('login.securityBindRequired')" width="500px" :close-on-click-modal="false" :close-on-press-escape="false" :show-close="false">
      <div class="text-center mb-4">
        <p class="text-lg font-medium mb-2">
          {{ $t('login.securityBindRequired') }}
        </p>
        <p class="text-gray-500 text-sm mb-2">
          {{ $t('login.bindingRequired') }}
        </p>
        <p class="text-orange-500 text-xs">
          {{ $t('login.bindingWarning') }}
        </p>
      </div>

      <el-tabs v-model="activeBindTab" class="security-bind-tabs">
        <!-- Passkey绑定 -->
        <el-tab-pane :label="$t('login.passkeyTab')" name="passkey">
          <div class="passkey-bind">
            <div class="text-center mb-6">
              <div class="passkey-icon mb-4">
                <el-icon size="64" color="#409EFF">
                  <Key />
                </el-icon>
              </div>
              <div class="instructions text-sm text-gray-600 mb-4">
                <p class="mb-2">
                  {{ $t('login.passkeyDescription') }}
                </p>
                <p class="mb-2">
                  {{ $t('login.passkeyInstructions1') }}
                </p>
                <p>{{ $t('login.passkeyInstructions2') }}</p>
              </div>
              <el-button v-if="!passkeyCreated" type="primary" size="large" :loading="creatingPasskey" @click="createPasskey">
                <el-icon class="mr-2"><Key /></el-icon>
                {{ $t('login.createPasskey') }}
              </el-button>
              <div v-else class="success-info">
                <el-icon class="text-green-500 mr-2" size="20">
                  <CircleCheck />
                </el-icon>
                <span class="text-green-600">{{ $t('login.passkeyCreatedSuccess') }}</span>
              </div>
            </div>
          </div>
        </el-tab-pane>
        <!-- 谷歌验证器绑定 -->
        <el-tab-pane :label="$t('login.googleAuthTab')" name="googleAuth">
          <div class="google-auth-bind">
            <div class="text-center mb-6">
              <div class="qr-code-container mb-4">
                <img v-if="qrCodeUrl" :src="qrCodeUrl" alt="QR Code" class="mx-auto border rounded-lg" style="width: 200px; height: 200px;">
              </div>
              <div class="secret-info mb-4">
                <p class="text-sm text-gray-600 mb-2">
                  {{ $t('login.secretKeyManualEntry') }}
                </p>
                <div class="bg-gray-100 p-2 rounded text-center font-mono text-sm">
                  {{ bindKeyInfo?.googleAuthInfo?.secretKey }}
                </div>
              </div>
              <div class="instructions text-sm text-gray-500">
                <p>{{ $t('login.googleAuthInstructions1') }}</p>
                <p>{{ $t('login.googleAuthInstructions2') }}</p>
                <p>{{ $t('login.googleAuthInstructions3') }}</p>
              </div>
            </div>
            <el-form :model="googleAuthForm" label-width="100px">
              <el-form-item :label="$t('login.verificationCodeLabel')" required>
                <el-input v-model="googleAuthForm.code" :placeholder="$t('login.googleAuthVerificationCode')" maxlength="6" clearable />
              </el-form-item>
            </el-form>
          </div>
        </el-tab-pane>
      </el-tabs>

      <template #footer>
        <div class="dialog-footer">
          <el-button @click="cancelBind">{{ $t('login.bindDialogFooterCancel') }}</el-button>
          <el-button type="primary" :loading="binding" @click="confirmBind">{{ $t('login.bindDialogFooterConfirm') }}</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { captcha, bindGoogleAuthByLogin, bindPasskeyByLogin } from '@/api/user'
import { reactive, ref, computed, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'
import { useI18n } from 'vue-i18n'
import { getDict } from '@/utils/dictionary'
import { Key, CircleCheck } from '@element-plus/icons-vue'

const { t, locale } = useI18n()

// 语言相关
const LanguageList = ref([])
const language = ref(locale.value)

// 登录表单相关
const loginForm = ref(null)
const picPath = ref('')
const needGoogleAuth = ref(false)

const loginFormData = reactive({
  loginType: 'AuthCode',
  username: localStorage.getItem('lastLoginUsername') || 'admin',
  password: '',
  captcha: '',
  captchaId: '',
  openCaptcha: false,
  googleAuthCode: '',
  domain: window.location.hostname
})

// 安全绑定相关
const showSecurityBindDialog = ref(false)
const activeBindTab = ref('passkey')
const bindKeyInfo = ref(null)
const binding = ref(false)

// 谷歌验证器绑定
const googleAuthForm = reactive({
  code: ''
})

// Passkey绑定
const passkeyCreated = ref(false)
const creatingPasskey = ref(false)
const passkeyData = ref(null)

// 当前登录方式状态
const currentLoginMode = ref('normal') // 'normal', 'passkey', 'googleAuth'

// 表单验证规则
const rules = computed(() => ({
  username: [
    { required: true, message: t('login.usernameRequired'), trigger: 'blur' },
    { min: 5, message: t('login.usernameRule'), trigger: 'blur' }
  ],
  password: (currentLoginMode.value === 'normal' && !needGoogleAuth.value) ? [
    { required: true, message: t('login.passwordRequired'), trigger: 'blur' },
    { min: 6, message: t('login.passwordRule'), trigger: 'blur' }
  ] : [],
  // googleAuthCode: needGoogleAuth.value ? [
  //   { required: true, message: t('login.googleAuthRequired'), trigger: 'blur' },
  //   { len: 6, message: t('login.validationCodeLength'), trigger: 'blur' }
  // ] : []
}))

const userStore = useUserStore()

// QR码URL计算属性
const qrCodeUrl = computed(() => {
  if (!bindKeyInfo.value?.googleAuthInfo?.qrCode) return ''

  // 如果是otpauth链接，生成QR码
  const qrText = bindKeyInfo.value.googleAuthInfo.qrCode
  return `https://api.qrserver.com/v1/create-qr-code/?size=200x200&data=${encodeURIComponent(qrText)}`
})

// 初始化
onMounted(async() => {
  try {
    const res = await getDict('Language')
    LanguageList.value = res
  } catch (error) {
    console.error('Failed to load language list:', error)
  }
})

// 语言切换
const handleLanguageChange = (val) => {
  locale.value = val
}

// 获取验证码
const getCaptcha = async() => {
  try {
    const res = await captcha()
    if (res.code === 0) {
      picPath.value = res.data.picPath
      loginFormData.captchaId = res.data.captchaId
    }
  } catch (error) {
    console.error('Failed to get captcha:', error)
  }
}

// 处理登录
const handleLogin = async() => {
  try {
    // 设置登录模式
    if (needGoogleAuth.value) {
      currentLoginMode.value = 'googleAuth'
    } else {
      currentLoginMode.value = 'normal'
    }

    const valid = await loginForm.value.validate()
    if (!valid) return

    const res = await userStore.LoginIn(loginFormData)

    switch (res.code) {
      case 0:
        // 登录成功
        localStorage.setItem('lastLoginUsername', loginFormData.username)
        break
      case 1:
        // ElMessage.error(res.message || t('login.loginFailed'))
        resetLoginState()
        await getCaptcha()
        break
      case 2:
        // ElMessage.error(res.message || t('login.googleAuthCodeError'))
        needGoogleAuth.value = true
        currentLoginMode.value = 'googleAuth' // 更新登录模式
        // 清空验证码输入框，让用户重新输入
        loginFormData.googleAuthCode = ''
        break
      case 3:
        // 需要绑定安全验证
        // console.log('Received bind data:', res.data)
        bindKeyInfo.value = res.data
        showSecurityBindDialog.value = true
        break
      default:
        // ElMessage.error(res.message || t('login.unknownError'))
        await getCaptcha()
    }
  } catch (error) {
    console.error('Login error:', error)
    // ElMessage.error(t('login.loginError'))
    await getCaptcha()
  }
}

// 处理Passkey登录
const handlePasskeyLogin = async() => {
  try {
    // 设置登录模式为passkey
    currentLoginMode.value = 'passkey'

    // 对于Passkey登录，只需要验证用户名
    if (!loginFormData.username) {
      ElMessage.error(t('login.enterUsername'))
      return
    }

    // 设置登录类型为Passkey
    const passkeyLoginData = {
      ...loginFormData,
      loginType: 'Passkey'
    }

    // 先尝试获取Passkey挑战
    const challengeRes = await userStore.LoginIn(passkeyLoginData)

    if ((challengeRes.code === 3 || challengeRes.code === 4) && challengeRes.data?.passkeyInfo) {
      // 获取到Passkey挑战信息，开始认证流程
      console.log('Passkey challenge received:', challengeRes.data)
      const passkeyInfo = challengeRes.data.passkeyInfo

      const publicKey = {
        challenge: base64urlToArrayBuffer(passkeyInfo.challenge),
        rpId: passkeyInfo.rpId,
        allowCredentials: [],
        userVerification: 'required',
        timeout: 60000
      }

      const credential = await navigator.credentials.get({ publicKey })

      if (!credential) {
        throw new Error('Passkey认证失败')
      }

      // 将认证结果添加到登录数据
      passkeyLoginData.passkey = JSON.stringify({
        id: credential.id,
        rawId: arrayBufferToBase64url(credential.rawId),
        type: credential.type,
        response: {
          authenticatorData: arrayBufferToBase64url(credential.response.authenticatorData),
          clientDataJSON: arrayBufferToBase64url(credential.response.clientDataJSON),
          signature: arrayBufferToBase64url(credential.response.signature),
          userHandle: credential.response.userHandle ? arrayBufferToBase64url(credential.response.userHandle) : null
        }
      })

      // 重新提交登录
      const finalRes = await userStore.LoginIn(passkeyLoginData)

      if (finalRes.code === 0) {
        localStorage.setItem('lastLoginUsername', loginFormData.username)
        ElMessage.success(t('login.passkeyLoginSuccess'))
      } else {
        // ElMessage.error(finalRes.message || t('login.passkeyLoginFailed'))
      }
    } else {
      // 处理其他响应
      switch (challengeRes.code) {
        case 0:
          localStorage.setItem('lastLoginUsername', loginFormData.username)
          ElMessage.success(t('login.loginSuccess'))
          break
        case 1:
          ElMessage.error(challengeRes.message || t('login.loginFailed'))
          break
        default:
          // ElMessage.error(challengeRes.message || t('login.passkeyLoginFailed'))
      }
    }
  } catch (error) {
    console.error('Passkey login error:', error)
    let errorMsg = t('login.passkeyLoginFailed')

    if (error.name === 'NotAllowedError') {
      errorMsg = t('login.passkeyUserCancelled')
    } else if (error.name === 'NotSupportedError') {
      errorMsg = t('login.passkeyNotSupported')
    } else if (error.message) {
      // errorMsg = `${t('login.passkeyLoginFailed')}: ${error.message}`
    }

    ElMessage.error(errorMsg)
  }
}

// Base64URL 工具函数
const base64urlToArrayBuffer = (b64url) => {
  try {
    // 如果是普通字符串，先转换为Base64URL
    if (typeof b64url === 'string' && !b64url.match(/^[A-Za-z0-9_-]+$/)) {
      // 如果不是Base64URL格式，生成随机字节
      const randomBytes = new Uint8Array(32)
      crypto.getRandomValues(randomBytes)
      return randomBytes.buffer
    }

    const pad = '='.repeat((4 - (b64url.length % 4)) % 4)
    const b64 = (b64url + pad).replace(/-/g, '+').replace(/_/g, '/')
    const raw = atob(b64)
    const arr = new Uint8Array(raw.length)
    for (let i = 0; i < raw.length; i++) {
      arr[i] = raw.charCodeAt(i)
    }
    return arr.buffer
  } catch (error) {
    console.warn('Base64URL decode failed, generating random challenge:', error)
    // 如果解码失败，生成随机字节
    const randomBytes = new Uint8Array(32)
    crypto.getRandomValues(randomBytes)
    return randomBytes.buffer
  }
}

const arrayBufferToBase64url = (buf) => {
  const bytes = new Uint8Array(buf)
  let str = ''
  for (let i = 0; i < bytes.length; i++) {
    str += String.fromCharCode(bytes[i])
  }
  return btoa(str).replace(/\+/g, '-').replace(/\//g, '_').replace(/=+$/g, '')
}

// 创建Passkey
const createPasskey = async() => {
  try {
    creatingPasskey.value = true

    if (!bindKeyInfo.value?.passkeyInfo) {
      ElMessage.error(t('login.passkeyMissingConfig'))
      return
    }

    const passkeyInfo = bindKeyInfo.value.passkeyInfo

    console.log('Passkey info received:', passkeyInfo)
    console.log('Challenge string:', passkeyInfo.challenge)
    console.log('Challenge is valid base64url:', /^[A-Za-z0-9_-]+$/.test(passkeyInfo.challenge))

    // 验证必要的字段
    if (!passkeyInfo.challenge || !passkeyInfo.rpId || !passkeyInfo.user?.id) {
      ElMessage.error(t('login.passkeyMissingConfig'))
      console.error('Missing passkey info:', {
        challenge: passkeyInfo.challenge,
        rpId: passkeyInfo.rpId,
        user: passkeyInfo.user,
        hasUserId: !!passkeyInfo.user?.id
      })
      return
    }

    const publicKey = {
      challenge: base64urlToArrayBuffer(passkeyInfo.challenge),
      rp: {
        id: passkeyInfo.rpId,
        name: passkeyInfo.rpName || 'BankInf'
      },
      user: {
        id: base64urlToArrayBuffer(passkeyInfo.user.id),
        name: passkeyInfo.user.name || loginFormData.username,
        displayName: passkeyInfo.user.displayName || passkeyInfo.user.name || loginFormData.username,
      },
      pubKeyCredParams: [
        { type: 'public-key', alg: -7 }, // ES256
        { type: 'public-key', alg: -257 }, // RS256
      ],
      timeout: 60000,
      attestation: 'none',
      authenticatorSelection: {
        residentKey: 'required',
        userVerification: 'required',
        authenticatorAttachment: 'platform',
      },
    }

    const credential = await navigator.credentials.create({ publicKey })

    if (!credential) {
      throw new Error('Failed to create credential')
    }

    passkeyData.value = {
      id: credential.id,
      rawId: arrayBufferToBase64url(credential.rawId),
      type: credential.type,
      response: {
        attestationObject: arrayBufferToBase64url(credential.response.attestationObject),
        clientDataJSON: arrayBufferToBase64url(credential.response.clientDataJSON),
      },
    }

    passkeyCreated.value = true
    ElMessage.success(t('login.passkeyCreateSuccess'))
  } catch (error) {
    console.error('Passkey creation failed:', error)
    let errorMsg = t('login.passkeyCreateFailed')

    if (error.name === 'NotSupportedError') {
      errorMsg = t('login.passkeyNotSupported')
    } else if (error.name === 'NotAllowedError') {
      errorMsg = t('login.passkeyUserCancelled')
    } else if (error.name === 'AbortError') {
      errorMsg = t('login.passkeyCreateFailed')
    } else if (error.message) {
      errorMsg = `${t('login.passkeyCreateFailed')}: ${error.message}`
    }

    ElMessage.error(errorMsg)
  } finally {
    creatingPasskey.value = false
  }
}

// 确认绑定
const confirmBind = async() => {
  try {
    binding.value = true

    // 验证是否有用户名和密码
    if (!loginFormData.username || !loginFormData.password) {
      ElMessage.error(t('login.enterUsernamePassword'))
      return
    }

    if (activeBindTab.value === 'googleAuth') {
      // 绑定谷歌验证器
      if (!googleAuthForm.code || googleAuthForm.code.length !== 6) {
        ElMessage.error(t('login.googleAuthcodeRequired'))
        return
      }

      const res = await bindGoogleAuthByLogin({
        username: loginFormData.username,
        password: loginFormData.password,
        code: googleAuthForm.code,
        secretKey: bindKeyInfo.value.googleAuthInfo.secretKey
      })

      if (res.code === 0) {
        ElMessage.success(t('login.bindGoogleAuthSuccess'))
        showSecurityBindDialog.value = false
        resetBindForm()
        // 重新登录
        await handleLogin()
      } else {
        ElMessage.error(res.message || '绑定失败')
      }
    } else if (activeBindTab.value === 'passkey') {
      // 绑定Passkey
      if (!passkeyCreated.value || !passkeyData.value) {
        ElMessage.error(t('login.passkeyRequiredFirst'))
        return
      }

      const res = await bindPasskeyByLogin({
        username: loginFormData.username,
        password: loginFormData.password,
        passkeyData: JSON.stringify(passkeyData.value)
      })

      if (res.code === 0) {
        ElMessage.success(t('login.passkeyBindSuccess'))
        showSecurityBindDialog.value = false
        resetBindForm()
        // 重新登录
        await handleLogin()
      } else {
        ElMessage.error(res.message || '绑定失败')
      }
    }
  } catch (error) {
    console.error('Bind error:', error)
    ElMessage.error(t('login.bindErrorOccurred'))
  } finally {
    binding.value = false
  }
}

// 取消绑定
const cancelBind = () => {
  ElMessage.warning(t('login.securityBindSkipped'))
  showSecurityBindDialog.value = false
  resetBindForm()
}

// 重置绑定表单
const resetBindForm = () => {
  activeBindTab.value = 'googleAuth'
  bindKeyInfo.value = null
  googleAuthForm.code = ''
  passkeyCreated.value = false
  passkeyData.value = null
}

// 重置登录状态
const resetLoginState = () => {
  currentLoginMode.value = 'normal'
  needGoogleAuth.value = false
  loginFormData.googleAuthCode = ''
}
</script>

<style scoped>
.security-bind-tabs :deep(.el-tabs__content) {
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

.google-auth-bind,
.passkey-bind {
  padding: 20px 10px;
}
</style>
