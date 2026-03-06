// 下载语言包
export function downloadLanguage(lang) {
  return request({
    url: '/system/translation/download',
    method: 'get',
    params: { language: lang }
  })
}

// 上传语言包
export function uploadLanguage(lang, file) {
  const formData = new FormData()
  formData.append('language', lang)
  formData.append('file', file)
  return request({
    url: '/system/translation/upload',
    method: 'post',
    data: formData,
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
} 