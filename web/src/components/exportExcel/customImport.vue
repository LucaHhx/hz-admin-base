<template>
  <div class="import-container">
    <el-upload
      v-if="btnAuth.importExcel"
      accept=".xlsx, .xls"
      :auto-upload="false"
      :show-file-list="true"
      :limit="1"
      :on-change="handleFileChange"
      :file-list="fileList"
      :on-remove="handleFileRemove"
    >
      <el-button
        type="primary"
        :icon="Upload"
      >
        {{ $t('common.import') }}
      </el-button>
    </el-upload>

    <el-button
      v-if="fileList.length"
      type="success"
      :disabled="importing"
      @click="importData('full')"
    >
      {{ importing ? $t('common.importing') : $t('common.importFull') }}
    </el-button>
    <el-button
      v-if="fileList.length"
      style="margin-left: 10px;"
      type="success"
      :disabled="importing"
      @click="importData('append')"
    >
      {{ importing ? $t('common.importing') : $t('common.importAppend') }}
    </el-button>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { Upload } from '@element-plus/icons-vue'
import * as XLSX from 'xlsx'
import { useI18n } from 'vue-i18n'
import { convertDateTime, convertDate } from '@/utils/time'

const { t } = useI18n()
import { useBtnAuth } from '@/utils/btnAuth'
// 按钮权限实例化
const btnAuth = useBtnAuth()

const props = defineProps({
  importFunction: {
    type: Function,
    required: true
  },
  columns: {
    type: Array,
    required: true
  }
})

const emit = defineEmits(['import-success'])

const fileList = ref([])
const importing = ref(false)

// 处理文件变化
const handleFileChange = (file) => {
  fileList.value = [file]
}

// 处理文件删除
const handleFileRemove = (file) => {
  fileList.value = []
}

// 解析Excel文件
const parseExcel = (file) => {
  return new Promise((resolve, reject) => {
    const reader = new FileReader()

    reader.onload = (e) => {
      try {
        const fileData = e.target.result
        const workbook = XLSX.read(fileData, { type: 'array' })
        const firstSheetName = workbook.SheetNames[0]
        const worksheet = workbook.Sheets[firstSheetName]
        const jsonData = XLSX.utils.sheet_to_json(worksheet, { header: 1 })

        // 检查是否有数据
        if (jsonData.length < 4) { // 至少需要key、label、type行和一行数据
          reject(new Error(t('common.noDataInFile')))
          return
        }

        // 获取key和type行
        const keys = jsonData[0]
        const types = jsonData[2] // 第三行是类型

        // 打印一下列信息看看
        console.log('Import keys:', keys)
        console.log('Import types:', types)

        // 创建字段和类型的映射
        const fieldMap = {}
        const typeMap = {}

        for (let i = 0; i < keys.length; i++) {
          const key = keys[i]
          const type = (types[i] || 'string').toLowerCase() // 默认为字符串类型

          if (key) {
            fieldMap[i] = key
            typeMap[i] = type
          }
        }

        // 确定数据起始行
        // 如果第四行是格式提示（包含"格式"字样），则从第五行开始
        let dataStartRow = 3
        const fourthRow = jsonData[3] || []
        const hasTips = fourthRow.some(cell =>
          typeof cell === 'string' &&
          (cell.includes('格式') || cell.includes('布尔值') || cell.includes('整数') || cell.includes('小数'))
        )

        if (hasTips) {
          dataStartRow = 4
        }

        // 转换数据（从数据起始行开始）
        const resultData = []
        for (let i = dataStartRow; i < jsonData.length; i++) {
          const row = jsonData[i]
          const item = {}

          // 跳过空行
          if (!row || row.every(cell => cell === undefined || cell === null || cell === '')) {
            continue
          }

          Object.keys(fieldMap).forEach(index => {
            const key = fieldMap[index]
            const type = typeMap[index]
            let value = row[index]

            // 处理空值，统一转为null
            if (value === undefined || value === null || value === '') {
              item[key] = null
              return
            }

            // 根据类型进行转换
            try {
              switch (type) {
                case 'boolean':
                case 'bool':
                  if (typeof value === 'string') {
                    value = value.toLowerCase()
                    value = value === 'true' || value === '1' || value === 'yes' || value === '是'
                  } else if (typeof value === 'number') {
                    value = value === 1
                  }
                  break
                case 'number':
                case 'int':
                case 'integer':
                case 'int64':
                case 'int32':
                case 'int16':
                case 'bigint':
                case 'uint':
                case 'float':
                case 'float32':
                case 'float64':
                case 'double':
                case 'decimal':
                case 'table':
                  // 所有浮点数类型统一处理
                  if (typeof value === 'string') {
                    // 尝试直接转换为浮点数
                    const floatValue = parseFloat(value)
                    if (!isNaN(floatValue)) {
                      value = floatValue
                    }
                  }
                  // 如果已经是数字则保持不变
                  break
                case 'json':
                  if (typeof value === 'string') {
                    try {
                      value = JSON.parse(value)
                    } catch (e) {
                      // 如果解析失败，保持原字符串
                    }
                  }
                  break
                default:
                  // 对于字符串类型，确保是字符串
                  if (typeof value !== 'string') {
                    value = String(value)
                  }
                  break
              }
            } catch (error) {
              console.error(`错误：转换字段 ${key} 值 ${value} 到类型 ${type} 失败:`, error)
              // 如果转换失败，保留原始值
            }

            // 设置属性值
            item[key] = value
          })

          if (Object.keys(item).length > 0) {
            resultData.push(item)
          }
        }

        console.log('Parsed import data:', resultData)
        resolve(resultData)
      } catch (error) {
        console.error('Parse Excel error:', error)
        reject(error)
      }
    }

    reader.onerror = (error) => {
      reject(error)
    }

    reader.readAsArrayBuffer(file.raw)
  })
}

// 导入数据
const importData = async(tep) => {
  if (fileList.value.length === 0) {
    ElMessage.warning(t('common.pleaseSelectFile'))
    return
  }

  importing.value = true

  try {
    const file = fileList.value[0]
    const data = await parseExcel(file)
    if (data.length === 0) {
      ElMessage.warning(t('common.noDataToImport'))
      importing.value = false
      return
    }

    // 调用导入函数
    const result = await props.importFunction({ list: data, type: tep })
    if (result.code === 0) {
      ElMessage.success(t('common.importSuccess'))
      fileList.value = []
      emit('import-success')
    } else {
      // ElMessage.error(result.msg || t('common.importFailed'))
    }
  } catch (error) {
    ElMessage.error(error.message || t('common.importFailed'))
  } finally {
    importing.value = false
  }
}
</script>

<style scoped>
.import-container {
  display: inline-flex;
  align-items: center;
  /* gap: 10px; */
}
.el-upload {
  margin-right: 5px;
}
</style>
