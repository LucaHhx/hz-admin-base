<template>
  <div
    id="app"
    class="bg-gray-50 text-slate-700 dark:text-slate-500 dark:bg-slate-800"
  >
    <ConfigProvider
      :theme="{
        algorithm: appStore.isDark ? darkAlgorithm : defaultAlgorithm,
        token: {
          // colorBgContainer: 'var(--el-bg-color)',
          colorBgElevated: 'var(--el-bg-color-overlay)',
          colorText: 'var(--el-text-color-primary)',
          colorTextSecondary: 'var(--el-text-color-regular)',
          colorBorder: 'var(--el-border-color)',
          colorPrimary: 'var(--el-color-primary)',
          colorPrimaryHover: 'var(--el-color-primary-light-3)',
          colorPrimaryActive: 'var(--el-color-primary-dark-2)',
          colorFillSecondary: 'var(--el-fill-color-light)',
          colorBgContainerDisabled: 'var(--el-disabled-bg-color)',
          colorTextDisabled: 'var(--el-text-color-disabled)',
          borderRadius: 4,
          wireframe: true
        }
      }"
    >
      <el-config-provider :locale="locale">
        <router-view />
      </el-config-provider>
    </ConfigProvider>
  </div>
</template>

<script setup>
import { ref, watchEffect } from 'vue'
import { useAppStore } from '@/pinia'
import { useI18n } from 'vue-i18n'
import zhCn from 'element-plus/dist/locale/zh-cn.mjs'
import enUs from 'element-plus/dist/locale/en.mjs'
import { ConfigProvider } from 'ant-design-vue'
import { theme } from 'ant-design-vue'
const { darkAlgorithm, defaultAlgorithm } = theme
const appStore = useAppStore()
const { locale: i18nLocale } = useI18n()
const locale = ref(zhCn)

watchEffect(() => {
  const lang = appStore.config.language
  i18nLocale.value = lang
  locale.value = lang === 'zh-CN' ? zhCn : enUs
})

defineOptions({
  name: 'App'
})
</script>
<style lang="scss">
  // 引入初始化样式
  #app {
    height: 100vh;
    overflow: hidden;
    font-weight: 400 !important;
  }
  .el-button {
    font-weight: 400 !important;
  }

  .gva-body-h {
    min-height: calc(100% - 3rem);
  }

  .gva-container {
    height: calc(100% - 2.5rem);
  }
  .gva-container2 {
    height: calc(100% - 4.5rem);
  }

  // .search-tooltip {
  //   background-color: #0d1729 !important;
  //   border-color: #363637 !important;
  // }
</style>
