<template>
  <div class="timezone-form-example">
    <h2>时区感知表单示例</h2>
    
    <el-card class="example-card">
      <template #header>
        <span>使用时区感知日期选择器的动态表单</span>
      </template>
      
      <DynamicFormSimple
        v-model="formData"
        :form-options="formOptions"
        struct-name="timezone-form"
        @submit="handleSubmit"
      >
        <template #additional>
          <el-form-item label="事件开始时间">
            <TimezoneDatePicker
              v-model="formData.eventStart"
              type="datetime"
              placeholder="选择事件开始时间"
              :show-timezone="true"
            />
          </el-form-item>
          
          <el-form-item label="事件结束时间">
            <TimezoneDatePicker
              v-model="formData.eventEnd"
              type="datetime"
              placeholder="选择事件结束时间"
              :show-timezone="true"
            />
          </el-form-item>
          
          <el-form-item label="截止日期（仅日期）">
            <TimezoneDatePicker
              v-model="formData.deadline"
              type="date"
              placeholder="选择截止日期"
              :show-timezone="false"
            />
          </el-form-item>
        </template>
      </DynamicFormSimple>
    </el-card>

    <el-card class="example-card">
      <template #header>
        <span>表单数据预览</span>
      </template>
      
      <el-descriptions :column="2" border>
        <el-descriptions-item label="任务名称">
          {{ formData.taskName || '未填写' }}
        </el-descriptions-item>
        <el-descriptions-item label="优先级">
          {{ formData.priority || '未选择' }}
        </el-descriptions-item>
        <el-descriptions-item label="是否紧急">
          {{ formData.isUrgent ? '是' : '否' }}
        </el-descriptions-item>
        <el-descriptions-item label="创建日期">
          {{ formatDate(formData.createDate) || '未选择' }}
        </el-descriptions-item>
        <el-descriptions-item label="截止日期">
          {{ formatDate(formData.deadline) || '未选择' }}
        </el-descriptions-item>
        <el-descriptions-item label="事件开始时间">
          {{ formatDateTime(formData.eventStart) || '未选择' }}
        </el-descriptions-item>
        <el-descriptions-item label="事件结束时间" span="2">
          {{ formatDateTime(formData.eventEnd) || '未选择' }}
        </el-descriptions-item>
      </el-descriptions>
    </el-card>

    <el-card class="example-card">
      <template #header>
        <span>当前时区信息</span>
      </template>
      
      <el-row :gutter="20">
        <el-col :span="8">
          <el-statistic title="当前时区" :value="currentTimezone" />
        </el-col>
        <el-col :span="8">
          <el-statistic title="本地时间" :value="currentLocalTime" />
        </el-col>
        <el-col :span="8">
          <el-statistic title="UTC时间" :value="currentUTCTime" />
        </el-col>
      </el-row>
    </el-card>

    <el-card class="example-card">
      <template #header>
        <span>表单配置说明</span>
      </template>
      
      <div class="config-info">
        <h4>配置特点：</h4>
        <ul>
          <li><strong>时区感知</strong>：所有日期时间字段都会使用 i18n 中配置的时区</li>
          <li><strong>自动格式化</strong>：根据字段类型自动选择合适的日期格式</li>
          <li><strong>多语言支持</strong>：日期格式会根据当前语言环境调整</li>
          <li><strong>兼容性</strong>：与原有的日期字段完全兼容，只需替换组件即可</li>
        </ul>
        
        <h4>支持的日期类型：</h4>
        <ul>
          <li><code>date</code>：仅日期选择</li>
          <li><code>datetime</code>：日期时间选择</li>
          <li><code>uintDate</code>：整数日期格式</li>
        </ul>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted, onUnmounted } from 'vue'
import { ElMessage } from 'element-plus'
import { useI18n } from 'vue-i18n'
import DynamicFormSimple from './DynamicFormSimple.vue'
import TimezoneDatePicker from './TimezoneDatePicker.vue'

