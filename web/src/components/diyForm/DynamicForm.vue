<template>
  <div class="dynamic-form-container">
    <el-form
      ref="formRef"
      :model="localFormData"
      :rules="formRules"
      label-position="right"
      label-width="auto"
      :inline="inline"
      :size="size"
    >
      <template v-for="option in formOptions" :key="option.name">
        <div v-if="loading" style="text-align: center; padding: 20px;">
          <el-icon class="is-loading">
            <Loading />
          </el-icon>
          <span style="margin-left: 8px;">加载配置中...</span>
        </div>
      </template>

      <template v-for="option in formOptions" v-if="!loading" :key="option.name">
        <el-form-item
          :label="option.name"
          :prop="option.name"
          :style="{ width: inline ? 'auto' : '100%' }"
        >
          <!-- String Type -->
          <template v-if="option.type === 'string'">
            <el-input
              v-model="localFormData[option.name]"
              :placeholder="`请输入${option.name}`"
              :disabled="disabled"
              clearable
              style="width: 100%"
            />
          </template>

          <!-- Text Type -->
          <template v-else-if="option.type === 'text'">
            <el-input
              v-model="localFormData[option.name]"
              :placeholder="`请输入${option.name}`"
              :disabled="disabled"
              type="textarea"
              :rows="3"
              clearable
              style="width: 100%"
            />
          </template>

          <!-- Int Type -->
          <template v-else-if="option.type === 'int'">
            <el-input-number
              v-model="localFormData[option.name]"
              :placeholder="`请输入${option.name}`"
              :disabled="disabled"
              :min="0"
              style="width: 100%"
            />
          </template>

          <!-- Float Type -->
          <template v-else-if="option.type === 'float'">
            <el-input-number
              v-model="localFormData[option.name]"
              :placeholder="`请输入${option.name}`"
              :disabled="disabled"
              :precision="2"
              :step="0.01"
              style="width: 100%"
            />
          </template>

          <!-- Bool Type -->
          <template v-else-if="option.type === 'bool'">
            <el-switch
              v-model="localFormData[option.name]"
              :disabled="disabled"
            />
          </template>

          <!-- Enum Type -->
          <template v-else-if="option.type === 'enum'">
            <el-select
              v-model="localFormData[option.name]"
              :placeholder="`请选择${option.name}`"
              :disabled="disabled"
              clearable
              style="width: 100%"
            >
              <el-option
                v-for="(optionItem, index) in option.options"
                :key="index"
                :label="optionItem"
                :value="optionItem"
              />
            </el-select>
          </template>

          <!-- Date Type -->
          <template v-else-if="option.type === 'date'">
            <el-date-picker
              v-model="localFormData[option.name]"
              type="date"
              :placeholder="`请选择${option.name}`"
              :disabled="disabled"
              style="width: 100%"
            />
          </template>

          <!-- Time Type -->
          <template v-else-if="option.type === 'time'">
            <el-time-picker
              v-model="localFormData[option.name]"
              :placeholder="`请选择${option.name}`"
              :disabled="disabled"
              style="width: 100%"
            />
          </template>

          <!-- Default fallback -->
          <template v-else>
            <el-input
              v-model="localFormData[option.name]"
              :placeholder="`请输入${option.name}`"
              :disabled="disabled"
              clearable
              style="width: 100%"
            />
          </template>
        </el-form-item>
      </template>

      <!-- Slot for additional form items -->
      <slot name="additional" :form-data="localFormData" />
    </el-form>

    <!-- Slot for form actions -->
    <div v-if="showActions" class="form-actions">
      <slot name="actions" :form-data="localFormData" :validate="validateForm" :reset="resetForm">
        <el-button @click="resetForm">重置</el-button>
        <el-button type="primary" @click="validateForm">提交</el-button>
      </slot>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { Loading } from '@element-plus/icons-vue'
import { getFormConfig } from '@/api/diyForm'
import SelectTool from '@/components/filterPop/selectTool.vue'
defineOptions({
  name: 'DynamicForm'
})

const props = defineProps({
  // 表单数据，支持 v-model
  modelValue: {
    type: Object,
    default: () => ({})
  },
  // 表单类型，从API获取配置
  formType: {
    type: String,
    required: true
  },
  // 表单大小
  size: {
    type: String,
    default: 'default'
  },
  // 是否内联显示
  inline: {
    type: Boolean,
    default: false
  },
  // 是否禁用
  disabled: {
    type: Boolean,
    default: false
  },
  // 是否显示操作按钮
  showActions: {
    type: Boolean,
    default: true
  },
  // 自定义验证规则
  customRules: {
    type: Object,
    default: () => ({})
  },
  // 是否自动加载配置
  autoLoad: {
    type: Boolean,
    default: true
  }
})

