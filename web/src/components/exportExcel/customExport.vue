<template>
  <el-dropdown
    v-if="btnAuth.exportExcel"
    trigger="click"
    @command="handleCommand"
  >
    <el-button type="primary" :icon="Download">
      {{ $t('common.export') }}
      <el-icon class="el-icon--right">
        <arrow-down />
      </el-icon>
    </el-button>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item v-if="btnAuth.exportExcel" command="exportAll">
          {{ $t('common.exportAll') }}
        </el-dropdown-item>
        <el-dropdown-item v-if="btnAuth.exportExcel" command="exportCurrent">
          {{ $t('common.exportCurrent') }}
        </el-dropdown-item>
        <el-dropdown-item
          v-if="btnAuth.exportTemplate"
          command="exportTemplate"
        >
          {{ $t('common.exportTemplate') }}
        </el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>

  <el-dialog
    v-model="dialogVisible"
    :title="$t('common.exportProgress')"
    width="30%"
  >
    <div class="export-progress">
      <el-progress :percentage="exportProgress" :format="progressFormat" />
      <p>{{ progressText }}</p>
    </div>
  </el-dialog>
</template>

<script setup>
import { ref, computed } from 'vue'
import { ElMessage } from 'element-plus'
import { Download, ArrowDown } from '@element-plus/icons-vue'
import { batchExport, generateTemplateData } from '@/utils/export'
import * as XLSX from 'xlsx'
import ExcelJS from 'exceljs'
import { useI18n } from 'vue-i18n'
import { requestFilterLabel, requestWithCache } from '@/utils/sysDataFilterCache'

const { t } = useI18n()

import { useBtnAuth } from '@/utils/btnAuth'
// 按钮权限实例化
const btnAuth = useBtnAuth()

// 日期格式化函数
const formatDate = (value) => {
  if (!value) return ''

  let date
  if (value instanceof Date) {
    date = value
  } else if (typeof value === 'string') {
    // 尝试解析ISO格式日期
    date = new Date(value)
  } else {
    return value
  }

  // 检查日期是否有效
  if (isNaN(date.getTime())) {
    return value
  }

  // 格式化为 yyyy-MM-dd
  const year = date.getFullYear()
  const month = String(date.getMonth() + 1).padStart(2, '0')
  const day = String(date.getDate()).padStart(2, '0')

  return `${year}-${month}-${day}`
}

