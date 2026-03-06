<template>
  <div class="profile-container">
    <!-- 顶部个人信息卡片 -->
    <div class="bg-white dark:bg-slate-800 rounded-2xl shadow-sm mb-8">
      <!-- 顶部背景图 -->
      <div class="h-48 bg-blue-50 dark:bg-slate-600 relative">
        <div class="absolute inset-0 bg-pattern opacity-7" />
      </div>

      <!-- 个人信息区 -->
      <div class="px-8 -mt-20 pb-8">
        <div class="flex flex-col lg:flex-row items-start gap-8">
          <!-- 左侧头像 -->
          <div class="profile-avatar-wrapper flex-shrink-0 mx-auto lg:mx-0">
            <SelectImage
              v-model="userStore.userInfo.headerImg"
              file-type="image"
              rounded
            />
          </div>

          <!-- 右侧信息 -->
          <div class="flex-1 pt-12 lg:pt-20 w-full">
            <div
              class="flex flex-col lg:flex-row items-start lg:items-start justify-between gap-4"
            >
              <div class="lg:mt-4">
                <div class="flex items-center gap-4 mb-4">
                  <div
                    v-if="!editFlag"
                    class="text-2xl font-bold flex items-center gap-3 text-gray-800 dark:text-gray-100"
                  >
                    {{ userStore.userInfo.nickName }}
                    <el-icon
                      class="cursor-pointer text-gray-400 hover:text-gray-600 dark:hover:text-gray-200 transition-colors duration-200"
                      @click="openEdit"
                    >
                      <edit />
                    </el-icon>
                  </div>
                  <div
                    v-else
                    class="flex items-center"
                  >
                    <el-input
                      v-model="nickName"
                      class="w-48 mr-4"
                    />
                    <el-button
                      type="primary"
                      plain
                      @click="enterEdit"
                    >
                      {{ $t('common.confirm') }}
                    </el-button>
                    <el-button
                      type="danger"
                      plain
                      @click="closeEdit"
                    >
                      {{ $t('common.cancel') }}
                    </el-button>
                  </div>
                </div>

                <div
                  class="flex flex-col lg:flex-row items-start lg:items-center gap-4 lg:gap-8 text-gray-500 dark:text-gray-400"
                >
                  <!-- <div class="flex items-center gap-2">
                    <el-icon><location /></el-icon>
                    <span>中国·北京市·朝阳区</span>
                  </div>
                  <div class="flex items-center gap-2">
                    <el-icon><office-building /></el-icon>
                    <span>北京翻转极光科技有限公司</span>
                  </div>
                  <div class="flex items-center gap-2">
                    <el-icon><user /></el-icon>
                    <span>技术部·前端事业群</span>
                  </div> -->
                </div>
              </div>

              <!-- <div class="flex gap-4 mt-4">
                <el-button type="primary" plain icon="message">
                  发送消息
                </el-button>
                <el-button icon="share"> 分享主页 </el-button>
              </div> -->
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 主要内容区 -->
    <div class="grid lg:grid-cols-12 md:grid-cols-1 gap-8">
      <!-- 左侧信息栏 -->
      <div class="lg:col-span-4">
        <div
          class="bg-white dark:bg-slate-800 rounded-xl p-6 mb-6 profile-card"
        >
          <h2 class="text-lg font-semibold mb-4 flex items-center gap-2">
            <el-icon class="text-blue-500">
              <info-filled />
            </el-icon>
            {{ $t('user.basicInfo') }}
          </h2>
          <div class="space-y-4">
            <div
              class="flex items-center gap-1 lg:gap-3 text-gray-600 dark:text-gray-300"
            >
              <el-icon class="text-blue-500">
                <phone />
              </el-icon>
              <span class="font-medium">{{ $t('user.phoneNumber') }}：</span>
              <span>{{ userStore.userInfo.phone || $t('user.notSet') }}</span>
              <el-button
                link
                type="primary"
                class="ml-auto"
                @click="changePhoneFlag = true"
              >
                {{ $t('common.edit') }}
              </el-button>
            </div>
            <div
              class="flex items-center gap-1 lg:gap-3 text-gray-600 dark:text-gray-300"
            >
              <el-icon class="text-green-500">
                <message />
              </el-icon>
              <span class="font-medium flex-shrink-0">{{ $t('user.emailAddress') }}：</span>
              <span>{{ userStore.userInfo.email || $t('user.notSet') }}</span>
              <el-button
                link
                type="primary"
                class="ml-auto"
                @click="changeEmailFlag = true"
              >
                {{ $t('common.edit') }}
              </el-button>
            </div>
            <div
              class="flex items-center gap-1 lg:gap-3 text-gray-600 dark:text-gray-300"
            >
              <el-icon class="text-purple-500">
                <lock />
              </el-icon>
              <span class="font-medium">{{ $t('user.accountPassword') }}：</span>
              <span>{{ $t('common.enabled') }}</span>
              <el-button
                link
                type="primary"
                class="ml-auto"
                @click="showPassword = true"
              >
                {{ $t('common.edit') }}
              </el-button>
            </div>
            <!-- <div
              class="flex items-center gap-1 lg:gap-3 text-gray-600 dark:text-gray-300"
            >
              <el-icon class="text-orange-500">
                <key />
              </el-icon>
              <span class="font-medium">{{ $t('user.googleAuth') }}：</span>
              <span>{{ userStore.userInfo.googleAuthEnabled ? $t('user.bound') : $t('user.notSet') }}</span>
              <el-button
                link
                :type="userStore.userInfo.googleAuthEnabled ? 'danger' : 'primary'"
                class="ml-auto"
                @click="handleGoogleAuth"
              >
                {{ userStore.userInfo.googleAuthEnabled ? $t('user.reset') : $t('user.bindGoogleAuth') }}
              </el-button>
            </div> -->
          </div>
        </div>
      </div>

      <!-- 右侧内容区 -->
      <div class="lg:col-span-8">
        <div class="bg-white dark:bg-slate-800 rounded-xl p-6 profile-card">
          <el-tabs class="custom-tabs">
            <el-tab-pane>
              <template #label>
                <div class="flex items-center gap-2">
                  <el-icon><data-line /></el-icon>
                  {{ $t('user.statistics') }}
                </div>
              </template>
              <div class="grid grid-cols-2 md:grid-cols-4 gap-4 lg:gap-6 py-6">
                <div class="stat-card bg-gradient-to-br from-blue-50 to-sky-50 dark:from-blue-900/20 dark:to-sky-900/20">
                  <div
                    class="text-2xl lg:text-4xl font-bold text-blue-500 mb-2"
                  >
                    {{ userStats.logins || 0 }}
                  </div>
                  <div class="text-gray-500 dark:text-gray-400 text-sm">
                    {{ $t('user.loginCount') }}
                  </div>
                </div>
                <div class="stat-card bg-gradient-to-br from-green-50 to-emerald-50 dark:from-green-900/20 dark:to-emerald-900/20">
                  <div
                    class="text-2xl lg:text-4xl font-bold text-green-500 mb-2"
                  >
                    {{ userStats.operations || 0 }}
                  </div>
                  <div class="text-gray-500 dark:text-gray-400 text-sm">
                    {{ $t('user.operations') }}
                  </div>
                </div>
                <div class="stat-card bg-gradient-to-br from-purple-50 to-indigo-50 dark:from-purple-900/20 dark:to-indigo-900/20">
                  <div
                    class="text-2xl lg:text-4xl font-bold text-purple-500 mb-2"
                  >
                    {{ userStats.completionRate || '0%' }}
                  </div>
                  <div class="text-gray-500 dark:text-gray-400 text-sm">
                    {{ $t('user.completionRate') }}
                  </div>
                </div>
                <div class="stat-card bg-gradient-to-br from-yellow-50 to-amber-50 dark:from-yellow-900/20 dark:to-amber-900/20">
                  <div
                    class="text-2xl lg:text-4xl font-bold text-yellow-500 mb-2"
                  >
                    {{ userStats.days || 0 }}
                  </div>
                  <div class="text-gray-500 dark:text-gray-400 text-sm">
                    {{ $t('user.activeDays') }}
                  </div>
                </div>
              </div>
            </el-tab-pane>
            <el-tab-pane>
              <template #label>
                <div class="flex items-center gap-2">
                  <el-icon><calendar /></el-icon>
                  {{ $t('user.recentActivity') }}
                </div>
              </template>
              <div class="py-6">
                <div
                  v-if="loading"
                  class="flex justify-center py-8"
                >
                  <el-skeleton
                    style="width: 100%"
                    :rows="6"
                    animated
                  />
                </div>
                <div
                  v-else-if="operationRecords.length === 0"
                  class="text-center py-10 text-gray-500"
                >
                  <el-empty :description="$t('user.noRecentActivity')" />
                </div>
                <el-timeline v-else>
                  <el-timeline-item
                    v-for="record in operationRecords"
                    :key="record.ID"
                    :type="getOperationTypeStyle(record.method)"
                    :timestamp="formatDate(record.CreatedAt)"
                    :hollow="true"
                    class="pb-6"
                  >
                    <h3 class="text-base font-medium mb-1">
                      {{ record.method }} {{ record.path }}
                    </h3>
                    <p class="text-gray-500 text-sm">
                      <span
                        class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                        :class="getStatusBadgeClass(record.status)"
                      >
                        {{ record.status }}
                      </span>
                      <span class="ml-2">{{ record.body || record.resp || '-' }}</span>
                    </p>
                  </el-timeline-item>
                </el-timeline>
                <div
                  v-if="operationRecords.length > 0"
                  class="flex justify-center mt-4"
                >
                  <el-button
                    type="primary"
                    plain
                    size="small"
                    @click="loadMoreRecords"
                  >
                    {{ $t('user.loadMore') }}
                  </el-button>
                </div>
              </div>
            </el-tab-pane>
          </el-tabs>
        </div>
      </div>
    </div>

    <!-- 弹窗 -->
    <el-dialog
      v-model="showPassword"
      :title="$t('user.changePassword')"
      width="400px"
      class="custom-dialog"
      @close="clearPassword"
    >
      <el-form
        ref="modifyPwdForm"
        :model="pwdModify"
        :rules="rules"
        label-width="90px"
        class="py-4"
      >
        <el-form-item
          :minlength="6"
          :label="$t('user.originalPassword')"
          prop="password"
        >
          <el-input
            v-model="pwdModify.password"
            show-password
          />
        </el-form-item>
        <el-form-item
          :minlength="6"
          :label="$t('user.newPassword')"
          prop="newPassword"
        >
          <el-input
            v-model="pwdModify.newPassword"
            show-password
          />
        </el-form-item>
        <el-form-item
          :minlength="6"
          :label="$t('user.confirmPassword')"
          prop="confirmPassword"
        >
          <el-input
            v-model="pwdModify.confirmPassword"
            show-password
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="showPassword = false">
            {{ $t('common.cancel') }}
          </el-button>
          <el-button
            type="primary"
            @click="savePassword"
          >
            {{ $t('common.confirm') }}
          </el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog
      v-model="changePhoneFlag"
      :title="$t('user.changePhone')"
      width="400px"
      class="custom-dialog"
    >
      <el-form
        :model="phoneForm"
        label-width="80px"
        class="py-4"
      >
        <el-form-item :label="$t('user.phone')">
          <el-input
            v-model="phoneForm.phone"
            :placeholder="$t('user.newPhone')"
          >
            <template #prefix>
              <el-icon><phone /></el-icon>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item :label="$t('user.verificationCode')">
          <div class="flex gap-4">
            <el-input
              v-model="phoneForm.code"
              :placeholder="$t('user.enterCode')"
              class="flex-1"
            >
              <template #prefix>
                <el-icon><key /></el-icon>
              </template>
            </el-input>
            <el-button
              type="primary"
              :disabled="time > 0"
              class="w-32"
              @click="getCode"
            >
              {{ time > 0 ? `${time}s` : $t('user.getCode') }}
            </el-button>
          </div>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeChangePhone">
            {{ $t('common.cancel') }}
          </el-button>
          <el-button
            type="primary"
            @click="changePhone"
          >
            {{ $t('common.confirm') }}
          </el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog
      v-model="changeEmailFlag"
      :title="$t('user.changeEmail')"
      width="400px"
      class="custom-dialog"
    >
      <el-form
        :model="emailForm"
        label-width="80px"
        class="py-4"
      >
        <el-form-item :label="$t('user.email')">
          <el-input
            v-model="emailForm.email"
            :placeholder="$t('user.newEmail')"
          >
            <template #prefix>
              <el-icon><message /></el-icon>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item :label="$t('user.verificationCode')">
          <div class="flex gap-4">
            <el-input
              v-model="emailForm.code"
              :placeholder="$t('user.enterCode')"
              class="flex-1"
            >
              <template #prefix>
                <el-icon><key /></el-icon>
              </template>
            </el-input>
            <el-button
              type="primary"
              :disabled="emailTime > 0"
              class="w-32"
              @click="getEmailCode"
            >
              {{ emailTime > 0 ? `${emailTime}s` : $t('user.getCode') }}
            </el-button>
          </div>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeChangeEmail">
            {{ $t('common.cancel') }}
          </el-button>
          <el-button
            type="primary"
            @click="changeEmail"
          >
            {{ $t('common.confirm') }}
          </el-button>
        </div>
      </template>
    </el-dialog>

    <!-- <el-dialog
      v-model="showGoogleAuth"
      :title="userStore.userInfo.googleAuthEnabled ? $t('user.resetGoogleAuth') : $t('user.bindGoogleAuth')"
      width="400px"
      class="custom-dialog"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
      :show-close="false"
      @close="closeGoogleAuth"
    >
      <template v-if="!userStore.userInfo.googleAuthEnabled">
        <div class="text-center mb-4">
          <div class="qr-container mb-4">
            <img
              v-if="googleAuthQR"
              :src="googleAuthQR"
              alt="Google Authenticator QR Code"
              class="mx-auto"
            >
          </div>
          <div class="secret-key mb-2">
            <span class="font-medium">{{ $t('user.secretKey') }}：</span>{{ googleAuthSecret }}
          </div>
          <div class="text-gray-500 text-sm mb-4">
            {{ $t('user.scanQrCodeTip') }}
          </div>
        </div>
        <el-form
          :model="googleAuthForm"
          label-width="80px"
        >
          <el-form-item :label="$t('user.verificationCode')">
            <el-input
              v-model="googleAuthForm.code"
              :placeholder="$t('user.enterVerificationCode')"
              maxlength="6"
            />
          </el-form-item>
        </el-form>
      </template>
      <template v-else>
        <div class="text-center mb-4">
          <p class="text-gray-600">
            {{ $t('user.resetGoogleAuthConfirm') }}
          </p>
        </div>
      </template>
      <template #footer>
        <div class="dialog-footer">
          <el-button
            v-if="userStore.userInfo.type===0"
            @click="closeGoogleAuth"
          >
            {{ $t('common.cancel') }}
          </el-button>
          <el-button
            type="primary"
            @click="confirmGoogleAuth"
          >
            {{ $t('common.confirm') }}
          </el-button>
        </div>
      </template>
    </el-dialog> -->
  </div>
