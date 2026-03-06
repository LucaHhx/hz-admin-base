import { DateTime } from 'luxon'

const serverTimezone = 'UTC'

export const convertDateTime = (
  dateString,
) => {
  // 参数验证
  if (!dateString || typeof dateString !== 'string') {
    console.error('无效的日期字符串:', dateString)
    return null
  }

  try {
    console.log('dateString', dateString)

    // 解析输入的日期字符串
    const dt = DateTime.fromFormat(dateString.trim(), 'yyyy-MM-dd HH:mm:ss', { zone: serverTimezone })

    console.log('dt', dt)

    // 如果解析失败，返回原字符串
    if (!dt.isValid) {
      return undefined
    }

    // 根据输出格式决定返回类型
    return dt.setZone(serverTimezone).toISO()
  } catch (error) {
    console.error('日期转换错误:', error.message)
    return undefined
  }
}

export const convertDate = (
  dateString,
) => {
  // 参数验证
  if (!dateString || typeof dateString !== 'string') {
    console.error('无效的日期字符串:', dateString)
    return null
  }

  try {
    // 解析输入的日期字符串
    const dt = DateTime.fromFormat(dateString.trim(), 'yyyy-MM-d', { zone: serverTimezone })

    console.log('dt', dt)

    // 如果解析失败，返回原字符串
    if (!dt.isValid) {
      return undefined
    }

    // 根据输出格式决定返回类型
    return dt.setZone(serverTimezone).toISO()
  } catch (error) {
    console.error('日期转换错误:', error.message)
    return undefined
  }
}
