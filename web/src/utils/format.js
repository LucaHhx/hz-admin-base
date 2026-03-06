import { formatTimeToStr } from '@/utils/date'
import { getDict } from '@/utils/dictionary'
import { ref } from 'vue'
import numeral from 'numeral'
import Decimal from 'decimal.js'

export const formatBoolean = (bool) => {
  if (bool !== null) {
    return bool ? '是' : '否'
  } else {
    return ''
  }
}

export const formatDate = (time) => {
  if (time !== null && time !== '') {
    var date = new Date(time)
    return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss')
  } else {
    return ''
  }
}

export const filterDict = (value, options) => {
  const rowLabel = options && options.filter((item) => item.value === value)
  return rowLabel && rowLabel[0] && rowLabel[0].label
}

export const filterDataSource = (dataSource, value) => {
  if (Array.isArray(value)) {
    return value.map((item) => {
      const rowLabel = dataSource && dataSource.find((i) => i.value === item)
      return rowLabel?.label
    })
  }
  const rowLabel = dataSource && dataSource.find((item) => item.value === value)
  return rowLabel?.label
}

export const getDictFunc = async(type) => {
  const dicts = await getDict(type)
  return dicts
}

const path =
  import.meta.env.VITE_BASE_PATH + ':' + import.meta.env.VITE_SERVER_PORT + '/'
export const ReturnArrImg = (arr) => {
  const imgArr = []
  if (arr instanceof Array) {
    // 如果是数组类型
    for (const arrKey in arr) {
      if (arr[arrKey].slice(0, 4) !== 'http') {
        imgArr.push(path + arr[arrKey])
      } else {
        imgArr.push(arr[arrKey])
      }
    }
  } else {
    // 如果不是数组类型
    if (arr?.slice(0, 4) !== 'http') {
      imgArr.push(path + arr)
    } else {
      imgArr.push(arr)
    }
  }
  return imgArr
}

export const returnArrImg = ReturnArrImg

export const onDownloadFile = (url) => {
  window.open(path + url)
}
const colorToHex = (u) => {
  const e = u.replace('#', '').match(/../g)
  for (let t = 0; t < 3; t++) e[t] = parseInt(e[t], 16)
  return e
}

const hexToColor = (u, e, t) => {
  const a = [u.toString(16), e.toString(16), t.toString(16)]
  for (let n = 0; n < 3; n++) a[n].length === 1 && (a[n] = `0${a[n]}`)
  return `#${a.join('')}`
}
const generateAllColors = (u, e) => {
  const t = colorToHex(u)
  const target = [10, 10, 30]
  for (let a = 0; a < 3; a++) t[a] = Math.floor(t[a] * (1 - e) + target[a] * e)
  return hexToColor(t[0], t[1], t[2])
}

const generateAllLightColors = (u, e) => {
  const t = colorToHex(u)
  const target = [240, 248, 255] // RGB for blue white color
  for (let a = 0; a < 3; a++) t[a] = Math.floor(t[a] * (1 - e) + target[a] * e)
  return hexToColor(t[0], t[1], t[2])
}

function addOpacityToColor(u, opacity) {
  const t = colorToHex(u)
  return `rgba(${t[0]}, ${t[1]}, ${t[2]}, ${opacity})`
}

export const setBodyPrimaryColor = (primaryColor, darkMode) => {
  let fmtColorFunc = generateAllColors
  if (darkMode === 'light') {
    fmtColorFunc = generateAllLightColors
  }

  document.documentElement.style.setProperty('--el-color-primary', primaryColor)
  document.documentElement.style.setProperty(
    '--el-color-primary-bg',
    addOpacityToColor(primaryColor, 0.4)
  )
  for (let times = 1; times <= 2; times++) {
    document.documentElement.style.setProperty(
      `--el-color-primary-dark-${times}`,
      fmtColorFunc(primaryColor, times / 10)
    )
  }
  for (let times = 1; times <= 10; times++) {
    document.documentElement.style.setProperty(
      `--el-color-primary-light-${times}`,
      fmtColorFunc(primaryColor, times / 10)
    )
  }
  document.documentElement.style.setProperty(
    `--el-menu-hover-bg-color`,
    addOpacityToColor(primaryColor, 0.2)
  )
}

const baseUrl = ref(import.meta.env.VITE_BASE_API)

export const getBaseUrl = () => {
  return baseUrl.value === '/' ? '' : baseUrl.value
}

