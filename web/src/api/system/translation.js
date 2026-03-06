import service from '@/utils/request'

// 获取翻译文件树
export const getFileTree = () => {
  return service({
    url: '/translation/tree',
    method: 'get'
  })
}

// 获取单个文件内容
export const getFileContent = (language, file) => {
  return service({
    url: '/translation/file',
    method: 'get',
    params: {
      language,
      file
    }
  })
}

// 获取翻译列表
export const getTranslationList = () => {
  return service({
    url: '/translation/list',
    method: 'get'
  })
}

// 更新翻译
export const updateTranslation = (data) => {
  return service({
    url: '/translation/update',
    method: 'post',
    data
  })
}

// 复制语言包
export const copyLanguage = (data) => {
  return service({
    url: '/translation/copy',
    method: 'post',
    data
  })
}

// 批量更新翻译
export const batchUpdate = (data) => {
  return service({
    url: '/translation/batch',
    method: 'post',
    data
  })
}

// 对比翻译
export const compareTranslation = (data) => {
  return service({
    url: '/translation/compare',
    method: 'post',
    data
  })
}

// 导出翻译
export const exportTranslation = (language) => {
  return service({
    url: `/translation/export?language=${language}`,
    method: 'get',
    responseType: 'blob'
  })
}

// 导入翻译
export const importTranslation = (language, file) => {
  const formData = new FormData()
  formData.append('language', language)
  formData.append('file', file)
  return service({
    url: '/translation/import',
    method: 'post',
    data: formData,
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
}

// 检查完整性
export const checkIntegrity = () => {
  return service({
    url: '/translation/check',
    method: 'get'
  })
}

// 导出语言包
export const exportLanguage = (data) => {
  return service({
    url: '/translation/export-language',
    method: 'post',
    data,
    responseType: 'blob'
  })
}

// 删除语言包
export const deleteLanguage = (data) => {
  return service({
    url: '/translation/delete-language',
    method: 'post',
    data
  })
}

// 下载语言包
export const downloadLanguage = (language) => {
  return service({
    url: `/translation/download-language?language=${language}`,
    method: 'get',
    responseType: 'blob'
  })
}

// 上传语言包
export const uploadLanguage = (language, file) => {
  const formData = new FormData()
  formData.append('language', language)
  formData.append('file', file)
  return service({
    url: '/translation/upload-language',
    method: 'post',
    data: formData,
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
}

// 获取翻译消息
export const getMessages = () => {
  return service({
    url: '/translation/messages',
    method: 'get',
    config: {
      donNotShowLoading: true
    }
  })
}