const emit = defineEmits(['update:modelValue', 'submit', 'reset', 'validate', 'config-loaded', 'config-error'])

// 表单引用
const formRef = ref()

// 表单配置选项（从API获取）
const formOptions = ref([])

// 加载状态
const loading = ref(false)

// 本地表单数据，用于双向绑定
const localFormData = ref({})

// 从API加载表单配置
const loadFormConfig = async() => {
  if (!props.formType) return

  try {
    loading.value = true
    const res = await getFormConfig(props.formType)

    if (res.code === 0) {
      formOptions.value = res.data || []
      emit('config-loaded', formOptions.value)
      initFormData()
    } else {
      ElMessage.error(res.message || '获取表单配置失败')
      emit('config-error', res.message)
    }
  } catch (error) {
    console.error('加载表单配置失败:', error)
    ElMessage.error('加载表单配置失败')
    emit('config-error', error)
  } finally {
    loading.value = false
  }
}

// 初始化表单数据
const initFormData = () => {
  const data = { ...props.modelValue }

  // 根据 formOptions 设置默认值
  formOptions.value.forEach(option => {
    if (!(option.name in data)) {
      data[option.name] = option.default_value !== undefined
        ? option.default_value
        : getDefaultValueByType(option.type)
    }
  })

  localFormData.value = data
}

// 根据类型获取默认值
const getDefaultValueByType = (type) => {
  const defaults = {
    string: '',
    text: '',
    int: 0,
    float: 0.0,
    bool: false,
    enum: '',
    date: null,
    time: null
  }
  return defaults[type] || ''
}

// 计算验证规则
const formRules = computed(() => {
  const rules = {}

  formOptions.value.forEach(option => {
    const fieldRules = []

    // 必填验证
    if (option.required) {
      fieldRules.push({
        required: true,
        message: `${option.name}不能为空`,
        trigger: ['blur', 'change']
      })
    }

    // 类型验证
    if (option.type === 'int' || option.type === 'float') {
      fieldRules.push({
        type: 'number',
        message: `${option.name}必须是数字`,
        trigger: ['blur', 'change']
      })
    }

    rules[option.name] = fieldRules
  })

  // 合并自定义规则
  return { ...rules, ...props.customRules }
})

// 监听本地数据变化，同步到父组件
watch(
  localFormData,
  (newValue) => {
    emit('update:modelValue', { ...newValue })
  },
  { deep: true }
)

// 监听外部数据变化，同步到本地
watch(
  () => props.modelValue,
  (newValue) => {
    if (JSON.stringify(newValue) !== JSON.stringify(localFormData.value)) {
      initFormData()
    }
  },
  { deep: true }
)

// 监听表单类型变化，重新加载配置
watch(
  () => props.formType,
  (newType) => {
    if (newType && props.autoLoad) {
      loadFormConfig()
    }
  },
  { immediate: true }
)

// 表单验证
const validateForm = async() => {
  if (!formRef.value) return false

  try {
    await formRef.value.validate()
    emit('validate', true, localFormData.value)
    emit('submit', localFormData.value)
    return true
  } catch (error) {
    console.error('表单验证失败:', error)
    emit('validate', false, localFormData.value)
    ElMessage.error('表单验证失败，请检查输入')
    return false
  }
}

// 重置表单
const resetForm = () => {
  if (formRef.value) {
    formRef.value.resetFields()
  }
  initFormData()
  emit('reset', localFormData.value)
}

// 清除验证
const clearValidate = () => {
  if (formRef.value) {
    formRef.value.clearValidate()
  }
}

// 手动加载配置
const reloadConfig = () => {
  return loadFormConfig()
}

// 暴露方法给父组件
defineExpose({
  validate: validateForm,
  resetFields: resetForm,
  clearValidate,
  reloadConfig,
  formData: localFormData,
  formOptions,
  loading
})
</script>

<style lang="scss" scoped>
.dynamic-form-container {
  .form-actions {
    display: flex;
    justify-content: flex-end;
    gap: 12px;
    margin-top: 20px;
    padding-top: 16px;
    border-top: 1px solid #e4e7ed;
  }

  .el-form-item {
    margin-bottom: 18px;
  }
}
</style>
