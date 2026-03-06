<template>
  <el-dialog
    v-model="visible"
    :title="$t('common.' + formType)"
    width="50%"
    :close-on-click-modal="false"
    :close-on-press-escape="false"
    @close="handleClose"
  >
    <div class="gva-form-box">
      <el-form :model="form" ref="formRef" label-position="right" :rules="rules" label-width="auto">
        <el-form-item :label="$t('business.order.columns.orderNo') + ':'" prop="orderNo">
          <el-input v-model="form.orderNo" :clearable="false" placeholder="please enter orderNo" :disabled="isView" />
        </el-form-item>
        <el-form-item :label="$t('business.order.columns.customerName') + ':'" prop="customerName">
          <el-input v-model="form.customerName" :clearable="false" placeholder="please enter customerName" :disabled="isView" />
        </el-form-item>
        <el-form-item :label="$t('business.order.columns.amount') + ':'" prop="amount">
          <el-input-number v-model="form.amount" :precision="2" :disabled="isView" :clearable="false" />
        </el-form-item>
        <el-form-item :label="$t('business.order.columns.status') + ':'" prop="status">
          <el-input-number v-model="form.status" :disabled="isView" :clearable="false" />
        </el-form-item>
        <el-form-item :label="$t('business.order.columns.remark') + ':'" prop="remark">
          <el-input v-model="form.remark" :clearable="false" placeholder="please enter remark" :disabled="isView" />
        </el-form-item>
      </el-form>
    </div>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="handleClose">{{ $t('common.cancel') }}</el-button>
        <el-button v-if="!isView" type="primary" :loading="loading" @click="handleSubmit">{{ $t('common.confirm') }}</el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>

import { createOrder, updateOrder, findOrder } from '@/api/tests/order'
import { ElMessage } from 'element-plus'
import { ref, computed } from 'vue'
import { convertDate, convertDateTime } from '@/utils/time'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()

const props = defineProps({
  callback: {
    type: Function,
    default: null
  }
})

defineOptions({
  name: 'OrderForm'
})

// 计算属性
const isView = computed(() => formType.value === 'view')
// 状态
const visible = ref(false)
// 加载状态
const loading = ref(false)
// 表单引用
const formRef = ref()
// 表单数据
const form = ref(getInitialForm())
// 表单类型
const formType = ref('create')

// 表单校验规则
const rules = {
  orderNo: [{ required: true, message: '', trigger: 'blur' }],
  customerName: [{ required: true, message: '', trigger: 'blur' }],
  amount: [{ required: true, message: '', trigger: 'blur' }],
}

// 获取初始表单数据
function getInitialForm() {
  return {
    orderNo: '',
    customerName: '',
    amount: 0,
    status: 0,
    remark: '',
  }
}

// 打开弹窗
async function openDialog(type, data) {
  visible.value = true
  formType.value = type
  if (formType.value === 'create') {
    form.value = getInitialForm()
  } else if (formType.value === 'copy') {
    const copyData = { ...data }
    delete copyData.ID
    delete copyData.createdAt
    delete copyData.updatedAt
    delete copyData.deletedAt
    form.value = copyData
  } else {
    try {
      const res = await findOrder({ ID: data.id })
      if (res.code === 0) {
        form.value = res.data
      } else {
        ElMessage.error('Failed to obtain data')
      }
    } catch (error) {
      console.error('Failed to retrieve data:', error)
      ElMessage.error('Failed to obtain data')
    }
  }
}

const itemFormat = (item) => {
  return item
}

// 提交表单
async function handleSubmit() {
  if (!formRef.value) return
  try {
    await formRef.value.validate()
    loading.value = true
    let api = null
    switch (formType.value) {
      case 'create':
        api = createOrder
        break
      case 'edit':
        api = updateOrder
        break
      case 'view':
        api = null
        break
      case 'copy':
        api = createOrder
        break
    }

    if (api) {
      const res = await api(itemFormat(form.value))
      if (res.code === 0) {
        ElMessage.success('operation is successful')
        props.onSuccess?.(form.value)
        handleClose()
      } else {
        ElMessage.error(res.message || 'operation failed')
      }
    }
  } catch (error) {
    console.error('Submission failed:', error)
    ElMessage.error('Submission failed')
  } finally {
    if (props.callback) {
      props.callback()
    }
    loading.value = false
  }
}

// 关闭弹窗
function handleClose() {
  visible.value = false
  form.value = getInitialForm()
  formRef.value?.resetFields()
}

// 暴露方法
defineExpose({
  openDialog
})
</script>

<style lang="scss" scoped>

</style>
