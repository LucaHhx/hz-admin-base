<template>
  <div class="timezone-date-picker">
    <el-date-picker
      v-model="localValue"
      :type="type"
      :placeholder="placeholder"
      :disabled="disabled"
      :size="size"
      :format="dateFormat"
      :value-format="valueFormat"
      style="width: 100%"
      @change="handleChange"
    />
    <!-- <div v-if="showTimezone && currentTimezone" class="timezone-info">
      <el-text size="small" type="info">
        时区: {{ currentTimezone }}
      </el-text>
    </div> -->
  </div>
</template>

<script setup>
import { ref, computed, watch, inject } from 'vue'
import { useI18n } from 'vue-i18n'

defineOptions({
  name: 'TimezoneDatePicker'
})

const props = defineProps({
  // v-model 支持
  modelValue: {
    type: [String, Date, Number],
    default: null
  },
  // 日期选择器类型
  type: {
    type: String,
    default: 'date', // date, datetime, daterange, datetimerange
    validator: (value) => {
      return ['date', 'datetime', 'daterange', 'datetimerange'].includes(value)
    }
  },
  // 占位符
  placeholder: {
    type: String,
    default: ''
  },
  // 是否禁用
  disabled: {
    type: Boolean,
    default: false
  },
  // 大小
  size: {
    type: String,
    default: 'default'
  },
  // 是否显示时区信息
  showTimezone: {
    type: Boolean,
    default: true
  },
  // 自定义时区（如果不使用全局时区）
  customTimezone: {
    type: String,
    default: ''
  },
  // 日期格式
  format: {
    type: String,
    default: ''
  },
  // 值格式
  valueFormat: {
    type: String,
    default: ''
  }
})

const emit = defineEmits(['update:modelValue', 'change'])

const { locale } = useI18n()

// 本地值
const localValue = ref(props.modelValue)

// 从i18n获取时区配置
const currentTimezone = computed(() => {
  if (props.customTimezone) {
    return props.customTimezone
  }

  // 尝试从全局配置获取时区
  try {
    const appStore = inject('appStore')
    if (appStore && appStore.timeZone) {
      console.log(appStore.timeZone)

      return appStore.timeZone
    }
  } catch (error) {
    console.warn('无法获取全局时区配置:', error)
  }

  // 默认时区
  return 'Asia/Shanghai'
})

// 日期格式
const dateFormat = computed(() => {
  if (props.format) {
    return props.format
  }

  const formatMap = {
    'date': 'YYYY-MM-DD',
    'datetime': 'YYYY-MM-DD HH:mm:ss',
    'daterange': 'YYYY-MM-DD',
    'datetimerange': 'YYYY-MM-DD HH:mm:ss'
  }

  return formatMap[props.type] || 'YYYY-MM-DD'
})

// 值格式
const valueFormat = computed(() => {
  if (props.valueFormat) {
    return props.valueFormat
  }

  const formatMap = {
    'date': 'YYYY-MM-DD',
    'datetime': 'YYYY-MM-DD HH:mm:ss',
    'daterange': 'YYYY-MM-DD',
    'datetimerange': 'YYYY-MM-DD HH:mm:ss'
  }

  return formatMap[props.type] || 'YYYY-MM-DD'
})

// 监听外部值变化
watch(
  () => props.modelValue,
  (newValue) => {
    localValue.value = newValue
  }
)

// 监听本地值变化
watch(
  localValue,
  (newValue) => {
    emit('update:modelValue', newValue)
  }
)

// 处理变化事件
const handleChange = (value) => {
  emit('change', value)
}

// 格式化显示日期（考虑时区）
const formatDateWithTimezone = (date) => {
  if (!date) return ''

  try {
    const dateObj = new Date(date)
    const formatter = new Intl.DateTimeFormat(locale.value, {
      timeZone: currentTimezone.value,
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
      hour: props.type.includes('datetime') ? '2-digit' : undefined,
      minute: props.type.includes('datetime') ? '2-digit' : undefined,
      second: props.type.includes('datetime') ? '2-digit' : undefined,
      hour12: false
    })

    return formatter.format(dateObj)
  } catch (error) {
    console.warn('日期格式化错误:', error)
    return date.toString()
  }
}

// 暴露方法
defineExpose({
  formatDateWithTimezone,
  currentTimezone
})
</script>

<style lang="scss" scoped>
.timezone-date-picker {
  .timezone-info {
    margin-top: 4px;
    padding: 2px 4px;
    background-color: #f5f7fa;
    border-radius: 4px;
    text-align: right;
  }
}
</style>
