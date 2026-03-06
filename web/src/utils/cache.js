import { ref } from 'vue'

// 缓存存储
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
export const setCache = (key, value, expireTime = 5 * 60 * 1000) => {
  cache.set(key, {
    data: value,
    expireTime: Date.now() + expireTime
  })
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

/**
 * 清除缓存
 * @param {string} key 缓存键
 */
export const clearCache = (key) => {
  cache.delete(key)
}

/**
 * 带缓存的请求函数
 * @param {string} key 缓存键
 * @param {Function} requestFn 请求函数
 * @param {number} expireTime 缓存过期时间（毫秒）
 * @returns {Promise<any>} 请求结果
 */
export const requestWithCache = async(key, requestFn, expireTime = 5 * 60 * 1000) => {
  // 检查是否有缓存且未过期
  if (!isCacheExpired(key)) {
    return getCache(key).data
  }

  // 检查是否有正在进行的请求
  if (pendingRequests.has(key)) {
    return pendingRequests.get(key)
  }

  // 创建新的请求
  const requestPromise = requestFn().then(res => {
    if (res.code === 0) {
      setCache(key, res.data, expireTime)
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
