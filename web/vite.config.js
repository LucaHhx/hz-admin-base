import { viteLogo } from './src/core/config'
import Banner from 'vite-plugin-banner'
import * as path from 'path'
import * as dotenv from 'dotenv'
import * as fs from 'fs'
import vuePlugin from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'
import VueFilePathPlugin from './vitePlugin/componentName/index.js'
import { svgBuilder } from 'vite-auto-import-svg'
import { AddSecret } from './vitePlugin/secret'
// @see https://cn.vitejs.dev/config/
export default ({ mode }) => {
  AddSecret('')
  const NODE_ENV = mode || 'development'
  const envFiles = [`.env.${NODE_ENV}`]
  for (const file of envFiles) {
    const envConfig = dotenv.parse(fs.readFileSync(file))
    for (const k in envConfig) {
      process.env[k] = envConfig[k]
    }
  }

  viteLogo(process.env)

  const timestamp = Date.parse(new Date())

  const optimizeDeps = {
    include: ['vue-i18n']
  }

  const alias = {
    '@': path.resolve(__dirname, './src'),
    vue$: 'vue/dist/vue.runtime.esm-bundler.js',
    'vue-i18n': 'vue-i18n/dist/vue-i18n.runtime.esm-bundler.js'
  }

  const esbuild = {
    target: 'esnext',
    drop: ['console', 'debugger']
  }

  const rollupOptions = {
    output: {
      entryFileNames: 'assets/087AC4D233B64EB0[name].[hash].js',
      chunkFileNames: 'assets/087AC4D233B64EB0[name].[hash].js',
      assetFileNames: 'assets/087AC4D233B64EB0[name].[hash].[ext]'
    }
  }

  const config = {
    base: '/', // 编译后js导入的资源路径
    root: './', // index.html文件所在位置
    publicDir: 'public', // 静态资源文件夹
    resolve: {
      alias
    },
    define: {
      'process.env': {
        VITE_DEFAULT_LOCALE: process.env.VITE_DEFAULT_LOCALE
      },
      __VUE_I18N_FULL_INSTALL__: false,
      __VUE_I18N_LEGACY_API__: false,
      __INTLIFY_PROD_DEVTOOLS__: false
    },
    css: {
      preprocessorOptions: {
        scss: {
          api: 'modern-compiler' // or "modern"
        }
      }
    },
    server: {
      host: '0.0.0.0',
      allowedHosts: true,
      // 如果使用docker-compose开发模式，设置为false
      open: true,
      port: process.env.VITE_CLI_PORT,
      proxy: {
        // 把key的路径代理到target位置
        // detail: https://cli.vuejs.org/config/#devserver-proxy
        [process.env.VITE_BASE_API]: {
          // 需要代理的路径   例如 '/api'
          target: `${process.env.VITE_BASE_PATH}:${process.env.VITE_SERVER_PORT}/`, // 代理到 目标路径
          changeOrigin: true,
          rewrite: (path) =>
            path.replace(new RegExp('^' + process.env.VITE_BASE_API), '')
        }
      }
    },
    build: {
      target: 'esnext',
      minify: 'esbuild',
      manifest: false,
      sourcemap: false,
      outDir: 'dist',
      rollupOptions,
      copyPublicDir: true,
      assetsDir: 'assets'
    },
    esbuild,
    optimizeDeps,
    plugins: [
      process.env.VITE_POSITION === 'open' &&
        vueDevTools({ launchEditor: process.env.VITE_EDITOR }),
      vuePlugin(),
      svgBuilder(['./src/assets/icons/', './src/plugin/'], '/', 'dist', 'assets', mode),
      [Banner(`\n Build based on hab \n Time : ${timestamp}`)],
      VueFilePathPlugin('./src/pathInfo.json'),
    ]
  }
  return config
}