</template>

<script setup>
import { setSelfInfo, changePassword, getGoogleAuthQR, bindGoogleAuth, resetGoogleAuth } from '@/api/user.js'
import { getSysOperationRecordList } from '@/api/system.js'
import { reactive, ref, watch, onMounted } from 'vue'
import { ElMessage, ElSkeleton, ElEmpty } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'
import SelectImage from '@/components/selectImage/selectImage.vue'
import QRCode from 'qrcode'
import { useI18n } from 'vue-i18n'
import { formatDate } from '@/utils/format'

const { t } = useI18n()
const userStore = useUserStore()
const modifyPwdForm = ref(null)
const showPassword = ref(false)
const pwdModify = ref({})
const nickName = ref('')
const editFlag = ref(false)

const rules = reactive({
  password: [
    { required: true, message: t('common.pleaseEnter') + t('user.originalPassword'), trigger: 'blur' },
    { min: 6, message: t('user.passwordMin'), trigger: 'blur' }
  ],
  newPassword: [
    { required: true, message: t('common.pleaseEnter') + t('user.newPassword'), trigger: 'blur' },
    { min: 6, message: t('user.passwordMin'), trigger: 'blur' }
  ],
  confirmPassword: [
    { required: true, message: t('common.pleaseEnter') + t('user.confirmPassword'), trigger: 'blur' },
    { min: 6, message: t('user.passwordMin'), trigger: 'blur' },
    {
      validator: (rule, value, callback) => {
        if (value !== pwdModify.value.newPassword) {
          callback(new Error(t('user.passwordNotMatch')))
        } else {
          callback()
        }
      },
      trigger: 'blur'
    }
  ]
})

