<template>
  <div class="simple-example">
    <h2>简化版动态表单示例</h2>
    
    <el-card class="example-card">
      <template #header>
        <span>基础用法 - 传入配置</span>
      </template>
      
      <DynamicFormSimple
        v-model="formData"
        :form-options="formOptions"
        :show-actions="true"
        @submit="handleSubmit"
        @reset="handleReset"
        @validate="handleValidate"
      >
        <template #additional="{ formData }">
          <el-form-item label="附加字段">
            <el-input v-model="additionalField" placeholder="这是一个额外的字段" />
          </el-form-item>
        </template>
        
        <template #actions="{ formData, validate, reset }">
          <el-button @click="reset">重置表单</el-button>
          <el-button type="primary" @click="validate">验证提交</el-button>
          <el-button type="success" @click="showFormData">查看数据</el-button>
        </template>
      </DynamicFormSimple>
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
          <pre class="json-display">{{ JSON.stringify(formOptions, null, 2) }}</pre>
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
          <DynamicFormSimple
            v-model="inlineFormData"
            :form-options="simpleFormOptions"
            :inline="true"
            :show-actions="false"
            size="small"
          />
        </el-col>
        
        <el-col :span="12">
          <h4>禁用表单</h4>
          <DynamicFormSimple
            v-model="disabledFormData"
            :form-options="simpleFormOptions"
            :disabled="true"
            :show-actions="false"
            size="small"
          />
        </el-col>
      </el-row>
    </el-card>
    
    <el-card class="example-card">
      <template #header>
        <span>动态配置生成器</span>
      </template>
      
      <div class="config-generator">
        <el-button type="primary" @click="generateUserConfig">生成用户表单配置</el-button>
        <el-button type="success" @click="generateProductConfig">生成商品表单配置</el-button>
        <el-button type="warning" @click="generateOrderConfig">生成订单表单配置</el-button>
      </div>
      
      <div v-if="generatedConfig.length > 0" style="margin-top: 20px;">
        <h4>生成的表单配置:</h4>
        <DynamicFormSimple
          v-model="generatedFormData"
          :form-options="generatedConfig"
          :show-actions="true"
        />
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import DynamicFormSimple from './DynamicFormSimple.vue'

defineOptions({
  name: 'SimpleExample'
})

const activeTab = ref('data')
const additionalField = ref('')

// 主表单数据
const formData = reactive({
  用户名: 'admin',
  年龄: 25,
  分数: 88.5,
  是否激活: true,
  用户类型: 'admin',
  生日: new Date(),
  工作时间: new Date(),
  个人简介: '这是一个测试用户'
})

// 表单配置选项 - 对应Go代码中的ItemOption
const formOptions = [
  {
    name: '用户名',
    type: 'string',
    required: true,
    default_value: ''
  },
  {
    name: '年龄',
    type: 'int',
    required: true,
    default_value: 0
  },
  {
    name: '分数',
    type: 'float',
    required: false,
    default_value: 0.0
  },
  {
    name: '是否激活',
    type: 'bool',
    required: false,
    default_value: false
  },
  {
    name: '用户类型',
    type: 'enum',
    required: true,
    options: ['admin', 'user', 'guest'],
    default_value: 'user'
  },
  {
    name: '生日',
    type: 'date',
    required: false,
    default_value: null
  },
  {
    name: '工作时间',
    type: 'time',
    required: false,
    default_value: null
  },
  {
    name: '个人简介',
    type: 'text',
    required: false,
    default_value: ''
  }
]

// 简单表单配置
const simpleFormOptions = [
  {
    name: '姓名',
    type: 'string',
    required: true,
    default_value: 'John'
  },
  {
    name: '年龄',
    type: 'int',
    required: false,
    default_value: 30
  },
  {
    name: '状态',
    type: 'bool',
    required: false,
    default_value: true
  }
]

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

// 动态生成的配置和数据
const generatedConfig = ref([])
const generatedFormData = reactive({})

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

const showFormData = () => {
  ElMessage.info('请查看控制台输出')
  console.log('当前表单数据:', formData)
  console.log('附加字段:', additionalField.value)
}

// 动态配置生成器
const generateUserConfig = () => {
  generatedConfig.value = [
    {
      name: '用户名',
      type: 'string',
      required: true,
      default_value: ''
    },
    {
      name: '邮箱',
      type: 'string',
      required: true,
      default_value: ''
    },
    {
      name: '年龄',
      type: 'int',
      required: false,
      default_value: 18
    },
    {
      name: '性别',
      type: 'enum',
      required: true,
      options: ['男', '女', '其他'],
      default_value: '男'
    },
    {
      name: '是否VIP',
      type: 'bool',
      required: false,
      default_value: false
    }
  ]
  resetGeneratedFormData()
  ElMessage.success('用户表单配置已生成!')
}

const generateProductConfig = () => {
  generatedConfig.value = [
    {
      name: '商品名称',
      type: 'string',
      required: true,
      default_value: ''
    },
    {
      name: '价格',
      type: 'float',
      required: true,
      default_value: 0.0
    },
    {
      name: '库存数量',
      type: 'int',
      required: true,
      default_value: 0
    },
    {
      name: '商品分类',
      type: 'enum',
      required: true,
      options: ['电子产品', '服装', '食品', '书籍'],
      default_value: '电子产品'
    },
    {
      name: '是否上架',
      type: 'bool',
      required: false,
      default_value: true
    },
    {
      name: '商品描述',
      type: 'text',
      required: false,
      default_value: ''
    }
  ]
  resetGeneratedFormData()
  ElMessage.success('商品表单配置已生成!')
}

const generateOrderConfig = () => {
  generatedConfig.value = [
    {
      name: '订单号',
      type: 'string',
      required: true,
      default_value: ''
    },
    {
      name: '订单金额',
      type: 'float',
      required: true,
      default_value: 0.0
    },
    {
      name: '下单日期',
      type: 'date',
      required: true,
      default_value: new Date()
    },
    {
      name: '订单状态',
      type: 'enum',
      required: true,
      options: ['待付款', '待发货', '已发货', '已完成', '已取消'],
      default_value: '待付款'
    },
    {
      name: '是否急单',
      type: 'bool',
      required: false,
      default_value: false
    },
    {
      name: '备注',
      type: 'text',
      required: false,
      default_value: ''
    }
  ]
  resetGeneratedFormData()
  ElMessage.success('订单表单配置已生成!')
}

const resetGeneratedFormData = () => {
  // 清空现有数据
  Object.keys(generatedFormData).forEach(key => {
    delete generatedFormData[key]
  })
  
  // 根据配置初始化数据
  generatedConfig.value.forEach(option => {
    generatedFormData[option.name] = option.default_value
  })
}
</script>

<style lang="scss" scoped>
.simple-example {
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
    
    .config-generator {
      display: flex;
      gap: 12px;
      flex-wrap: wrap;
    }
  }
  
  h4 {
    margin-bottom: 10px;
    color: #606266;
  }
}
</style>