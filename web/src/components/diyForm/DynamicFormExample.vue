<template>
  <div class="dynamic-form-example">
    <h2>动态表单示例</h2>

    <el-card class="example-card">
      <template #header>
        <span>基础用法</span>
      </template>

      <DynamicForm
        v-model="formData"
        form-type="user-profile"
        :show-actions="true"
        @submit="handleSubmit"
        @reset="handleReset"
        @validate="handleValidate"
        @config-loaded="onConfigLoaded"
        @config-error="onConfigError"
      >
        <template #additional>
          <el-form-item label="附加字段">
            <el-input v-model="additionalField" placeholder="这是一个额外的字段" />
          </el-form-item>
        </template>

        <template #actions="{ validate, reset }">
          <el-button @click="reset">重置表单</el-button>
          <el-button type="primary" @click="validate">验证提交</el-button>
          <el-button type="success" @click="showFormData">查看数据</el-button>
        </template>
      </DynamicForm>
    </el-card>

    <el-card class="example-card">
      <template #header>
        <span>表单配置和数据</span>
      </template>

      <el-tabs v-model="activeTab">
        <el-tab-pane label="当前表单数据" name="data">
          <pre class="json-display">{{ JSON.stringify(formData, null, 2) }}</pre>
        </el-tab-pane>

        <el-tab-pane label="表单配置" name="options">
          <pre class="json-display">{{ JSON.stringify(loadedFormOptions, null, 2) }}</pre>
        </el-tab-pane>

        <el-tab-pane label="Go代码映射" name="go">
          <div class="code-block">
            <h4>对应的Go代码:</h4>
            <pre class="go-code">{{ goCodeExample }}</pre>
          </div>
        </el-tab-pane>
      </el-tabs>
    </el-card>

    <el-card class="example-card">
      <template #header>
        <span>其他配置示例</span>
      </template>

      <el-row :gutter="20">
        <el-col :span="12">
          <h4>内联表单</h4>
          <DynamicForm
            v-model="inlineFormData"
            form-type="simple-form"
            :inline="true"
            :show-actions="false"
            size="small"
          />
        </el-col>

        <el-col :span="12">
          <h4>禁用表单</h4>
          <DynamicForm
            v-model="disabledFormData"
            form-type="simple-form"
            :disabled="true"
            :show-actions="false"
            size="small"
          />
        </el-col>
      </el-row>
    </el-card>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import DynamicForm from './DynamicForm.vue'

defineOptions({
  name: 'DynamicFormExample'
})

const activeTab = ref('data')
const additionalField = ref('')

// 主表单数据
const formData = reactive({
  username: 'admin',
  age: 25,
  score: 88.5,
  isActive: true,
  userType: 'admin',
  birthday: new Date(),
  workTime: new Date()
})

// 加载的表单配置
const loadedFormOptions = ref([])

const inlineFormData = reactive({
  姓名: 'John',
  年龄: 30,
  状态: true
})

const disabledFormData = reactive({
  姓名: 'Jane',
  年龄: 25,
  状态: false
})

// Go代码示例
const goCodeExample = `// 使用Go代码创建对应的表单配置
func CreateUserForm() []diyfrom.ItemOption {
    return []diyfrom.ItemOption{
        diyfrom.NewStrItem("用户名", true),
        diyfrom.NewIntItem("年龄", true),
        diyfrom.NewFloatItem("分数", false),
        diyfrom.NewBoolItem("是否激活", false),
        diyfrom.NewEnumItem("用户类型", true, "admin", "user", "guest"),
        diyfrom.NewDateItem("生日", false),
        diyfrom.NewTimeItem("工作时间", false),
        diyfrom.NewTextItem("个人简介", false),
    }
}`

// 事件处理
const handleSubmit = (data) => {
  ElMessage.success('表单提交成功!')
  console.log('提交的数据:', data)
}

const handleReset = (data) => {
  ElMessage.info('表单已重置')
  console.log('重置后的数据:', data)
}

const handleValidate = (isValid, data) => {
  if (isValid) {
    ElMessage.success('表单验证通过!')
  } else {
    ElMessage.error('表单验证失败!')
  }
  console.log('验证结果:', isValid, '数据:', data)
}

const onConfigLoaded = (config) => {
  loadedFormOptions.value = config
  ElMessage.success('表单配置加载成功!')
  console.log('加载的配置:', config)
}

const onConfigError = (error) => {
  ElMessage.error('加载表单配置失败!')
  console.error('配置加载错误:', error)
}

const showFormData = () => {
  ElMessage.info('请查看控制台输出')
  console.log('当前表单数据:', formData)
  console.log('附加字段:', additionalField.value)
}
</script>

<style lang="scss" scoped>
.dynamic-form-example {
  padding: 20px;

  h2 {
    margin-bottom: 20px;
    color: #303133;
  }

  .example-card {
    margin-bottom: 20px;

    .json-display {
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

    .code-block {
      h4 {
        margin-bottom: 10px;
        color: #409eff;
      }

      .go-code {
        background: #2d2d2d;
        color: #f8f8f2;
        padding: 16px;
        border-radius: 4px;
        font-family: 'Courier New', monospace;
        font-size: 12px;
        overflow-x: auto;
        max-height: 300px;
        overflow-y: auto;
      }
    }
  }

  h4 {
    margin-bottom: 10px;
    color: #606266;
  }
}
</style>
