import axios from 'axios' // 引入axios
import { ElMessage, ElMessageBox } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'
import router from '@/router/index'
import { ElLoading } from 'element-plus'

let i18nInstance = null

export const setI18n = (i18n) => {
  i18nInstance = i18n
}

const service = axios.create({
  baseURL: import.meta.env.VITE_BASE_API,
  timeout: 99999
})
let activeAxios = 0
let timer
let loadingInstance
const showLoading = (
  option = {
    target: null
  }
) => {
  const loadDom = document.getElementById('hab-base-load-dom')
  activeAxios++
  if (timer) {
    clearTimeout(timer)
  }
  timer = setTimeout(() => {
    if (activeAxios > 0) {
      if (!option.target) option.target = loadDom
      loadingInstance = ElLoading.service(option)
    }
  }, 400)
}

const closeLoading = () => {
  activeAxios--
  if (activeAxios <= 0) {
    clearTimeout(timer)
    loadingInstance && loadingInstance.close()
  }
}
// http request 拦截器
service.interceptors.request.use(
  (config) => {
    if (!config.donNotShowLoading) {
      showLoading(config.loadingOption)
    }
    const userStore = useUserStore()
    config.headers = {
      'Content-Type': 'application/json',
      'x-token': userStore.token,
      'x-user-id': userStore.userInfo.ID,
      ...config.headers
    }
    return config
  },
  (error) => {
    if (!error?.config?.donNotShowLoading) {
      closeLoading()
    }
    ElMessage({
      showClose: true,
      message: error,
      type: 'error'
    })
    return error
  }
)

// http response 拦截器
service.interceptors.response.use(
  (response) => {
    const userStore = useUserStore()
    if (!response.config.donNotShowLoading) {
      closeLoading()
    }
    if (response.headers['new-token']) {
      userStore.setToken(response.headers['new-token'])
    }
    if (response.data.code === undefined || response.data.code === 0 || response.headers.success === 'true') {
      if (response.headers.msg) {
        response.data.msg = decodeURI(response.headers.msg)
      }
      return response.data
    } else {
      let msg = i18nInstance?.global.t('error.' + response.data.code + '.' + response.data.msg)
      if (msg === 'error.' + response.data.code + '.error') {
        msg = response.data.data || response.data.msg
        console.log(response.data.data, response.data.msg, msg)
      }
      // console.log(response.data, i18nInstance?.global.t('error.' + response.data.code + '.' + response.data.msg))
      ElMessage({
        showClose: true,
        message: msg,
        type: 'error'
      })
      return response.data.msg ? response.data : response
    }
  },
  (error) => {
    if (!error?.config?.donNotShowLoading) {
      closeLoading()
    }

    if (!error.response) {
      ElMessageBox.confirm(
        `
        <p>${i18nInstance?.global?.t('common.request.detectError')}</p>
        <p>${error}</p>
        `,
        i18nInstance?.global?.t('common.request.apiError'),
        {
          dangerouslyUseHTMLString: true,
          distinguishCancelAndClose: true,
          confirmButtonText: i18nInstance?.global?.t('common.request.tryLater'),
          cancelButtonText: i18nInstance?.global?.t('cancel')
        }
      )
      return
    }

    switch (error.response.status) {
      case 500:
        ElMessageBox.confirm(
          `
        <p>${error}</p>
        <p>${i18nInstance.global.t('common.request.error500')}</p>
        `,
          i18nInstance.global.t('common.request.apiError'),
          {
            dangerouslyUseHTMLString: true,
            distinguishCancelAndClose: true,
            confirmButtonText: i18nInstance.global.t('common.request.clearCache'),
            cancelButtonText: i18nInstance.global.t('cancel')
          }
        ).then(() => {
          const userStore = useUserStore()
          userStore.ClearStorage()
          router.push({ name: 'Login', replace: true })
        })
        break
      case 404:
        ElMessageBox.confirm(
          `
          <p>${i18nInstance.global.t('common.request.error404')}${error}</p>
          <p>${i18nInstance.global.t('common.request.error404Detail')}</p>
          `,
          i18nInstance.global.t('common.request.apiError'),
          {
            dangerouslyUseHTMLString: true,
            distinguishCancelAndClose: true,
            confirmButtonText: i18nInstance.global.t('common.request.iKnow'),
            cancelButtonText: i18nInstance.global.t('cancel')
          }
        )
        break
      case 401:
        ElMessageBox.confirm(
          `
          <p>${i18nInstance.global.t('common.request.invalidToken')}</p>
          <p>${i18nInstance.global.t('common.request.error401')}${error}</p>
          `,
          i18nInstance.global.t('common.request.identityInfo'),
          {
            dangerouslyUseHTMLString: true,
            distinguishCancelAndClose: true,
            confirmButtonText: i18nInstance.global.t('common.request.relogin'),
            cancelButtonText: i18nInstance.global.t('cancel')
          }
        ).then(() => {
          const userStore = useUserStore()
          userStore.ClearStorage()
          router.push({ name: 'Login', replace: true })
        })
        break
    }

    return error
  }
)
export default service