export const CreateUUID = () => {
  let d = new Date().getTime()
  if (window.performance && typeof window.performance.now === 'function') {
    d += performance.now()
  }
  return '00000000-0000-0000-0000-000000000000'.replace(/0/g, (c) => {
    const r = (d + Math.random() * 16) % 16 | 0 // d是随机种子
    d = Math.floor(d / 16)
    return (c === '0' ? r : (r & 0x3) | 0x8).toString(16)
  })
}

/**
 * 根据指定格式模板格式化字符串
 * @param {string} string - 要格式化的字符串
 * @param {string} format - 格式模板，使用{%s}作为字符串占位符，{i}作为索引占位符
 * @returns {string} 格式化后的字符串
 *
 * 示例:
 * formatString('hello', '{s}!') => 'hello!'
 * formatString('hello', '#{i}: {s}') => '#0: hello'
 * formatString(['a', 'b'], '{s}') => 'a,b'
 * formatString(['a', 'b'], '{0},{1}') => 'a,b'
 */
export const formatString = (string, format) => {
  if (!format) {
    format = '{%s}'
  }
  if (string === null || string === undefined) {
    return ''
  }

  // 替换{s}占位符为实际字符串
  let result = format.replace('{%s}', string)
  // 替换{i}占位符为索引0（适用于单个字符串）
  result = result.replace('{%s}', '0')
  return result
}

// 格式化数字为带千分位逗号和小数点的格式
export const formatNumber = (number, format) => {
  if (!format) {
    format = '0.[00]'
  }
  if (number === null || number === undefined || isNaN(number)) {
    return '0.00'
  }
  try {
    if (!format.split) {
      return numeral(number).format(format)
    }
    const split = format.split(':')
    if (!split || split.length === 1) {
      return numeral(number).format(format)
    }
    var [prefix, suffix] = split
    var value = new Decimal(number)

    // 处理前缀中的数学运算
    if (prefix !== null && prefix !== undefined && prefix !== '') {
      // 处理除法运算 /100 等
      if (prefix.startsWith('/')) {
        const divisor = new Decimal(prefix.substring(1))
        if (!divisor.isZero()) {
          value = value.div(divisor)
        }
      } else if (prefix.startsWith('*')) { // 处理乘法运算 *100 等
        const multiplier = new Decimal(prefix.substring(1))
        value = value.mul(multiplier)
      } else if (prefix.startsWith('+')) { // 处理加法运算 +100 等
        const addend = new Decimal(prefix.substring(1))
        value = value.plus(addend)
      } else if (prefix.startsWith('-')) { // 处理减法运算 -100 等
        const subtrahend = new Decimal(prefix.substring(1))
        value = value.minus(subtrahend)
      }
    }

    if (suffix === null || suffix === undefined) {
      suffix = '0.[00]'
    }
    // console.log(value.toNumber(), numeral(value.toNumber()).format(suffix))

    return numeral(value.toNumber()).format(suffix)
  } catch (error) {
    console.error('Error in formatNumber:', error)
    return numeral(number).format(format)
  }
}

/**
 * 根据格式化后的值和format反向计算原始值
 * @param {string|number} formattedValue - 格式化后的值（支持带千分位和不带千分位的数值）
 * @param {string} format - 原始format字符串
 * @returns {string} 计算后的原始值（字符串格式）
 */
export const unformatNumber = (formattedValue, format) => {
  if (!format) {
    return formattedValue?.toString() || '0'
  }

  try {
    if (typeof formattedValue === 'number') {
      return formattedValue.toString()
    }

    const cleanValue = formattedValue.toString().replace(/,/g, '')
    let value = new Decimal(cleanValue)

    if (value.isNaN()) {
      return '0'
    }

    const prefix = format.split(':')[0]

    if (!prefix || !['/', '*', '+', '-'].some(op => prefix.startsWith(op))) {
      return value.toString()
    }

    if (prefix.startsWith('/')) {
      const divisor = new Decimal(prefix.substring(1))
      if (!divisor.isZero()) {
        value = value.mul(divisor)
      }
    } else if (prefix.startsWith('*')) {
      const multiplier = new Decimal(prefix.substring(1))
      if (!multiplier.isZero()) {
        value = value.div(multiplier)
      }
    } else if (prefix.startsWith('+')) {
      const addend = new Decimal(prefix.substring(1))
      value = value.minus(addend)
    } else if (prefix.startsWith('-')) {
      const subtrahend = new Decimal(prefix.substring(1))
      value = value.plus(subtrahend)
    }

    return value.toString()
  } catch (error) {
    console.error('Error in unformatNumber:', error)
    return formattedValue?.toString() || '0'
  }
}
