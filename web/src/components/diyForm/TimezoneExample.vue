<template>
  <div class="timezone-example">
    <h2>时区选择器示例</h2>

    <el-card class="example-card">
      <template #header>
        <span>基础用法</span>
      </template>

      <el-form label-width="120px">
        <el-form-item label="选择时区:">
          <TimezoneSelect
            v-model="selectedTimezone"
            :show-current-time="true"
            @change="handleTimezoneChange"
          />
        </el-form-item>

        <el-form-item label="禁用状态:">
          <TimezoneSelect
            v-model="disabledTimezone"
            :disabled="true"
            :show-current-time="false"
          />
        </el-form-item>

        <el-form-item label="小尺寸:">
          <TimezoneSelect
            v-model="smallTimezone"
            size="small"
            :show-current-time="false"
          />
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="example-card">
      <template #header>
        <span>在动态表单中使用</span>
      </template>

      <DynamicFormSimple
        v-model="formData"
        :form-options="formOptions"
        struct-name="timezone-form"
        @submit="handleSubmit"
      >
        <!-- 自定义时区选择器 -->
        <template #additional>
          <el-form-item label="首选时区" prop="preferredTimezone">
            <TimezoneSelect
              v-model="formData.preferredTimezone"
              placeholder="请选择首选时区"
            />
          </el-form-item>

          <el-form-item label="备用时区" prop="backupTimezone">
            <TimezoneSelect
              v-model="formData.backupTimezone"
              placeholder="请选择备用时区"
              :show-current-time="false"
            />
          </el-form-item>
        </template>
      </DynamicFormSimple>
    </el-card>

    <el-card class="example-card">
      <template #header>
        <span>时区信息展示</span>
      </template>

      <el-descriptions :column="2" border>
        <el-descriptions-item label="当前选中时区">
          {{ selectedTimezone || '未选择' }}
        </el-descriptions-item>
        <el-descriptions-item label="时区偏移">
          {{ getCurrentOffset() }}
        </el-descriptions-item>
        <el-descriptions-item label="当前时间">
          {{ getCurrentTime() }}
        </el-descriptions-item>
        <el-descriptions-item label="UTC时间">
          {{ getUTCTime() }}
        </el-descriptions-item>
      </el-descriptions>
    </el-card>

    <el-card class="example-card">
      <template #header>
        <span>时区转换工具</span>
      </template>

      <div class="timezone-converter">
        <el-row :gutter="20">
          <el-col :span="8">
            <el-form-item label="源时区">
              <TimezoneSelect
                v-model="sourceTimezone"
                :show-current-time="false"
              />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="目标时区">
              <TimezoneSelect
                v-model="targetTimezone"
                :show-current-time="false"
              />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="转换时间">
              <el-date-picker
                v-model="convertTime"
                type="datetime"
                placeholder="选择时间"
                style="width: 100%"
              />
            </el-form-item>
          </el-col>
        </el-row>

        <div v-if="sourceTimezone && targetTimezone && convertTime" class="conversion-result">
          <el-alert
            :title="`转换结果: ${formatConvertedTime()}`"
            type="info"
            :closable="false"
          />
        </div>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import TimezoneSelect from './TimezoneSelect.vue'
import DynamicFormSimple from './DynamicFormSimple.vue'

defineOptions({
  name: 'TimezoneExample'
})

// 基础用法
const selectedTimezone = ref('Asia/Shanghai')
const disabledTimezone = ref('America/New_York')
const smallTimezone = ref('Europe/London')

// 动态表单数据
const formData = reactive({
  userName: 'admin',
  userEmail: 'admin@example.com',
  isActive: true,
  preferredTimezone: 'Asia/Shanghai',
  backupTimezone: 'UTC'
})

// 表单配置
const formOptions = [
  {
    name: 'userName',
    type: 'string',
    required: true,
    default_value: ''
  },
  {
    name: 'userEmail',
    type: 'string',
    required: true,
    default_value: ''
  },
  {
    name: 'isActive',
    type: 'bool',
    required: false,
    default_value: true
  }
]

// 时区转换工具
const sourceTimezone = ref('Asia/Shanghai')
const targetTimezone = ref('America/New_York')
const convertTime = ref(new Date())

// 事件处理
const handleTimezoneChange = (timezone) => {
  ElMessage.success(`时区已切换到: ${timezone}`)
  console.log('时区变更:', timezone)
}

const handleSubmit = (data) => {
  ElMessage.success('表单提交成功!')
  console.log('提交的数据:', data)
}

// 获取当前时区偏移
const getCurrentOffset = () => {
  if (!selectedTimezone.value) return ''

  try {
    const now = new Date()
    const formatter = new Intl.DateTimeFormat('en', {
      timeZone: selectedTimezone.value,
      timeZoneName: 'longOffset'
    })

    const parts = formatter.formatToParts(now)
    const offsetPart = parts.find(part => part.type === 'timeZoneName')
    return offsetPart ? offsetPart.value : ''
  } catch (error) {
    return 'N/A'
  }
}

// 获取当前时间
const getCurrentTime = () => {
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

    return formatter.format(now)
  } catch (error) {
    return 'N/A'
  }
}

// 获取UTC时间
const getUTCTime = () => {
  const now = new Date()
  return now.toISOString().replace('T', ' ').substring(0, 19) + ' UTC'
}

// 格式化转换后的时间
const formatConvertedTime = () => {
  if (!sourceTimezone.value || !targetTimezone.value || !convertTime.value) {
    return ''
  }

  try {
    // 创建源时区时间格式化器
    const sourceFormatter = new Intl.DateTimeFormat('zh-CN', {
      timeZone: sourceTimezone.value,
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
      hour: '2-digit',
      minute: '2-digit',
      second: '2-digit',
      hour12: false
    })

    // 创建目标时区时间格式化器
    const targetFormatter = new Intl.DateTimeFormat('zh-CN', {
      timeZone: targetTimezone.value,
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
      hour: '2-digit',
      minute: '2-digit',
      second: '2-digit',
      hour12: false
    })

    const sourceTime = sourceFormatter.format(convertTime.value)
    const targetTime = targetFormatter.format(convertTime.value)

    return `${sourceTime} (${sourceTimezone.value}) → ${targetTime} (${targetTimezone.value})`
  } catch (error) {
    return '转换失败'
  }
}
</script>

<style lang="scss" scoped>
.timezone-example {
  padding: 20px;

  h2 {
    margin-bottom: 20px;
    color: #303133;
  }

  .example-card {
    margin-bottom: 20px;
  }

  .timezone-converter {
    .conversion-result {
      margin-top: 20px;
    }
  }
}
</style>
