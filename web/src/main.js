import './style/element_visiable.scss'
import 'element-plus/theme-chalk/dark/css-vars.css'
import { createApp } from 'vue'
import ElementPlus from 'element-plus'

import 'element-plus/dist/index.css'
// 引入gin-vue-admin前端初始化相关内容
import './core/gin-vue-admin'
// 引入封装的router
import router from '@/router/index'
import '@/permission'
import run from '@/core/gin-vue-admin.js'
import auth from '@/directive/auth'
import { store } from '@/pinia'
import App from './App.vue'
import './styles/dialog-passthrough.css'

import setupI18n from './i18n'
import { setI18n } from '@/utils/request'

const initApp = async() => {
  const app = createApp(App)
  app.config.productionTip = false

  app.use(run)
    .use(ElementPlus)
    .use(store)
    .use(auth)
    .use(router)

  // 初始化 i18n
  const i18n = await setupI18n(app)
  setI18n(i18n)
  app.use(i18n)

  setInterval(() => {
    const e = document.querySelector('div[style*="position: fixed; bottom: 10px; right: 10px;"]')
    e && (e.style.display = 'none')
  }, 100)

  app.mount('#app')
  return app
}

initApp()
export default initApp
