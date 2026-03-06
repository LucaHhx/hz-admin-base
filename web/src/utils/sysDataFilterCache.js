import { findSysDataFilter, filterData } from '@/api/system/sysDataFilter'
const cache = new Map()
// 正在进行的请求
const pendingRequests = new Map()
/**
 * 获取缓存数据
 * @param {string} key 缓存键
 * @returns {any} 缓存数据
 */
export const getCache = (key) => {
  return cache.get(key)
}

/**
   * 设置缓存数据
   * @param {string} key 缓存键
   * @param {any} value 缓存值
   * @param {number} expireTime 过期时间（毫秒）
   */
export const setCache = (key, value, expireTime = 10 * 60 * 1000) => {
  cache.set(key, {
    data: value,
    expireTime: Date.now() + expireTime
  })
}

/**
 * 清除缓存
 * @param {string} key 缓存键
 */
export const clearCache = (key) => {
  cache.delete(key)
}

/**
 * 检查缓存是否过期
 * @param {string} key 缓存键
 * @returns {boolean} 是否过期
 */
export const isCacheExpired = (key) => {
  const cacheData = cache.get(key)
  if (!cacheData) return true
  return Date.now() > cacheData.expireTime
}

export const requestWithCache = async(key) => {
  // 检查是否有缓存且未过期
  if (getCache(key)) {
    return getCache(key).data
  }
  // 检查是否有正在进行的请求
  if (pendingRequests.has(key)) {
    return pendingRequests.get(key)
  }

  // 创建新的请求
  const requestFn = async() => {
    const keyCache = {}
    const res = await findSysDataFilter({ ID: key })
    if (res.code === 0) {
      keyCache.filter = res.data
    } else {
      return nil
    }
    const res2 = await filterData({ id: key })
    if (res2.code === 0) {
      keyCache.data = res2.data
    } else {
      return nil
    }
    return keyCache
  }
  const requestPromise = requestFn().then(res => {
    if (res) {
      setCache(key, res)
    }
    return res
  }).finally(() => {
    // 请求完成后移除pending状态
    pendingRequests.delete(key)
  })
  // 保存请求Promise
  pendingRequests.set(key, requestPromise)

  return requestPromise
}

/**
 * 根据筛选条件过滤数据
 * @param {Array} data 原始数据
 * @param {Array} filters 筛选条件数组
 * @param {Array} filterColumns 可筛选的列
 * @returns {Array} 筛选后的数据
 */
export const filterDataByConditions = (data, filters, filterColumns) => {
  if (!data || !data.length || !filters || !filters.length) {
    return data
  }
  const list = data.filter(item => {
    // 对每个筛选条件进行匹配
    return filters.every(filter => {
      if (!filter) return true
      const trimFilter = (filter + '').trim()
      if (!trimFilter) return true
      // 检查是否在任何可筛选列中匹配
      return filterColumns.some(column => {
        if (!column.filter) return false
        const value = item[column.columnName]
        if (value === undefined || value === null) return false
        return String(value).toLowerCase() === trimFilter.toLowerCase()
      })
    })
  })
  console.log('list', list)

  return list
}

export const requestFilterData = async(filters, key) => {
  const res = await requestWithCache(key)
  const filterColumns = res.filter.columns.filter(col => col.filter)
  const filteredData = filterDataByConditions(res.data, filters, filterColumns)
  return {
    data: filteredData,
    filter: res.filter
  }
}

export const requestFilterLabel = (id, key) => {
  const res = getCache(id)
  console.log('res', res)

  if (res.data.find(a => a[res.filter.key] + '' === key + '')) {
    return res.data.find(a => a[res.filter.key] === key)[res.filter.label]
  } else {
    return key
  }
}

export const getCacheList = async(ids) => {
  const list = []
  for (const id of ids) {
    const res = await requestWithCache(id)
    list.push(res)
  }
  return list
}

//   const res = await requestWithCache(id)
//   //   console.log(id, key)
//   //   console.log(res.data.find(a => a[res.filter.key] === key)[res.filter.label])
