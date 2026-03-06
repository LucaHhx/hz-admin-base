<template>
  <div class="user-info-container">
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>{{ $t('user.personalSettings') }}</span>
        </div>
      </template>
      <el-form
        :model="form"
        label-width="120px"
      >
        <el-form-item :label="$t('user.avatar')">
          <el-upload
            class="avatar-uploader"
            action="/upload"
            :show-file-list="false"
            :on-success="handleAvatarSuccess"
            :before-upload="beforeAvatarUpload"
          >
            <template v-if="form.headerImg">
              <img
                :src="form.headerImg"
                class="avatar"
              >
            </template>
            <template v-else>
              <el-icon class="avatar-uploader-icon">
                <Plus />
              </el-icon>
            </template>
          </el-upload>
        </el-form-item>
        <el-form-item :label="$t('user.nickname')">
          <el-input v-model="form.nickName" />
        </el-form-item>
        <el-form-item :label="$t('user.phone')">
          <el-input v-model="form.phone" />
        </el-form-item>
        <el-form-item :label="$t('user.email')">
          <el-input v-model="form.email" />
        </el-form-item>
        <el-form-item :label="$t('user.language')">
          <el-select
            v-model="form.language"
            @change="handleLanguageChange"
          >
            <el-option
              label="简体中文"
              value="zh-CN"
            />
            <el-option
              label="English"
              value="en-US"
            />
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('user.googleAuth')">
          <template v-if="!form.googleAuthEnabled">
            <el-button
              type="primary"
              @click="handleBindGoogleAuth"
            >
              {{ $t('user.bindGoogleAuth') }}
            </el-button>
          </template>
          <template v-else>
            <el-tag type="success">
              {{ $t('user.bound') }}
            </el-tag>
            <el-button
              type="warning"
              class="ml-2"
              @click="handleResetGoogleAuth"
            >
              {{ $t('user.reset') }}
            </el-button>
          </template>
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            @click="submitForm"
          >
            {{ $t('common.save') }}
          </el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- 绑定谷歌验证器对话框 -->
    <el-dialog
      v-model="googleAuthDialog.visible"
      :title="$t('user.bindGoogleAuth')"
      width="500px"
      @open="handleDialogOpen"
    >
      <template v-if="googleAuthDialog.step === 1">
        <div class="qr-container">
          <img
            v-if="qrCodeUrl"
            :src="qrCodeUrl"
            alt="Google Authenticator QR Code"
          >
        </div>
        <div class="secret-key">
          {{ $t('user.secretKey') }}: {{ googleAuthDialog.secretKey }}
        </div>
        <div class="tips">
          {{ $t('user.scanQrCodeTip') }}
        </div>
        <div class="dialog-footer">
          <el-button
            type="primary"
            @click="googleAuthDialog.step = 2"
          >
            {{ $t('common.next') }}
          </el-button>
        </div>
      </template>
      <template v-else>
        <el-form
          :model="googleAuthDialog.form"
          label-width="120px"
        >
          <el-form-item :label="$t('user.verificationCode')">
            <el-input
              v-model="googleAuthDialog.form.code"
              :placeholder="$t('user.enterVerificationCode')"
            />
          </el-form-item>
        </el-form>
        <div class="dialog-footer">
          <el-button @click="googleAuthDialog.step = 1">
            {{ $t('common.previous') }}
          </el-button>
          <el-button
            type="primary"
            @click="confirmBindGoogleAuth"
          >
            {{ $t('common.confirm') }}
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { reactive, onMounted, ref, watch, onActivated } from 'vue'
import { ElMessage } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'
import { useI18n } from 'vue-i18n'
import QRCode from 'qrcode'
import { getUserInfo, setSelfInfo, getGoogleAuthQR, bindGoogleAuth, resetGoogleAuth } from '@/api/user'

const { t } = useI18n()
const qrCodeUrl = ref('')

const form = reactive({
  nickName: '',
  headerImg: '',
  phone: '',
  email: '',
  language: 'zh-CN',
  googleAuthEnabled: false
})

const googleAuthDialog = reactive({
  visible: false,
  step: 1,
  qrCode: '',
  secretKey: '',
  form: {
    code: ''
  }
})

// 监听对话框显示状态
watch(() => googleAuthDialog.visible, (newVal) => {
  if (newVal && googleAuthDialog.step === 1 && googleAuthDialog.qrCode) {
    // 使用 qrcode 生成二维码
    QRCode.toDataURL(googleAuthDialog.qrCode, {
      width: 200,
      margin: 2,
      color: {
        dark: '#000000',
        light: '#ffffff'
      }
    }).then((url) => {
      qrCodeUrl.value = url
    }).catch(() => {
      ElMessage.error(t('user.qrCodeGenerationFailed'))
    })
  }
})

// 获取用户信息
const loadUserInfo = async() => {
  try {
    const res = await getUserInfo()
    console.log('获取用户信息:', res)

    const info = res.data.userInfo
    Object.assign(form, {
      nickName: info.nickName,
      headerImg: info.headerImg,
      phone: info.phone,
      email: info.email,
      language: info.language || 'zh-CN',
      googleAuthEnabled: info.googleAuthEnabled
    })
  } catch {
    ElMessage.error(t('common.getFailed'))
  }
}

