<template>
  <div class="backend-timezone-example">
    <h2>后端时区数据示例</h2>

    <el-card class="example-card">
      <template #header>
        <div class="card-header">
          <span>从后端获取时区数据</span>
          <el-button type="primary" @click="reloadTimezones">重新加载</el-button>
        </div>
      </template>

      <el-form label-width="120px">
        <el-form-item label="时区选择:">
          <TimezoneSelect
            ref="timezoneRef"
            v-model="selectedTimezone"
            :show-current-time="true"
            @change="handleTimezoneChange"
          />
        </el-form-item>
      </el-form>

      <el-divider content-position="left">
        加载的时区数据
      </el-divider>

      <el-table
        :data="timezoneTableData"
        border
        style="width: 100%"
        max-height="400"
      >
        <el-table-column prop="value" label="时区标识" width="200" />
        <el-table-column prop="name" label="时区名称" width="200" />
        <el-table-column prop="region" label="地区" width="100" />
        <el-table-column prop="offset" label="偏移量" width="100" />
      </el-table>
    </el-card>

    <el-card class="example-card">
      <template #header>
        <span>预期的后端数据格式</span>
      </template>

      <div class="code-block">
        <h4>API 响应格式:</h4>
        <pre class="api-format">{{ apiFormat }}</pre>

        <h4>数据结构说明:</h4>
        <ul class="data-structure">
          <li><strong>value</strong>: 时区标识符 (例如: 'Asia/Shanghai')</li>
          <li><strong>name</strong>: 时区显示名称 (例如: '中国标准时间')</li>
          <li><strong>offset</strong>: UTC偏移量 (例如: '+08:00')</li>
          <li><strong>region</strong>: 地区分组 (例如: '亚洲')</li>
        </ul>
      </div>
    </el-card>

    <el-card class="example-card">
      <template #header>
        <span>后端API接口说明</span>
      </template>

      <el-descriptions :column="1" border>
        <el-descriptions-item label="获取所有时区">
          <code>GET /timezone/list</code>
          <p>返回所有可用的时区列表</p>
        </el-descriptions-item>
        <el-descriptions-item label="按地区获取时区">
          <code>GET /timezone/region?region=亚洲</code>
          <p>根据地区筛选时区</p>
        </el-descriptions-item>
        <el-descriptions-item label="获取服务器时区">
          <code>GET /timezone/server</code>
          <p>获取当前服务器的时区设置</p>
        </el-descriptions-item>
        <el-descriptions-item label="时区转换">
          <code>POST /timezone/convert</code>
          <p>在不同时区之间转换时间</p>
        </el-descriptions-item>
      </el-descriptions>
    </el-card>

    <el-card class="example-card">
      <template #header>
        <span>实际使用示例</span>
      </template>

      <DynamicFormSimple
        v-model="userFormData"
        :form-options="userFormOptions"
        struct-name="user-settings"
        @submit="handleUserSubmit"
      >
        <template #additional>
          <el-form-item label="首选时区" prop="timezone">
            <TimezoneSelect
              v-model="userFormData.timezone"
              placeholder="请选择时区"
            />
          </el-form-item>

          <el-form-item label="通知时区" prop="notificationTimezone">
            <TimezoneSelect
              v-model="userFormData.notificationTimezone"
              placeholder="请选择通知时区"
              :show-current-time="false"
            />
          </el-form-item>
        </template>
      </DynamicFormSimple>
    </el-card>
  </div>
</template>

<script setup>
import { ref, computed, reactive, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import TimezoneSelect from './TimezoneSelect.vue'
import DynamicFormSimple from './DynamicFormSimple.vue'

defineOptions({
  name: 'BackendTimezoneExample'
})

const timezoneRef = ref()
const selectedTimezone = ref('Asia/Shanghai')

// 用户表单数据
const userFormData = reactive({
  username: 'admin',
  email: 'admin@example.com',
  language: 'zh-CN',
  timezone: 'Asia/Shanghai',
  notificationTimezone: 'UTC'
})

// 用户表单配置
const userFormOptions = [
  {
    name: 'username',
    type: 'string',
    required: true,
    default_value: ''
  },
  {
    name: 'email',
    type: 'string',
    required: true,
    default_value: ''
  },
  {
    name: 'language',
    type: 'enum',
    required: true,
    options: ['zh-CN', 'en-US', 'ja-JP'],
    default_value: 'zh-CN'
  }
]

// 时区表格数据
const timezoneTableData = computed(() => {
  return timezoneRef.value?.timezones?.value || []
})

// API 格式示例
const apiFormat = `{
  "code": 0,
  "message": "success",
  "data": [
    {
      "value": "Asia/Shanghai",
      "name": "中国标准时间", 
      "offset": "+08:00",
      "region": "亚洲"
    },
    {
      "value": "Asia/Tokyo",
      "name": "日本标准时间",
      "offset": "+09:00", 
      "region": "亚洲"
    },
    {
      "value": "Europe/London",
      "name": "英国时间",
      "offset": "+00:00",
      "region": "欧洲"
    },
    {
      "value": "America/New_York",
      "name": "美国东部时间",
      "offset": "-05:00",
      "region": "北美"
    }
  ]
}`

// 事件处理
const handleTimezoneChange = (timezone) => {
  ElMessage.success(`时区已切换到: ${timezone}`)
  console.log('选择的时区:', timezone)
}

const handleUserSubmit = (data) => {
  ElMessage.success('用户设置已保存!')
  console.log('用户设置:', data)
}

const reloadTimezones = async() => {
  if (timezoneRef.value) {
    await timezoneRef.value.loadTimezones()
    ElMessage.success('时区数据已重新加载')
  }
}
</script>

<style lang="scss" scoped>
.backend-timezone-example {
  padding: 20px;

  h2 {
    margin-bottom: 20px;
    color: #303133;
  }

  .example-card {
    margin-bottom: 20px;

    .card-header {
      display: flex;
      justify-content: space-between;
      align-items: center;
    }
  }

  .code-block {
    .api-format {
      background: #f5f7fa;
      padding: 16px;
      border-radius: 4px;
      font-family: 'Courier New', monospace;
      font-size: 12px;
      color: #606266;
      overflow-x: auto;
      max-height: 300px;
      overflow-y: auto;
    }

    .data-structure {
      margin: 16px 0;
      padding-left: 20px;

      li {
        margin: 8px 0;
        line-height: 1.6;

        strong {
          color: #409eff;
        }
      }
    }
  }
}
</style>