defineOptions({
  name: 'TimezoneFormExample'
})

const { locale } = useI18n()

// 表单数据
const formData = reactive({
  taskName: '项目任务示例',
  priority: 'high',
  isUrgent: true,
  createDate: new Date(),
  deadline: null,
  eventStart: null,
  eventEnd: null
})

// 表单配置
const formOptions = [
  {
    name: 'taskName',
    type: 'string',
    required: true,
    default_value: ''
  },
  {
    name: 'priority',
    type: 'enum',
    required: true,
    options: ['low', 'medium', 'high'],
    default_value: 'medium'
  },
  {
    name: 'isUrgent',
    type: 'bool',
    required: false,
    default_value: false
  },
  {
    name: 'createDate',
    type: 'date',
    required: true,
    default_value: new Date()
  }
]

// 当前时区
const currentTimezone = computed(() => {
  // 尝试获取i18n中的时区设置
  return 'Asia/Shanghai' // 这里应该从i18n配置中获取
})

// 当前本地时间
const currentLocalTime = ref('')
// 当前UTC时间  
const currentUTCTime = ref('')

// 定时器
let timeTimer = null

// 更新时间显示
const updateTime = () => {
  const now = new Date()
  
  // 格式化本地时间
  try {
    const localFormatter = new Intl.DateTimeFormat(locale.value, {
      timeZone: currentTimezone.value,
      hour: '2-digit',
      minute: '2-digit',
      second: '2-digit',
      hour12: false
    })
    currentLocalTime.value = localFormatter.format(now)
  } catch (error) {
    currentLocalTime.value = now.toLocaleTimeString()
  }
  
  // 格式化UTC时间
  currentUTCTime.value = now.toISOString().substring(11, 19)
}

// 格式化日期
const formatDate = (date) => {
  if (!date) return ''
  
  try {
    const dateObj = new Date(date)
    const formatter = new Intl.DateTimeFormat(locale.value, {
      timeZone: currentTimezone.value,
      year: 'numeric',
      month: '2-digit',
      day: '2-digit'
    })
    return formatter.format(dateObj)
  } catch (error) {
    return date.toString()
  }
}

// 格式化日期时间
const formatDateTime = (datetime) => {
  if (!datetime) return ''
  
  try {
    const dateObj = new Date(datetime)
    const formatter = new Intl.DateTimeFormat(locale.value, {
      timeZone: currentTimezone.value,
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
      hour: '2-digit',
      minute: '2-digit',
      second: '2-digit',
      hour12: false
    })
    return formatter.format(dateObj)
  } catch (error) {
    return datetime.toString()
  }
}

// 事件处理
const handleSubmit = (data) => {
  ElMessage.success('表单提交成功！')
  console.log('提交的数据:', data)
  
  // 显示时区感知的时间信息
  console.log('时区感知的格式化结果:')
  console.log('创建日期:', formatDate(data.createDate))
  console.log('截止日期:', formatDate(data.deadline))
  console.log('开始时间:', formatDateTime(data.eventStart))
  console.log('结束时间:', formatDateTime(data.eventEnd))
}

// 生命周期
onMounted(() => {
  updateTime()
  timeTimer = setInterval(updateTime, 1000)
})

onUnmounted(() => {
  if (timeTimer) {
    clearInterval(timeTimer)
  }
})
</script>

<style lang="scss" scoped>
.timezone-form-example {
  padding: 20px;
  
  h2 {
    margin-bottom: 20px;
    color: #303133;
  }
  
  .example-card {
    margin-bottom: 20px;
  }
  
  .config-info {
    h4 {
      margin: 16px 0 8px 0;
      color: #409eff;
    }
    
    ul {
      margin: 8px 0;
      padding-left: 20px;
      
      li {
        margin: 8px 0;
        line-height: 1.6;
        
        strong, code {
          color: #e6a23c;
        }
      }
    }
  }
}
</style>