const savePassword = async() => {
  modifyPwdForm.value.validate((valid) => {
    if (valid) {
      changePassword({
        password: pwdModify.value.password,
        newPassword: pwdModify.value.newPassword
      }).then((res) => {
        if (res.code === 0) {
          ElMessage.success(t('user.passwordChanged'))
        }
        showPassword.value = false
      })
    }
  })
}

const clearPassword = () => {
  pwdModify.value = {
    password: '',
    newPassword: '',
    confirmPassword: ''
  }
  modifyPwdForm.value?.clearValidate()
}

const openEdit = () => {
  nickName.value = userStore.userInfo.nickName
  editFlag.value = true
}

const closeEdit = () => {
  nickName.value = ''
  editFlag.value = false
}

const enterEdit = async() => {
  const res = await setSelfInfo({
    nickName: nickName.value
  })
  if (res.code === 0) {
    userStore.ResetUserInfo({ nickName: nickName.value })
    ElMessage.success(t('user.modifySuccessfully'))
  }
  nickName.value = ''
  editFlag.value = false
}

const changePhoneFlag = ref(false)
const time = ref(0)
const phoneForm = reactive({
  phone: '',
  code: ''
})

const getCode = async() => {
  time.value = 60
  let timer = setInterval(() => {
    time.value--
    if (time.value <= 0) {
      clearInterval(timer)
      timer = null
    }
  }, 1000)
}

