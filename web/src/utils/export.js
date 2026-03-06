// import { useI18n } from 'vue-i18n'
// export const { t } = useI18n()
// console.log('window.i18n', t)
// const { t } = useI18n()
import { getMessagesInfo } from '@/utils/lolocale'
/**
 * 批量导出数据工具函数
 */

/**
 * 分批导出数据
 * @param {Function} fetchFunction - 获取数据的API函数
 * @param {Object} searchInfo - 搜索参数
 * @param {Number} batchSize - 每批次的数据量，默认1000
 * @param {Function} onProgress - 进度回调函数
 * @param {Function} onComplete - 导出完成回调函数
 * @param {Function} processData - 处理每批数据的函数
 */
export const batchExport = async(
  fetchFunction,
  searchInfo = {},
  batchSize = 1000,
  onProgress = null,
  onComplete = null,
  processData = null
) => {
  try {
    // 克隆搜索条件，避免修改原对象
    const searchParams = JSON.parse(JSON.stringify(searchInfo))

    // 确保有页码信息
    if (!searchParams.pageInfo) {
      searchParams.pageInfo = {
        page: 1,
        pageSize: batchSize
      }
    } else {
      // 保留原始页码信息，但修改每页数量
      searchParams.pageInfo.pageSize = batchSize
    }

    // 第一次请求获取总数
    const firstBatch = await fetchFunction(searchParams)
    if (firstBatch.code !== 0) {
      throw new Error(firstBatch.msg || '获取数据失败')
    }

    // 计算总页数
    const total = firstBatch.data.total
    const totalPages = Math.ceil(total / batchSize)

    // 存储所有数据
    let allData = [...firstBatch.data.list]

    // 处理第一批数据
    if (processData) {
      processData(firstBatch.data.list)
    }

    // 报告进度
    if (onProgress) {
      onProgress(1, totalPages, allData.length, total)
    }

    // 如果有多页，继续获取其余页面
    for (let page = 2; page <= totalPages; page++) {
      searchParams.pageInfo.page = page
      const result = await fetchFunction(searchParams)

      if (result.code === 0) {
        allData = [...allData, ...result.data.list]

        // 处理每批数据
        if (processData) {
          processData(result.data.list)
        }

        // 报告进度
        if (onProgress) {
          onProgress(page, totalPages, allData.length, total)
        }
      } else {
        throw new Error(result.msg || `获取第${page}页数据失败`)
      }
    }

    // 完成回调
    if (onComplete) {
      onComplete(allData, total)
    }

    return {
      success: true,
      data: allData,
      total: total
    }
  } catch (error) {
    console.error('批量导出数据失败:', error)
    return {
      success: false,
      error: error.message || '导出失败'
    }
  }
}

/**
 * 根据columns生成Excel模板数据
 * @param {Array} columns - 列定义
 * @returns {Array} - 模板数据
 */
export const generateTemplateData = (columns) => {
  // 获取表头信息
  const filteredColumns = columns
    .filter(col => col.isShow !== false)
    .sort((a, b) => a.sort - b.sort)

  // 第一行是key
  const keys = filteredColumns.map(col =>
    col.jsonName || col.dataIndex || col.prop || col.field || col.columnName || '')

  // 第二行是label
  const labels = filteredColumns.map(col => {
    const msgLabel = getMessagesInfo('business.' + col.structName + '.columns.' + col.jsonName)
    // 尝试获取标签，兼容国际化和直接的标签
    if (col.structName && col.jsonName && msgLabel) {
      try {
        return msgLabel
      } catch {
        return col.label || col.title || col.columnName || ''
      }
    } else {
      return col.label || col.title || col.columnName || ''
    }
  })

  // 第三行是type
  const types = filteredColumns.map(col => col.type || col.dataType || 'string')

  // 第四行是格式提示（用于导入模板）
  const formatHints = types.map(type => {
    const lowerType = (type || '').toLowerCase()

    if (lowerType === 'date') {
      return '日期格式: yyyy-MM-dd'
    } else if (lowerType === 'datetime') {
      return '日期时间格式: yyyy-MM-dd HH:mm:ss'
    } else if (lowerType === 'boolean' || lowerType === 'bool') {
      return '布尔值: true/false 或 是/否'
    } else if (
      lowerType.includes('int') ||
      lowerType === 'number' ||
      lowerType === 'uint' ||
      lowerType === 'enum'
    ) {
      return '整数'
    } else if (
      lowerType === 'float' ||
      lowerType === 'float32' ||
      lowerType === 'float64' ||
      lowerType === 'double' ||
      lowerType === 'decimal'
    ) {
      return '小数'
    } else {
      return ''
    }
  })

  return [keys, labels, types, formatHints]
}