// 日期时间格式化函数
const formatDateTime = (value) => {
  if (!value) return ''

  let date
  if (value instanceof Date) {
    date = value
  } else if (typeof value === 'string') {
    // 尝试解析ISO格式日期
    date = new Date(value)
  } else {
    return value
  }

  // 检查日期是否有效
  if (isNaN(date.getTime())) {
    return value
  }

  // 格式化为 yyyy-MM-dd HH:mm:ss
  const year = date.getFullYear()
  const month = String(date.getMonth() + 1).padStart(2, '0')
  const day = String(date.getDate()).padStart(2, '0')
  const hours = String(date.getHours()).padStart(2, '0')
  const minutes = String(date.getMinutes()).padStart(2, '0')
  const seconds = String(date.getSeconds()).padStart(2, '0')

  return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`
}

const props = defineProps({
  fetchFunction: {
    type: Function,
    required: true
  },
  searchInfo: {
    type: Object,
    default: () => ({})
  },
  columns: {
    type: Array,
    required: true
  },
  fileName: {
    type: String,
    default: 'export'
  },
  batchSize: {
    type: Number,
    default: 1000
  }
})

// 导出进度相关
const dialogVisible = ref(false)
const exportProgress = ref(0)
const currentPage = ref(0)
const totalPages = ref(0)
const currentCount = ref(0)
const totalCount = ref(0)

// 格式化进度
const progressFormat = (percentage) => {
  return `${Math.floor(percentage)}%`
}

// 进度文本
const progressText = computed(() => {
  return t('common.exportProgressText', {
    current: currentPage.value,
    total: totalPages.value,
    count: currentCount.value,
    totalCount: totalCount.value
  })
})

// 更新进度
const updateProgress = (page, pages, count, total) => {
  currentPage.value = page
  totalPages.value = pages
  currentCount.value = count
  totalCount.value = total
  exportProgress.value = Math.floor((page / pages) * 100)
}

// 使用 ExcelJS 导出包含图片的 Excel
const exportToExcelWithImages = async(data, filename, headers) => {
  const workbook = new ExcelJS.Workbook()
  const worksheet = workbook.addWorksheet('Sheet1')

  // 设置列
  worksheet.columns = headers.map((h) => ({
    header: h.label,
    key: h.key,
    width: h.type.toLowerCase() === 'image' ? 20 : 15
  }))

  // 设置表头行高
  worksheet.getRow(1).height = 20

  // 添加数据行并处理图片
  for (let rowIndex = 0; rowIndex < data.length; rowIndex++) {
    const row = data[rowIndex]
    const rowData = {}

    // 先添加非图片数据
    for (const h of headers) {
      const value = row[h.key]
      const type = h.type.toLowerCase()

      if (type !== 'image') {
        // 处理非图片类型的值
        if (value === null || value === undefined) {
          rowData[h.key] = ''
        } else {
          switch (type) {
            case 'boolean':
            case 'bool':
              rowData[h.key] = value ? 'true' : 'false'
              break
            case 'date':
            case 'datetime':
              rowData[h.key] = formatDateTime(value)
              break
            case 'number':
            case 'int':
            case 'integer':
            case 'int64':
            case 'int32':
            case 'int16':
            case 'uint':
              if (typeof value === 'number') {
                rowData[h.key] = Math.round(value)
              } else if (typeof value === 'string' && !isNaN(parseInt(value, 10))) {
                rowData[h.key] = parseInt(value, 10)
              } else {
                rowData[h.key] = value
              }
              break
            case 'float':
            case 'float32':
            case 'float64':
            case 'double':
            case 'decimal':
              if (typeof value === 'number') {
                rowData[h.key] = value
              } else if (typeof value === 'string' && !isNaN(parseFloat(value))) {
                rowData[h.key] = parseFloat(value)
              } else {
                rowData[h.key] = value
              }
              break
            case 'json':
              rowData[h.key] = JSON.stringify(value)
              break
            default:
              rowData[h.key] = value
          }
        }
      } else {
        rowData[h.key] = '' // 图片列先留空
      }
    }

    const excelRow = worksheet.addRow(rowData)
    excelRow.height = 80 // 设置行高以容纳图片

    // 处理图片
    for (let colIndex = 0; colIndex < headers.length; colIndex++) {
      const h = headers[colIndex]
      const type = h.type.toLowerCase()

      if (type === 'image') {
        const value = row[h.key]
        if (value && typeof value === 'string' && value.trim() !== '') {
          try {
            let base64Data = value
            let imageFormat = 'png' // 默认格式

            // 如果包含 data:image 前缀,提取 base64 数据和格式
            if (value.startsWith('data:image')) {
              base64Data = value.split(',')[1]
              imageFormat = value.split(';')[0].split('/')[1]
            } else {
              // 纯 base64 数据,尝试从数据头部推断格式
              // PNG: iVBORw0KGgo
              // JPEG: /9j/
              // GIF: R0lGODlh or R0lGODdh
              if (base64Data.startsWith('iVBORw0KGgo')) {
                imageFormat = 'png'
              } else if (base64Data.startsWith('/9j/')) {
                imageFormat = 'jpeg'
              } else if (base64Data.startsWith('R0lGOD')) {
                imageFormat = 'gif'
              }
            }

            // 添加图片到工作簿
            const imageId = workbook.addImage({
              base64: base64Data,
              extension: imageFormat
            })

            // 将图片插入到单元格
            worksheet.addImage(imageId, {
              tl: { col: colIndex, row: rowIndex + 1 }, // +1 因为表头占第一行
              ext: { width: 100, height: 75 }
            })
          } catch (error) {
            console.error('Failed to add image:', error)
          }
        }
      }
    }
  }

  // 导出文件
  const buffer = await workbook.xlsx.writeBuffer()
  const blob = new Blob([buffer], { type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet' })
  const url = window.URL.createObjectURL(blob)
  const link = document.createElement('a')
  link.href = url
  link.download = `${filename}.xlsx`
  link.click()
  window.URL.revokeObjectURL(url)

  // 关闭进度对话框
  dialogVisible.value = false
  ElMessage.success(t('common.exportSuccess'))
}

// 使用 XLSX 导出不含图片的 Excel
const exportToExcelSimple = async(data, filename, headers) => {
  // 创建工作表数据，第一行是key，第二行是label，第三行是type
  const wsData = [
    headers.map((h) => h.key), // 第一行是key
    headers.map((h) => h.label), // 第二行是label
    headers.map((h) => h.type), // 第三行是type
    ...data.map((row) =>
      headers.map((h) => {
        // 获取数据值
        const value = row[h.key]
        const type = h.type.toLowerCase()

        // 处理不同类型的值
        if (value === null || value === undefined) {
          return ''
        }

        // 处理不同类型的格式化
        switch (type) {
          case 'boolean':
          case 'bool':
            return value ? 'true' : 'false'
          case 'date':
            return formatDateTime(value)
          case 'datetime':
            return formatDateTime(value)
          case 'number':
          case 'int':
          case 'integer':
          case 'int64':
          case 'int32':
          case 'int16':
          case 'uint':
            if (typeof value === 'number') {
              return Math.round(value)
            } else if (typeof value === 'string' && !isNaN(parseInt(value, 10))) {
              return parseInt(value, 10)
            }
            return value
          case 'float':
          case 'float32':
          case 'float64':
          case 'double':
          case 'decimal':
            // 确保浮点数格式正确
            if (typeof value === 'number') {
              return value
            } else if (
              typeof value === 'string' &&
                !isNaN(parseFloat(value))
            ) {
              return parseFloat(value)
            }
            return value
          case 'json':
            return JSON.stringify(value)
          default:
            return value
        }
      })
    )
  ]

  // 创建工作簿
  const wb = XLSX.utils.book_new()
  const ws = XLSX.utils.aoa_to_sheet(wsData)

  // 添加工作表到工作簿
  XLSX.utils.book_append_sheet(wb, ws, 'Sheet1')

  // 保存文件
  XLSX.writeFile(wb, `${filename}.xlsx`)

  // 关闭进度对话框
  dialogVisible.value = false
  ElMessage.success(t('common.exportSuccess'))
}

// 导出 Excel
const exportToExcel = async(data, filename) => {
  const filterCache = {}
  for (const col of props.columns) {
    if (col.filterId) {
      filterCache[col.filterId + ''] = await requestWithCache(col.filterId)
    }
  }
  console.log('filterCache', filterCache)
  // 转换数据
  const headers = props.columns
    .filter((col) => col.isShow !== false)
    .sort((a, b) => a.sort - b.sort)
    .map((col) => {
      // 尝试从不同的属性中获取key
      const key =
          col.jsonName ||
          col.dataIndex ||
          col.prop ||
          col.field ||
          col.columnName ||
          ''

      // 尝试获取标签，兼容国际化和直接的标签
      let label
      if (col.structName && col.jsonName) {
        try {
          label = t('business.' + col.structName + '.columns.' + col.jsonName)
        } catch {
          label = col.label || col.title || col.columnName || key
        }
      } else {
        label = col.label || col.title || col.columnName || key
      }

      // 获取类型，如果没有则使用"string"
      const type = col.type || col.dataType || 'string'

      return { key, label, type, structName: col.structName, columnName: col.columnName, filterId: col.filterId }
    })

  console.log('Export headers:', headers)

  // 检查是否有图片类型的列
  const hasImageColumn = headers.some((h) => h.type.toLowerCase() === 'image')

  // 根据是否有图片选择不同的导出方式
  if (hasImageColumn) {
    await exportToExcelWithImages(data, filename, headers)
  } else {
    await exportToExcelSimple(data, filename, headers)
  }
}

// 导出全部数据
const exportAllData = async() => {
  dialogVisible.value = true
  exportProgress.value = 0

  const result = await batchExport(
    props.fetchFunction,
    {},
    props.batchSize,
    updateProgress
  )

  if (result.success) {
    exportToExcel(result.data, props.fileName)
  } else {
    ElMessage.error(result.error || t('common.exportFailed'))
    dialogVisible.value = false
  }
}

// 导出当前条件的数据
const exportCurrentData = async() => {
  dialogVisible.value = true
  exportProgress.value = 0

  // 获取当前查询条件下的所有数据
  const result = await batchExport(
    props.fetchFunction,
    props.searchInfo,
    props.batchSize,
    updateProgress
  )

  if (result.success) {
    exportToExcel(result.data, props.fileName)
  } else {
    ElMessage.error(result.error || t('common.exportFailed'))
    dialogVisible.value = false
  }
}

// 导出模板
const exportTemplate = () => {
  const templateData = generateTemplateData(props.columns)

  // 创建工作簿
  const wb = XLSX.utils.book_new()
  const ws = XLSX.utils.aoa_to_sheet(templateData)

  // 添加工作表到工作簿
  XLSX.utils.book_append_sheet(wb, ws, 'Sheet1')

  // 保存文件
  XLSX.writeFile(wb, `${props.fileName}_template.xlsx`)

  ElMessage.success(t('common.exportTemplateSuccess'))
}

// 处理下拉菜单命令
const handleCommand = (command) => {
  switch (command) {
    case 'exportAll':
      exportAllData()
      break
    case 'exportCurrent':
      exportCurrentData()
      break
    case 'exportTemplate':
      exportTemplate()
      break
  }
}
</script>

<style scoped>
  .export-progress {
    text-align: center;
  }
</style>
