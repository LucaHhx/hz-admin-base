<template>
  <el-dialog
    v-model="visible"
    :title="$t('user.bindGoogleAuth')"
    width="500px"
    :close-on-click-modal="false"
    :close-on-press-escape="false"
    :show-close="false"
  >
    <template v-if="step === 1">
      <div class="text-center mb-4">
        <div class="qr-container mb-4">
          <img
            v-if="qrCodeUrl"
            :src="qrCodeUrl"
            alt="Google Authenticator QR Code"
            class="mx-auto"
          >
        </div>
        <div class="secret-key mb-2">
          <span class="font-medium">{{ $t('user.secretKey') }}：</span>{{ secretKey }}
        </div>
        <div class="text-gray-500 text-sm mb-4">
          {{ $t('user.scanQrCodeTip') }}
        </div>
      </div>
      <div class="dialog-footer text-center">
        <el-button
          type="primary"
          @click="step = 2"
        >
          {{ $t('common.next') }}
        </el-button>
      </div>
    </template>
    <template v-else>
      <el-form
        :model="form"
        label-width="120px"
      >
        <el-form-item
          :label="$t('user.verificationCode')"
          prop="code"
        >
          <el-input
            v-model="form.code"
            :placeholder="$t('user.enterVerificationCode')"
            maxlength="6"
          />
        </el-form-item>
      </el-form>
      <div class="dialog-footer text-center">
        <el-button @click="step = 1">
          {{ $t('common.previous') }}
        </el-button>
        <el-button
          type="primary"
          @click="confirmBind"
        >
          {{ $t('common.confirm') }}
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>
import { ref, reactive, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { useI18n } from 'vue-i18n'
import QRCode from 'qrcode'
import { getGoogleAuthQR, bindGoogleAuth } from '@/api/user'

const { t } = useI18n()

const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['update:modelValue', 'success'])

const visible = ref(props.modelValue)
const step = ref(1)
const qrCodeUrl = ref('')
const secretKey = ref('')
const form = reactive({
  code: ''
})

watch(() => props.modelValue, (val) => {
  visible.value = val
  if (val) {
    initGoogleAuth()
  }
})

watch(() => visible.value, (val) => {
  emit('update:modelValue', val)
  if (!val) {
    resetForm()
  }
})

const resetForm = () => {
  step.value = 1
  qrCodeUrl.value = ''
  secretKey.value = ''
  form.code = ''
}

const initGoogleAuth = async() => {
  try {
    const res = await getGoogleAuthQR()
    if (res.code === 0) {
      secretKey.value = res.data.secretKey
      qrCodeUrl.value = await QRCode.toDataURL(res.data.qrCode)
    }
  } catch {
    ElMessage.error(t('user.getGoogleAuthFailed'))
  }
}

const confirmBind = async() => {
  if (!form.code) {
    ElMessage.warning(t('user.enterVerificationCode'))
    return
  }

  try {
    const res = await bindGoogleAuth({
      code: form.code,
      secretKey: secretKey.value
    })
    if (res.code === 0) {
      ElMessage.success(t('user.bindSuccess'))
      emit('success')
      visible.value = false
    }
  } catch {
    ElMessage.error(t('user.bindFailed'))
  }
}
</script>

<style scoped>
.qr-container {
  display: flex;
  justify-content: center;
  align-items: center;
}
.qr-container img {
  max-width: 200px;
}
</style>
