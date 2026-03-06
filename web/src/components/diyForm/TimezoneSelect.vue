<template>
  <div class="timezone-select">
    <el-select
      v-model="selectedTimezone"
      :placeholder="placeholder"
      :size="size"
      :disabled="disabled"
      :loading="loading"
      filterable
      clearable
      style="width: 100%"
      @change="handleChange"
      @focus="loadTimezones"
    >
      <el-option-group
        v-for="group in timezoneGroups"
        :key="group.label"
        :label="group.label"
      >
        <el-option
          v-for="timezone in group.options"
          :key="timezone.value"
          :label="timezone.label"
          :value="timezone.value"
        >
          <span class="timezone-option">
            <span class="timezone-name">{{ timezone.label }}</span>
            <span class="timezone-offset">{{ timezone.offset }}</span>
          </span>
        </el-option>
      </el-option-group>
    </el-select>
    
    <!-- 当前时间显示 -->
    <div v-if="showCurrentTime && selectedTimezone" class="current-time">
      <el-text size="small" type="info">
        {{ formatCurrentTime() }}
      </el-text>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { getTimezoneList } from '@/api/timezone'

defineOptions({
  name: 'TimezoneSelect'
})

const props = defineProps({
  // v-model 支持
  modelValue: {
    type: String,
    default: ''
  },
  // 占位符
  placeholder: {
    type: String,
    default: '请选择时区'
  },
  // 大小
  size: {
    type: String,
    default: 'default'
  },
  // 是否禁用
  disabled: {
    type: Boolean,
    default: false
  },
  // 是否显示当前时间
  showCurrentTime: {
    type: Boolean,
    default: true
  },
  // 默认时区（如果未提供modelValue）
  defaultTimezone: {
    type: String,
    default: 'Asia/Shanghai'
  }
})

const emit = defineEmits(['update:modelValue', 'change'])

// 选中的时区
const selectedTimezone = ref('')

// 时区数据（从后端获取）
const timezones = ref([])

// 加载状态
const loading = ref(false)

// 从后端加载时区数据
const loadTimezones = async () => {
  if (timezones.value.length > 0) return
  
  try {
    loading.value = true
    const res = await getTimezoneList()
    
    if (res.code === 0) {
      timezones.value = res.data || []
    } else {
      ElMessage.error(res.message || '获取时区列表失败')
      // 使用默认时区数据作为后备
      timezones.value = getDefaultTimezones()
    }
  } catch (error) {
    console.error('加载时区数据失败:', error)
    ElMessage.error('加载时区数据失败，使用默认数据')
    // 使用默认时区数据作为后备
    timezones.value = getDefaultTimezones()
  } finally {
    loading.value = false
  }
}

// 默认时区数据（后备方案）
const getDefaultTimezones = () => {
  return [
    { value: 'Asia/Shanghai', name: '中国标准时间', offset: '+08:00', region: '亚洲' },
    { value: 'Asia/Tokyo', name: '日本标准时间', offset: '+09:00', region: '亚洲' },
    { value: 'Europe/London', name: '英国时间', offset: '+00:00', region: '欧洲' },
    { value: 'America/New_York', name: '美国东部时间', offset: '-05:00', region: '北美' },
    { value: 'UTC', name: '协调世界时', offset: '+00:00', region: 'UTC' }
  ]
}

// 按地区分组的时区选项
const timezoneGroups = computed(() => {
  const groups = {}
  
  timezones.value.forEach(tz => {
    if (!groups[tz.region]) {
      groups[tz.region] = []
    }
    groups[tz.region].push({
      value: tz.value,
      label: tz.name,
      offset: tz.offset
    })
  })
  
  // 转换为数组格式
  return Object.keys(groups).map(region => ({
    label: region,
    options: groups[region].sort((a, b) => a.label.localeCompare(b.label, 'zh-CN'))
  }))
})

// 初始化
onMounted(async () => {
  selectedTimezone.value = props.modelValue || props.defaultTimezone
  await loadTimezones()
})

// 监听外部值变化
watch(
  () => props.modelValue,
  (newValue) => {
    selectedTimezone.value = newValue
  }
)

// 监听选择变化
watch(
  selectedTimezone,
  (newValue) => {
    emit('update:modelValue', newValue)
  }
)

// 处理选择变化
const handleChange = (value) => {
  emit('change', value)
}

// 格式化当前时间
const formatCurrentTime = () => {
  if (!selectedTimezone.value) return ''
  
  try {
    const now = new Date()
    const formatter = new Intl.DateTimeFormat('zh-CN', {
      timeZone: selectedTimezone.value,
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
      hour: '2-digit',
      minute: '2-digit',
      second: '2-digit',
      hour12: false
    })
    
    return `当前时间: ${formatter.format(now)}`
  } catch (error) {
    console.warn('时区格式化错误:', error)
    return ''
  }
}

// 获取时区偏移
const getTimezoneOffset = (timezone) => {
  const tz = timezones.value.find(t => t.value === timezone)
  return tz ? tz.offset : '+00:00'
}

// 暴露方法
defineExpose({
  getTimezoneOffset,
  formatCurrentTime,
  loadTimezones,
  timezones
})
</script>

<style lang="scss" scoped>
.timezone-select {
  .timezone-option {
    display: flex;
    justify-content: space-between;
    align-items: center;
    width: 100%;
    
    .timezone-name {
      flex: 1;
    }
    
    .timezone-offset {
      color: #909399;
      font-size: 12px;
      margin-left: 8px;
    }
  }
  
  .current-time {
    margin-top: 4px;
    padding: 2px 4px;
    background-color: #f5f7fa;
    border-radius: 4px;
    text-align: center;
  }
}

:deep(.el-select-dropdown__item) {
  padding: 8px 20px;
}
</style>