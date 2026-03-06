import { useAppStore } from '@/pinia'

export const getMessagesInfo = (key) => {
  try {
    const appStore = useAppStore()
    const messages = appStore.getMessages()

    const currentLanguage = appStore.config.language

    // 检查 messages 和 language 是否已加载
    if (!messages || !currentLanguage || !messages[currentLanguage]) {
      return null
    }

    // 从messages中获取当前语言的翻译对象
    let result = messages[currentLanguage]

    // 如果key为空，直接返回null
    if (!key) {
      return null
    }

    // 将键按照点号分割成数组并逐层访问对象
    const keys = key.split('.')
    for (const k of keys) {
      if (result && typeof result === 'object') {
        result = result[k]
      } else {
        return undefined // 如果中间某层不存在，返回undefined
      }
    }

    return result
  } catch (error) {
    console.error('Error in getMessagesInfo:', error)
    return null
  }
}
