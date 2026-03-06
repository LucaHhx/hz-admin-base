import fs from 'fs-extra'
import path from 'path'

export function transformLocalesPlugin() {
  return {
    name: 'vite-plugin-transform-locales',
    closeBundle() {
      const srcDir = path.resolve(process.cwd(), 'src/locales')
      const destDir = path.resolve(process.cwd(), 'dist/src/locales')

      // 使用回调方式而不是 async/await
      fs.copy(srcDir, destDir)
        .then(() => {
          console.log('Locales directory copied successfully!')
        })
        .catch((err) => {
          console.error('Error copying locales directory:', err)
        })
    }
  }
}
