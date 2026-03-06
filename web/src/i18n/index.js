import { createI18n } from 'vue-i18n'
import axios from 'axios'
const defaultLocale = 'zh-CN'

import { useAppStore } from '@/pinia'
import { ref } from 'vue'

const timeZone = ref('')
// 从后端加载翻译内容
const loadMessages = async() => {
  try {
    const appStore = useAppStore()
    const response = await axios.get('/api/translation/messages')

    appStore.setMessages(response.data.data.messages)
    timeZone.value = response.data.data.timeZone
    return appStore.getMessages()
  } catch (error) {
    console.error('加载语言包时发生错误:', error)
    return {}
  }
}
export const setupI18n = async(app) => {
  // 加载翻译内容
  const messages = await loadMessages()
  const i18n = createI18n({
    legacy: false, // 使用 Composition API 模式
    locale: defaultLocale, // 使用环境变量中的默认语言
    fallbackLocale: 'zh-CN', // 默认语言
    messages,
    silentTranslationWarn: true, // 禁用翻译警告
    silentFallbackWarn: true, // 禁用回退警告
    missingWarn: false, // 禁用缺失警告
    fallbackWarn: false, // 禁用回退警告
    allowComposition: true, // 允许组合式API
    runtimeOnly: true, // 仅运行时
    datetimeFormats: {
      'zh-CN': {
        short: {
          year: 'numeric',
          month: '2-digit',
          day: '2-digit',
          timeZone: timeZone.value
        },
        medium: {
          year: 'numeric',
          month: '2-digit',
          day: '2-digit',
          hour: '2-digit',
          minute: '2-digit',
          timeZone: timeZone.value
        },
        long: {
          year: 'numeric',
          month: '2-digit',
          day: '2-digit',
          hour: '2-digit',
          minute: '2-digit',
          second: '2-digit',
          timeZone: timeZone.value
        },
        full: {
          year: 'numeric',
          month: 'long',
          day: 'numeric',
          weekday: 'long',
          hour: '2-digit',
          minute: '2-digit',
          second: '2-digit',
          timeZone: timeZone.value
        },
        date: {
          year: 'numeric',
          month: '2-digit',
          day: '2-digit',
          timeZone: timeZone.value
        }
      },
      'en-US': {
        short: {
          year: 'numeric',
          month: '2-digit',
          day: '2-digit',
          timeZone: timeZone.value
        },
        medium: {
          year: 'numeric',
          month: '2-digit',
          day: '2-digit',
          hour: '2-digit',
          minute: '2-digit',
          timeZone: timeZone.value
        },
        long: {
          year: 'numeric',
          month: '2-digit',
          day: '2-digit',
          hour: '2-digit',
          minute: '2-digit',
          second: '2-digit',
          timeZone: timeZone.value
        },
        full: {
          year: 'numeric',
          month: 'long',
          day: 'numeric',
          weekday: 'long',
          hour: '2-digit',
          minute: '2-digit',
          second: '2-digit',
          timeZone: timeZone.value
        },
        date: {
          year: 'numeric',
          month: '2-digit',
          day: '2-digit',
          timeZone: timeZone.value
        }
      }
    },
    getMessagesInfo: () => {
      return messages
    }
  })

  // 扩展全局 $d 方法
  app.config.globalProperties.$serverDate = (date, format = 'long', timeZone = 'Asia/Bangkok') => {
    if (!date) return ''
    try {
      let dateObj
      switch (typeof date) {
        case 'number':
          if (date.toString().length === 8) {
            dateObj = new Date(date.toString().slice(0, 4) + '-' + date.toString().slice(4, 6) + '-' + date.toString().slice(6, 8))
          } else {
            dateObj = new Date(date * 1000)
          }
          break
        case 'string':
          dateObj = new Date(date)
          break
        default:
          dateObj = new Date(date)
      }
      // console.log(date, dateObj)

      const newData = new Date(dateObj.toISOString())
      return i18n.global.d(newData, format)
      // return dateObj.toISOString()
    } catch (error) {
      console.error('日期格式化错误:', error)
      return ''
    }
  }

  app.use(i18n)
  i18n.global.timeZone = timeZone.value // 设置时区
  return i18n
}

export default setupI18n