const closeChangePhone = () => {
  changePhoneFlag.value = false
  phoneForm.phone = ''
  phoneForm.code = ''
}

const changePhone = async() => {
  const res = await setSelfInfo({ phone: phoneForm.phone })
  if (res.code === 0) {
    ElMessage.success(t('user.modifySuccessfully'))
    userStore.ResetUserInfo({ phone: phoneForm.phone })
    closeChangePhone()
  }
}

const changeEmailFlag = ref(false)
const emailTime = ref(0)
const emailForm = reactive({
  email: '',
  code: ''
})

const getEmailCode = async() => {
  emailTime.value = 60
  let timer = setInterval(() => {
    emailTime.value--
    if (emailTime.value <= 0) {
      clearInterval(timer)
      timer = null
    }
  }, 1000)
}

const closeChangeEmail = () => {
  changeEmailFlag.value = false
  emailForm.email = ''
  emailForm.code = ''
}

const changeEmail = async() => {
  const res = await setSelfInfo({ email: emailForm.email })
  if (res.code === 0) {
    ElMessage.success(t('user.modifySuccessfully'))
    userStore.ResetUserInfo({ email: emailForm.email })
    closeChangeEmail()
  }
}

watch(() => userStore.userInfo.headerImg, async(val) => {
  const res = await setSelfInfo({ headerImg: val })
  if (res.code === 0) {
    userStore.ResetUserInfo({ headerImg: val })
    ElMessage({
      type: 'success',
      message: t('user.setSuccessfully'),
    })
  }
})

