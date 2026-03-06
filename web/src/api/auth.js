import service from '@/utils/request'

// 新的认证API - 按设计规范

/**
 * 查询用户安全状态
 * @param {string} username - 用户名
 * @returns {Promise}
 */
export const securityState = (data) => {
  return service({
    url: '/auth/security-state',
    method: 'post',
    data
  })
}

/**
 * 密码验证
 * @param {string} username - 用户名
 * @param {string} password - 密码
 * @returns {Promise}
 */
export const passwordVerify = (data) => {
  return service({
    url: '/auth/password/verify',
    method: 'post',
    data
  })
}

/**
 * TOTP登录
 * @param {string} username - 用户名
 * @param {string} code - TOTP验证码
 * @returns {Promise}
 */
export const totpLogin = (data) => {
  return service({
    url: '/auth/totp/login',
    method: 'post',
    data
  })
}

/**
 * 获取Passkey登录选项
 * @param {string} username - 用户名（可选）
 * @returns {Promise}
 */
export const passkeyAssertionOptions = (data) => {
  return service({
    url: '/auth/passkey/assertion/options',
    method: 'post',
    data
  })
}

/**
 * 验证Passkey登录
 * @param {Object} data - Passkey验证数据
 * @returns {Promise}
 */
export const passkeyAssertionVerify = (data) => {
  return service({
    url: '/auth/passkey/assertion/verify',
    method: 'post',
    data
  })
}

/**
 * 初始化TOTP绑定
 * @returns {Promise}
 */
export const totpBindInit = (data = {}) => {
  return service({
    url: '/auth/totp/bind/init',
    method: 'post',
    data
  })
}

/**
 * 验证TOTP绑定
 * @param {string} code - TOTP验证码
 * @param {string} secret - 密钥
 * @returns {Promise}
 */
export const totpBindVerify = (data) => {
  return service({
    url: '/auth/totp/bind/verify',
    method: 'post',
    data
  })
}

/**
 * 获取Passkey绑定选项
 * @param {string} displayName - 设备显示名称
 * @returns {Promise}
 */
export const passkeyAttestationOptions = (data) => {
  return service({
    url: '/auth/passkey/attestation/options',
    method: 'post',
    data
  })
}

/**
 * 验证Passkey绑定
 * @param {Object} data - Passkey绑定数据
 * @returns {Promise}
 */
export const passkeyAttestationVerify = (data) => {
  return service({
    url: '/auth/passkey/attestation/verify',
    method: 'post',
    data
  })
}

// 恢复码相关（预留）
export const generateRecoveryCodes = () => {
  return service({
    url: '/auth/recovery-codes/generate',
    method: 'post'
  })
}

export const consumeRecoveryCode = (data) => {
  return service({
    url: '/auth/recovery-codes/consume',
    method: 'post',
    data
  })
}

// 登出和刷新
export const logout = () => {
  return service({
    url: '/auth/logout',
    method: 'post'
  })
}

export const refreshToken = (data) => {
  return service({
    url: '/auth/token/refresh',
    method: 'post',
    data
  })
}