// 提交表单
const submitForm = async() => {
  try {
    await setSelfInfo({
      nickName: form.nickName,
      headerImg: form.headerImg,
      phone: form.phone,
      email: form.email,
      language: form.language
    })
    ElMessage.success(t('common.saveSuccess'))
    loadUserInfo()
  } catch {
    ElMessage.error(t('common.saveFailed'))
  }
}

// 处理对话框打开事件
const handleDialogOpen = async() => {
  if (googleAuthDialog.step === 1 && googleAuthDialog.qrCode) {
    try {
      // 生成二维码 URL
      const url = await QRCode.toDataURL(googleAuthDialog.qrCode, {
        width: 200,
        margin: 2,
        color: {
          dark: '#000000',
          light: '#ffffff'
        }
      })
      qrCodeUrl.value = url
    } catch {
      ElMessage.error(t('user.qrCodeGenerationFailed'))
    }
  }
}

// 检查谷歌验证器状态
const checkGoogleAuthStatus = async() => {
  try {
    const res = await getUserInfo()
    const info = res.data.userInfo
    console.log('检查谷歌验证器状态:', info.googleAuthEnabled)
    console.log('当前弹窗状态:', googleAuthDialog.visible)

    if (!info.googleAuthEnabled) {
      console.log('未绑定谷歌验证器，准备弹窗')
      // 重置弹窗状态
      googleAuthDialog.visible = false
      googleAuthDialog.step = 1
      googleAuthDialog.form.code = ''
      // 使用 nextTick 确保状态更新后再显示弹窗
      // await nextTick()
      handleBindGoogleAuth()
    }
  } catch (error) {
    console.error('检查谷歌验证器状态失败:', error)
    ElMessage.error(t('common.getFailed'))
  }
}

// 绑定谷歌验证器
const handleBindGoogleAuth = async() => {
  try {
    console.log('开始获取谷歌验证器二维码')
    const res = await getGoogleAuthQR()
    googleAuthDialog.qrCode = res.data.qrCode
    googleAuthDialog.secretKey = res.data.secretKey
    googleAuthDialog.visible = true
    console.log('弹窗已显示')
  } catch (error) {
    console.error('获取谷歌验证器二维码失败:', error)
    ElMessage.error(t('user.getGoogleAuthFailed'))
  }
}

// 确认绑定谷歌验证器
const confirmBindGoogleAuth = async() => {
  try {
    const result = await bindGoogleAuth({
      code: googleAuthDialog.form.code,
      secretKey: googleAuthDialog.secretKey
    })
    if (result.code === 0) {
      ElMessage.success(t('user.bindSuccess'))
      googleAuthDialog.visible = false
      loadUserInfo()
    } else {
      ElMessage.error(t('user.bindFailed'))
    }
  } catch {
    ElMessage.error(t('user.bindFailed'))
  }
}

// 重置谷歌验证器
const handleResetGoogleAuth = async() => {
  try {
    await resetGoogleAuth()
    ElMessage.success(t('user.resetSuccess'))
    loadUserInfo()
  } catch {
    ElMessage.error(t('user.resetFailed'))
  }
}

// 切换语言
const handleLanguageChange = async(value) => {
  try {
    await setSelfInfo({
      language: value
    })
    ElMessage.success(t('common.saveSuccess'))
    // 刷新页面以应用新语言
    window.location.reload()
  } catch {
    ElMessage.error(t('common.saveFailed'))
  }
}

// 头像上传相关方法
const handleAvatarSuccess = (res) => {
  form.headerImg = res.data.url
}

const beforeAvatarUpload = (file) => {
  const isImage = file.type.startsWith('image/')
  const isLt2M = file.size / 1024 / 1024 < 2

  if (!isImage) {
    ElMessage.error(t('user.imageFormatError'))
  }
  if (!isLt2M) {
    ElMessage.error(t('user.imageSizeError'))
  }
  return isImage && isLt2M
}

defineOptions({
  name: 'UserInfo'
})

onMounted(() => {
  console.log('组件已挂载')
  loadUserInfo()
  checkGoogleAuthStatus()
})

onActivated(() => {
  console.log('组件已激活')
  checkGoogleAuthStatus()
})
</script>

<style scoped>
.user-info-container {
    padding: 20px;
}
.avatar-uploader {
    border: 1px dashed #d9d9d9;
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
    width: 178px;
    height: 178px;
}
.avatar-uploader:hover {
    border-color: #409EFF;
}
.avatar-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    text-align: center;
    line-height: 178px;
}
.avatar {
    width: 178px;
    height: 178px;
    display: block;
}
.qr-container {
    text-align: center;
    margin-bottom: 20px;
}
.qrcode {
    display: inline-block;
}
.secret-key {
    text-align: center;
    margin-bottom: 10px;
    font-weight: bold;
}
.tips {
    text-align: center;
    color: #666;
    margin-bottom: 20px;
}
.ml-2 {
    margin-left: 8px;
}
.dialog-footer {
    margin-top: 20px;
    text-align: center;
}
</style>