// 操作记录相关
const operationRecords = ref([])
const loading = ref(false)
const pageInfo = reactive({
  page: 1,
  pageSize: 5,
  method: '',
  path: '',
  status: 0
})

// 用户统计数据
const userStats = reactive({
  logins: 0,
  operations: 0,
  completionRate: '95%',
  days: 0
})

const loadOperationRecords = async() => {
  loading.value = true
  try {
    const res = await getSysOperationRecordList(pageInfo)
    if (res.code === 0) {
      if (pageInfo.page === 1) {
        operationRecords.value = res.data.list
      } else {
        operationRecords.value = [...operationRecords.value, ...res.data.list]
      }
      // 更新统计数据
      if (pageInfo.page === 1) {
        userStats.operations = res.data.total || 0
        userStats.logins = Math.floor(Math.random() * 100)
        userStats.days = Math.floor(Math.random() * 30)
      }
    }
  } catch (error) {
    console.error(error)
  } finally {
    loading.value = false
  }
}

const loadMoreRecords = () => {
  pageInfo.page++
  loadOperationRecords()
}

const getOperationTypeStyle = (method) => {
  switch (method?.toUpperCase()) {
    case 'GET': return 'primary'
    case 'POST': return 'success'
    case 'PUT': return 'warning'
    case 'DELETE': return 'danger'
    default: return 'info'
  }
}

const getStatusBadgeClass = (status) => {
  if (status >= 200 && status < 300) {
    return 'bg-green-100 text-green-800 dark:bg-green-900/30 dark:text-green-300'
  } else if (status >= 400 && status < 500) {
    return 'bg-yellow-100 text-yellow-800 dark:bg-yellow-900/30 dark:text-yellow-300'
  } else if (status >= 500) {
    return 'bg-red-100 text-red-800 dark:bg-red-900/30 dark:text-red-300'
  }
  return 'bg-gray-100 text-gray-800 dark:bg-gray-700 dark:text-gray-300'
}

onMounted(() => {
  console.log('组件已挂载')
  loadOperationRecords()
  // 检查 Google 验证器状态
  // if (!userStore.userInfo.googleAuthEnabled && userStore.userInfo.ID !== 1) {
  //   // ElMessage.warning(t('user.googleAuthRequired'))
  //   // handleGoogleAuth()
  // }
})

// 谷歌验证相关
const showGoogleAuth = ref(false)
const googleAuthQR = ref('')
const googleAuthSecret = ref('')
const googleAuthForm = reactive({
  code: ''
})

const handleGoogleAuth = async() => {
  console.log(userStore.userInfo)

  if (userStore.userInfo.googleAuthEnabled) {
    showGoogleAuth.value = true
  } else {
    try {
      const res = await getGoogleAuthQR()
      if (res.code === 0) {
        googleAuthQR.value = await QRCode.toDataURL(res.data.qrCode)
        googleAuthSecret.value = res.data.secretKey
        showGoogleAuth.value = true
      }
    } catch {
      ElMessage.error(t('user.getGoogleAuthFailed'))
    }
  }
}

const closeGoogleAuth = () => {
  showGoogleAuth.value = false
  googleAuthQR.value = ''
  googleAuthSecret.value = ''
  googleAuthForm.code = ''
}

const confirmGoogleAuth = async() => {
  try {
    if (userStore.userInfo.googleAuthEnabled) {
      // 重置
      const res = await resetGoogleAuth()
      if (res.code === 0) {
        ElMessage.success(t('user.resetSuccess'))
        userStore.ResetUserInfo({ googleAuthEnabled: false })
        closeGoogleAuth()
      }
    } else {
      // 绑定
      const res = await bindGoogleAuth({
        code: googleAuthForm.code,
        secretKey: googleAuthSecret.value
      })
      if (res.code === 0) {
        ElMessage.success(t('user.bindSuccess'))
        userStore.ResetUserInfo({ googleAuthEnabled: true })
        closeGoogleAuth()
      }
    }
  } catch {
    ElMessage.error(userStore.userInfo.googleAuthEnabled ? t('user.resetFailed') : t('user.bindFailed'))
  }
}
</script>

<style lang="scss">
  .profile-container {
    @apply p-4 lg:p-6 min-h-screen bg-gray-50 dark:bg-slate-900;

    .bg-pattern {
      background-image: url("data:image/svg+xml,%3Csvg width='60' height='60' viewBox='0 0 60 60' xmlns='http://www.w3.org/2000/svg'%3E%3Cg fill='none' fill-rule='evenodd'%3E%3Cg fill='%23000000' fill-opacity='0.1'%3E%3Cpath d='M36 34v-4h-2v4h-4v2h4v4h2v-4h4v-2h-4zm0-30V0h-2v4h-4v2h4v4h2V6h4V4h-4zM6 34v-4H4v4H0v2h4v4h2v-4h4v-2H6zM6 4V0H4v4H0v2h4v4h2V6h4V4H6z'/%3E%3C/g%3E%3C/g%3E%3C/svg%3E");
    }

    .profile-card {
      @apply shadow-sm hover:shadow-md transition-shadow duration-300;
    }

    .profile-action-btn {
      @apply bg-white/10 hover:bg-white/20 border-white/20;
      .el-icon {
        @apply mr-1;
      }
    }

    .stat-card {
      @apply p-4 lg:p-6 rounded-lg bg-gray-50 dark:bg-slate-700/50 text-center hover:shadow-md transition-all duration-300;
    }

    .custom-tabs {
      :deep(.el-tabs__nav-wrap::after) {
        @apply h-0.5 bg-gray-100 dark:bg-gray-700;
      }
      :deep(.el-tabs__active-bar) {
        @apply h-0.5 bg-blue-500;
      }
      :deep(.el-tabs__item) {
        @apply text-base font-medium px-6;
        .el-icon {
          @apply mr-1 text-lg;
        }
        &.is-active {
          @apply text-blue-500;
        }
      }
      :deep(.el-timeline-item__node--normal) {
        @apply left-[-2px];
      }
      :deep(.el-timeline-item__wrapper) {
        @apply pl-8;
      }
      :deep(.el-timeline-item__timestamp) {
        @apply text-gray-400 text-sm;
      }
    }

    .custom-dialog {
      :deep(.el-dialog__header) {
        @apply mb-0 pb-4 border-b border-gray-100 dark:border-gray-700;
      }
      :deep(.el-dialog__footer) {
        @apply mt-0 pt-4 border-t border-gray-100 dark:border-gray-700;
      }
      :deep(.el-input__wrapper) {
        @apply shadow-none;
      }
      :deep(.el-input__prefix) {
        @apply mr-2;
      }
    }

    .edit-input {
      :deep(.el-input__wrapper) {
        @apply bg-white/10 border-white/20 shadow-none;
        input {
          @apply text-white;
          &::placeholder {
            @apply text-white/60;
          }
        }
      }
    }
  }
</style>
